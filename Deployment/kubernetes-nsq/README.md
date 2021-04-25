Note:
    
    1) At the moment this guide does not show a secure installation of scoretrak and is only for dev purposes. 

Requirements: 

    1) Install Helm
    2) ENSURE THAT TIME BETWEEN NODES IS SYNCHRONIZED!
    You can use something like https://github.com/geerlingguy/ansible-role-ntp to automate the process
    

1) Kubernetes Dashboard
    1) Install Metrics:
        ```
        kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/download/v0.3.6/components.yaml
        ```
    
    2) Install kubernetes dashboard (https://github.com/kubernetes/dashboard):
        ```
        kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.3/aio/deploy/recommended.yaml
        ```
    
    3) Next, run 
        ```
        kubectl proxy
        ```
    
    4) Now access Dashboard by navigating to:
        ```
        http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/
        ```
    
    5) Create admin user:
        ```
        kubectl apply -f admin-user.yml
        ```
    
    6) Extract admin token and paste it into web GUI:
        ```
        kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')
        ```
       in case the token expires, you will need to fetch the same way
    
2) Service Mesh. We will need to configure service mesh to allow for greater monitoring of the traffic. In addition, it will allow our application to utilize grpc web, with client browsers, that at the moment do not natively support grpc protocol.
    1) Follow setup process to get istio up and running here: https://istio.io/latest/docs/setup/getting-started/
    2) Make sure you don't forget the https://istio.io/latest/docs/setup/getting-started/#dashboard section to enable kiali dashboard
    3) https://github.com/venilnoronha/grpc-web-istio-demo
    4) Run `kubectl edit configmaps istio-sidecar-injector -n istio-system` and add following configuration to `neverInjectSelector`:
    ```
      neverInjectSelector:
      - matchExpressions:
        - {key: job-name, operator: Exists}
    ```
    This allows us to skip sidecar injection for job containers
3) Database
    1) Download cockroachdb helm charts
       ```
       helm repo add cockroachdb https://charts.cockroachdb.com/
       ```
    2) Update helm charts:
       ```
       helm repo update
       ```
    3) Create persistent volume claims:
       ```
       bash CreateCRDBPersistentVolumeAndClaims.sh
       ```
    4) Install cockroach using helm (Don't wait for the command to finish, and move to step 5):
       ```
       helm install scoretrak --values cockroach-helm-values.yml cockroachdb/cockroachdb
       ```

    5) Make sure you see 4 CSRs when before proceeding to the next step by running: `kubectl get csr`.
       
    6) We need to deploy cluster in secure mode when using istio due to this issue(https://github.com/cockroachdb/cockroach/issues/19667). Approve CSRs made by cockroachdb server:
        ```
        kubectl certificate approve default.node.scoretrak-cockroachdb-0
        kubectl certificate approve default.node.scoretrak-cockroachdb-1
        kubectl certificate approve default.node.scoretrak-cockroachdb-2
        sleep 20;
        kubectl certificate approve default.client.root
        ```
        Make sure you see that every instance has 1/1 ready, like this:
        ```
        NAME                               READY   STATUS      RESTARTS   AGE
        scoretrak-cockroachdb-0            1/1     Running     0          50s
        scoretrak-cockroachdb-1            1/1     Running     0          50s
        scoretrak-cockroachdb-2            1/1     Running     0          50s
        scoretrak-cockroachdb-init-4qm9c   0/1     Completed   0          50s        
        ```
        (If that is not the case, you might have to redo everything starting with step 3. Make sure to follow this guide https://www.cockroachlabs.com/docs/stable/orchestrate-cockroachdb-with-kubernetes.html#stop-the-cluster to delete the failed deployment using helm. Make sure to delete secrets.)
        
        If you would like to check cockroach web console, you could run:
           
        ```
        kubectl port-forward service/scoretrak-cockroachdb-public 8080:8080
        ```
        to port forward remote port to your local port. Same applies for port   which is cockroachdb port.
        
    7) Setup Databases:
       1) Run:
        ```
        kubectl create -f client-secure.yaml
        ```
       2) Run:
       ```
       kubectl exec -it cockroachdb-client-secure \
        -- ./cockroach sql \
        --certs-dir=/cockroach-certs \
        --host=scoretrak-cockroachdb-public
       ```
       3) Generate secure password, and run following command(make sure to replace SOME_SECURE_PASSWORD) Execute:
       ```
       CREATE DATABASE IF NOT EXISTS scoretrak;
       ALTER USER root WITH PASSWORD 'SOME_SECURE_PASSWORD';
       \q;
       ```
       4) Inside of config.yml replace SOME_SECURE_PASSWORD with your generated password
       5) Run: 
        ```
       kubectl delete -f client-secure.yaml
        ```
       6) Run (Note: this is a workaround to create pass certs to scoretrak https://github.com/kubernetes/kubernetes/issues/66020):
       ```
       kubectl apply -f cockroachdb-cert-loader-api.yaml
       ```
       We will use the generated above key when deploying scoretrak in later sections
    
   
4) Deploy NSQ
    1) We will use a modified version of the following repo https://github.com/adrianchifor/k8s-nsq. To do that, just run:
       ```
       kubectl create -f nsq.yaml
       ```
       Likewise, if you would like to access the admin panel you can execute:
       ```
       kubectl port-forward $(kubectl get pods | grep nsqadmin |  awk '{print $1}') 4171:4171
       ```
    
5) Setting up scoretrak
    1) Create Config map:
       ```
       kubectl create configmap scoretrak-config --from-file=./config.yml
       ```
    2) Deploy Scoretrak master component:
       ```
       kubectl create -f scoretrak.yaml
       ```
    3) Run `kubectl apply -f istio.yaml` to deploy necessary gateways, and virtual services


6) Configuring istio.
    0) Change `hosts` in istio.yaml, and specify the exact hostname of your engine website. (You can leave it blank for dev.)
    1) Simply Run:
       ```
       kubectl apply -f istio.yaml
       ```
    2) If you would like to configure ingress port type:
       ```
       kubectl edit svc istio-ingressgateway -n istio-system
       ```
       Change the value of the port 80
       
   

7) (Optional) Configure External Load Balancer. (This step will depend heavily on your environment, hence this step is up to the devs to implement)
   Notes:
    - Make sure idle time out is set to something large (Ex: 5 hours). This is needed because ScoreTrak utilizes gRPC with server push.
    - Make sure your proxy supports HTTP2.0
8) Labels.
    1) Using the dashboard, add labels to desired nodes. The labels should follow the format:
    scoretrak_worker: <YOUR_LABEL_VALUE>, for instance scoretrak_worker: internal, where internal can represet a set of workers responsible for scoring internal services.
    
    2) Go to scoretrak web's "Service Groups" section, and add a new entry with the label specified above.
    3) Once added, go to Dashboard, and ensure Daemonset named <YOUR_LABEL_VALUE>-<YOUR_SERVICE_GROUP_NAME> was created
