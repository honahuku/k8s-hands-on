# k8s-scale-test
事前に [manifest/metric-server](manifest/metric-server/base/README.md) をデプロイしておく

```bash
kubectl run -i --tty load-generator --rm --image=busybox:1.28 --restart=Never -- /bin/sh -c "while sleep 0.01; do wget -q -O- http://php-apache.k8s-scale-test; done"
```
