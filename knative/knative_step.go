package knative

import (
	"fmt"
	"github.com/KnativeMesh/deployer/utils"
	"os/exec"
)

func InstallKnative(source utils.Reader) error {
	fmt.Println("Step 1. Installing Knative")
	fmt.Println("Do you need to install knative? Y/N")
	answ := source.ReadRune()
	switch answ {
	case 'Y', 'y':
		fmt.Println("User has chosen to install Knative")
		err := installKnative(source)
		if err != nil {
			fmt.Printf("Knative installation was unsuccessfull because of err %+v \n", err.Error())
			return err
		}
		fmt.Println("Knative was successfully installed")
	case 'N', 'n':
		fmt.Println("Knative won't be installed")
	default:
		fmt.Println("Knative won't be installed")
	}
	return nil
}

func installKnative(source utils.Reader) error {
	fmt.Println("User has chosen to install Knative")
	fmt.Println("Please choose version to install. Available versions are [1.7, 1.8, 1.9]. By default version 1.8 will be installed")
	version := source.ReadString()
	instance := KnativeVersionFabric(version)

	err := CreateKnativeInCluster(instance)
	if err != nil {
		return err
	}
	return nil
}

func CreateKnativeInCluster(version Version) error {
	fmt.Println("Installing knative-serving crds...")
	cmd := exec.Command("kubectl", "apply", "-f", version.GetServingCRDs())
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println("Installing knative-serving core...")
	cmd = exec.Command("kubectl", "apply", "-f", version.GetServingCore())
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println("Installing knative-eventing crds...")
	cmd = exec.Command("kubectl", "apply", "-f", version.GetEventingCRDs())
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println("Installing knative-eventin core...")
	cmd = exec.Command("kubectl", "apply", "-f", version.GetEventingCore())
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
