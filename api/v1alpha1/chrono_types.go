/*
Copyright 2023.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ChronoSpec defines the desired state of Chrono
type ChronoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Chrono. Edit chrono_types.go to remove/update
	Foo      string      `json:"foo,omitempty"`
	Schedule string      `json:"schedule,omitempty"`
	PodTemplate PodTemplate `json:"podTemplate,omitempty"`
}

// PodTemplate defines the pod template specification for Chrono
type PodTemplate struct {
	Spec PodSpec `json:"spec,omitempty"`
}

// PodSpec defines the pod specification for Chrono
type PodSpec struct {
	Template PodTemplateSpec `json:"template,omitempty"`
	RestartPolicy string `json:"restartPolicy,omitempty"`
}

// PodTemplateSpec defines the template for the pod
type PodTemplateSpec struct {
	Spec PodSpecSpec `json:"spec,omitempty"`
}

// PodSpecSpec defines the specification for the pod template
type PodSpecSpec struct {
	Containers []Container `json:"containers,omitempty"`
}

// Container defines the container specification
type Container struct {
	Name            string   `json:"name,omitempty"`
	Image           string   `json:"image,omitempty"`
	ImagePullPolicy string   `json:"imagePullPolicy,omitempty"`
	Command         []string `json:"command,omitempty"`
}

// ChronoStatus defines the observed state of Chrono
type ChronoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Chrono is the Schema for the chronoes API
type Chrono struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChronoSpec   `json:"spec,omitempty"`
	Status ChronoStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ChronoList contains a list of Chrono
type ChronoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Chrono `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Chrono{}, &ChronoList{})
}

