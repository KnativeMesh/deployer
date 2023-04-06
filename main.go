package main

import (
	"fmt"
	"github.com/KnativeMesh/deployer/knative"
	"github.com/KnativeMesh/deployer/utils"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func main() {
	fmt.Println("*** Welcome to Knative-Mesh-Deployer ***")
	fmt.Printf("This is an application for convenient migration from microservices to serverless\n\n\n")
	source := utils.NewStdinReader()

	config, err := buildKubernetesConfig(source)
	if err != nil {
		fmt.Printf("Can't build appropriate kubernetes config. Please check yout kubeconfig file or provide another one \n")
		fmt.Printf("Err %+v \n", err.Error())
		return
	}
	cmd := exec.Command("kubectl", "config", "use-context", "rancher-desktop")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	knative.InstallKnative(config, source)
}

func buildKubernetesConfig(source utils.Reader) (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := ""
		fmt.Println("Can't find authentiphicated kubernetes context, trying find kubeconfig in current folder")
		kubeconfig, fileErr := findKubeconfigInFolder()
		if fileErr != nil {
			fmt.Println("Can't find kubeconfig in curret folder, please provide path to kubeconfig")
			kubeconfig = source.ReadString()
		}
		fmt.Printf("Using configuration from '%s' \n", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}

func findKubeconfigInFolder() (string, error) {
	libRegEx, _ := regexp.Compile("(\\w*config)")
	configName := ""
	fileError := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			abs, _ := filepath.Abs(info.Name())
			configName = abs
			return nil
		}
		return nil
	})
	if fileError != nil {
		return "", fileError
	}
	return configName, nil
}
