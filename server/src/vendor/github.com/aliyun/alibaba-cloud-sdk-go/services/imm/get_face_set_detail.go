package imm

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

// GetFaceSetDetail invokes the imm.GetFaceSetDetail API synchronously
// api document: https://help.aliyun.com/api/imm/getfacesetdetail.html
func (client *Client) GetFaceSetDetail(request *GetFaceSetDetailRequest) (response *GetFaceSetDetailResponse, err error) {
	response = CreateGetFaceSetDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetFaceSetDetailWithChan invokes the imm.GetFaceSetDetail API asynchronously
// api document: https://help.aliyun.com/api/imm/getfacesetdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFaceSetDetailWithChan(request *GetFaceSetDetailRequest) (<-chan *GetFaceSetDetailResponse, <-chan error) {
	responseChan := make(chan *GetFaceSetDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFaceSetDetail(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetFaceSetDetailWithCallback invokes the imm.GetFaceSetDetail API asynchronously
// api document: https://help.aliyun.com/api/imm/getfacesetdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFaceSetDetailWithCallback(request *GetFaceSetDetailRequest, callback func(response *GetFaceSetDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFaceSetDetailResponse
		var err error
		defer close(result)
		response, err = client.GetFaceSetDetail(request)
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

// GetFaceSetDetailRequest is the request struct for api GetFaceSetDetail
type GetFaceSetDetailRequest struct {
	*requests.RpcRequest
	Marker          string `position:"Query" name:"Marker"`
	Project         string `position:"Query" name:"Project"`
	SetId           string `position:"Query" name:"SetId"`
	ReturnAttribute string `position:"Query" name:"ReturnAttribute"`
}

// GetFaceSetDetailResponse is the response struct for api GetFaceSetDetail
type GetFaceSetDetailResponse struct {
	*responses.BaseResponse
	RequestId   string            `json:"RequestId" xml:"RequestId"`
	SetId       string            `json:"SetId" xml:"SetId"`
	NextMarker  string            `json:"NextMarker" xml:"NextMarker"`
	FaceDetails []FaceDetailsItem `json:"FaceDetails" xml:"FaceDetails"`
}

// CreateGetFaceSetDetailRequest creates a request to invoke GetFaceSetDetail API
func CreateGetFaceSetDetailRequest() (request *GetFaceSetDetailRequest) {
	request = &GetFaceSetDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetFaceSetDetail", "imm", "openAPI")
	return
}

// CreateGetFaceSetDetailResponse creates a response to parse from GetFaceSetDetail response
func CreateGetFaceSetDetailResponse() (response *GetFaceSetDetailResponse) {
	response = &GetFaceSetDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}