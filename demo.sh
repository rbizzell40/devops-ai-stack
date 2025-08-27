#!/bin/bash

# Demo Script: Trigger MCP Workflow via n8n webhook

echo "🔧 Triggering MCP workflow via webhook..."

# Simulate Prometheus alert or external trigger
curl -X POST http://localhost:5678/webhook/mcp-alert -H "Content-Type: application/json" -d '{
  "alertname": "HighMemoryUsage",
  "instance": "auth-api",
  "description": "Memory usage exceeded 90%"
}'

echo -e "\n✅ Alert sent to n8n workflow."

# Optionally check n8n logs (assuming you are using kubectl)
echo "📜 Fetching n8n logs..."
kubectl logs deployment/n8n -n devops

# Check result of Kubernetes command (e.g., restart)
echo "🔍 Verifying if deployment was restarted..."
kubectl rollout status deployment/auth-api -n default
