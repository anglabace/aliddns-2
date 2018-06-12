package cdn

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

func (client *Client) ClearUserDomainBlackList(request *ClearUserDomainBlackListRequest) (response *ClearUserDomainBlackListResponse, err error) {
	response = CreateClearUserDomainBlackListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ClearUserDomainBlackListWithChan(request *ClearUserDomainBlackListRequest) (<-chan *ClearUserDomainBlackListResponse, <-chan error) {
	responseChan := make(chan *ClearUserDomainBlackListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ClearUserDomainBlackList(request)
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

func (client *Client) ClearUserDomainBlackListWithCallback(request *ClearUserDomainBlackListRequest, callback func(response *ClearUserDomainBlackListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ClearUserDomainBlackListResponse
		var err error
		defer close(result)
		response, err = client.ClearUserDomainBlackList(request)
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

type ClearUserDomainBlackListRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerAccount  string           `position:"Query" name:"OwnerAccount"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type ClearUserDomainBlackListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateClearUserDomainBlackListRequest() (request *ClearUserDomainBlackListRequest) {
	request = &ClearUserDomainBlackListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "ClearUserDomainBlackList", "", "")
	return
}

func CreateClearUserDomainBlackListResponse() (response *ClearUserDomainBlackListResponse) {
	response = &ClearUserDomainBlackListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
