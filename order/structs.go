package order

// Activate struct contains response from API ActivateOrder.
type Activate struct {
	Status string `json:"status"`
	Data   struct {
		Message    string `json:"message"`
		StatusCode int    `json:"statusCode"`
	} `json:"data"`
}

// ActivateParams struct contains request parameter for API ActivateOrder.
type ActivateParams struct {
	Active  int `json:"active" validate:"required"`
	AgentID int `json:"agentId"`
}

// CancelledOrder struct contains response from API CancelOrder.
type CancelledOrder struct {
	Status string `json:"status"`
	Data   struct {
		Title      string `json:"title"`
		Message    string `json:"message"`
		StatusCode int    `json:"statusCode"`
	} `json:"data"`
}

// DomesticOrder struct contains response from API CreateDomesticOrder.
type DomesticOrder struct {
	Status string `json:"status"`
	Data   struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		StatusCode int    `json:"statusCode"`
		ID         string `json:"id"`
	} `json:"data"`
}

// DomesticOrderParams struct contains request parameter for API CreateDomesticOrder.
type DomesticOrderParams struct {
	Origin                int        `json:"o" validate:"required"`
	Destination           int        `json:"d" validate:"required"`
	Length                float64    `json:"l" validate:"required"`
	Width                 float64    `json:"w" validate:"required"`
	Height                float64    `json:"h" validate:"required"`
	WeightTotal           float64    `json:"wt" validate:"required"`
	Value                 float64    `json:"v" validate:"required"`
	RateID                int        `json:"rateID" validate:"required"`
	ConsigneeName         string     `json:"consigneeName" validate:"required"`
	ConsigneePhoneNumber  string     `json:"consigneePhoneNumber" validate:"required"`
	ConsignerName         string     `json:"consignerName"`
	ConsignerPhoneNumber  string     `json:"consignerPhoneNumber"`
	OriginAddress         string     `json:"originAddress" validate:"required"`
	OriginDirection       string     `json:"originDirection" validate:"required"`
	DestinationAddress    string     `json:"destinationAddress" validate:"required"`
	DestinationDirection  string     `json:"destinationDirection" validate:"required"`
	ItemName              []ItemName `json:"itemName" validate:"required"`
	Contents              string     `json:"contents" validate:"required"`
	UseInsurance          int        `json:"useInsurance"`
	ExternalID            string     `json:"externalID"`
	PaymentType           string     `json:"paymentType"`
	PackageType           int        `json:"packageType" validate:"required"`
	COD                   int        `json:"cod"`
	OriginCoordinate      *string    `json:"originCoord"`
	DestinationCoordinate *string    `json:"destinationCoord"`
}

// InternationalOrder struct contains response from API CreateInternationalOrder.
type InternationalOrder struct {
	Status string `json:"status"`
	Data   struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		StatusCode int    `json:"statusCode"`
		ID         string `json:"id"`
	} `json:"data"`
}

// InternationalOrderParams struct contains request parameter for API CreateInternationalOrder.
type InternationalOrderParams struct {
	Origin               int        `json:"o" validate:"required"`
	Destination          int        `json:"d" validate:"required"`
	Length               float64    `json:"l" validate:"required"`
	Width                float64    `json:"w" validate:"required"`
	Height               float64    `json:"h" validate:"required"`
	WeightTotal          float64    `json:"wt" validate:"required"`
	Value                float64    `json:"v" validate:"required"`
	RateID               int        `json:"rateID" validate:"required"`
	ConsigneeName        string     `json:"consigneeName" validate:"required"`
	ConsigneePhoneNumber string     `json:"consigneePhoneNumber" validate:"required"`
	ConsignerName        string     `json:"consignerName"`
	ConsignerPhoneNumber string     `json:"consignerPhoneNumber"`
	OriginAddress        string     `json:"originAddress" validate:"required"`
	OriginDirection      string     `json:"originDirection" validate:"required"`
	DestinationAddress   string     `json:"destinationAddress" validate:"required"`
	DestinationDirection string     `json:"destinationDirection" validate:"required"`
	DestinationArea      string     `json:"destinationArea"`
	DestinationSuburb    string     `json:"destinationSuburb"`
	DestinationCity      string     `json:"destinationCity"`
	DestinationProvince  string     `json:"destinationProvince"`
	DestinationPostCode  string     `json:"destinationPostcode"`
	ItemName             []ItemName `json:"itemName" validate:"required"`
	Contents             string     `json:"contents" validate:"required"`
	UseInsurance         int        `json:"useInsurance"`
	ExternalID           string     `json:"externalID"`
	PaymentType          string     `json:"paymentType"`
	PackageType          int        `json:"packageType" validate:"required"`
}

// DetailOrder struct contains response from API GetOrderDetail.
type DetailOrder struct {
	Status string `json:"status"`
	Data   struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Order   struct {
			Tracking []tracking `json:"tracking"`
			Detail   struct {
				ID            string `json:"_id"`
				TrackingID    string `json:"id"`
				GroupID       int    `json:"groupID"`
				SpecialID     string `json:"specialID"`
				ExternalID    string `json:"externalID"`
				LabelChecksum string `json:"labelChecksum"`
				Consigner     struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					PhoneNumber string `json:"phoneNumber"`
				}
				Consignee struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					PhoneNumber string `json:"phoneNumber"`
				}
				BatchID        int            `json:"batchID"`
				AWBNumber      string         `json:"awbNumber"`
				StickerNumber  string         `json:"stickerNumber"`
				ShipmentStatus shipmentStatus `json:"shipmentStatus"`
				InternalStatus shipmentStatus `json:"internalStatus"`
				ExternalStatus shipmentStatus `json:"externalStatus"`
				Origin         origin         `json:"origin"`
				Destination    origin         `json:"destination"`
				Package        struct {
					ItemType  string `json:"itemType"`
					Contents  string `json:"contents"`
					Price     item   `json:"price"`
					ItemName  int    `json:"itemName"`
					Dimension struct {
						Length item `json:"length"`
						Width  item `json:"width"`
						Height item `json:"height"`
					} `json:"dimension"`
					Weight        item   `json:"weight"`
					CubicalWeight item   `json:"cubicalWeight"`
					Fragile       int    `json:"fragile"`
					Type          int    `json:"type"`
					IsConfirmed   int    `json:"isConfirmed"`
					PictureURL    string `json:"pictureURL"`
					Details       []struct {
						ItemID    int    `json:"itemID"`
						ItemName  string `json:"itemName"`
						ItemPrice int    `json:"itemPrice"`
						ItemQTY   int    `json:"itemQTY"`
					} `json:"details"`
				} `json:"package"`
				Driver struct {
					ID                 int    `json:"id"`
					Name               string `json:"name"`
					PhoneNumber        string `json:"phoneNumber"`
					VehicleType        string `json:"vehicleType"`
					VehicleNumber      string `json:"vehicleNumber"`
					IsPaymentCollected int    `json:"isPaymentCollected"`
					Feedback           struct {
						Score   int    `json:"score"`
						Comment string `json:"comment"`
					} `json:"feedback"`
					AgentID     int    `json:"agentID"`
					AgentName   string `json:"agentName"`
					AgentCityID int    `json:"agentCityID"`
				} `json:"driver"`
				Courier struct {
					ID           int    `json:"id"`
					Name         string `json:"name"`
					RateID       int    `json:"rate_id"`
					RateName     string `json:"rate_name"`
					ShipmentType int    `json:"shipmentType"`
					ActualID     int    `json:"actualID"`
					MinDay       int    `json:"min_day"`
					MaxDay       int    `json:"max_day"`
					Rate         item   `json:"rate"`
					ActualRate   struct {
						ID int `json:"id"`
						item
					} `json:"actualRate"`
				} `json:"courier"`
				Rates struct {
					Shipment        item `json:"shipment"`
					PaidShipment    item `json:"paidShipment"`
					ActualShipment  item `json:"actualShipment"`
					Insurance       item `json:"insurance"`
					PaidInsurance   item `json:"paidInsurance"`
					ActualInsurance item `json:"actualInsurance"`
					EscrowCost      item `json:"escrowCost"`
					FulfillmentCost item `json:"fulfillmentCost"`
					ItemPrice       item `json:"itemPrice"`
					Discount        item `json:"discount"`
				} `json:"rates"`
				UseInsurance    int    `json:"useInsurance"`
				IsLabelPrinted  int    `json:"isLabelPrinted"`
				PaymentType     string `json:"paymentType"`
				Source          string `json:"source"`
				IsActive        int    `json:"isActive"`
				IsAutoTrack     int    `json:"isAutoTrack"`
				IsCustomAWB     int    `json:"isCustomAWB"`
				ReadyTime       string `json:"readyTime"`
				PickUpTime      string `json:"pickUpTime"`
				CreationDate    string `json:"creationDate"`
				ActiveDate      string `json:"activeDate"`
				LastUpdatedDate string `json:"lastUpdatedDate"`
				CreatedBy       int    `json:"createdBy"`
				ShipmentArea    string `json:"shipmentArea"`
				COD             struct {
					Order        int `json:"order"`
					FreeDelivery int `json:"freeDelivery"`
				} `json:"cod"`
				Voucher          string `json:"voucher"`
				Whitelabel       string `json:"whitelabel"`
				JOBNumber        string `json:"JOBNumber"`
				Domain           string `json:"domain"`
				IsAutomateOrder  int    `json:"isAutomateOrder"`
				Channel          string `json:"channel"`
				IsHubless        int    `json:"isHubless"`
				ZonaID           int    `json:"zonaID"`
				IsJobMarketplace int    `json:"isJobMarketplace"`
				RequestID        string `json:"requestID"`
				PartitionKey     int    `json:"partitionKey"`
				HandlerIDs       []int  `json:"handler_ids"`
			} `json:"detail"`
		} `json:"order"`
		StatusCode int `json:"statusCode"`
	} `json:"data"`
}

// UpdatedOrder struct contains response from API UpdateOrder.
type UpdatedOrder struct {
	Status string `json:"status"`
	Data   struct {
		Title           string `json:"title"`
		Content         string `json:"content"`
		CorrectedFields struct {
			Weight              float64 `json:"weight"`
			VolumeWeight        float64 `json:"volumeWeight"`
			Length              float64 `json:"length"`
			Height              float64 `json:"height"`
			Width               float64 `json:"width"`
			CompulsoryInsurance int     `json:"compulsoryInsurance"`
			Insurance           int     `json:"insurance"`
			FinalRate           int     `json:"finalRate"`
		} `json:"correctedFields"`
		StatusCode int `json:"statusCode"`
	} `json:"data"`
}

// UpdateOrderParams struct contains request parameter for API UpdateOrder.
type UpdateOrderParams struct {
	Length      float64 `json:"l" validate:"required"`
	Width       float64 `json:"w" validate:"required"`
	Height      float64 `json:"h" validate:"required"`
	WeightTotal float64 `json:"wt" validate:"required"`
}

// ItemName struct contains list of items for API CreateOrder.
type ItemName struct {
	Name  string `json:"name"`
	Qty   int    `json:"qty"`
	Value int    `json:"value"`
}

// item struct contains item description for API GetOrderDetail.
type item struct {
	Value float64 `json:"value"`
	UOM   string  `json:"UoM"`
}

// origin struct contains data of origin & destination for API GetOrderDetail.
type origin struct {
	ID           int    `json:"id"`
	Address      string `json:"address"`
	Direction    string `json:"direction"`
	PostCode     string `json:"postcode"`
	StopID       int    `json:"stopID"`
	AreaID       int    `json:"areaID"`
	SuburbID     int    `json:"suburbID"`
	AreaName     string `json:"areaName"`
	SuburbName   string `json:"suburbName"`
	CityID       int    `json:"cityID"`
	CityName     string `json:"cityName"`
	ProvinceID   int    `json:"provinceID"`
	ProvinceName string `json:"provinceName"`
	CountryID    int    `json:"countryID"`
	CountryName  string `json:"countryName"`
	Latitude     string `json:"lat"`
	Longitude    string `json:"lng"`
}

// shipmentStatus struct contains description of shipment status for API GetOrderDetail.
type shipmentStatus struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StatusCode  int    `json:"statusCode"`
	UpdatedBy   string `json:"updatedBy"`
	UpdateDate  string `json:"updateDate"`
	IsCanceled  int    `json:"isCanceled"`
}

// tracking struct contains list of tracking data for API GetOrderDetail.
type tracking struct {
	ID             string        `json:"id"`
	OrderID        string        `json:"orderID"`
	UniqueID       string        `json:"uniqueID"`
	TrackStatus    trackStatus   `json:"trackStatus"`
	InternalStatus trackStatus   `json:"internalStatus"`
	ExternalStatus trackStatus   `json:"externalStatus"`
	LogisticStatus []trackStatus `json:"logisticStatus"`
	CreatedBy      string        `json:"createdBy"`
	CreatedDate    string        `json:"createdDate"`
}

// trackStatus struct contains data of tracking status for tracking struct.
type trackStatus struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
