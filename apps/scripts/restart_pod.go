// restart_pod.go
package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	namespace := flag.String("namespace", "default", "Kubernetes namespace")
	podName := flag.String("pod", "", "Pod name")
	flag.Parse()

	if *podName == "" {
		log.Fatal("❌ Pod name is required using --pod")
	}

	kubeconfig := filepath.Join(homeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("❌ Error loading kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("❌ Error creating clientset: %v", err)
	}

	fmt.Printf("🔁 Restarting pod: %s in namespace: %s\n", *podName, *namespace)
	err = clientset.CoreV1().Pods(*namespace).Delete(*podName, metav1.DeleteOptions{})
	if err != nil {
		log.Fatalf("❌ Error deleting pod: %v", err)
	}
	fmt.Println("✅ Pod restart triggered successfully.")
}

func homeDir() string {
	return filepath.Dir("/root")
}
