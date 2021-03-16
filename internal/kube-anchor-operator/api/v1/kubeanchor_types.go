/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KubeAnchorSpec defines the desired state of KubeAnchor
type KubeAnchorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// +kubebuilder:validation:UniqueItems=true
	// +kubebuilder:printcolumn:name="Resource types",type=String,description="K8s built-in resource types or CRDs (Deployment,Statefulset, Service etc)"
	// +optional
	ResourceTypes []string `json:"resourceTypes,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	// +kubebuilder:printcolumn:name="Namespaces",type=String,description="Namespace(s) where KubeAnchor will be applied"
	// +optional
	Namespaces []string `json:"namespaces,omitempty"`
	// +optional
	// +kubebuilder:printcolumn:name="Enabled",type=String,description="Whether KubeAnchor is enabled or not, default disabled"
	Enabled bool `json:"enabled,omitempty"`
	// +kubebuilder:printcolumn:name="Effective From",type=String,description="Effective from date/time (eg:2021-05-01 15:04:05)"
	// +optional
	EffectiveFrom string `json:"effectiveFrom,omitempty"`
	// +optional
	// +kubebuilder:printcolumn:name="Effective To",type=String,description="Effective to date/time (eg:2021-05-01 18:04:05)"
	EffectiveTo string `json:"effectiveTo,omitempty"`
}

// KubeAnchorStatus defines the observed state of KubeAnchor
type KubeAnchorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=ka
// +kubebuilder:subresource:status
// KubeAnchor is the Schema for the kubeanchors API
type KubeAnchor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeAnchorSpec   `json:"spec,omitempty"`
	Status KubeAnchorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubeAnchorList contains a list of KubeAnchor
type KubeAnchorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeAnchor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubeAnchor{}, &KubeAnchorList{})
}
