package ups

type Location struct {
	City              string
	StateProvinceCode string
	PostalCode        string
	CountryCode       string
}

type UPSTokenResponse struct {
	TokenType    string `json:"token_type"`
	IssuedAt     string `json:"issued_at"`
	ClientID     string `json:"client_id"`
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	ExpiresIn    string `json:"expires_in"`
	RefreshCount string `json:"refresh_count"`
	Status       string `json:"status"`
	Response     struct {
		Errors []UPSErrorResponse
	} `json:"response"`
}

type UPSErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

/*	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ Quantum View ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~	*/

type UPSQuantumViewResponse struct {
	QuantumViewResponse *QuantumViewQuantumViewResponse `json:"quantumViewResponse,omitempty"`
	QuantumViewkError   *struct {
		Error []UPSErrorResponse `json:"errors"`
	} `json:"response,omitempty"`
}

type QuantumViewQuantumViewResponse struct {
	Response          QuantumViewResponse `json:"response"`
	QuantumViewEvents QuantumViewEvents   `json:"quantumViewEvents"`
	Bookmark          string              `json:"bookmark"`
}

type QuantumViewResponse struct {
	TransactionReference      TransactionReferenceResponse `json:"transactionReference"`
	ResponseStatusCode        string                       `json:"responseStatusCode"`
	ResponseStatusDescription string                       `json:"responseStatusDescription"`
	Error                     []QuantumViewError           `json:"error"`
}

type TransactionReferenceResponse struct {
	CustomerContext string `json:"customerContext"`
	XpciVersion     string `json:"xpciVersion"`
	ToolVersion     string `json:"toolVersion"`
}

type QuantumViewError struct {
	ErrorSeverity       string          `json:"errorSeverity"`
	ErrorCode           string          `json:"errorCode"`
	ErrorDescription    string          `json:"errorDescription"`
	MinimumRetrySeconds string          `json:"minimumRetrySeconds"`
	ErrorLocation       []ErrorLocation `json:"errorLocation"`
	ErrorDigest         []string        `json:"errorDigest"`
}

type ErrorLocation struct {
	ErrorLocationElementName   string `json:"errorLocationElementName"`
	ErrorLocationAttributeName string `json:"errorLocationAttributeName"`
}

type QuantumViewEvents struct {
	SubscriberID       string                          `json:"subscriberID"`
	SubscriptionEvents []QuantumViewSubscriptionEvents `json:"subscriptionEvents"`
}

type QuantumViewSubscriptionEvents struct {
	Name               string                        `json:"name"`
	Number             string                        `json:"number"`
	SubscriptionStatus QuantumViewSubscriptionStatus `json:"subscriptionStatus"`
	DateRange          QuantumViewDateRange          `json:"dateRange"`
	SubscriptionFiles  []QuantumViewSubscriptionFile `json:"subscriptionFile"`
}

type QuantumViewSubscriptionStatus struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type QuantumViewDateRange struct {
	BeginDate string `json:"beginDate"`
	EndDate   string `json:"endDate"`
}

type QuantumViewSubscriptionFile struct {
	FileName   string                `json:"fileName"`
	StatusType QuantumViewStatusType `json:"statusType"`
	Manifest   []interface{}         `json:"manifest"`
	Origin     []interface{}         `json:"origin"`
	Exception  []interface{}         `json:"exception"`
	Delivery   []interface{}         `json:"delivery"`
	Generic    []interface{}         `json:"generic"`
}

type QuantumViewStatusType struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type QuantumViewRequestParams struct {
	BeginDateTime    string
	EndDateTime      string
	Bookmark         string
	SubscriptionName string
}

type UPSQuantumViewRequest struct {
	QuantumViewRequest QuantumViewQuantumViewRequest `json:"quantumViewRequest"`
}

type QuantumViewQuantumViewRequest struct {
	Request             QuantumViewRequest             `json:"request"`
	SubscriptionRequest QuantumViewSubscriptionRequest `json:"subscriptionRequest"`
	Bookmark            string                         `json:"bookmark"`
}

type QuantumViewRequest struct {
	RequestAction        string                      `json:"requestAction"`
	TransactionReference TransactionReferenceRequest `json:"transactionReference"`
}

type TransactionReferenceRequest struct {
	CustomerContext string `json:"customerContext"`
}

type QuantumViewSubscriptionRequest struct {
	Name          string                   `json:"name"`
	DateTimeRange QuantumViewDateTimeRange `json:"dateTimeRange"`
}

type QuantumViewDateTimeRange struct {
	BeginDateTime string `json:"beginDateTime"`
	EndDateTime   string `json:"endDateTime"`
}

/*	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ RATING ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~	*/

type UPSRatingRequest struct {
	RateRequest RatingRateRequest `json:"RateRequest"`
}

type RatingRateRequest struct {
	Request  RatingRequest     `json:"Request"`
	Shipment RatingUPSShipment `json:"Shipment"`
}

type RatingRequest struct {
	RequestOption string `json:"RequestOption"`
}

type RatingUPSShipment struct {
	ShipmentRatingOptions RatingShipmentRatingOptions `json:"ShipmentRatingOptions"`
	Shipper               RatingShipper               `json:"Shipper"`
	ShipTo                RatingShipTo                `json:"ShipTo"`
	Service               RatingService               `json:"Service"`
	Package               RatingPackage               `json:"Package"`
	ShipmentTotalWeight   RatingPackageWeight         `json:"ShipmentTotalWeight"`
}

type RatingShipmentRatingOptions struct {
	NegotiatedRatesIndicator string `json:"NegotiatedRatesIndicator"`
}

type RatingShipper struct {
	Name          string        `json:"Name"`
	ShipperNumber string        `json:"ShipperNumber"`
	Address       RatingAddress `json:"Address"`
}
type RatingShipTo struct {
	Name          string        `json:"Name"`
	AttentionName string        `json:"AttentionName"`
	Address       RatingAddress `json:"Address"`
}

type RatingService struct {
	Code        string `json:"Code"`
	Description string `json:"Description"`
}

type RatingPackage struct {
	PackagingType RatingPackagingType `json:"PackagingType"`
	Dimensions    RatingDimensions    `json:"Dimensions"`
	PackageWeight RatingPackageWeight `json:"PackageWeight"`
}

type RatingPackageWeight struct {
	UnitOfMeasurement RatingUnitOfMeasurement `json:"UnitOfMeasurement"`
	Weight            string                  `json:"Weight"`
}

type RatingAddress struct {
	AddressLine                 []string `json:"AddressLine"`
	City                        string   `json:"City"`
	StateProvinceCode           string   `json:"StateProvinceCode"`
	PostalCode                  string   `json:"PostalCode"`
	CountryCode                 string   `json:"CountryCode"`
	ResidentialAddressIndicator string   `json:"ResidentialAddressIndicator"`
}

type RatingPackagingType struct {
	Code string `json:"Code"`
}

type RatingDimensions struct {
	UnitOfMeasurement RatingUnitOfMeasurement `json:"UnitOfMeasurement"`
	Length            string                  `json:"Length"`
	Width             string                  `json:"Width"`
	Height            string                  `json:"Height"`
}

type RatingUnitOfMeasurement struct {
	Code string `json:"Code"`
}

type UPSRatingResponse struct {
	RateResponse RatingRateResponse `json:"RateResponse,omitempty"`
	RatingError  *struct {
		Error []RatingErrorResponse `json:"errors"`
	} `json:"response,omitempty"`
}

type RatingRateResponse struct {
	Response      RatingResponse      `json:"Response"`
	RatedShipment RatingRatedShipment `json:"RatedShipment"`
}

type RatingResponse struct {
	ResponseStatus       RatingResponseStatus  `json:"ResponseStatus"`
	Alert                []RatingAlertResponse `json:"Alert"`
	AlertDetail          []RatingAlertResponse `json:"AlertDetail"`
	TransactionReference string                `json:"TransactionReference"`
}

type RatingResponseStatus struct {
	Code        string `json:"Code"`
	Description string `json:"Description"`
}

type RatingAlertResponse struct {
	Code                    string      `json:"Code"`
	Description             string      `json:"Description"`
	ElementLevelInformation interface{} `json:"elementLevelInformation"`
}

type RatingTransactionReferenceResponse struct {
	CustomerContext string `json:"CustomerContext"`
}

type RatingErrorResponse struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

type RatingRatedShipment struct {
	TotalCharges          RatingTotalCharges          `json:"TotalCharges"`
	NegotiatedRateCharges RatingNegotiatedRateCharges `json:"NegotiatedRateCharges"`
}

type RatingTotalCharges struct {
	CurrencyCode  string `json:"CurrencyCode"`
	MonetaryValue string `json:"MonetaryValue"`
}

type RatingNegotiatedRateCharges struct {
	TotalCharge RatingTotalCharges `json:"TotalCharge"`
}

/*	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ Time In Transit ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~	*/

type UPSTimeInTransitRequest struct {
	OriginCountryCode            string `json:"originCountryCode,omitempty"`
	OriginStateProvince          string `json:"originStateProvince,omitempty"`
	OriginCityName               string `json:"originCityName,omitempty"`
	OriginTownName               string `json:"originTownName,omitempty"`
	OriginPostalCode             string `json:"originPostalCode,omitempty"`
	DestinationCountryCode       string `json:"destinationCountryCode,omitempty"`
	DestinationStateProvince     string `json:"destinationStateProvince,omitempty"`
	DestinationCityName          string `json:"destinationCityName,omitempty"`
	DestinationTownName          string `json:"destinationTownName,omitempty"`
	DestinationPostalCode        string `json:"destinationPostalCode,omitempty"`
	Weight                       string `json:"weight,omitempty"`
	WeightUnitOfMeasure          string `json:"weightUnitOfMeasure,omitempty"`
	ShipmentContentsValue        string `json:"shipmentContentsValue,omitempty"`
	ShipmentContentsCurrencyCode string `json:"shipmentContentsCurrencyCode,omitempty"`
	BillType                     string `json:"billType,omitempty"`
	ShipDate                     string `json:"shipDate,omitempty"`
	ShipTime                     string `json:"shipTime,omitempty"`
	ResidentialIndicator         string `json:"residentialIndicator,omitempty"`
	AvvFlag                      bool   `json:"avvFlag,omitempty"`
	NumberOfPackages             string `json:"numberOfPackages,omitempty"`
}

type UPSTimeInTransitResponse struct {
	ValidationList      *TITValidationListResponse `json:"validationList,omitempty"`
	DestinationPickList *[]TITPickListResponse     `json:"destinationPickList,omitempty"`
	OriginPickList      *[]TITPickListResponse     `json:"originPickList,omitempty"`
	EMSResponse         *TITEMSResponse            `json:"emsResponse,omitempty"`
	TITError            *struct {
		Error []UPSErrorResponse `json:"errors"`
	} `json:"response,omitempty"`
}

type TITValidationListResponse struct {
	InvalidFieldList      []string `json:"invalidFieldList"`
	InvalidFieldListCodes []string `json:"invalidFieldListCodes"`
	DestinationAmbiguous  bool     `json:"destinationAmbiguous"`
	OriginAmbiguous       bool     `json:"originAmbiguous"`
}

type TITPickListResponse struct {
	CountryName    string `json:"countryName"`
	CountryCode    string `json:"countryCode"`
	StateProvince  string `json:"stateProvince"`
	City           string `json:"city"`
	Town           string `json:"town"`
	PostalCode     string `json:"postalCode"`
	PostalCodeLow  string `json:"postalCodeLow"`
	PostalCodeHigh string `json:"postalCodeHigh"`
}

type TITEMSResponse struct {
	ShipDate                     string               `json:"shipDate"`
	ShipTime                     string               `json:"shipTime"`
	ServiceLevel                 string               `json:"serviceLevel"`
	BillType                     string               `json:"billType"`
	DutyType                     string               `json:"dutyType"`
	ResidentialIndicator         string               `json:"residentialIndicator"`
	DestinationCountryName       string               `json:"destinationCountryName"`
	DestinationCountryCode       string               `json:"destinationCountryCode"`
	DestinationPostalCode        string               `json:"destinationPostalCode"`
	DestinationPostalCodeLow     string               `json:"destinationPostalCodeLow"`
	DestinationPostalCodeHigh    string               `json:"destinationPostalCodeHigh"`
	DestinationStateProvince     string               `json:"destinationStateProvince"`
	DestinationCityName          string               `json:"destinationCityName"`
	OriginCountryName            string               `json:"originCountryName"`
	OriginCountryCode            string               `json:"originCountryCode"`
	OriginPostalCode             string               `json:"originPostalCode"`
	OriginPostalCodeLow          string               `json:"originPostalCodeLow"`
	OriginPostalCodeHigh         string               `json:"originPostalCodeHigh"`
	OriginStateProvince          string               `json:"originStateProvince"`
	OriginCityName               string               `json:"originCityName"`
	Weight                       string               `json:"weight"`
	WeightUnitOfMeasure          string               `json:"weightUnitOfMeasure"`
	ShipmentContentsValue        string               `json:"shipmentContentsValue"`
	ShipmentContentsCurrencyCode string               `json:"shipmentContentsCurrencyCode"`
	GuaranteeSuspended           bool                 `json:"guaranteeSuspended"`
	NumberOfServices             int                  `json:"numberOfServices"`
	Services                     []TITServiceResponse `json:"services"`
}

type TITServiceResponse struct {
	ServiceLevel            string `json:"serviceLevel"`
	ServiceLevelDescription string `json:"serviceLevelDescription"`
	ShipDate                string `json:"shipDate"`
	DeliveryDate            string `json:"deliveryDate"`
	CommitTime              string `json:"commitTime"`
	DeliveryTime            string `json:"deliveryTime"`
	DeliveryDayOfWeek       string `json:"deliveryDayOfWeek"`
	NextDayPickupIndicator  string `json:"nextDayPickupIndicator"`
	SaturdayPickupIndicator string `json:"saturdayPickupIndicator"`
	SaturdayDeliveryDate    string `json:"saturdayDeliveryDate"`
	SaturdayDeliveryTime    string `json:"saturdayDeliveryTime"`
	ServiceRemarksText      string `json:"serviceRemarksText"`
	GuaranteeIndicator      string `json:"guaranteeIndicator"`
	TotalTransitDays        int    `json:"totalTransitDays"`
	BusinessTransitDays     int    `json:"businessTransitDays"`
	RestDaysCount           int    `json:"restDaysCount"`
	HolidayCount            int    `json:"holidayCount"`
	DelayCount              int    `json:"delayCount"`
	PickupDate              string `json:"pickupDate"`
	PickupTime              string `json:"pickupTime"`
	CSTCCutoffTime          string `json:"cstccutoffTime"`
	PODDate                 string `json:"poddate"`
	PODDays                 int    `json:"poddays"`
}

type TITError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

/*	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ Track ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~	*/

type UPSTrackResponse struct {
	TrackResponse *TrackResponse `json:"trackResponse,omitempty"`
	TrackError    *struct {
		Error []UPSErrorResponse `json:"errors"`
	} `json:"response"`
}

type TrackResponse struct {
	Shipments []TrackShipmentResponse `json:"shipment,omitempty"`
}

type TrackShipmentResponse struct {
	InquiryNumber string                 `json:"inquiryNumber"`
	Warnings      []TrackWarningResponse `json:"warnings"`
	UserRelations []string               `json:"userRelation"`
	Packages      []TrackPackageResponse `json:"package"`
}

type TrackWarningResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type TrackPackageResponse struct {
	AccessPointInformation struct {
		PickupByDate string `json:"pickupByDate"`
	} `json:"accessPointInformation"`
	Activity                []TrackActivityResponse                `json:"activity"`
	AdditionalAttributes    []string                               `json:"additionalAttributes"`
	AdditionalServices      []string                               `json:"additionalServices"`
	AlternateTrackingNumber []TrackAlternateTrackingNumberResponse `json:"alternateTrackingNumber"`
	CurrentStatus           TrackStatusResponse                    `json:"currentStatus"`
	DeliveryDate            []TrackDeliveryDateResponse            `json:"deliveryDate"`
	DeliveryInformation     TrackDeliveryInformationResponse       `json:"deliveryInformation"`
	DeliveryTime            TrackDeliveryTimeResponse              `json:"deliveryTime"`
	Milestones              []TrackMilestoneResponse               `json:"milestones"`
	PackageAddress          []TrackPackageAddressResponse          `json:"packageAddress"`
	PackageCount            int                                    `json:"packageCount"`
	PaymentInformation      []TrackPaymentInformationResponse      `json:"paymentInformation"`
	ReferenceNumber         []TrackAlternateTrackingNumberResponse `json:"referenceNumber"`
	Service                 TrackServiceResponse                   `json:"service"`
	StatusCode              string                                 `json:"statusCode"`
	StatusDescription       string                                 `json:"statusDescription"`
	SuppressionIndicators   string                                 `json:"suppressionIndicators"`
	TrackingNumber          string                                 `json:"trackingNumber"`
	Weight                  TrackPackageWeightResponse             `json:"weight"`
}

type TrackActivityResponse struct {
	Date      string                `json:"date"`
	GmtDate   string                `json:"gmtDate"`
	GmtOffset string                `json:"gmtOffset"`
	GmtTime   string                `json:"gmtTime"`
	Location  TrackLocationResponse `json:"location"`
	Status    TrackStatusResponse   `json:"status"`
	Time      string                `json:"time"`
}

type TrackLocationResponse struct {
	Address TrackLocationAddressResponse `json:"address"`
	SLIC    string                       `json:"slic"`
}

type TrackLocationAddressResponse struct {
	AddressLine1  string `json:"addressLine1"`
	AddressLine2  string `json:"addressLine2"`
	AddressLine3  string `json:"addressLine3"`
	City          string `json:"city"`
	Country       string `json:"country"`
	CountryCode   string `json:"countryCode"`
	PostalCode    string `json:"postalCode"`
	StateProvince string `json:"stateProvince"`
}

type TrackAlternateTrackingNumberResponse struct {
	Number string `json:"number"`
	Type   string `json:"type"`
}

type TrackStatusResponse struct {
	Code                      string `json:"code"`
	Description               string `json:"description"`
	SimplifiedTextDescription string `json:"simplifiedTextDescription"`
	StatusCode                string `json:"statusCode"`
	Type                      string `json:"type"`
}

type TrackDeliveryDateResponse struct {
	Date string `json:"date"`
	Type string `json:"type"`
}

type TrackDeliveryInformationResponse struct {
	Location   string `json:"location"`
	ReceivedBy string `json:"receivedBy"`
	Signature  struct {
		Image string `json:"image"`
	} `json:"signature"`
	Pod struct {
		Content string `json:"content"`
	} `json:"pod"`
}

type TrackDeliveryTimeResponse struct {
	EndTime   string `json:"endTime"`
	StartTime string `json:"startTime"`
	Type      string `json:"type"`
}

type TrackMilestoneResponse struct {
	Category       string `json:"category"`
	Code           string `json:"code"`
	Current        string `json:"current"`
	Description    string `json:"description"`
	LinkedActivity string `json:"linkedActivity"`
	State          string `json:"state"`
	SubMilestone   struct {
		Category string `json:"category"`
	} `json:"subMilestone"`
}

type TrackPackageAddressResponse struct {
	Address       TrackLocationAddressResponse `json:"address"`
	AttentionName string                       `json:"attentionName"`
	Name          string                       `json:"name"`
	Type          string                       `json:"type"`
}

type TrackPaymentInformationResponse struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	ID            string `json:"id"`
	Paid          string `json:"paid"`
	PaymentMethod string `json:"paymentMethod"`
	Type          string `json:"type"`
}

type TrackServiceResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	LevelCode   string `json:"levelCode"`
}

type TrackPackageWeightResponse struct {
	UnitOfMeasurement string `json:"unitOfMeasurement"`
	Weight            string `json:"weight"`
}
