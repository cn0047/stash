package main

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	f := "/Users/user/.kube/config"
	n := ""

	c := getClient(f)
	//c := getInC`lusterClientSet()

	podsList(c, n)
	//podsGet(c, n)
	//podsCreate(c, n)
	//hrdCreate(c, n)
}

func getInClusterClientSet() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return c
}

func getClient(configFile string) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", configFile)
	if err != nil {
		panic(err.Error())
	}

	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return c
}

func podsList(c *kubernetes.Clientset, n string) {
	ctx := context.TODO()
	pods, err := c.CoreV1().Pods(n).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(fmt.Errorf("podsList failure, error: %w", err))
	}

	fmt.Printf("pods:\n")
	for _, p := range pods.Items {
		fmt.Printf("  - %v\n", p.Name)
	}
}

func podsGet(c *kubernetes.Clientset, n string) {
	ctx := context.TODO()
	p, err := c.CoreV1().Pods(n).Get(ctx, "ksh-pod-2", metav1.GetOptions{})
	if err != nil {
		panic(fmt.Errorf("podsList failure, error: %w", err))
	}

	fmt.Printf("[podsGet] found pod: %+v \n", p.Name)
}

func podsCreate(c *kubernetes.Clientset, n string) {
	ctx := context.TODO()
	p := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ksh-pod-2",
			Namespace: n,
			Labels:    map[string]string{"app": "ksh"},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{{
				Name:            "klog-pod-2",
				Image:           "cn007b/alpine",
				ImagePullPolicy: v1.PullAlways,
				Command: []string{
					"sh",
					"-c",
					"while true; do curl -i -XPOST 'https://realtimelog.herokuapp.com:443/rkc8q6llprn' -H 'Content-Type: application/json' -d '{\"code\":\"100\", \"status\": \"ok from pod\", \"env\":\"'$APP_ENV'\"}'; sleep 5; done",
				},
			}},
		},
	}
	p, err := c.CoreV1().Pods(n).Create(ctx, p, metav1.CreateOptions{})
	if err != nil {
		panic(fmt.Errorf("podsList failure, error: %w", err))
	}

	fmt.Printf("[podsGet] pod created: %+v \n", p.Name)
}

func hrdCreate(c *kubernetes.Clientset, n string) {
	// Deployment.
	d := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:   "hrd-deployment",
			Labels: map[string]string{"app": "hrd"},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RollingUpdateDeploymentStrategyType,
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": "hrd"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": "hrd"},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "hrd-container",
							Image: "ghcr.io/thepkg/hrd:v1.1.2",
							Ports: []corev1.ContainerPort{{ContainerPort: 80}},
							Env: []corev1.EnvVar{
								{Name: "APP_ENV", Value: "local"},
								{Name: "TEST_LABEL", Value: "xtest"},
								{Name: "APP_ORCHESTRATOR", Value: "k8s"},
							},
						},
					},
				},
			},
		},
	}
	_, err := c.AppsV1().Deployments("default").Create(context.TODO(), d, metav1.CreateOptions{})
	if err != nil {
		panic(fmt.Errorf("failed to create deployment: %v", err))
	}
	fmt.Println("Deployment created successfully!")

	// Service.
	s := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:   "hrd-service",
			Labels: map[string]string{"app": "hrd"},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{"app": "hrd"},
			Type:     corev1.ServiceTypeLoadBalancer,
			Ports: []corev1.ServicePort{
				{Protocol: corev1.ProtocolTCP, Port: 8080, TargetPort: intStrPtr(8080)},
			},
		},
	}
	_, err = c.CoreV1().Services("default").Create(context.TODO(), s, metav1.CreateOptions{})
	if err != nil {
		panic(fmt.Errorf("failed to create service: %v", err))
	}
	fmt.Println("Service created successfully!")
}

func int32Ptr(i int32) *int32 {
	return &i
}

func intStrPtr(i int) intstr.IntOrString {
	return intstr.IntOrString{IntVal: int32(i)}
}
