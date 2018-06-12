package ecs

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

func (client *Client) DetachKeyPair(request *DetachKeyPairRequest) (response *DetachKeyPairResponse, err error) {
	response = CreateDetachKeyPairResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DetachKeyPairWithChan(request *DetachKeyPairRequest) (<-chan *DetachKeyPairResponse, <-chan error) {
	responseChan := make(chan *DetachKeyPairResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachKeyPair(request)
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

func (client *Client) DetachKeyPairWithCallback(request *DetachKeyPairRequest, callback func(response *DetachKeyPairResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachKeyPairResponse
		var err error
		defer close(result)
		response, err = client.DetachKeyPair(request)
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

type DetachKeyPairRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string           `position:"Query" name:"InstanceIds"`
	KeyPairName          string           `position:"Query" name:"KeyPairName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DetachKeyPairResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	TotalCount  string `json:"TotalCount" xml:"TotalCount"`
	FailCount   string `json:"FailCount" xml:"FailCount"`
	KeyPairName string `json:"KeyPairName" xml:"KeyPairName"`
	Results     struct {
		Result []struct {
			InstanceId string `json:"InstanceId" xml:"InstanceId"`
			Success    string `json:"Success" xml:"Success"`
			Code       string `json:"Code" xml:"Code"`
			Message    string `json:"Message" xml:"Message"`
		} `json:"Result" xml:"Result"`
	} `json:"Results" xml:"Results"`
}

func CreateDetachKeyPairRequest() (request *DetachKeyPairRequest) {
	request = &DetachKeyPairRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DetachKeyPair", "ecs", "openAPI")
	return
}

func CreateDetachKeyPairResponse() (response *DetachKeyPairResponse) {
	response = &DetachKeyPairResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
