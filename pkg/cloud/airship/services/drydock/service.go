/*
Copyright 2018 The Kubernetes Authors.

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

package drydock

import (
	"os"

	"github.com/go-openapi/strfmt"
	"github.com/keleustes/cluster-api-provider-airship/pkg/cloud/airship/actuators"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/kubekit99/airship-go-api/drydock/client"
)

// Service holds a collection of interfaces.
// The interfaces are broken down like this to group functions together.
// One alternative is to have a large list of functions from the ec2 client.
type Service struct {
	scope         *actuators.Scope
	airshipclient *apiclient.Drydock
}

// NewService returns a new service given the api clients.
func NewService(scope *actuators.Scope) *Service {
	// create the transport
	transport := httptransport.New(os.Getenv("TODOLIST_HOST"), "", nil)

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	return &Service{
		scope:         scope,
		airshipclient: client,
	}
}
