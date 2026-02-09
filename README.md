# DevOps AI Stack
**Local-First MLOps & AIOps Platform with Agentic LLM-Driven Remediation**

---

## 📌 What This Platform Is

This project is a **production-oriented, local-first DevOps AI platform** designed to automate
**first-level incident response** using Kubernetes-native GitOps practices and on-cluster LLMs.

It focuses on:
- Alert triage
- Log summarization
- Root cause reasoning
- Safe, automated remediation
- Human-in-the-loop escalation
- Incident documentation (RCA / postmortem)

All without sending data to external AI services.

---

## 🧠 Why These Components Were Chosen

This stack is intentionally opinionated. Each component solves a specific problem in a way that
scales cleanly from local development to production.

### 🧭 Argo CD — GitOps Control Plane

**Why Argo CD**
- Declarative, auditable deployments
- Clear separation between desired vs actual state
- First-class Kubernetes-native GitOps

**What it does here**
- Bootstraps the entire platform using an App-of-Apps pattern
- Manages infrastructure, apps, and models from Git
- Enforces consistency across environments

ArgoCD is the backbone that makes everything reproducible.

---

### 🤖 Ollama — Secure Local LLM Runtime

**Why Ollama**
- Runs fully on-cluster
- No external API calls
- Supports multiple models concurrently
- Simple HTTP API
- Easy to cache models locally

**What it does here**
- Executes LLM inference for automation workflows
- Hosts multiple specialized models (reasoning vs summarization)
- Acts as the AI decision engine for remediation

Ollama allows you to treat AI like infrastructure, not a SaaS dependency.

---

### 🔁 n8n — Automation & Orchestration Engine

**Why n8n**
- Visual, auditable workflows
- Easy API integrations
- Kubernetes-friendly
- Ideal for event-driven automation

**What it does here**
- Receives alerts from Prometheus
- Calls Ollama with structured prompts
- Branches logic based on AI output
- Triggers remediation scripts
- Sends notifications
- Creates tickets and documentation

n8n is the glue between signals, AI, and actions.

---

### 📊 Prometheus — Monitoring & Alerting

**Why Prometheus**
- Industry standard for Kubernetes observability
- Pull-based, reliable metrics
- Powerful alerting model

**What it does here**
- Generates alerts that trigger AI workflows
- Provides metrics for system health
- (Planned) Model-level metrics for AI observability

Prometheus supplies the signals that activate the AI.

---

### 📈 Grafana — Visualization & Insight

**Why Grafana**
- Best-in-class dashboards
- Works natively with Prometheus
- Useful for both humans and AI context

**What it does here**
- Visualizes infrastructure health
- Displays alert history
- (Planned) Shows AI decision metrics

Grafana provides visibility and confidence.

---

### 🌐 Ingress NGINX — Traffic Entry Point

**Why Ingress NGINX**
- Battle-tested ingress controller
- Wide ecosystem support
- Works well locally and in production

**What it does here**
- Exposes ArgoCD, n8n, Grafana, Ollama (optional)
- Provides a consistent access layer

---

### 🔐 cert-manager — TLS Automation

**Why cert-manager**
- Automates certificate lifecycle
- Kubernetes-native CRDs
- Production-ready TLS management

**What it does here**
- Manages TLS for ingress endpoints
- Enables secure internal and external traffic

---

### 💬 Slack & Telegram Connectors — Human Notification

**Why both**
- Slack for enterprise on-call workflows
- Telegram for lightweight / mobile alerts

**What they do here**
- Notify humans of incidents
- Provide AI-generated summaries
- Enable escalation when automation stops

---

## 📁 Repository Structure

```
.
├── applications/              # ArgoCD Application definitions
├── argocd/                    # App-of-Apps bootstrap
├── models/
│   └── ollama/                # GitOps-managed model pull Jobs
├── mcp-agent/                 # Remediation logic & scripts
├── devops-ai-stack-helmified/ # Optional internal Helm charts
├── secrets/                   # ExternalSecrets definitions
├── kind/                      # Local cluster config
├── k3s/                       # Edge cluster config
└── README.md
```

---

## 🤖 Model Strategy (GitOps)

Models are treated as **versioned infrastructure artifacts**.

### Current Models

| Model | Purpose |
|------|--------|
| llama3:8b | Deep reasoning & remediation decisions |
| phi3:mini | Fast alert classification & summaries |

### Why GitOps for Models
- Full auditability
- Deterministic startup
- No manual pulls
- Easy rollback

Models are pulled via Kubernetes Jobs and cleaned up using TTL.

---

## 🧪 Execution Flow

1. Prometheus fires an alert
2. n8n workflow starts
3. Ollama model is selected
4. AI analyzes context
5. Remediation is proposed or executed
6. Notifications are sent
7. Incident documentation is generated

---

## 🏁 Why This Architecture Matters

This platform demonstrates:
- Practical AIOps (not demos)
- Secure, local-first AI usage
- GitOps discipline
- Clear separation of concerns
- A realistic path to production

It is designed to be **explained, defended, and extended**.

---

