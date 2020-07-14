package platforming

type Config struct {
	Use    string `default:"none"`
	Docker struct {
		Name    string `default:"scoretrak"`
		Host    string `default:"unix:///var/run/docker.sock"`
		Network string `default:"default"`
	}
	Kubernetes struct {
		Namespace string `default:"default"`
	}
}
