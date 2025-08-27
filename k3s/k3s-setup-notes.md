# 🐳 K3s Setup for DevOps AI Stack

## 1. Install K3s (Single Node)

```bash
curl -sfL https://get.k3s.io | sh -s - --disable traefik
```
- `traefik` is disabled so you can use your own Ingress controller (NGINX or Traefik)

## 2. Get Kubeconfig

```bash
sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
sudo chown $(id -u):$(id -g) ~/.kube/config
```

## 3. Install Helm & ArgoCD

```bash
helm repo add argo https://argoproj.github.io/argo-helm
helm upgrade --install argo-cd argo/argo-cd -n argocd --create-namespace
```

## 4. Install Your DevOps Stack

```bash
kubectl apply -f argo-apps/apps-of-apps.yaml
```

## 5. Set Up Ingress, TLS, and Cert Manager

Install `ingress-nginx`, `cert-manager`, then expose n8n and ollama services.