package location

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

func TestGetCountries(t *testing.T) {
	countries, err := GetCountries()

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(countries, "", "\t")

	fmt.Println("Get Countries:")
	fmt.Println(string(s))
}

func TestGetProvinces(t *testing.T) {
	provinces, err := GetProvinces()

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(provinces, "", "\t")

	fmt.Println("Get Provinces:")
	fmt.Println(string(s))
}

func TestGetCities(t *testing.T) {
	cities, err := GetCities(9)

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(cities, "", "\t")

	fmt.Println("Get Cities:")
	fmt.Println(string(s))
}

func TestGetOriginCities(t *testing.T) {
	cities, err := GetOriginCities()

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(cities, "", "\t")

	fmt.Println("Get Origin Cities:")
	fmt.Println(string(s))
}

func TestGetSuburbs(t *testing.T) {
	suburbs, err := GetSuburbs(80)

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(suburbs, "", "\t")

	fmt.Println("Get Suburbs:")
	fmt.Println(string(s))
}

func TestGetAreas(t *testing.T) {
	areas, err := GetAreas(1330)

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(areas, "", "\t")

	fmt.Println("Get Areas:")
	fmt.Println(string(s))
}

func TestSearchLocations(t *testing.T) {
	locations, err := SearchLocation("Tirtajaya")

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(locations, "", "\t")

	fmt.Println("Search Locations:")
	fmt.Println(string(s))
}
