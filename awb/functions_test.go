package awb

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/denifrahman/shipper-go"
	"github.com/denifrahman/shipper-go/order"
	"github.com/joho/godotenv"
)

func init() {
	var err = godotenv.Load("../.env")

	if err != nil {
		panic("Error loading .env file.")
	}

	productionMode, errParse := strconv.ParseBool(os.Getenv("PRODUCTION_MODE"))

	if errParse != nil {
		productionMode = false
	}

	shipper.Conf.SetAPIKey(os.Getenv("API_KEY")).SetProductionMode(productionMode)
}

func TestGenerateAWB(t *testing.T) {
	externalID := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	domesticOrder, err := order.CreateDomesticOrder(&order.DomesticOrderParams{
		Origin:               12921,
		Destination:          4645,
		Length:               10,
		Width:                10,
		Height:               10,
		WeightTotal:          0.4,
		Value:                100000,
		RateID:               59,
		ConsigneeName:        "Peoorang",
		ConsigneePhoneNumber: "089899878987",
		ConsignerName:        "Peorang",
		ConsignerPhoneNumber: "089891891818",
		OriginAddress:        "Mangga Dua Selatan",
		OriginDirection:      "Just follow the road",
		DestinationAddress:   "Pasar Baru",
		DestinationDirection: "Lurus terus",
		ItemName: []order.ItemName{
			{
				Name:  "Baju",
				Qty:   1,
				Value: 100000,
			},
		},
		Contents:              "Barang mudah pecah",
		UseInsurance:          0,
		ExternalID:            externalID,
		PaymentType:           "cash",
		PackageType:           1,
		COD:                   0,
		OriginCoordinate:      "-6.308033944807303,106.73339847804874",
		DestinationCoordinate: "49.020733179213,12.114381752908",
	})

	if err != nil {
		t.Error(err.Error())
	}

	generatedAWB, err := Generate(&GenerateParams{
		OID: domesticOrder.Data.ID,
	})

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(generatedAWB, "", "\t")

	fmt.Println("Generated AWB:")
	fmt.Println(string(s))
}
