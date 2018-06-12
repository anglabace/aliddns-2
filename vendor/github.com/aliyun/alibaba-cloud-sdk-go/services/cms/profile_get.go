package cms

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

func (client *Client) ProfileGet(request *ProfileGetRequest) (response *ProfileGetResponse, err error) {
	response = CreateProfileGetResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ProfileGetWithChan(request *ProfileGetRequest) (<-chan *ProfileGetResponse, <-chan error) {
	responseChan := make(chan *ProfileGetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ProfileGet(request)
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

func (client *Client) ProfileGetWithCallback(request *ProfileGetRequest, callback func(response *ProfileGetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ProfileGetResponse
		var err error
		defer close(result)
		response, err = client.ProfileGet(request)
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

type ProfileGetRequest struct {
	*requests.RpcRequest
	UserId requests.Integer `position:"Query" name:"UserId"`
}

type ProfileGetResponse struct {
	*responses.BaseResponse
	ErrorCode                int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage             string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success                  bool   `json:"Success" xml:"Success"`
	RequestId                string `json:"RequestId" xml:"RequestId"`
	UserId                   int    `json:"UserId" xml:"UserId"`
	AutoInstall              bool   `json:"AutoInstall" xml:"AutoInstall"`
	EnableInstallAgentNewECS bool   `json:"EnableInstallAgentNewECS" xml:"EnableInstallAgentNewECS"`
}

func CreateProfileGetRequest() (request *ProfileGetRequest) {
	request = &ProfileGetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "ProfileGet", "cms", "openAPI")
	return
}

func CreateProfileGetResponse() (response *ProfileGetResponse) {
	response = &ProfileGetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
