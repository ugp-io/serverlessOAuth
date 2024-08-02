package ups_test

import (
	//"time"

	"testing"
	"time"

	//"reflect"

	ups "github.com/ugp-io/serverless/ugp.upsOAuth"
)

func TestTimeInTransit(t *testing.T) {
	tables := []struct {
		from           ups.Location
		to             ups.Location
		inHandsTime    time.Time
		transId        string
		transactionSrc string
	}{
		{
			ups.Location{
				City:              "Ann Arbor",
				StateProvinceCode: "MI",
				PostalCode:        "48103",
				CountryCode:       "US",
			},
			ups.Location{
				City:              "Toronto",
				StateProvinceCode: "ON",
				PostalCode:        "M8X 2X9",
				CountryCode:       "CA",
			},
			time.Date(2018, time.January, 9, 23, 59, 0, 0, time.UTC),
			"test_transId",
			"test_transactionSrc",
		},
	}
	for _, table := range tables {
		ups.TimeInTransit(table.from, table.to, table.inHandsTime, table.transId, table.transactionSrc)
	}
}
