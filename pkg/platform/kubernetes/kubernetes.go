package kubernetes

import (
	"context"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/worker"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Kubernetes struct {
	Client *kubernetes.Clientset
	// Namespace in which worker nodes are to be deployed.
	Namespace string
	Config    config.StaticConfig
}

// https://stackoverflow.com/questions/53283347/how-to-get-current-namespace-of-an-in-cluster-go-kubernetes-client
func Namespace() string {
	// This way assumes you've set the POD_NAMESPACE environment variable using the downward API.
	// This check has to be done first for backwards compatibility with the way InClusterConfig was originally set up
	if ns, ok := os.LookupEnv("POD_NAMESPACE"); ok {
		return ns
	}

	// Fall back to the namespace associated with the service account token, if available
	if data, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace"); err == nil {
		if ns := strings.TrimSpace(string(data)); len(ns) > 0 {
			return ns
		}
	}

	return "default"
}

func NewKubernetes(cfg config.StaticConfig) (d *Kubernetes, err error) {
	c, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, err
	}

	namespace := cfg.Platform.Kubernetes.Namespace
	if namespace == "" {
		namespace = Namespace()
	}

	return &Kubernetes{Client: clientset, Namespace: namespace, Config: cfg}, nil
}

func (k *Kubernetes) DeployWorkers(ctx context.Context, info worker.Info) error {
	name := info.Label + "-" + info.Topic
	labels := map[string]string{"scoretrak_worker": info.Label}
	encWorkerCfg, err := util.GenerateEncodedWorkerCfg(k.Config, info)
	if err != nil {
		return err
	}
	_, err = k.Client.AppsV1().DaemonSets(k.Namespace).Create(ctx,
		&appv1.DaemonSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: k.Namespace,
			},
			Spec: appv1.DaemonSetSpec{
				Selector: &metav1.LabelSelector{
					MatchLabels: labels,
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Name:      name,
						Namespace: k.Namespace,
						Labels:    labels,
					},
					Spec: corev1.PodSpec{
						Tolerations:  []corev1.Toleration{{Key: "node-role.kubernetes.io/master", Effect: corev1.TaintEffectPreferNoSchedule}},
						NodeSelector: labels,
						Containers: []corev1.Container{
							{
								Name:  "worker",
								Image: util.Image,
								Args:  []string{"worker", "--encoded-config", encWorkerCfg},
							},
						},
					},
				},
			},
		}, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (k *Kubernetes) RemoveWorkers(ctx context.Context, info worker.Info) error {
	return k.Client.AppsV1().DaemonSets(k.Namespace).Delete(ctx, info.Label+"-"+info.Topic, metav1.DeleteOptions{})
}
