package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindChart = "Chart"
)

// NewChartCRD returns a new custom resource definition for Chart.
// This might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: charts.application.giantswarm.io
//     spec:
//       group: application.giantswarm.io
//       scope: Namespaced
//       version: v1alpha1
//       names:
//         kind: Chart
//         plural: charts
//         singular: chart
//
func NewChartCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "charts.application.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "application.giantswarm.io",
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "Chart",
				Plural:   "charts",
				Singular: "chart",
			},
			Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
				Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
			},
		},
	}
}

func NewChartTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: version,
		Kind:       kindChart,
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Chart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ChartSpec   `json:"spec"`
	Status            ChartStatus `json:"status" yaml:"status"`
}

type ChartSpec struct {
	// Config is the config to be applied when the chart is deployed.
	Config ChartSpecConfig `json:"config" yaml:"config"`
	// Name is the name of the Helm chart to be deployed.
	// e.g. kubernetes-prometheus
	Name string `json:"name" yaml:"name"`
	// KubeConfig is the kubeconfig to connect to the cluster when deploying
	// the app.
	KubeConfig ChartSpecKubeConfig `json:"kubeConfig" yaml:"kubeConfig"`
	// Namespace is the namespace where the chart should be deployed.
	// e.g. monitoring
	Namespace string `json:"namespace" yaml:"namespace"`
	// TarballURL is the URL for the Helm chart tarball to be deployed.
	// e.g. https://path/to/prom-1-0-0.tgz"
	TarballURL string `json:"tarballURL" yaml:"tarballURL"`
}

type ChartSpecConfig struct {
	// ConfigMap references a config map containing values that should be
	// applied to the chart.
	ConfigMap ChartSpecConfigConfigMap `json:"configMap" yaml:"configMap"`
	// Secret references a secret containing secret values that should be
	// applied to the chart.
	Secret ChartSpecConfigSecret `json:"secret" yaml:"secret"`
}

type ChartSpecConfigConfigMap struct {
	// Name is the name of the config map containing chart values to apply,
	// e.g. prometheus-chart-values.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the values config map,
	// e.g. monitoring.
	Namespace string `json:"namespace" yaml:"namespace"`
	// ResourceVersion is the Kubernetes resource version of the configmap.
	// Used to detect if the configmap has changed, e.g. 12345.
	ResourceVersion string `json:"resourceVersion" yaml:"resourceVersion"`
}

type ChartSpecConfigSecret struct {
	// Name is the name of the secret containing chart values to apply,
	// e.g. prometheus-chart-secret.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the secret,
	// e.g. kube-system.
	Namespace string `json:"namespace" yaml:"namespace"`
	// ResourceVersion is the Kubernetes resource version of the secret.
	// Used to detect if the secret has changed, e.g. 12345.
	ResourceVersion string `json:"resourceVersion" yaml:"resourceVersion"`
}

type ChartSpecKubeConfig struct {
	// Secret references a secret containing the kubconfig.
	Secret ChartSpecKubeConfigSecret `json:"secret" yaml:"secret"`
}

type ChartSpecKubeConfigSecret struct {
	// Name is the name of the secret containing the kubeconfig,
	// e.g. chart-operator-kubeconfig.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the secret containing the kubeconfig,
	// e.g. giantswarm.
	Namespace string `json:"namespace" yaml:"namespace"`
}

type ChartStatus struct {
	// Status is the status of the deployed chart,
	// e.g. DEPLOYED.
	Status string `json:"status" yaml:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ChartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Chart `json:"items"`
}
