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
    
2) Database
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
       bash CreateCockroachDBPersistentVolumeAndClaims.sh
       ```
    3) Install cockroach using helm:
       ```
       helm install scoretrak --values cockroach-helm-values.yml cockroachdb/cockroachdb
       ```
       If you would like to check cockroach web console, you could run:
       ```
       kubectl port-forward service/scoretrak-cockroachdb-public 8080:8080
       ```
       to port forward remote port to your local port. Same applies for port   which is cockroachdb port.
   
    4) Setup Databases
       ```
        kubectl create -f setup-db-job.yaml
       ```
   
3) Deploy NSQ
    1) We will use a modified version of the following repo https://github.com/adrianchifor/k8s-nsq. To do that, just run:
       ```
       kubectl create -f deploy-nsq.yaml
       ```
       Likewise, if you would like to access the admin panel you can execute:
       ```
       kubectl port-forward $(kubectl get pods | grep nsqadmin |  awk '{print $1}') 4171:4171
       ```
    
4) Setting up scoretrak
    1) Create service role for scoretrak master
       ```
       kubectl create -f scoretrak-role.yaml
       ```
    2) Create Config map:
       ```
       kubectl create configmap scoretrak-config --from-file=./config.yml
       ```
    3) Deploy Scoretrak master component:
       ```
       kubectl create -f scoretrak.yaml
       ```
       If you would like to access scoretrak, you can do so by running:
       ```
       kubectl port-forward service/scoretrak 33333:33333
       ```
    4) Deploy Scoretrak web component
       ```
       kubectl create -f scoretrak-web.yaml
       ```
       Access the web component via port forwarding like so open up the browser and input
       any node IP followed by port 30080
     
5) (Optional) Configure External Load Balancer. (This step will depend heavily on your environment, hence this step is up to the devs to implement)

6) Labels.
    1) Using the dashboard, add labels to desired nodes. The labels should follow the format:
    scoretrak_worker: <YOUR_LABEL_VALUE>, for instance scoretrak_worker: internal, where internal can represet a set of workers responsible for scoring internal services.
    
    2) Go to scoretrak web's "Service Groups" section, and add a new entry with the label specified above.
    3) Once added, go to Dashboard, and ensure Daemonset named <YOUR_LABEL_VALUE>-<YOUR_SERVICE_GROUP_NAME> was created