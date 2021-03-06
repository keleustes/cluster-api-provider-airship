/*
Copyright 2019 The Kubernetes authors.

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

package controller

import (
	airshipmachine "github.com/keleustes/cluster-api-provider-airship/pkg/cloud/airship/actuators/machine"
	clientset "sigs.k8s.io/cluster-api/pkg/client/clientset_generated/clientset"
	capimachine "sigs.k8s.io/cluster-api/pkg/controller/machine"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

//+kubebuilder:rbac:groups=airship.airshipit.org,resources=airshipmachineproviderspecs;airshipmachineproviderstatuses,verbs=get;list;watch;create;update;patch;delete
func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, func(m manager.Manager) error {
		cfg := m.GetConfig()
		cs, err := clientset.NewForConfig(cfg)
		if err != nil {
			panic(err)
		}
		params := airshipmachine.ActuatorParams{Client: cs.ClusterV1alpha1()}
		machineActuator := airshipmachine.NewActuator(params)
		return capimachine.AddWithActuator(m, machineActuator)
	})
}
