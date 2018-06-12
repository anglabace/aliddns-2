package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) CreateDampPolicy(request *CreateDampPolicyRequest) (response *CreateDampPolicyResponse, err error) {
	response = CreateCreateDampPolicyResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateDampPolicyWithChan(request *CreateDampPolicyRequest) (<-chan *CreateDampPolicyResponse, <-chan error) {
	responseChan := make(chan *CreateDampPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDampPolicy(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) CreateDampPolicyWithCallback(request *CreateDampPolicyRequest, callback func(response *CreateDampPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDampPolicyResponse
		var err error
		defer close(result)
		response, err = client.CreateDampPolicy(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type CreateDampPolicyRequest struct {
	*requests.RpcRequest
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	TimeRules            string           `position:"Query" name:"TimeRules"`
	PolicyName           string           `position:"Query" name:"PolicyName"`
	ActionRules          string           `position:"Query" name:"ActionRules"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Handlers             string           `position:"Query" name:"Handlers"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceRules          string           `position:"Query" name:"SourceRules"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
}

type CreateDampPolicyResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	PolicyId   string `json:"PolicyId" xml:"PolicyId"`
	PolicyName string `json:"PolicyName" xml:"PolicyName"`
}

func CreateCreateDampPolicyRequest() (request *CreateDampPolicyRequest) {
	request = &CreateDampPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateDampPolicy", "rds", "openAPI")
	return
}

func CreateCreateDampPolicyResponse() (response *CreateDampPolicyResponse) {
	response = &CreateDampPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
