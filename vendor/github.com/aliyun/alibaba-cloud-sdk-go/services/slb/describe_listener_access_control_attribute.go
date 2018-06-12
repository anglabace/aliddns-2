package slb

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

func (client *Client) DescribeListenerAccessControlAttribute(request *DescribeListenerAccessControlAttributeRequest) (response *DescribeListenerAccessControlAttributeResponse, err error) {
	response = CreateDescribeListenerAccessControlAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeListenerAccessControlAttributeWithChan(request *DescribeListenerAccessControlAttributeRequest) (<-chan *DescribeListenerAccessControlAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeListenerAccessControlAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeListenerAccessControlAttribute(request)
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

func (client *Client) DescribeListenerAccessControlAttributeWithCallback(request *DescribeListenerAccessControlAttributeRequest, callback func(response *DescribeListenerAccessControlAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeListenerAccessControlAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeListenerAccessControlAttribute(request)
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

type DescribeListenerAccessControlAttributeRequest struct {
	*requests.RpcRequest
	Tags                 string           `position:"Query" name:"Tags"`
	ListenerPort         requests.Integer `position:"Query" name:"ListenerPort"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeListenerAccessControlAttributeResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	AccessControlStatus string `json:"AccessControlStatus" xml:"AccessControlStatus"`
	SourceItems         string `json:"SourceItems" xml:"SourceItems"`
}

func CreateDescribeListenerAccessControlAttributeRequest() (request *DescribeListenerAccessControlAttributeRequest) {
	request = &DescribeListenerAccessControlAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeListenerAccessControlAttribute", "slb", "openAPI")
	return
}

func CreateDescribeListenerAccessControlAttributeResponse() (response *DescribeListenerAccessControlAttributeResponse) {
	response = &DescribeListenerAccessControlAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
