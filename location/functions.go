package location

import (
	"context"
	"fmt"

	"github.com/denifrahman/shipper-go"
)

// GetAreas retrieves areas based on submitted suburb ID.
func GetAreas(areaIDS string) (Areas, error) {
	return GetAreasWithContext(context.Background(), areaIDS)
}

// GetAreasWithContext retrieves areas based on submitted suburb ID.
func GetAreasWithContext(ctx context.Context, areaIDS string) (Areas, error) {
	var endpoint = shipper.Conf.BaseURL + "/location/areas"
	var responseStruct = AreasV3{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
		AdditionalQuery: map[string]interface{}{
			"area_ids": fmt.Sprint(areaIDS),
		},
	}, &responseStruct)

	return responseStruct.ToAreas(), err
}

// GetCities retrieves cities based on province ID.
func GetCities(provinceID int) (Cities, error) {
	return GetCitiesWithContext(context.Background(), provinceID)
}

// GetCitiesWithContext retrieves cities based on province ID with context.
func GetCitiesWithContext(ctx context.Context, provinceID int) (Cities, error) {
	var endpoint = shipper.Conf.BaseURL + "/location/cities"
	var responseStruct = CitiesV3{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
		AdditionalQuery: map[string]interface{}{
			"province_id": fmt.Sprint(provinceID),
		},
	}, &responseStruct)

	return responseStruct.ToCities(), err
}

// GetCountries retrieves country data in a list.
func GetCountries() (Countries, error) {
	return GetCountriesWithContext(context.Background())
}

// GetCountriesWithContext retrieves country data in a list with context.
func GetCountriesWithContext(ctx context.Context) (Countries, error) {
	var endpoint = shipper.Conf.BaseURL + "/location/countries"
	var responseStruct = CountriesV3{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
	}, &responseStruct)

	return responseStruct.ToCountries(), err
}

// GetOriginCities retrieves provinces in which Shipper provides pickup service.
func GetOriginCities() (Cities, error) {
	return GetOriginCitiesWithContext(context.Background())
}

// GetOriginCitiesWithContext retrieves provinces in which Shipper provides pickup service with context.
func GetOriginCitiesWithContext(ctx context.Context) (Cities, error) {
	var endpoint = shipper.Conf.BaseURL + "/location/cities"
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
	var endpoint = shipper.Conf.BaseURL + "/location/provinces"
	var responseStruct = ProvincesV3{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:            ctx,
		HTTPMethod:     "GET",
		Endpoint:       endpoint,
		AdditionalBody: nil,
	}, &responseStruct)

	return responseStruct.ToProvince(), err
}

// GetSuburbs retrieves suburbs based on submitted city ID.
func GetSuburbs(cityID int) (Suburbs, error) {
	return GetSuburbsWithContext(context.Background(), cityID)
}

// GetSuburbsWithContext retrieves suburbs based on submitted city ID with context.
func GetSuburbsWithContext(ctx context.Context, cityID int) (Suburbs, error) {
	var endpoint = shipper.Conf.BaseURL + "/location/city/" + fmt.Sprint(cityID) + "/suburbs"
	var responseStruct = SuburbsV3{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
	}, &responseStruct)

	return responseStruct.ToSuburbs(), err
}

// SearchLocation retrieves every area, suburb, and city whose names include the submitted substring (including postcode).
func SearchLocation(keyword string) (Locations, error) {
	return SearchLocationWithContext(context.Background(), keyword)
}

// SearchLocationWithContext retrieves every area, suburb, and city whose names include the submitted substring (including postcode) with context.
func SearchLocationWithContext(ctx context.Context, keyword string) (Locations, error) {
	var endpoint = shipper.Conf.BaseURL + "/location"
	var responseStruct = LocationsV3{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		AdditionalQuery: map[string]interface{}{
			"keyword": keyword,
		},
		Endpoint: endpoint,
	}, &responseStruct)

	return responseStruct.ToLocations(), err
}
