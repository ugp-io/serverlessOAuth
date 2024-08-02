package ups

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func TimeInTransit(from Location, to Location, inHandsTime time.Time, transId, transactionSrc string) (UPSTimeInTransitResponse, error) {

	var response UPSTimeInTransitResponse
	accessCode, err := requestToken()
	if err != nil {
		return response, err
	}

	requestBytes, err := json.Marshal(UPSTimeInTransitRequest{
		OriginCountryCode:            from.CountryCode,
		OriginStateProvince:          from.StateProvinceCode,
		OriginCityName:               from.City,
		OriginPostalCode:             from.PostalCode,
		DestinationCountryCode:       to.CountryCode,
		DestinationStateProvince:     to.StateProvinceCode,
		DestinationCityName:          to.City,
		DestinationPostalCode:        to.PostalCode,
		ShipDate:                     inHandsTime.Format("20060102"),
		ShipTime:                     inHandsTime.Format("150405"),
		Weight:                       "1",
		WeightUnitOfMeasure:          "LBS",
		ShipmentContentsValue:        "1",
		ShipmentContentsCurrencyCode: "USD",
	})
	if err != nil {
		return response, err
	}

	httpReq, err := http.NewRequest(
		"POST",
		// "https://onlinetools.ups.com/api/shipments/v1/transittimes",
		"https://wwwcie.ups.com/api/shipments/v1/transittimes",
		bytes.NewBuffer(requestBytes))
	if err != nil {
		return response, err
	}

	httpReq.Header.Add("Authorization", "Bearer "+accessCode)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("transId", transId)
	httpReq.Header.Add("transactionSrc", transactionSrc)

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

	return response, err
}
