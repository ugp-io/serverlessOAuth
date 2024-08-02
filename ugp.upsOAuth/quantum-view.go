package ups

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func QVEvents(params QuantumViewRequestParams) (UPSQuantumViewResponse, error) {

	var response UPSQuantumViewResponse
	accessCode, err := requestToken()
	if err != nil {
		return response, err
	}

	request := UPSQuantumViewRequest{
		QuantumViewRequest: QuantumViewQuantumViewRequest{
			Request: QuantumViewRequest{
				RequestAction: "QVEvents",
			},
			SubscriptionRequest: QuantumViewSubscriptionRequest{
				Name: params.SubscriptionName,
				DateTimeRange: QuantumViewDateTimeRange{
					BeginDateTime: params.BeginDateTime,
					EndDateTime:   params.EndDateTime,
				},
			},
			Bookmark: params.Bookmark,
		},
	}

	bolB, _ := json.Marshal(request)
	fmt.Println("Request:\n", string(bolB))

	requestBytes, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	httpReq, errNewRequest := http.NewRequest(
		"POST",
		// "https://onlinetools.ups.com/api/quantumview/v2/events",
		"https://wwwcie.ups.com/api/quantumview/v2/events",
		bytes.NewBuffer(requestBytes))
	if errNewRequest != nil {
		return response, errNewRequest
	}

	httpReq.Header.Add("Authorization", "Bearer "+accessCode)
	httpReq.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	//Decode
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	bolB, _ = json.Marshal(response)
	fmt.Println("Response:\n", string(bolB))

	//Error
	if response.QuantumViewResponse != nil && len(response.QuantumViewResponse.Response.Error) > 0 {
		for _, errorC := range response.QuantumViewResponse.Response.Error {
			err := errors.New(errorC.ErrorDescription)
			if err != nil {
				return response, err
			}
		}
	}

	return response, nil
}
