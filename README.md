# Shipper API Go Client

Unofficial library for access [Shipper](http://shipper.id) API from applications written with Go.

- [Installation](#installation)
- [Usage](#usage)
  - [Set the API Key](#set-the-api-key)
  - [Set the Production Mode](#set-the-production-mode)
- [Available Methods and Examples](#available-methods-and-examples)
  - [Locations](#locations)
    - [Get Countries](#get-countries)
    - [Get Provinces](#get-provinces)
    - [Get Cities](#get-cities)
    - [Get Origin Cities](#get-origin-cities)
    - [Get Suburbs](#get-suburbs)
    - [Get Areas](#get-areas)
    - [Search Location](#search-location)
  - [Rates](#rates)
    - [Get Domestic Rates](#get-domestic-rates)
    - [Get International Rates](#get-international-rates)
  - [Orders](#orders)
    - [Create Domestic Order](#create-domestic-order)
    - [Create International Order](#create-international-order)
    - [Get Tracking ID](#get-tracking-id)
    - [Activate Order](#activate-order)
    - [Get Order Detail](#get-order-detail)
    - [Update Order](#update-order)
    - [Cancel Order](#cancel-order)
  - [Pickup Orders](#pickup-orders)
    - [Create Pickup Request](#create-pickup-request)
    - [Cancel Pickup Request](#cancel-pickup-request)
    - [Get Agents by Suburb](#get-agents-by-suburb)
  - [Furthermore](#furthermore)
    - [Get All Tracking Status](#get-all-tracking-status)
    - [Generate AWB Number](#generate-awb-number)
- [Test](#test)
- [Contributing](#contributing)

---

## Installation

Install shipper-go using Go Module by following command:

```bash
go get github.com/yusufthedragon/shipper-go
```

Then you import it by following code:

```go
import "github.com/yusufthedragon/shipper-go"
```

## Usage

### Set the API Key

Configure package with your account's API key obtained from Shipper.

```go
shipper.Conf.SetAPIKey("API_KEY")
```

### Set the Production Mode

When deploying your application to production, you may want to change API Endpoint as well by setting `SetProductionMode` to `true`.

```go
shipper.Conf.SetProductionMode(true)
// or chain it with SetAPIKey method
shipper.Conf.SetAPIKey("API_KEY").SetProductionMode(true)
```

## Available Methods and Examples

### Locations

You need to import `location` package first by following code:

```go
import "github.com/yusufthedragon/shipper-go/location"
```

#### Get Countries

Retrieve country data in a list.

```go
func GetCountries()
```

Usage example:

```go
var countries, err = location.GetCountries()

if err != nil {
    panic(err.Error())
}

fmt.Printf("List Countries: %+v\n", countries)
```

#### Get Provinces

Retrieve all provinces in Indonesia in a list.

```go
func GetProvinces()
```

Usage example:

```go
var provinces, err = location.GetProvinces()

if err != nil {
    panic(err.Error())
}

fmt.Printf("List Provinces: %+v\n", provinces)
```

#### Get Cities

Retrieve cities based on submitted province ID.

```go
func GetCities(provinceID int)
```

Usage example:

```go
var cities, err = location.GetCities(9)

if err != nil {
    panic(err.Error())
}

fmt.Printf("List Cities: %+v\n", cities)
```

#### Get Origin Cities

Retrieve provinces in which Shipper provides pickup service.

```go
func GetOriginCities()
```

Usage example:

```go
var originCities, err = location.GetOriginCities()

if err != nil {
    panic(err.Error())
}

fmt.Printf("List Origin Cities: %+v\n", originCities)
```

#### Get Suburbs

Retrieve suburbs based on submitted city ID.

```go
func GetSuburbs(cityID int)
```

Usage example:

```go
var suburbs, err = location.GetSuburbs(80)

if err != nil {
    panic(err.Error())
}

fmt.Printf("List Suburbs: %+v\n", suburbs)
```

#### Get Areas

Retrieve areas based on submitted suburb ID.

```go
func GetAreas(suburbID int)
```

Usage example:

```go
var areas, err = location.GetAreas(1330)

if err != nil {
    panic(err.Error())
}

fmt.Printf("List Areas: %+v\n", areas)
```

#### Search Location

Retrieve every area, suburb, and city whose names include the submitted substring (including postcode).

```go
func GetAreas(substring string)
```

Usage example:

```go
var searchLocation, err = location.SearchLocation("Sukmajaya")

if err != nil {
    panic(err.Error())
}

fmt.Printf("Searched Location: %+v\n", searchLocation)
```

### Rates

You need to import `rates` package first by following code:

```go
import "github.com/yusufthedragon/shipper-go/rates"
```

#### Get Domestic Rates

```go
func GetDomesticRates(params *DomesticRatesParams)
```

Usage example:

```go
var domesticRates, err = rates.GetDomesticRates(&rates.DomesticRatesParams{
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
    panic(err.Error())
}

fmt.Printf("Domestic Rates: %+v\n", domesticRates)
```

#### Get International Rates

```go
func GetInternationalRates(params *InternationalRatesParams)
```

Usage example:

```go
var internationalRates, err = rates.GetInternationalRates(&rates.InternationalRatesParams{
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
    panic(err.Error())
}

fmt.Printf("International Rates: %+v\n", internationalRates)
```

### Orders

You need to import `order` package first by following code:

```go
import "github.com/yusufthedragon/shipper-go/order"
```

#### Create Domestic Order

```go
func CreateDomesticOrder(params *DomesticOrderParams)
```

Usage example:

```go
var domesticOrder, err = order.CreateDomesticOrder(&order.DomesticOrderParams{
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
	ItemName: []ItemName{
		{
			Name:  "Baju",
			Qty:   1,
			Value: 100000,
		},
	},
	Contents:              "Barang mudah pecah",
	UseInsurance:          0,
	ExternalID:            "",
	PaymentType:           "cash",
	PackageType:           1,
	COD:                   0,
	OriginCoordinate:      "-6.308033944807303,106.73339847804874",
	DestinationCoordinate: "49.020733179213,12.114381752908",
})

if err != nil {
    panic(err.Error())
}

fmt.Printf("Domestic Order: %+v\n", domesticOrder)
```

#### Create International Order

```go
func CreateInternationalOrder(params *InternationalOrderParams)
```

Usage example:

```go
var internationalOrder, err = order.CreateInternationalOrder(&order.InternationalOrderParams{
    Origin:               4819,
    Destination:          12,
    Length:               10,
    Width:                10,
    Height:               10,
    WeightTotal:          0.4,
    Value:                150000,
    RateID:               327,
    ConsigneeName:        "Peoorang",
    ConsigneePhoneNumber: "089899878987",
    ConsignerName:        "Peorang",
    ConsignerPhoneNumber: "089891891818",
    OriginAddress:        "Mangga Dua Selatan",
    OriginDirection:      "Just follow the road",
    DestinationAddress:   "Orchard Road 101",
    DestinationDirection: "Lurus terus",
    DestinationArea:      "Singapore",
    DestinationSuburb:    "Singapore",
    DestinationCity:      "Singapore",
    DestinationProvince:  "Singapore",
    DestinationPostCode:  "111111",
    ItemName: []ItemName{
        {
            Name:  "Baju",
            Qty:   1,
            Value: 150000,
        },
    },
    Contents:     "Barang mudah pecah",
    UseInsurance: 0,
    PackageType:  2,
    PaymentType:  "cash",
    ExternalID:   ""
})

if err != nil {
    panic(err.Error())
}

fmt.Printf("International Order: %+v\n", internationalOrder)
```

#### Get Tracking ID

Retrieve tracking ID of the order with the provided ID.

```go
func GetTrackingID(orderID string)
```

Usage example:

```go
var trackingID, err = order.GetTrackingID("5f259130a172cf001222f533")

if err != nil {
    panic(err.Error())
}

fmt.Printf("Tracking ID: %+v\n", trackingID)
```

#### Activate Order

Activate/Deactivate an order. Such activation will initiate Shipper's pickup process.

```go
func ActivateOrder(orderID string, params *ActivateParams)
```

Usage example:

```go
var activatedOrder, err = order.ActivateOrder("5f259130a172cf001222f533", &order.ActivateParams{
    Active: 1,
})

if err != nil {
    panic(err.Error())
}

fmt.Printf("Activate Order: %+v\n", activatedOrder)
```

#### Get Order Detail

Retrieve an order's detail. Date format is UTC time.

```go
func GetOrderDetail(orderID string)
```

Usage example:

```go
var orderDetail, err = order.GetOrderDetail("5f259130a172cf001222f533")

if err != nil {
    panic(err.Error())
}

fmt.Printf("Detail Order: %+v\n", orderDetail)
```

#### Update Order

Update an order's package's weight and dimension.

```go
func UpdateOrder(orderID string, params *UpdateOrderParams)
```

Usage example:

```go
var updateOrder, err = order.UpdateOrder("5f259130a172cf001222f533", &order.UpdateOrderParams{
    Length:      5,
    Width:       5,
    Height:      5,
    WeightTotal: 1,
})

if err != nil {
    panic(err.Error())
}

fmt.Printf("Updated Order: %+v\n", updateOrder)
```

#### Cancel Order

Cancel an order.

```go
func CancelOrder(orderID string)
```

Usage example:

```go
var cancelOrder, err = order.CancelOrder("5f259130a172cf001222f533")

if err != nil {
    panic(err.Error())
}

fmt.Printf("Cancel Order: %+v\n", cancelOrder)
```

### Pickup Orders

You need to import `pickup` package first by following code:

```go
import "github.com/yusufthedragon/shipper-go/pickup"
```

#### Create Pickup Request

Assign agent and activate orders.

```go
func CreatePickup(params *CreatePickupParams)
```

Usage examples

```go
var createPickup, err = pickup.CreatePickup(&pickup.CreatePickupParams{
    OrderIDs:   []string{"5e45538"},
    DatePickup: "2020-08-11 10:30:00",
    AgentID:    1432,
})

if err != nil {
    panic(err.Error())
}

fmt.Printf("Created Pickup: %+v\n", createPickup)
```

#### Cancel Pickup Request

Cancel pickup request.

```go
func CancelPickup(params *CancelPickupParams)
```

Usage example:

```go
cancelPickup, err := pickup.CancelPickup(&pickup.CancelPickupParams{
    OrderIDs: []string{"5e45538"},
})

if err != nil {
    panic(err.Error())
}

fmt.Printf("Cancelled Pickup: %+v\n", cancelPickup)
```

#### Get Agents by Suburb

Get agent by origin suburb ID.

```go
func GetAgents(suburbID int)
```

Usage example:

```go
agents, err := pickup.GetAgents(1330)

if err != nil {
    panic(err.Error())
}

fmt.Printf("List Agents: %+v\n", agents)
```

### Furthermore

#### Get All Tracking Status

You need to import `tracking` package first by following code:

```go
import "github.com/yusufthedragon/shipper-go/tracking"
```

```go
func GetAllStatus()
```

Usage example:

```go
allTrackingStatus, err := tracking.GetAllStatus()

if err != nil {
    panic(err.Error())
}

fmt.Printf("List Tracking Status: %+v\n", allTrackingStatus)
```

#### Generate AWB Number

Generate AWB from related logistic, in case that AWB number in order is not generated yet when order sent.

You need to import `awb` package first by following code:

```go
import "github.com/yusufthedragon/shipper-go/awb"
```

```go
func Generate(params *GenerateParams)
```

Usage example:

```go
generatedAWB, err := awb.Generate(&awb.GenerateParams{
    OID: "5f259130a172cf001222f533",
})

if err != nil {
    panic(err.Error())
}

fmt.Printf("AWB Number: %+v\n", generatedAWB)
```

## Test

Rename the `.env.example` file to `.env` and set the API Key with the one you obtained from Shipper. After that, you can run the test by following command:

```bash
go test -v ./...
```

## Contributing

For any requests, bugs, or comments, please open an [issue](https://github.com/yusufthedragon/shipper-go/issues) or [submit a pull request](https://github.com/yusufthedragon/shipper-go/pulls).
