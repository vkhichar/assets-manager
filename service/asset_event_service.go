package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/vkhichar/assets-manager/config"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/domain"
)

type AssetEventService interface {
	PostCreateAssetEvent(ctx context.Context, asset *domain.Asset) (string, error)
	PostUpdateAssetEvent(ctx context.Context, asset *domain.Asset) (string, error)
}

type assetEventService struct {
}

func NewAssetEventService() AssetEventService {
	return &assetEventService{}
}

func (assetEvSrv *assetEventService) PostCreateAssetEvent(ctx context.Context, asset *domain.Asset) (string, error) {
	//Request struct Object
	reqObj := contract.CreateAssetEventReq{
		EventType: "asset",
		Data:      asset,
	}

	//convert struct to byte array(json marshal)
	requestBody, err := json.Marshal(reqObj)
	if err != nil {
		fmt.Printf("Asset Event Service: Error while marshaling %s", err.Error())
		return "", err
	}
	//convert bytes array to bytes reader
	reqReader := bytes.NewReader(requestBody)

	//create a http request by using the bytes reader
	req, err := http.NewRequest("POST", config.GetAssetServiceURL()+"/events", reqReader)

	if err != nil {
		fmt.Printf("Asset Event Service: Error during http request %s", err.Error())
		return "", err
	}

	//create a new http client to make http request
	var netClient = &http.Client{
		Timeout: time.Second * 3, //set timeout
	}

	//send the http request and get the http response
	response, err := netClient.Do(req)

	if err != nil {
		fmt.Printf("Asset Event service: Request Timeout %s: Taking more than %v seconds", err.Error(), response)
		return "", err
	}

	//convert http response body to bytes array
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Asset Event Service : error while converting into byte stream: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var respObj contract.CreateAssetEventResp

	//unmarshal bytes array to the desired contract type
	err = json.Unmarshal(body, &respObj)
	if err != nil {
		fmt.Printf("Event service: Error while json unmarshal. Error: %s", err.Error())
		return "", err
	}

	eventId := strconv.Itoa(respObj.Id)

	//send the id from the response object
	return eventId, nil

}

func (assetEvSrv *assetEventService) PostUpdateAssetEvent(ctx context.Context, asset *domain.Asset) (string, error) {
	//Request struct Object
	reqObj := contract.UpdateAssetEventReq{
		EventType: "asset",
		Data:      asset,
	}

	//convert struct to byte array(json marshal)
	requestBody, err := json.Marshal(reqObj)
	if err != nil {
		fmt.Printf("Asset Event Service: Error while marshaling %s", err.Error())
		return "", err
	}
	//convert bytes array to bytes reader
	reqReader := bytes.NewReader(requestBody)

	//create a http request by using the bytes reader
	req, err := http.NewRequest("POST", config.GetAssetServiceURL()+"/events", reqReader)

	if err != nil {
		fmt.Printf("Asset Event Service: Error during http request %s", err.Error())
		return "", err
	}

	//create a new http client to make http request
	var netClient = &http.Client{
		Timeout: time.Second * 3, //set timeout
	}
	req.Header.Add("Content-type", "application/json")
	//send the http request and get the http response
	response, err := netClient.Do(req)

	if err != nil {
		fmt.Printf("Asset Event server: Request Timeout %s: Taking more than %v seconds", err.Error(), response)
		return "", err
	}

	//convert http response body to bytes array
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Asset Event Service : error while converting into byte stream: %s", err.Error())
		return "", err
	}

	var respObj contract.UpdateAssetEventResp

	//unmarshal bytes array to the desired contract type
	err = json.Unmarshal(body, &respObj)
	if err != nil {
		fmt.Printf("Event service: Error while json unmarshal. Error: %s", err.Error())
		return "", err
	}

	eventId := strconv.Itoa(respObj.Id)

	//send the id from the response object
	return eventId, nil

}
