package ups

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Track(transId, transactionSrc, trackingNumber string) ([]TrackActivityResponse, error) {

	var response UPSTrackResponse
	var activities []TrackActivityResponse
	accessCode, err := requestToken()
	if err != nil {
		return activities, err
	}

	httpReq, err := http.NewRequest(
		"GET",
		// "https://onlinetools.ups.com/api/track/v1/details/"+trackingNumber,
		"https://wwwcie.ups.com/api/track/v1/details/"+trackingNumber,
		nil)
	if err != nil {
		return activities, err
	}

	httpReq.Header.Add("Authorization", "Bearer "+accessCode)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("transId", transId)
	httpReq.Header.Add("transactionSrc", transactionSrc)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return activities, err
	}
	defer resp.Body.Close()

	//Decode
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response", err)
		return activities, err
	}

	//Format
	if resp.StatusCode == 200 {

		bolB, _ := json.Marshal(response)
		fmt.Println(string(bolB))

		for _, shipment := range response.TrackResponse.Shipments {
			for _, trackPackage := range shipment.Packages {
				if trackPackage.TrackingNumber == trackingNumber {
					activities = trackPackage.Activity
					break
				}
			}
		}
	}

	return activities, err
}
