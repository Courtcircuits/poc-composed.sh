package main

import "fmt"

func main() {
	docker_compose := openFile("docker-compose.yml")
	compose := parse(docker_compose)

	for name, svc := range compose.Services {
		deployDeployment(svc, name)
		deployService(svc, name)

		fmt.Println("Deployed", name)
	}
	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()

	// // use the current context in kubeconfig
	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// // create the clientset
	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	// deployment := &appsv1.Deployment{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name: "demo-deployment",
	// 	},
	// 	Spec: appsv1.DeploymentSpec{
	// 		Replicas: int32Ptr(2),
	// 		Selector: &metav1.LabelSelector{
	// 			MatchLabels: map[string]string{
	// 				"app": "demo",
	// 			},
	// 		},
	// 		Template: apiv1.PodTemplateSpec{
	// 			ObjectMeta: metav1.ObjectMeta{
	// 				Labels: map[string]string{
	// 					"app": "demo",
	// 				},
	// 			},
	// 			Spec: apiv1.PodSpec{
	// 				Containers: []apiv1.Container{
	// 					{
	// 						Name:  "web",
	// 						Image: "nginx:1.12",
	// 						Ports: []apiv1.ContainerPort{
	// 							{
	// 								Name:          "http",
	// 								Protocol:      apiv1.ProtocolTCP,
	// 								ContainerPort: 80,
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// fmt.Println("creating deployment...")
	// result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

}

func int32Ptr(i int32) *int32 { return &i }
