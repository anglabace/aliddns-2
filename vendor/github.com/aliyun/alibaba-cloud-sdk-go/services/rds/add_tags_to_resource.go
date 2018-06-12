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

func (client *Client) AddTagsToResource(request *AddTagsToResourceRequest) (response *AddTagsToResourceResponse, err error) {
	response = CreateAddTagsToResourceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddTagsToResourceWithChan(request *AddTagsToResourceRequest) (<-chan *AddTagsToResourceResponse, <-chan error) {
	responseChan := make(chan *AddTagsToResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddTagsToResource(request)
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

func (client *Client) AddTagsToResourceWithCallback(request *AddTagsToResourceRequest, callback func(response *AddTagsToResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddTagsToResourceResponse
		var err error
		defer close(result)
		response, err = client.AddTagsToResource(request)
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

type AddTagsToResourceRequest struct {
	*requests.RpcRequest
	Tags                 string           `position:"Query" name:"Tags"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ProxyId              string           `position:"Query" name:"proxyId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tag5Key              string           `position:"Query" name:"Tag.5.key"`
	Tag5Value            string           `position:"Query" name:"Tag.5.value"`
	Tag3Key              string           `position:"Query" name:"Tag.3.key"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Tag1Key              string           `position:"Query" name:"Tag.1.key"`
	Tag2Key              string           `position:"Query" name:"Tag.2.key"`
	Tag1Value            string           `position:"Query" name:"Tag.1.value"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Tag4Value            string           `position:"Query" name:"Tag.4.value"`
	Tag3Value            string           `position:"Query" name:"Tag.3.value"`
	Tag2Value            string           `position:"Query" name:"Tag.2.value"`
	Tag4Key              string           `position:"Query" name:"Tag.4.key"`
}

type AddTagsToResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateAddTagsToResourceRequest() (request *AddTagsToResourceRequest) {
	request = &AddTagsToResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "AddTagsToResource", "rds", "openAPI")
	return
}

func CreateAddTagsToResourceResponse() (response *AddTagsToResourceResponse) {
	response = &AddTagsToResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
