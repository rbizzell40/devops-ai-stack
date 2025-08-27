# ✅ Environment Compatibility Matrix

| Component        | KIND (Local) | K3s (Edge/VM) | Notes                        |
|------------------|--------------|---------------|------------------------------|
| ArgoCD           | ✅           | ✅            | GitOps works on both         |
| Helm Charts      | ✅           | ✅            | Unified via values.yaml      |
| Ingress-NGINX    | ✅           | ✅            | HostPorts mapped in KIND     |
| Cert-Manager     | ✅           | ✅            | Needs ClusterIssuer setup    |
| Ollama           | ✅           | ✅            | Port 11434 exposed           |
| n8n              | ✅           | ✅            | Accessible via Ingress       |
| Prometheus       | ✅           | ✅            | Full monitoring supported    |
| Telegram/Slack   | ✅           | ✅            | Secrets managed via ESO      |
| ESO (Secrets)    | ✅           | ✅            | Fake or real backends        |
| MCP Actions      | ✅           | ✅            | Tested via demo.sh script    |