package main

import (
	"encoding/json"

	"fmt"
	"github.com/ghodss/yaml"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"net/http"
	v1 "k8s.io/api/rbac/v1" // TODO. Named it v1 as typecasting issues. Fix this. Earlier was rbacv1
)

func main() {
	fmt.Println("main()::RedHat Openshift APIs Demo ....")


	filePath := "/Users/sshridhar/code/nonHWorks/gitlab/openshiftApisDemo/resources/openshift/request201.yml"
	fmt.Println("main()::Reading NamespaceCrp file : ", filePath)
	nsCrp, nsCrpErr := readAndMapCrpFile(filePath)
	if nsCrpErr != nil {
		log.Fatalf("error: %v", nsCrpErr)
	}
	fmt.Println("main():: Done reading NamespaceCrp file")


	// Instantiate loader for kubeconfig file.
	fmt.Println("main()::Reading kubeconfig")
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

	// Create a Kubernetes core/v1 client.
	coreClient, err := corev1client.NewForConfig(restconfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("main()::Created client\n")

	// INSTANTIATE THE TEMPLATE.

	// 2. Create Secrets

	// 3. Create Roles
	// Ref : https://github.com/openshift/local-storage-operator/blob/master/pkg/controller/controller.go

	// Create rbac client
	rbacClient, rbacClientErr := rbacv1client.NewForConfig(restconfig)
	if rbacClientErr != nil {
		panic(err)
	}

	svcAccsInfoList := nsCrp.Spec.Permissions[0].Serviceaccounts
	//fmt.Println("main():: secAccInfoList : %s", secAccountsInfoList)

	fmt.Printf("main()::Creating passed-in Service Account(s) ..... ")
	for svcAccsItr := 0; svcAccsItr < len(svcAccsInfoList); svcAccsItr++ {
		// TODO : SWAPAN check for returned error
		//createServiceAccount(coreClient, namespace, svcAccsInfoList[svcAccsItr].Name, svcAccsInfoList[svcAccsItr].Labels)
		fmt.Println("aa")
		// Create associated Role(s) and Role Bindings(s).
		// TODO : SWAPAN check for returned error before proceeding
		for rolesItr :=0; rolesItr < len(svcAccsInfoList[svcAccsItr].Roles); rolesItr ++{
			fmt.Println("bb")
			rbacClient.RESTClient()
			//createRole(rbacClient, namespace, &(svcAccsInfoList[svcAccsItr].Roles[rolesItr]))
			//createRoleBinding(rbacClient, namespace, svcAccsInfoList[svcAccsItr].Name, &(svcAccsInfoList[svcAccsItr].Roles[rolesItr]))
		}
    //
		//fmt.Println(secAccInfo)
	}


	/*
	var roleName = "prog-role1"

	if roleErr != nil {
		panic(err)
	}
	fmt.Println(role)

	*/

	// 4. Create Svc Account
	//var svcAccName = "prog-sc"
	//var svcAccLabels = map[string]string{
	//	"key1": "value1",
	//	"key2": "value2",
	//}

	//createServiceAccount(coreClient, namespace, svcAccName, svcAccLabels)

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


func readAndMapCrpFile(path string) (NamespaceCrp, error) {
	// Read roles file
	namespaceCrp := NamespaceCrp{}
	/*data, err := ReadFile(path)
	if err != nil {
	}*/

	var data = `apiVersion: compute.cloud.cloudera.io/v1alpha1
kind: NamespaceCRP
metadata:
  name: <tenant>.UNCLASSIFIED.<env>.ml-bio.ml-infra
  clusterId: cluster-id-200
spec:
  namespace:
    name: oc-svc-acc-create # SWAPAN
    labels:
      environment: dev
      key2: value2
  quota:
    min:
      memory: 10G
      cpu: 2
      nvidia.com/gpu: 1
    max:
      memory: 100G
      cpu: 10
      nvidia.com/gpu: 5
  permissions:
    - serviceaccounts:
        - name: ml-service-account
          labels:
            Key1: value1
            Key2: value2
          roles:
            - apiVersion: rbac.authorization.k8s.io/v1
              kind: Role
              metadata:
                namespace: oc-svc-acc-create # remove SWAPAN
                name : access-role-1
              rules:
                - apiGroups: [""]
                  resources: ["configmaps", "secrets", "services"]
                  verbs: ["get", "list", "create", "update", "delete"]
                - apiGroups: ["apiextensions.k8s.io"]
                  resources: ["customresourcedefinitions"]
                  verbs: ["get", "list", "watch", "create", "update", "delete"]
            - apiVersion: rbac.authorization.k8s.io/v1
              kind: ClusterRole
              metadata:
                namespace: oc-svc-acc-create # remove SWAPAN
                name : access-role-2
              rules:
                - apiGroups: [""]
                  resources: ["configmaps", "secrets", "services"]
                  verbs: ["get", "list", "create", "update", "delete"]
                - apiGroups: ["apiextensions.k8s.io"]
                  resources: ["customresourcedefinitions"]
                  verbs: ["get", "list", "watch", "create", "update", "delete"]
  resources:`

	//err = yaml.Unmarshal(data, &namespaceCrp)
	//return namespaceCrp, err

	// TODO : Swapan Check that we are converting to JSON first.
	// Ref : https://github.com/kubernetes/client-go/issues/193
	j2, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
  unMarErr := json.Unmarshal(j2, &namespaceCrp)
	if unMarErr != nil {
		panic(unMarErr)
	}

	return namespaceCrp, err
}

// TODO : SWAPAN Check Do we have RoleBinding only.. or we create ClusterRoleBinding also sometimes ?
func createRoleBinding(rbacClient *rbacv1client.RbacV1Client, namespace string, svcAccName string, role *v1.Role) {
	fmt.Printf("\n				main()::Creating RoleBinding  : '%s-role-binding' under Namespace : '%s' for Svc Acc : %s", role.Name, namespace, svcAccName)
	_, roleErr := rbacClient.RoleBindings(namespace).Create(&v1.RoleBinding{
		// TODO : What about TypeMeta.. Add it.
		ObjectMeta: metav1.ObjectMeta{
			Name: role.Name + "-role-binding", // TODO : SWAPAN check
			Namespace: namespace,
			// TODO : Labels ?
		},
		// TODO : SWAPAN Remove ServiceAccount hardcoding below
		Subjects: []v1.Subject{
			{
				Kind: "ServiceAccount",
				Name: svcAccName,
				APIGroup: "",
			}, // SWAPAN TODO : Keep it empty ? Where do we read these ?

		},
		RoleRef: v1.RoleRef{
			Kind: role.Kind,
			Name: role.Name,
			APIGroup: "rbac.authorization.k8s.io", // TODO : SWAPAN Remove hardcoding
		},
	})
	if roleErr != nil {
		fmt.Println("Error creating role Binding")
		fmt.Println(roleErr)
		panic(roleErr)
	}
	fmt.Printf("\n				main()::Creating RoleBinding  : '%s-role-binding' under Namespace : '%s' for Svc Acc : %s", role.Name, namespace, svcAccName)
}


// TODO : Tie the fn.s so that their type/object only can call them rather than having plain fns and passing the client
// Create
func createRole(rbacClient *rbacv1client.RbacV1Client, namespace string, role *v1.Role) {
	fmt.Printf("\n			main()::Creating Role  : '%s' under Namespace : '%s'", role.Name, namespace)
	_, roleErr := rbacClient.Roles(namespace).Create(role)
	if roleErr != nil {
		fmt.Println("Error creating role")
		fmt.Println(roleErr)
		panic(roleErr)
	}
	fmt.Printf("\n			main()::Created Role  : '%s' under Namespace : '%s'", role.Name, namespace)
}



// Create a Service Account
func createServiceAccount(coreClient *corev1client.CoreV1Client, namespace string, svcAccName string, svcAccLabels map[string]string) {
	fmt.Printf("\n		main()::Creating Svc Account : '%s'", svcAccName)

  // TODO : What about TypeMeta.. Add it.
	svcAcc, err := coreClient.ServiceAccounts(namespace).Create(&corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: svcAccName,
			Namespace: namespace,
			Labels: svcAccLabels,
		},
	})
	if err != nil {
		fmt.Println("\n		main()::Received error while creating Svc Account.\n")
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("		main()::Created Svc Account : '%s'\n\n", svcAccName)
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
