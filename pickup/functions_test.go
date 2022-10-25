package pickup

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

func TestCreatePickup(t *testing.T) {
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
		ExternalID:            strconv.FormatInt(time.Now().UTC().UnixNano(), 10),
		PaymentType:           "cash",
		PackageType:           1,
		COD:                   0,
		OriginCoordinate:      "-6.308033944807303,106.73339847804874",
		DestinationCoordinate: "49.020733179213,12.114381752908",
	})

	if err != nil {
		t.Error(err.Error())
	}

	trackingID, err := order.GetTrackingID(domesticOrder.Data.ID)

	if err != nil {
		t.Error(err.Error())
	}

	createdPickup, err := CreatePickup(&CreatePickupParams{
		OrderIDs:   []string{trackingID.Data.ID},
		DatePickup: "2021-01-23 11:00:00",
		AgentID:    3538,
	})

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(createdPickup, "", "\t")

	fmt.Println("Created Pickup:")
	fmt.Println(string(s))
}

func TestCancelPickup(t *testing.T) {
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
		ExternalID:            strconv.FormatInt(time.Now().UTC().UnixNano(), 10),
		PaymentType:           "cash",
		PackageType:           1,
		COD:                   0,
		OriginCoordinate:      "-6.308033944807303,106.73339847804874",
		DestinationCoordinate: "49.020733179213,12.114381752908",
	})

	if err != nil {
		t.Error(err.Error())
	}

	trackingID, err := order.GetTrackingID(domesticOrder.Data.ID)

	if err != nil {
		t.Error(err.Error())
	}

	cancelPickup, err := CancelPickup(&CancelPickupParams{
		OrderIDs: []string{trackingID.Data.ID},
	})

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(cancelPickup, "", "\t")

	fmt.Println("Cancel Pickup:")
	fmt.Println(string(s))
}

func TestGetAgents(t *testing.T) {
	agents, err := GetAgents(1330)

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(agents, "", "\t")

	fmt.Println("Agents:")
	fmt.Println(string(s))
}
