/*
Copyright 2021.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
type ClientConfig struct {
	BearToken string `json:"bearToken,omitempty"`
	// CertData holds PEM-encoded bytes.
	CertData []byte `json:"certData,omitempty"`
	// KeyData holds PEM-encoded bytes.
	KeyData []byte `json:"keyData,omitempty"`
	// CAData holds PEM-encoded bytes.
	CAData     []byte `json:"cAData,omitempty"`
	SecretName string `json:"secretName,omitempty"`
}

// ManagedClusterSpec defines the desired state of ManagedCluster
type ManagedClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ClientConfig ClientConfig `json:"clientConfig"`
}

// ManagedClusterStatus defines the observed state of ManagedCluster
type ManagedClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ManagedCluster is the Schema for the managedclusters API
type ManagedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedClusterSpec   `json:"spec,omitempty"`
	Status ManagedClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ManagedClusterList contains a list of ManagedCluster
type ManagedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedCluster `json:"items"`
}

// func init() {
// 	SchemeBuilder.Register(&ManagedCluster{}, &ManagedClusterList{})
// }
