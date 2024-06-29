package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path"
)

var (
	namespace = "kube-system"
	timeout   = int64(1)
)

func GetClientSet() (*kubernetes.Clientset, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	config, err := clientcmd.BuildConfigFromFlags("", path.Join(home, ".kube/config"))
	if err != nil {
		panic(err.Error())
	}

	return kubernetes.NewForConfig(config)
}

func main() {
	// Start with the ClientSet
	client, err := GetClientSet()
	if err != nil {
		panic(err)
	}

	// Fetch configmap from kube-system namespace
	cms, err := client.CoreV1().
		ConfigMaps(namespace).
		List(
			context.Background(),
			metav1.ListOptions{Limit: 3, TimeoutSeconds: &timeout},
		)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Listing the Configmaps", namespace)
	for _, cm := range cms.Items {
		fmt.Println("- ", cm.Name)
	}
}
