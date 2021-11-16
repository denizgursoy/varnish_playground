# varnish_playground


curl -XGET localhost:8090/cached

curl -XGET localhost:8090/uncached

curl -XPURGE localhost:8090/cached 

https://github.com/mittwald/kube-httpcache

https://kruyt.org/varnish-kuberenets/

https://ibm.github.io/varnish-operator/installation.html


k delete -f varnish-k8s.yaml
docker build -t varnish-backend:0.0.1 .
k apply -f varnish-k8s.yaml
k port-forward svc/backend-service   8091:80
http://localhost:8091/cached