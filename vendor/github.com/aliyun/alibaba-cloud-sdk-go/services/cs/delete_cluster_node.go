package cs

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

func (client *Client) DeleteClusterNode(request *DeleteClusterNodeRequest) (response *DeleteClusterNodeResponse, err error) {
	response = CreateDeleteClusterNodeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteClusterNodeWithChan(request *DeleteClusterNodeRequest) (<-chan *DeleteClusterNodeResponse, <-chan error) {
	responseChan := make(chan *DeleteClusterNodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteClusterNode(request)
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

func (client *Client) DeleteClusterNodeWithCallback(request *DeleteClusterNodeRequest, callback func(response *DeleteClusterNodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteClusterNodeResponse
		var err error
		defer close(result)
		response, err = client.DeleteClusterNode(request)
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

type DeleteClusterNodeRequest struct {
	*requests.RoaRequest
	ReleaseInstance string `position:"Query" name:"releaseInstance"`
	ClusterId       string `position:"Path" name:"ClusterId"`
	Force           string `position:"Query" name:"force"`
	Ip              string `position:"Path" name:"Ip"`
}

type DeleteClusterNodeResponse struct {
	*responses.BaseResponse
}

func CreateDeleteClusterNodeRequest() (request *DeleteClusterNodeRequest) {
	request = &DeleteClusterNodeRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DeleteClusterNode", "/clusters/[ClusterId]/ip/[Ip]", "", "")
	request.Method = requests.DELETE
	return
}

func CreateDeleteClusterNodeResponse() (response *DeleteClusterNodeResponse) {
	response = &DeleteClusterNodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
