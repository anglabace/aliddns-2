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

func (client *Client) DescribeBackupsForSecurity(request *DescribeBackupsForSecurityRequest) (response *DescribeBackupsForSecurityResponse, err error) {
	response = CreateDescribeBackupsForSecurityResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeBackupsForSecurityWithChan(request *DescribeBackupsForSecurityRequest) (<-chan *DescribeBackupsForSecurityResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupsForSecurityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupsForSecurity(request)
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

func (client *Client) DescribeBackupsForSecurityWithCallback(request *DescribeBackupsForSecurityRequest, callback func(response *DescribeBackupsForSecurityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupsForSecurityResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupsForSecurity(request)
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

type DescribeBackupsForSecurityRequest struct {
	*requests.RpcRequest
	EndTime              string           `position:"Query" name:"EndTime"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	BackupStatus         string           `position:"Query" name:"BackupStatus"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	TargetAliBid         string           `position:"Query" name:"TargetAliBid"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	BackupLocation       string           `position:"Query" name:"BackupLocation"`
	BackupId             string           `position:"Query" name:"BackupId"`
	BackupMode           string           `position:"Query" name:"BackupMode"`
	TargetAliUid         string           `position:"Query" name:"TargetAliUid"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DescribeBackupsForSecurityResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	TotalRecordCount string `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       string `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  string `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalBackupSize  int    `json:"TotalBackupSize" xml:"TotalBackupSize"`
	Items            struct {
		Backup []struct {
			BackupId                  string `json:"BackupId" xml:"BackupId"`
			DBInstanceId              string `json:"DBInstanceId" xml:"DBInstanceId"`
			BackupStatus              string `json:"BackupStatus" xml:"BackupStatus"`
			BackupStartTime           string `json:"BackupStartTime" xml:"BackupStartTime"`
			BackupEndTime             string `json:"BackupEndTime" xml:"BackupEndTime"`
			BackupType                string `json:"BackupType" xml:"BackupType"`
			BackupMode                string `json:"BackupMode" xml:"BackupMode"`
			BackupMethod              string `json:"BackupMethod" xml:"BackupMethod"`
			BackupDownloadURL         string `json:"BackupDownloadURL" xml:"BackupDownloadURL"`
			BackupIntranetDownloadURL string `json:"BackupIntranetDownloadURL" xml:"BackupIntranetDownloadURL"`
			BackupLocation            string `json:"BackupLocation" xml:"BackupLocation"`
			BackupExtractionStatus    string `json:"BackupExtractionStatus" xml:"BackupExtractionStatus"`
			BackupScale               string `json:"BackupScale" xml:"BackupScale"`
			BackupDBNames             string `json:"BackupDBNames" xml:"BackupDBNames"`
			TotalBackupSize           int    `json:"TotalBackupSize" xml:"TotalBackupSize"`
			BackupSize                int    `json:"BackupSize" xml:"BackupSize"`
			HostInstanceID            string `json:"HostInstanceID" xml:"HostInstanceID"`
			StoreStatus               string `json:"StoreStatus" xml:"StoreStatus"`
		} `json:"Backup" xml:"Backup"`
	} `json:"Items" xml:"Items"`
}

func CreateDescribeBackupsForSecurityRequest() (request *DescribeBackupsForSecurityRequest) {
	request = &DescribeBackupsForSecurityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeBackupsForSecurity", "rds", "openAPI")
	return
}

func CreateDescribeBackupsForSecurityResponse() (response *DescribeBackupsForSecurityResponse) {
	response = &DescribeBackupsForSecurityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
