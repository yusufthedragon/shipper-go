package rates

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/yusufthedragon/shipper-go"
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

func TestGetDomesticRates(t *testing.T) {
	domesticRates, err := GetDomesticRates(&DomesticRatesParams{
		Origin:                12921,
		Destination:           4645,
		Length:                10,
		Width:                 10,
		Height:                10,
		WeightTotal:           0.4,
		Value:                 70000,
		Type:                  1,
		COD:                   1,
		Order:                 0,
		OriginCoordinate:      "-6.308033944807303,106.73339847804874",
		DestinationCoordinate: "49.020733179213,12.114381752908",
	})

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(domesticRates, "", "\t")

	fmt.Println("Domestic Rates:")
	fmt.Println(string(s))
}

func TestInternationalRates(t *testing.T) {
	internationalRates, err := GetInternationalRates(&InternationalRatesParams{
		Origin:      4819,
		Destination: 12,
		Length:      10,
		Width:       10,
		Height:      10,
		WeightTotal: 0.4,
		Value:       150000,
		Type:        2,
		Order:       0,
	})

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(internationalRates, "", "\t")

	fmt.Println("International Rates:")
	fmt.Println(string(s))
}
