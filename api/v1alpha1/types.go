/*
SPDX-FileCopyrightText: 2026 SAP SE or an SAP affiliate company and nginx-ingress-controller-cop contributors
SPDX-License-Identifier: Apache-2.0
*/

package v1alpha1

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	"github.com/sap/component-operator-runtime/pkg/component"
	componentoperatorruntimetypes "github.com/sap/component-operator-runtime/pkg/types"
)

// +kubebuilder:pruning:PreserveUnknownFields

// NginxIngressControllerSpec defines the desired state of NginxIngressController.
type NginxIngressControllerSpec struct {
	apiextensionsv1.JSON `json:"-"`
}

// NginxIngressControllerStatus defines the observed state of NginxIngressController.
type NginxIngressControllerStatus struct {
	component.Status `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +genclient

// NginxIngressController is the Schema for the nginxingresscontrollers API.
type NginxIngressController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NginxIngressControllerSpec `json:"spec,omitempty"`
	// +kubebuilder:default={"observedGeneration":-1}
	Status NginxIngressControllerStatus `json:"status,omitempty"`
}

var _ component.Component = &NginxIngressController{}

// +kubebuilder:object:root=true

// NginxIngressControllerList contains a list of NginxIngressController.
type NginxIngressControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NginxIngressController `json:"items"`
}

func (s *NginxIngressControllerSpec) ToUnstructured() map[string]any {
	result, err := runtime.DefaultUnstructuredConverter.ToUnstructured(s)
	if err != nil {
		panic(err)
	}
	if namespace, ok := result["namespace"]; ok {
		if _, ok := namespace.(string); !ok {
			panic("spec.namespace set but is not a string")
		}
	}
	if name, ok := result["name"]; ok {
		if _, ok := name.(string); !ok {
			panic("spec.name set but is not a string")
		}
	}
	return result
}

func (c *NginxIngressController) GetDeploymentNamespace() string {
	if namespace, ok := c.Spec.ToUnstructured()["namespace"]; ok && namespace.(string) != "" {
		return namespace.(string)
	}
	return c.Namespace
}

func (c *NginxIngressController) GetDeploymentName() string {
	if name, ok := c.Spec.ToUnstructured()["name"]; ok && name.(string) != "" {
		return name.(string)
	}
	return c.Name
}

func (c *NginxIngressController) GetSpec() componentoperatorruntimetypes.Unstructurable {
	return &c.Spec
}

func (c *NginxIngressController) GetStatus() *component.Status {
	return &c.Status.Status
}

func init() {
	SchemeBuilder.Register(&NginxIngressController{}, &NginxIngressControllerList{})
}
