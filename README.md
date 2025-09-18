<<<<<<< HEAD
# 📊 DevOps AI Stack Demo: Narrated Walkthrough

This guide supports the `devops-ai-stack-demo.pptx` presentation with talking points and explanations for each slide.

---

## 🎬 Slide 1: Title
**Welcome!** Today we’re introducing the **DevOps AI Stack**, an automation-first platform designed for secure, intelligent, and resilient infrastructure management.

---

## ⚙️ Slide 2: Stack Overview

**What's included?**

- **ArgoCD** – GitOps deployment & drift detection.
- **Helm** – All apps are defined and deployed via Helm.
- **n8n** – The orchestration brain to trigger and chain actions.
- **Ollama (LLM)** – Responds to prompts and suggests system actions.
- **Prometheus + Grafana** – Observability and metrics.
- **Ingress Controller** – NGINX or Traefik for exposing services.
- **Cert-Manager** – Automatic TLS.
- **ESO** – Secure secret handling from any backend.
- **Slack & Telegram** – Alerting & feedback loop for human operators.

---

## 🚨 Slide 3: Use Case – Automated Incident Response

Imagine a **high memory alert** fires:

1. Prometheus triggers a **webhook**.
2. **n8n** receives the event and sends prompt to Ollama.
3. LLM recommends a **restart of a failing deployment**.
4. n8n runs a **Kubernetes command** to fix the issue.
5. Alert sent to **Slack & Telegram** for visibility.

💡 No human intervention needed — but humans are notified.

---

## 🧱 Slide 4: Deployment Environments

This stack works **locally or at the edge**:

- **KIND** – Great for testing, CI pipelines, and local demos.
- **K3s** – Lightweight Kubernetes for edge, IoT, or VMs.

Compatible with Ingress, TLS, ESO, GitOps.

---

## 🧪 Slide 5: Demo Script

Run the full MCP pipeline using:

```bash
./demo.sh
```

- Fires a webhook (simulating alert)
- n8n routes it to LLM and action chain
- Kubernetes actions applied
- Notifications sent

Use ArgoCD or `kubectl` to observe results.

---

## 🔐 Slide 6: Security & Observability

- All ingress traffic is encrypted (TLS via cert-manager)
- Secrets are injected securely via **External Secrets Operator (ESO)**
- Prometheus captures metrics
- Grafana provides visualization dashboards
- **GitOps ensures repeatability + auditability**

---

## 💡 Slide 7: Why This Stack?

- Works **entirely local** – no cloud dependencies
- Uses **LLMs to act**, not just alert
- Built for **production GitOps** and **local demoing**
- Modular, pluggable, scalable
- Built with **security, observability**, and **automation** in mind

---

🎤 You’re ready to demo the future of AI-driven infrastructure.
=======
# devops-ai-stack
>>>>>>> bd7b7e88bafa0f04f1dce9fadc86922d11423df0


kubectl -n argocd get secret argocd-initial-admin-secret \
-o jsonpath="{.data.password}" | base64 -d && echo
