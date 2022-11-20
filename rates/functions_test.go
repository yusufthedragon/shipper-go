package rates

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/denifrahman/shipper-go"
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

func TestGetDomesticRates(t *testing.T) {
	domesticRates, err := GetDomesticRates(&DomesticRatesParams{
		Origin:         30151,
		Destination:    30019,
		Length:         10,
		Width:          10,
		Height:         10,
		WeightTotal:    0.4,
		Value:          70000,
		Type:           1,
		COD:            false,
		Order:          0,
		ForOrder:       true,
		OriginLat:      "-7.2593277",
		OriginLng:      "112.767352",
		DestinationLat: "-7.3260971",
		DestinationLng: "112.8135304",
	})

	if err != nil {
		fmt.Println(err)
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
