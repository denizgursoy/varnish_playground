# varnish_playground


curl -XGET localhost:8090/cached

curl -XGET localhost:8090/uncached

curl -XPURGE localhost:8090/cached 

https://github.com/mittwald/kube-httpcache

https://kruyt.org/varnish-kuberenets/

https://ibm.github.io/varnish-operator/installation.html



docker build -t varnish-backend:0.0.1 .
