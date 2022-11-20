package pickup

import (
	"encoding/json"
	"fmt"
	"github.com/denifrahman/shipper-go"
	"github.com/denifrahman/shipper-go/order"
	"strconv"
	"testing"
)

func init() {
	//var err = godotenv.Load("../.env")

	//if err != nil {
	//	panic("Error loading .env file.")
	//}

	productionMode, errParse := strconv.ParseBool("false")

	if errParse != nil {
		productionMode = false
	}

	shipper.Conf.SetAPIKey("tjHxw6FsElI7gB2mfSfqYBw6a6VDm6KXqYFyXYrNXKI8kOpqwLFyZlVxtggWnGdw").SetProductionMode(productionMode)
}

func TestCreatePickup(t *testing.T) {
	slots, err := GetPickupTimeSlots("Asia/Jakarta")
	if err != nil {
		return
	}
	fmt.Println(slots.Data.Data[0].StartTime)
	createdPickup, err := CreatePickup(&CreatePickupParams{
		OrderIDs:   []string{"22B4V84JQNQQZ"},
		DatePickup: slots.Data.Data[0].StartTime,
	})

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(createdPickup, "", "\t")

	fmt.Println("Created Pickup:")
	fmt.Println(string(s))
}

func TestGetPickupTimeSlots(t *testing.T) {
	timeslot, err := GetPickupTimeSlots("Asia/Jakarta")
	if err != nil {
		return
	}
	s, _ := json.MarshalIndent(timeslot, "", "\t")

	fmt.Println("Get Time Slots:")
	fmt.Println(string(s))
}

func TestCancelPickup(t *testing.T) {
	//domesticOrder, err := order.CreateDomesticOrder(&order.DomesticOrderParams{
	//	Origin:               12921,
	//	Destination:          4645,
	//	Length:               10,
	//	Width:                10,
	//	Height:               10,
	//	WeightTotal:          0.4,
	//	Value:                100000,
	//	RateID:               59,
	//	ConsigneeName:        "Peoorang",
	//	ConsigneePhoneNumber: "089899878987",
	//	ConsignerName:        "Peorang",
	//	ConsignerPhoneNumber: "089891891818",
	//	OriginAddress:        "Mangga Dua Selatan",
	//	OriginDirection:      "Just follow the road",
	//	DestinationAddress:   "Pasar Baru",
	//	DestinationDirection: "Lurus terus",
	//	ItemName: []order.ItemName{
	//		{
	//			Name:  "Baju",
	//			Qty:   1,
	//			Value: 100000,
	//		},
	//	},
	//	Contents:     "Barang mudah pecah",
	//	UseInsurance: true,
	//	ExternalID:   strconv.FormatInt(time.Now().UTC().UnixNano(), 10),
	//	PaymentType:  "cash",
	//	PackageType:  1,
	//	COD:          false,
	//})
	//
	//if err != nil {
	//	t.Error(err.Error())
	//}

	trackingID, err := order.GetOrderDetail("22BED8ZYKDWNE")
	fmt.Println(trackingID.Data.Order.Detail.PickUpCode)
	if err != nil {
		t.Error(err.Error())
	}

	cancelPickup, err := CancelPickup(&CancelPickupParams{
		PickupCode: trackingID.Data.Order.Detail.PickUpCode,
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
