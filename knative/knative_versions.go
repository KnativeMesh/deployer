package knative

import "fmt"

type Version interface {
	GetServingCRDs() string
	GetServingCore() string
	GetEventingCRDs() string
	GetEventingCore() string
}

func KnativeVersionFabric(version string) Version {
	switch version {
	case "1.8", "8":
		fmt.Println("Chosen version is 1.8")
		return VersionKnative18{}
	case "1.7", "7":
		fmt.Println("Chosen version is 1.7")
		return VersionKnative17{}
	case "1.9", "9":
		fmt.Println("Chosen version is 1.9")
		return VersionKnative19{}
	default:
		fmt.Println("Chosen version is 1.8")
		return VersionKnative18{}
	}
	return nil
}

type VersionKnative18 struct {
}

func (v VersionKnative18) GetServingCRDs() string {
	return "https://github.com/knative/serving/releases/download/knative-v1.8.6/serving-crds.yaml"
}

func (v VersionKnative18) GetServingCore() string {
	return "https://github.com/knative/serving/releases/download/knative-v1.8.6/serving-core.yaml"
}

func (v VersionKnative18) GetEventingCRDs() string {
	return "https://github.com/knative/eventing/releases/download/knative-v1.8.8/eventing-crds.yaml"
}

func (v VersionKnative18) GetEventingCore() string {
	return "https://github.com/knative/eventing/releases/download/knative-v1.8.8/eventing-core.yaml"
}

type VersionKnative17 struct {
}

func (v VersionKnative17) GetServingCRDs() string {
	return "https://github.com/knative/serving/releases/download/knative-v1.7.4/serving-crds.yaml"
}

func (v VersionKnative17) GetServingCore() string {
	return "https://github.com/knative/serving/releases/download/knative-v1.7.4/serving-core.yaml"
}

func (v VersionKnative17) GetEventingCRDs() string {
	return "https://github.com/knative/eventing/releases/download/knative-v1.7.8/eventing-crds.yaml"
}

func (v VersionKnative17) GetEventingCore() string {
	return "https://github.com/knative/eventing/releases/download/knative-v1.7.8/eventing-core.yaml"
}

type VersionKnative19 struct {
}

func (v VersionKnative19) GetServingCRDs() string {
	return "https://github.com/knative/serving/releases/download/knative-v1.9.3/serving-crds.yaml"
}

func (v VersionKnative19) GetServingCore() string {
	return "https://github.com/knative/serving/releases/download/knative-v1.9.3/serving-core.yaml"
}

func (v VersionKnative19) GetEventingCRDs() string {
	return "https://github.com/knative/eventing/releases/download/knative-v1.9.7/eventing-crds.yaml"
}

func (v VersionKnative19) GetEventingCore() string {
	return "https://github.com/knative/eventing/releases/download/knative-v1.9.7/eventing-core.yaml"
}
