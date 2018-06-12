package vpc

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

func (client *Client) DescribeBandwidthPackagePublicIpMonitorData(request *DescribeBandwidthPackagePublicIpMonitorDataRequest) (response *DescribeBandwidthPackagePublicIpMonitorDataResponse, err error) {
	response = CreateDescribeBandwidthPackagePublicIpMonitorDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeBandwidthPackagePublicIpMonitorDataWithChan(request *DescribeBandwidthPackagePublicIpMonitorDataRequest) (<-chan *DescribeBandwidthPackagePublicIpMonitorDataResponse, <-chan error) {
	responseChan := make(chan *DescribeBandwidthPackagePublicIpMonitorDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBandwidthPackagePublicIpMonitorData(request)
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

func (client *Client) DescribeBandwidthPackagePublicIpMonitorDataWithCallback(request *DescribeBandwidthPackagePublicIpMonitorDataRequest, callback func(response *DescribeBandwidthPackagePublicIpMonitorDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBandwidthPackagePublicIpMonitorDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeBandwidthPackagePublicIpMonitorData(request)
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

type DescribeBandwidthPackagePublicIpMonitorDataRequest struct {
	*requests.RpcRequest
	EndTime              string           `position:"Query" name:"EndTime"`
	StartTime            string           `position:"Query" name:"StartTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Period               requests.Integer `position:"Query" name:"Period"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AllocationId         string           `position:"Query" name:"AllocationId"`
}

type DescribeBandwidthPackagePublicIpMonitorDataResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	MonitorDatas struct {
		MonitorData []struct {
			RX                   int    `json:"RX" xml:"RX"`
			TX                   int    `json:"TX" xml:"TX"`
			ReceivedBandwidth    int    `json:"ReceivedBandwidth" xml:"ReceivedBandwidth"`
			TransportedBandwidth int    `json:"TransportedBandwidth" xml:"TransportedBandwidth"`
			Flow                 int    `json:"Flow" xml:"Flow"`
			Bandwidth            int    `json:"Bandwidth" xml:"Bandwidth"`
			Packets              int    `json:"Packets" xml:"Packets"`
			TimeStamp            string `json:"TimeStamp" xml:"TimeStamp"`
		} `json:"MonitorData" xml:"MonitorData"`
	} `json:"MonitorDatas" xml:"MonitorDatas"`
}

func CreateDescribeBandwidthPackagePublicIpMonitorDataRequest() (request *DescribeBandwidthPackagePublicIpMonitorDataRequest) {
	request = &DescribeBandwidthPackagePublicIpMonitorDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeBandwidthPackagePublicIpMonitorData", "vpc", "openAPI")
	return
}

func CreateDescribeBandwidthPackagePublicIpMonitorDataResponse() (response *DescribeBandwidthPackagePublicIpMonitorDataResponse) {
	response = &DescribeBandwidthPackagePublicIpMonitorDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
