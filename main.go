package main

import (
	"encoding/json"
	//"encoding/json"
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
	fmt.Println("main()::RedHat Openshift APIs Demo ....")

	fmt.Println("main()::=======================================================\n")
	// Instantiate loader for kubeconfig file.
	fmt.Println("main()::Reading kubeconfig ..: ")
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	fmt.Println("main()::Done reading kubeconfig\n")

	// Determine the Namespace referenced by the current context in the kubeconfig file.
	namespace, _, err := kubeconfig.Namespace()
	if err != nil {
		panic(err)
	}
	fmt.Printf("main()::Current Namespace : %s\n", namespace)

	// Get a rest.Config from the kubeconfig file.  This will be passed into all
	// the client objects we create.
	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("main()::Created clientConfig\n")
	fmt.Println(restconfig)

	// Create a Kubernetes core/v1 client.
	coreclient, err := corev1client.NewForConfig(restconfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("main()::Created client\n")

	// INSTANTIATE THE TEMPLATE.

	// 2. Create Secrets

	// 3. Create Roles
	// Ref : https://github.com/openshift/local-storage-operator/blob/master/pkg/controller/controller.go


	// 4. Create Svc Account
	var svcAccName = "prog-sc"
	var svcAccLabels = map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	createServiceAccount(coreclient, namespace, svcAccName, svcAccLabels)

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


	fmt.Println("main():: Done. Returning ...")
}

// Create a Service Account
func createServiceAccount(coreclient *corev1client.CoreV1Client, namespace string, svcAccName string, svcAccLabels map[string]string) {
	fmt.Printf("main()::Creating Svc Account : %s", svcAccName)


	svcAcc, err := coreclient.ServiceAccounts(namespace).Create(&corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: svcAccName,
			Namespace: namespace,
			Labels: svcAccLabels,
		},
	})
	if err != nil {
		fmt.Println("\nmain()::Received error while creating Svc Account.")
		panic(err)
	}
	fmt.Printf("main()::Created Svc Account : %s\n", svcAccName)
	fmt.Println(svcAcc)
}

// Lists Service Accounts
func listServiceAccounts(coreclient *corev1client.CoreV1Client, namespace string) {
	//coreclient.Pods(namespace).Get("aa", metav1.GetOptions{})
  fmt.Printf("Listing ServiceAccounts in NS : %s", namespace)

	svcList, err := coreclient.ServiceAccounts(namespace).List(
		metav1.ListOptions{},
	)
	if (err != nil) {
		fmt.Println("main()::Got error while listing Svc Account...")
		panic(err)
	}
	fmt.Println("Listing Svc Accounts : ")
	fmt.Println(svcList)
	svcListMar, _ := json.Marshal(svcList)
	fmt.Println(svcListMar)
}
