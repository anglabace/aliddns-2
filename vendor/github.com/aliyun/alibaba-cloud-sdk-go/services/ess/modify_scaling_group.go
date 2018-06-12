package ess

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

func (client *Client) ModifyScalingGroup(request *ModifyScalingGroupRequest) (response *ModifyScalingGroupResponse, err error) {
	response = CreateModifyScalingGroupResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyScalingGroupWithChan(request *ModifyScalingGroupRequest) (<-chan *ModifyScalingGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyScalingGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyScalingGroup(request)
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

func (client *Client) ModifyScalingGroupWithCallback(request *ModifyScalingGroupRequest, callback func(response *ModifyScalingGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyScalingGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyScalingGroup(request)
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

type ModifyScalingGroupRequest struct {
	*requests.RpcRequest
	ActiveScalingConfigurationId string           `position:"Query" name:"ActiveScalingConfigurationId"`
	MinSize                      requests.Integer `position:"Query" name:"MinSize"`
	ScalingGroupName             string           `position:"Query" name:"ScalingGroupName"`
	OwnerId                      requests.Integer `position:"Query" name:"OwnerId"`
	ScalingGroupId               string           `position:"Query" name:"ScalingGroupId"`
	MaxSize                      requests.Integer `position:"Query" name:"MaxSize"`
	ResourceOwnerAccount         string           `position:"Query" name:"ResourceOwnerAccount"`
	RemovalPolicy1               string           `position:"Query" name:"RemovalPolicy.1"`
	RemovalPolicy2               string           `position:"Query" name:"RemovalPolicy.2"`
	ResourceOwnerId              requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DefaultCooldown              requests.Integer `position:"Query" name:"DefaultCooldown"`
	OwnerAccount                 string           `position:"Query" name:"OwnerAccount"`
}

type ModifyScalingGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyScalingGroupRequest() (request *ModifyScalingGroupRequest) {
	request = &ModifyScalingGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "ModifyScalingGroup", "ess", "openAPI")
	return
}

func CreateModifyScalingGroupResponse() (response *ModifyScalingGroupResponse) {
	response = &ModifyScalingGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
