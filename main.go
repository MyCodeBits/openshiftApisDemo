package main

import (
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	// rbacv1 "k8s.io/api/rbac/v1"
)

func main() {
	fmt.Println("Openshift APIs Demo ....")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	fmt.Println("=======================================================")
	fmt.Println("Entered Main() ...")
	// Instantiate loader for kubeconfig file.
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	fmt.Println("Got kubeconfig ..: ")
	fmt.Println(kubeconfig)

	// Determine the Namespace referenced by the current context in the
	// kubeconfig file.
	namespace, _, err := kubeconfig.Namespace()
	if err != nil {
		panic(err)
	}
	fmt.Println("Got Namespace ..: ")
	fmt.Println(namespace)

	// Get a rest.Config from the kubeconfig file.  This will be passed into all
	// the client objects we create.
	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Got restconfig ..: ")
	fmt.Println(restconfig)

	// Create a Kubernetes core/v1 client.
	coreclient, err := corev1client.NewForConfig(restconfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("Got client ::")
	fmt.Println(coreclient)

	// INSTANTIATE THE TEMPLATE.

	// 2. Create Secrets

	// 3. Create Roles
	// Ref : https://github.com/openshift/local-storage-operator/blob/master/pkg/controller/controller.go


	// 4. Create Svc Account
	var svcAccName = "test_sc"
	var labels = map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println("Creating Svc Account")

	svcAcc, err := coreclient.ServiceAccounts(namespace).Create(&corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: svcAccName,
			Namespace: namespace,
			Labels: labels,
		},
	})
	if err != nil {
		fmt.Println("Got error while creating Svc Account...")
		panic(err)
	}
	fmt.Println(svcAcc)
	fmt.Println("Returning ...")

	/*
		// To set Template parameters, create a Secret holding overridden parameters
		// and their values.
		secret, err := coreclient.Secrets(namespace).Create(&corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name: "parameters",
			},
			StringData: map[string]string{
				"MEMORY_LIMIT": "1024Mi",
			},
		})
		if err != nil {
			panic(err)
		}

		fmt.Println(secret)
	*/

	log.Fatal(http.ListenAndServe(":8081", nil))


}
