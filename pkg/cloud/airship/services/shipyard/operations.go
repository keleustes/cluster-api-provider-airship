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

package shipyard

import (
// "github.com/kubekit99/airship-go-api/shipyard/services/shipyard"
)

// From Shipyard API Definition
// func (a *Client) CommitConfigDocs(params *CommitConfigDocsParams) (*CommitConfigDocsOK, error)
// func (a *Client) GetConfig(params *GetConfigParams) (*GetConfigOK, error)
// func (a *Client) GetHealth(params *GetHealthParams) (*GetHealthOK, error)
// func (a *Client) GetNoteDetails(params *GetNoteDetailsParams) (*GetNoteDetailsOK, error)
// func (a *Client) GetSiteStatuses(params *GetSiteStatusesParams) (*GetSiteStatusesOK, error)
// func (a *Client) GetVersions(params *GetVersionsParams) (*GetVersionsOK, error)
// func (a *Client) GetWFActionByID(params *GetWFActionByIDParams) (*GetWFActionByIDOK, error)
// func (a *Client) GetWFActionStepByID(params *GetWFActionStepByIDParams) (*GetWFActionStepByIDOK, error)
// func (a *Client) GetWFActionStepLogsByID(params *GetWFActionStepLogsByIDParams) (*GetWFActionStepLogsByIDOK, error)
// func (a *Client) GetWFActionValidationByID(params *GetWFActionValidationByIDParams) (*GetWFActionValidationByIDOK, error)
// func (a *Client) GetWorkflowByID(params *GetWorkflowByIDParams) (*GetWorkflowByIDOK, error)
// func (a *Client) GetWorkflows(params *GetWorkflowsParams) (*GetWorkflowsOK, error)
// func (a *Client) InvokeHelmTests(params *InvokeHelmTestsParams) (*InvokeHelmTestsOK, error)
// func (a *Client) RetrieveConfigDocsClearTextByCollectionID(params *RetrieveConfigDocsClearTextByCollectionIDParams) (*RetrieveConfigDocsClearTextByCollectionIDOK, error)
// func (a *Client) RetrieveConfigDocsStatus(params *RetrieveConfigDocsStatusParams) (*RetrieveConfigDocsStatusOK, error)
// func (a *Client) RetrieveRenderedCleartextConfigDocs(params *RetrieveRenderedCleartextConfigDocsParams) (*RetrieveRenderedCleartextConfigDocsOK, error)
// func (a *Client) SendControlToWFAction(params *SendControlToWFActionParams) (*SendControlToWFActionOK, error)

// WaitForSomething
func (s *Service) WaitForSomething(thething string) (string, error) {
	return thething, nil
}
