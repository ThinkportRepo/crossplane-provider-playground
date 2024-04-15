/*
Copyright 2022 The Crossplane Authors.

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

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// IPv4NetworkParameters are the configurable fields of a IPv4Network.
type IPv4NetworkParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// IPv4NetworkObservation are the observable fields of a IPv4Network.
type IPv4NetworkObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A IPv4NetworkSpec defines the desired state of a IPv4Network.
type IPv4NetworkSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IPv4NetworkParameters `json:"forProvider"`
}

// A IPv4NetworkStatus represents the observed state of a IPv4Network.
type IPv4NetworkStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IPv4NetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A IPv4Network is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,xacinfoblox}
type IPv4Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IPv4NetworkSpec   `json:"spec"`
	Status IPv4NetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPv4NetworkList contains a list of IPv4Network
type IPv4NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPv4Network `json:"items"`
}

// IPv4Network type metadata.
var (
	IPv4NetworkKind             = reflect.TypeOf(IPv4Network{}).Name()
	IPv4NetworkGroupKind        = schema.GroupKind{Group: Group, Kind: IPv4NetworkKind}.String()
	IPv4NetworkKindAPIVersion   = IPv4NetworkKind + "." + SchemeGroupVersion.String()
	IPv4NetworkGroupVersionKind = SchemeGroupVersion.WithKind(IPv4NetworkKind)
)

func init() {
	SchemeBuilder.Register(&IPv4Network{}, &IPv4NetworkList{})
}
