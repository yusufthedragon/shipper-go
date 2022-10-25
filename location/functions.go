package location

import (
	"context"
	"fmt"

	"github.com/denifrahman/shipper-go"
)

// GetAreas retrieves areas based on submitted suburb ID.
func GetAreas(suburbID int) (Areas, error) {
	return GetAreasWithContext(context.Background(), suburbID)
}

// GetAreasWithContext retrieves areas based on submitted suburb ID.
func GetAreasWithContext(ctx context.Context, suburbID int) (Areas, error) {
	var endpoint = shipper.Conf.BaseURL + "/areas"
	var responseStruct = Areas{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
		AdditionalQuery: map[string]interface{}{
			"suburb": fmt.Sprint(suburbID),
		},
	}, &responseStruct)

	return responseStruct, err
}

// GetCities retrieves cities based on province ID.
func GetCities(provinceID int) (Cities, error) {
	return GetCitiesWithContext(context.Background(), provinceID)
}

// GetCitiesWithContext retrieves cities based on province ID with context.
func GetCitiesWithContext(ctx context.Context, provinceID int) (Cities, error) {
	var endpoint = shipper.Conf.BaseURL + "/cities"
	var responseStruct = Cities{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
		AdditionalQuery: map[string]interface{}{
			"province": fmt.Sprint(provinceID),
		},
	}, &responseStruct)

	return responseStruct, err
}

// GetCountries retrieves country data in a list.
func GetCountries() (Countries, error) {
	return GetCountriesWithContext(context.Background())
}

// GetCountriesWithContext retrieves country data in a list with context.
func GetCountriesWithContext(ctx context.Context) (Countries, error) {
	var endpoint = shipper.Conf.BaseURL + "/countries"
	var responseStruct = Countries{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
	}, &responseStruct)

	return responseStruct, err
}

// GetOriginCities retrieves provinces in which Shipper provides pickup service.
func GetOriginCities() (Cities, error) {
	return GetOriginCitiesWithContext(context.Background())
}

// GetOriginCitiesWithContext retrieves provinces in which Shipper provides pickup service with context.
func GetOriginCitiesWithContext(ctx context.Context) (Cities, error) {
	var endpoint = shipper.Conf.BaseURL + "/cities"
	var responseStruct = Cities{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
		AdditionalQuery: map[string]interface{}{
			"origin": "all",
		},
	}, &responseStruct)

	return responseStruct, err
}

// GetProvinces retrieves all provinces in Indonesia in a list.
func GetProvinces() (Provinces, error) {
	return GetProvincesWithContext(context.Background())
}

// GetProvincesWithContext retrieves all provinces in Indonesia in a list with context.
func GetProvincesWithContext(ctx context.Context) (Provinces, error) {
	var endpoint = shipper.Conf.BaseURL + "/provinces"
	var responseStruct = Provinces{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:            ctx,
		HTTPMethod:     "GET",
		Endpoint:       endpoint,
		AdditionalBody: nil,
	}, &responseStruct)

	return responseStruct, err
}

// GetSuburbs retrieves suburbs based on submitted city ID.
func GetSuburbs(cityID int) (Suburbs, error) {
	return GetSuburbsWithContext(context.Background(), cityID)
}

// GetSuburbsWithContext retrieves suburbs based on submitted city ID with context.
func GetSuburbsWithContext(ctx context.Context, cityID int) (Suburbs, error) {
	var endpoint = shipper.Conf.BaseURL + "/suburbs"
	var responseStruct = Suburbs{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
		AdditionalQuery: map[string]interface{}{
			"city": fmt.Sprint(cityID),
		},
	}, &responseStruct)

	return responseStruct, err
}

// SearchLocation retrieves every area, suburb, and city whose names include the submitted substring (including postcode).
func SearchLocation(substring string) (Locations, error) {
	return SearchLocationWithContext(context.Background(), substring)
}

// SearchLocationWithContext retrieves every area, suburb, and city whose names include the submitted substring (including postcode) with context.
func SearchLocationWithContext(ctx context.Context, substring string) (Locations, error) {
	var endpoint = shipper.Conf.BaseURL + "/details/" + substring
	var responseStruct = Locations{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
	}, &responseStruct)

	return responseStruct, err
}
