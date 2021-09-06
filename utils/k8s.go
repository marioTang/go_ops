package utils

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func K8s() {
//
//	var (
//		k8sconfig = flag.String("k8sconfig", "./config", "kubernetes auth config") //使用kubeconfig配置文件进行集群权限认证
//		config    *rest.Config
//		err       error
//	)
//	flag.Parse()
//
//	config, err = clientcmd.BuildConfigFromFlags("", *k8sconfig)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(config)
//	// 从指定的config创建一个新的clientset
//	clientset, err := kubernetes.NewForConfig(config)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	} else {
//		fmt.Println("connect kubernetes cluster success.")
//	}
//	fmt.Println(clientset.StorageV1())
//	//pods, err := clientset.CoreV1().Pods("applets-prod").List(context.TODO(),metav1.ListOptions{})
//	pods, err := clientset.CoreV1().Services("applets-prod").List(context.TODO(),metav1.ListOptions{})
//
//	//fmt.Println(pods)
//
//	//pods, err := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
//
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//
////fmt.Println(pods)
//
//
//	for _,pod := range pods.Items {
//
//
//		//	fmt.Println(pod.ObjectMeta.Name,pod.Status.Phase,pod.Name)
//		fmt.Println(pod.Status)
//
//		//fmt.Printf("\nThere are %d pods in namespaces %s\n", len(pods.Items), namespace)
//		fmt.Printf("Name: %s, Status: %s, CreateTime: %s\n", pod.ObjectMeta.Name, pod.Status.Phase, pod.ObjectMeta.CreationTimestamp)
//
//
//
//	}

}


func Newk8s() {

	// 配置 k8s 集群外 kubeconfig 配置文件，默认位置 $HOME/.kube/config

	var (
		kubeconfig = flag.String("k8sconfig", "./config", "kubernetes auth config") //使用kubeconfig配置文件进行集群权限认证
		config    *rest.Config
		err       error
	)
	flag.Parse()

	config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)
	// 从指定的config创建一个新的clientset
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}
	for {
		// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PodsGetter 接口方法 Pods(namespace string) 返回 PodInterface
		// PodInterface 接口拥有操作 Pod 资源的方法，例如 Create、Update、Get、List 等方法
		// 注意：Pods() 方法中 namespace 不指定则获取 Cluster 所有 Pod 列表
		pods, err := clientset.CoreV1().Pods("applets-prod").List(context.TODO(),metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the k8s cluster\n", len(pods.Items))

		// 获取指定 namespace 中的 Pod 列表信息
		namespace := "kubeless"
		pods, err = clientset.CoreV1().Pods("applets-prod").List(context.TODO(),metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("\nThere are %d pods in namespaces %s\n", len(pods.Items), namespace)
		for _, pod := range pods.Items {
			fmt.Printf("Name: %s, Status: %s, CreateTime: %s\n", pod.ObjectMeta.Name, pod.Status.Phase, pod.ObjectMeta.CreationTimestamp)
		}

		


		// 获取指定 namespaces 和 podName 的详细信息，使用 error handle 方式处理错误信息
		//namespace = "kubeless"
		//podName := "get-java-5ff45cd65d-2frkx"
		//pod, err := clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
		//if errors.IsNotFound(err) {
		//	fmt.Printf("Pod %s in namespace %s not found\n", podName, namespace)
		//} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		//	fmt.Printf("Error getting pod %s in namespace %s: %v\n",
		//		podName, namespace, statusError.ErrStatus.Message)
		//} else if err != nil {
		//	panic(err.Error())
		//} else {
		//	fmt.Printf("\nFound pod %s in namespace %s\n", podName, namespace)
		//	maps := map[string]interface{}{
		//		"Name":        pod.ObjectMeta.Name,
		//		"Namespaces":  pod.ObjectMeta.Namespace,
		//		"NodeName":    pod.Spec.NodeName,
		//		"Annotations": pod.ObjectMeta.Annotations,
		//		"Labels":      pod.ObjectMeta.Labels,
		//		"SelfLink":    pod.ObjectMeta.SelfLink,
		//		"Uid":         pod.ObjectMeta.UID,
		//		"Status":      pod.Status.Phase,
		//		"IP":          pod.Status.PodIP,
		//		"Image":       pod.Spec.Containers[0].Image,
		//	}
		//	prettyPrint(maps)
		//}

		//time.Sleep(10 * time.Second)
	}

}



