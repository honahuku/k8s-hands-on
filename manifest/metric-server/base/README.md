# metric-server
kind で metric-server を動かすための設定

argocd をすでにデプロイしていれば sync する  
していなければ以下を実行  
```bash
kubectl apply -k manifest/metric-server
```
