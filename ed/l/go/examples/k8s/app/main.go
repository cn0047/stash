package main

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	f := ""
	n := ""

	c := getOutOfClusterClientSet(f)
	//c := getInClusterClientSet()

	//podsList(c, n)
	//podsGet(c, n)
	podsCreate(c, n)
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

func getOutOfClusterClientSet(configFile string) *kubernetes.Clientset {
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

	fmt.Printf("[podsList] found %+v pods. \n", len(pods.Items))
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
