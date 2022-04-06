package utils

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	bootstrapsecretutil "k8s.io/cluster-bootstrap/util/secrets"
	"path/filepath"
)

func GetTokenSecret() map[string]string {
	// get kube-system bootstrap-token-<token id> from k8s

	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	selector := fields.Set{
		"type": "bootstrap.kubernetes.io/token",
	}.AsSelector().String()
	options := metav1.ListOptions{FieldSelector: selector}
	secretList, err := clientset.CoreV1().Secrets("kube-system").List(context.TODO(), options)
	if err != nil {
		panic(err.Error())
	}
	sigs := map[string]string{}
	for _, secret := range secretList.Items {
		if secret.Type == "bootstrap.kubernetes.io/token" {
			sigs[bootstrapsecretutil.GetData(&secret, "token-id")] = bootstrapsecretutil.GetData(&secret, "token-secret")
		}
	}

	return sigs
}
