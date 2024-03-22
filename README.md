# k8s-hands-on

## bootstrap
クラスタが出来てからの bootstrap手順  
kind でクラスタを作成する場合はこっち → [manifest/kind/README.md](manifest/kind/README.md)  

```bash
aqua i -l
# argocd をデプロイ
kubectl apply -k manifest/argocd/overlays/production
# 初期パスワードを取得
argocd admin initial-password -n argocd
# argocd namespace にある argocd-server を port-forward する
k9s
# 他の apps を全部展開
kubectl apply -k manifest/apps/overlays/production
# 以降はリポジトリにコミットで同期される
```
最初から apps をデプロイすると 
