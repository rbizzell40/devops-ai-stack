# restart_pod.py
import argparse
from kubernetes import client, config

def restart_pod(namespace, pod_name):
    config.load_kube_config()
    v1 = client.CoreV1Api()

    print(f"Deleting pod '{pod_name}'in namespace '{namespace}'...")
    try:
        v1.delete_namespaced_pod(name=pod_name, namespace=namespace)
        print("✅ Pod restart triggered successfully.")
    except Exception as e:
        print(f"❌ Failed to restart pod: {e}")

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--namespace", required=True, help="Kubernetes namespace")
    parser.add_argument("--pod", required=True, help="Pod name")
    args = parser.parse_args()

    restart_pod(args.namespace,args.p)

def restart_pod(namespace, pod_name):
    config.load_kube_config()
    v1 = client.CoreV1Api()

    print(f"Deleting pod '{pod_name}'in namespace '{namespace}'...")
    try:
        v1.delete_namespaced_pod(name=pod_name, namespace=namespace)
        print("✅ Pod restart triggered successfully.")
    except Exception as e:
        print(f"❌ Failed to restart pod: {e}")

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--namespace", required=True, help="Kubernetes namespace")
    parser.add_argument("--pod", required=True, help="Pod name")
    args = parser.parse_args()

    restart_pod(args.namespace,args.pod)