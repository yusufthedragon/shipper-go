package rates

// courier struct contains the detail of available courier from API GetRates.
type courier struct {
	Name                string  `json:"name"`
	LogoURL             string  `json:"logo_url"`
	RateID              int     `json:"rate_id"`
	ShowID              int     `json:"show_id"`
	RateName            string  `json:"rate_name"`
	StopOrigin          int     `json:"stop_origin"`
	StopDestination     int     `json:"stop_destination"`
	Weight              float64 `json:"weight"`
	VolumeWeight        float64 `json:"volumeWeight"`
	FinalWeight         float64 `json:"finalWeight"`
	ItemPrice           int     `json:"itemPrice"`
	FinalRate           int     `json:"finalRate"`
	InsuranceRate       int     `json:"insuranceRate"`
	CompulsoryInsurance int     `json:"compulsory_insurance"`
	Liability           int     `json:"liability"`
	Discount            int     `json:"discount"`
	MinDay              int     `json:"min_day"`
	MaxDay              int     `json:"max_day"`
	PickupAgent         int     `json:"pickup_agent"`
}

// DomesticRates struct contains response from API GetDomesticRates.
type DomesticRates struct {
	Status string `json:"status"`
	Data   struct {
		Title           string `json:"title"`
		Content         string `json:"content"`
		Rule            string `json:"rule"`
		OriginArea      string `json:"originArea"`
		DestinationArea string `json:"destinationArea"`
		Rates           struct {
			Logistic struct {
				Express  []courier `json:"express"`
				Regular  []courier `json:"regular"`
				Trucking []courier `json:"trucking"`
				SameDay  []courier `json:"same day"`
				Instant  []courier `json:"instant"`
			} `json:"logistic"`
		} `json:"rates"`
		StatusCode int `json:"statusCode"`
	} `json:"data"`
}

// DomesticRatesParams struct contains request parameter for API GetDomesticRates.
type DomesticRatesParams struct {
	Origin                int     `json:"o" validate:"required"`
	Destination           int     `json:"d" validate:"required"`
	Length                float64 `json:"l" validate:"required"`
	Width                 float64 `json:"w" validate:"required"`
	Height                float64 `json:"h" validate:"required"`
	WeightTotal           float64 `json:"wt" validate:"required"`
	Value                 float64 `json:"v" validate:"required"`
	Type                  int     `json:"type"`
	COD                   int     `json:"cod"`
	Order                 int     `json:"order"`
	OriginCoordinate      string  `json:"originCoord"`
	DestinationCoordinate string  `json:"destinationCoord"`
}

// InternationalRates struct contains response from API GetInternationalRates.
type InternationalRates struct {
	Status string `json:"status"`
	Data   struct {
		Title           string `json:"title"`
		Content         string `json:"content"`
		Rule            string `json:"rule"`
		OriginArea      string `json:"originArea"`
		DestinationArea string `json:"destinationArea"`
		Rates           struct {
			Logistic struct {
				International []courier `json:"international"`
			} `json:"logistic"`
		} `json:"rates"`
		StatusCode int `json:"statusCode"`
	} `json:"data"`
}

// InternationalRatesParams struct contains request parameter for API GetInternationalRates.
type InternationalRatesParams struct {
	Origin      int     `json:"o" validate:"required"`
	Destination int     `json:"d" validate:"required"`
	Length      float64 `json:"l" validate:"required"`
	Width       float64 `json:"w" validate:"required"`
	Height      float64 `json:"h" validate:"required"`
	WeightTotal float64 `json:"wt" validate:"required"`
	Value       float64 `json:"v" validate:"required"`
	Type        int     `json:"type"`
	Order       int     `json:"order"`
}
