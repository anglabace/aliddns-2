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

func (client *Client) CreateUploadPathForSQLServer(request *CreateUploadPathForSQLServerRequest) (response *CreateUploadPathForSQLServerResponse, err error) {
	response = CreateCreateUploadPathForSQLServerResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateUploadPathForSQLServerWithChan(request *CreateUploadPathForSQLServerRequest) (<-chan *CreateUploadPathForSQLServerResponse, <-chan error) {
	responseChan := make(chan *CreateUploadPathForSQLServerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUploadPathForSQLServer(request)
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

func (client *Client) CreateUploadPathForSQLServerWithCallback(request *CreateUploadPathForSQLServerRequest, callback func(response *CreateUploadPathForSQLServerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUploadPathForSQLServerResponse
		var err error
		defer close(result)
		response, err = client.CreateUploadPathForSQLServer(request)
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

type CreateUploadPathForSQLServerRequest struct {
	*requests.RpcRequest
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBName               string           `position:"Query" name:"DBName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type CreateUploadPathForSQLServerResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	InternetFtpServer string `json:"InternetFtpServer" xml:"InternetFtpServer"`
	InternetPort      int    `json:"InternetPort" xml:"InternetPort"`
	IntranetFtpserver string `json:"IntranetFtpserver" xml:"IntranetFtpserver"`
	Intranetport      int    `json:"Intranetport" xml:"Intranetport"`
	UserName          string `json:"UserName" xml:"UserName"`
	Password          string `json:"Password" xml:"Password"`
	FileName          string `json:"FileName" xml:"FileName"`
}

func CreateCreateUploadPathForSQLServerRequest() (request *CreateUploadPathForSQLServerRequest) {
	request = &CreateUploadPathForSQLServerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateUploadPathForSQLServer", "rds", "openAPI")
	return
}

func CreateCreateUploadPathForSQLServerResponse() (response *CreateUploadPathForSQLServerResponse) {
	response = &CreateUploadPathForSQLServerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
