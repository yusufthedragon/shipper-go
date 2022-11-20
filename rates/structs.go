package rates

import (
	"github.com/denifrahman/shipper-go/shareds/structs"
	"strings"
)

// courier struct contains the detail of available courier from API GetRates.
type courier struct {
	Name             string  `json:"name"`
	LogoURL          string  `json:"logo_url"`
	RateID           int     `json:"rate_id"`
	ShowID           int     `json:"show_id"`
	RateName         string  `json:"rate_name"`
	StopOrigin       int     `json:"stop_origin"`
	StopDestination  int     `json:"stop_destination"`
	Weight           float64 `json:"weight"`
	VolumeWeight     float64 `json:"volumeWeight"`
	FinalWeight      float64 `json:"finalWeight"`
	ItemPrice        int     `json:"itemPrice"`
	FinalRate        int     `json:"finalRate"`
	InsuranceRate    int     `json:"insuranceRate"`
	Liability        int     `json:"liability"`
	Discount         int     `json:"discount"`
	DiscountValue    int     `json:"discount_value"`
	DiscountedPrice  int     `json:"discounted_price"`
	MinDay           int     `json:"min_day"`
	MaxDay           int     `json:"max_day"`
	MustUseInsurance bool    `json:"must_use_insurance"`
	Currency         string  `json:"currency"`
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

// DomesticRatesV3 struct contains response from API GetDomesticRates.
type DomesticRatesV3 struct {
	Metadata structs.Metadata `json:"metadata"`
	Data     struct {
		Origin struct {
			AreaId       int     `json:"area_id"`
			AreaName     string  `json:"area_name"`
			SuburbId     int     `json:"suburb_id"`
			SuburbName   string  `json:"suburb_name"`
			CityId       int     `json:"city_id"`
			CityName     string  `json:"city_name"`
			ProvinceId   int     `json:"province_id"`
			ProvinceName string  `json:"province_name"`
			CountryId    int     `json:"country_id"`
			CountryName  string  `json:"country_name"`
			Lat          float64 `json:"lat"`
			Lng          float64 `json:"lng"`
		} `json:"origin"`
		Destination struct {
			AreaId       int     `json:"area_id"`
			AreaName     string  `json:"area_name"`
			SuburbId     int     `json:"suburb_id"`
			SuburbName   string  `json:"suburb_name"`
			CityId       int     `json:"city_id"`
			CityName     string  `json:"city_name"`
			ProvinceId   int     `json:"province_id"`
			ProvinceName string  `json:"province_name"`
			CountryId    int     `json:"country_id"`
			CountryName  string  `json:"country_name"`
			Lat          float64 `json:"lat"`
			Lng          float64 `json:"lng"`
		} `json:"destination"`
		Pricings []struct {
			Logistic struct {
				Id          int    `json:"id"`
				Name        string `json:"name"`
				LogoUrl     string `json:"logo_url"`
				Code        string `json:"code"`
				CompanyName string `json:"company_name"`
			} `json:"logistic"`
			Rate struct {
				Id              int    `json:"id"`
				Name            string `json:"name"`
				Type            string `json:"type"`
				Description     string `json:"description"`
				FullDescription string `json:"full_description"`
				IsHubless       bool   `json:"is_hubless"`
			} `json:"rate"`
			Weight           float64 `json:"weight"`
			Volume           float64 `json:"volume"`
			VolumeWeight     float64 `json:"volume_weight"`
			FinalWeight      float64 `json:"final_weight"`
			MinDay           int     `json:"min_day"`
			MaxDay           int     `json:"max_day"`
			UnitPrice        int     `json:"unit_price"`
			TotalPrice       int     `json:"total_price"`
			Discount         float64 `json:"discount"`
			DiscountValue    int     `json:"discount_value"`
			DiscountedPrice  int     `json:"discounted_price"`
			InsuranceFee     int     `json:"insurance_fee"`
			MustUseInsurance bool    `json:"must_use_insurance"`
			LiabilityValue   int     `json:"liability_value"`
			FinalPrice       int     `json:"final_price"`
			Currency         string  `json:"currency"`
			InsuranceApplied bool    `json:"insurance_applied"`
			BasePrice        int     `json:"base_price"`
			SurchargeFee     int     `json:"surcharge_fee"`
		} `json:"pricings"`
	} `json:"data"`
	Pagination struct {
		CurrentPage     int      `json:"current_page"`
		CurrentElements int      `json:"current_elements"`
		TotalPages      int      `json:"total_pages"`
		TotalElements   int      `json:"total_elements"`
		SortBy          []string `json:"sort_by"`
	} `json:"pagination"`
}

// ToDomesticRates convert V3 to V1
func (r *DomesticRatesV3) ToDomesticRates() DomesticRates {

	response := DomesticRates{}

	for _, val := range r.Data.Pricings {
		rate := &courier{
			Name:             val.Logistic.Name,
			LogoURL:          val.Logistic.LogoUrl,
			RateID:           val.Rate.Id,
			ShowID:           val.Rate.Id,
			RateName:         val.Rate.Name,
			StopOrigin:       0,
			StopDestination:  0,
			Weight:           val.Weight,
			VolumeWeight:     val.VolumeWeight,
			FinalWeight:      val.FinalWeight,
			ItemPrice:        val.UnitPrice,
			FinalRate:        val.FinalPrice,
			InsuranceRate:    val.InsuranceFee,
			Liability:        val.LiabilityValue,
			Discount:         val.DiscountedPrice,
			DiscountValue:    val.DiscountValue,
			DiscountedPrice:  val.DiscountedPrice,
			MinDay:           val.MinDay,
			MaxDay:           val.MaxDay,
			MustUseInsurance: val.MustUseInsurance,
			Currency:         val.Currency,
		}

		switch strings.ToLower(val.Rate.Type) {
		case "express":
			response.Data.Rates.Logistic.Express = append(response.Data.Rates.Logistic.Express, *rate)
		case "regular":
			response.Data.Rates.Logistic.Regular = append(response.Data.Rates.Logistic.Regular, *rate)
		case "instant":
			response.Data.Rates.Logistic.Instant = append(response.Data.Rates.Logistic.Instant, *rate)
		case "same day":
			response.Data.Rates.Logistic.SameDay = append(response.Data.Rates.Logistic.SameDay, *rate)
		case "trucking":
			response.Data.Rates.Logistic.Trucking = append(response.Data.Rates.Logistic.Trucking, *rate)
		}
		response.Data.Title = "Domestic Rates"
		response.Data.DestinationArea = r.Data.Destination.AreaName
		response.Data.OriginArea = r.Data.Origin.AreaName
		response.Data.StatusCode = r.Metadata.HttpStatusCode
	}

	return response
}

// DomesticRatesParams struct contains request parameter for API GetDomesticRates.
type DomesticRatesParams struct {
	Origin         int     `json:"o" validate:"required"`
	Destination    int     `json:"d" validate:"required"`
	Length         float64 `json:"l" validate:"required"`
	Width          float64 `json:"w" validate:"required"`
	Height         float64 `json:"h" validate:"required"`
	WeightTotal    float64 `json:"wt" validate:"required"`
	Value          float64 `json:"v" validate:"required"`
	Type           int     `json:"type"`
	COD            bool    `json:"cod"`
	Order          int     `json:"order"`
	ForOrder       bool    `json:"for_order"`
	OriginLat      string  `json:"origin_lat"`
	OriginLng      string  `json:"origin_lng"`
	DestinationLat string  `json:"destination_lat"`
	DestinationLng string  `json:"destination_lng"`
}

// DomesticRatesParamsV3 struct contains request parameter for API GetDomesticRates.
type DomesticRatesParamsV3 struct {
	Destination struct {
		AreaId   int    `json:"area_id"`
		Lat      string `json:"lat"`
		Lng      string `json:"lng"`
		SuburbId int    `json:"suburb_id"`
	} `json:"destination"`
	Origin struct {
		AreaId   int    `json:"area_id"`
		Lat      string `json:"lat"`
		Lng      string `json:"lng"`
		SuburbId int    `json:"suburb_id"`
	} `json:"origin"`
	Cod       bool    `json:"cod"`
	ForOrder  bool    `json:"for_order"`
	Height    float64 `json:"height"`
	ItemValue float64 `json:"item_value"`
	Length    float64 `json:"length"`
	Weight    float64 `json:"weight"`
	Width     float64 `json:"width"`
}

// ToDomesticRatesParamsV3 convert v1 to v3
func (r *DomesticRatesParams) ToDomesticRatesParamsV3() DomesticRatesParamsV3 {
	return DomesticRatesParamsV3{
		Destination: struct {
			AreaId   int    `json:"area_id"`
			Lat      string `json:"lat"`
			Lng      string `json:"lng"`
			SuburbId int    `json:"suburb_id"`
		}{
			AreaId:   r.Destination,
			Lat:      r.DestinationLat,
			Lng:      r.DestinationLng,
			SuburbId: 0,
		},
		Origin: struct {
			AreaId   int    `json:"area_id"`
			Lat      string `json:"lat"`
			Lng      string `json:"lng"`
			SuburbId int    `json:"suburb_id"`
		}{
			AreaId:   r.Origin,
			Lat:      r.OriginLat,
			Lng:      r.OriginLng,
			SuburbId: 0,
		},
		Cod:       r.COD,
		ForOrder:  r.ForOrder,
		Height:    r.Height,
		ItemValue: r.Value,
		Length:    r.Length,
		Weight:    r.WeightTotal,
		Width:     r.Width,
	}
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
