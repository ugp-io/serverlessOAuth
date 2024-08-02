package ups

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	//"fmt"
)

func Rating(from Location, to Location, serviceCode string, length string, width string, height string, weight string) (int64, error) {

	var rate int64
	var response UPSRatingResponse
	accessCode, err := requestToken()
	if err != nil {
		return rate, err
	}

	request := UPSRatingRequest{
		RateRequest: RatingRateRequest{
			Request: RatingRequest{
				RequestOption: "Rate",
			},
			Shipment: RatingUPSShipment{
				ShipmentRatingOptions: RatingShipmentRatingOptions{
					NegotiatedRatesIndicator: "",
				},
				Shipper: RatingShipper{
					ShipperNumber: "35W170",
					Address: RatingAddress{
						City:              from.City,
						StateProvinceCode: from.StateProvinceCode,
						PostalCode:        from.PostalCode,
						CountryCode:       from.CountryCode,
					},
				},
				ShipTo: RatingShipTo{
					Address: RatingAddress{
						City:              to.City,
						StateProvinceCode: to.StateProvinceCode,
						PostalCode:        to.PostalCode,
						CountryCode:       to.CountryCode,
					},
				},
				Service: RatingService{
					Code: serviceCode,
				},
				Package: RatingPackage{
					PackagingType: RatingPackagingType{
						Code: "02",
					},
					Dimensions: RatingDimensions{
						UnitOfMeasurement: RatingUnitOfMeasurement{
							Code: "IN",
						},
						Length: length,
						Width:  width,
						Height: height,
					},
					PackageWeight: RatingPackageWeight{
						UnitOfMeasurement: RatingUnitOfMeasurement{
							Code: "Lbs",
						},
						Weight: weight,
					},
				},
				ShipmentTotalWeight: RatingPackageWeight{
					UnitOfMeasurement: RatingUnitOfMeasurement{
						Code: "Lbs",
					},
					Weight: weight,
				},
			},
		},
	}

	requestBytes, err := json.Marshal(request)
	if err != nil {
		return rate, err
	}

	httpReq, err := http.NewRequest(
		"POST",
		// "https://onlinetools.ups.com/api/rating/v2403/Rate",
		"https://wwwcie.ups.com/api/rating/v2403/Rate",
		bytes.NewBuffer(requestBytes))
	if err != nil {
		return rate, err
	}

	httpReq.Header.Add("Authorization", "Bearer "+accessCode)
	httpReq.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return rate, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return rate, err
	}

	bolB, _ := json.Marshal(response)
	fmt.Println(string(bolB))

	//Format
	if resp.StatusCode == 200 {

		if response.RateResponse.RatedShipment.NegotiatedRateCharges.TotalCharge.MonetaryValue != "" {
			stringRate := strings.Replace(response.RateResponse.RatedShipment.NegotiatedRateCharges.TotalCharge.MonetaryValue, ".", "", 1)
			rate, err = strconv.ParseInt(stringRate, 10, 64)
		} else {
			if response.RateResponse.RatedShipment.TotalCharges.MonetaryValue != "" {
				stringRate := strings.Replace(response.RateResponse.RatedShipment.TotalCharges.MonetaryValue, ".", "", 1)
				rate, err = strconv.ParseInt(stringRate, 10, 64)
			}
		}
	}

	// Error
	if response.RatingError != nil {
		err := errors.New(response.RatingError.Error[0].Message)
		if err != nil {
			return rate, err
		}
	}

	return rate, err
}
