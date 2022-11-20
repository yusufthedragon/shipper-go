package location

import "github.com/denifrahman/shipper-go/shareds/structs"

// Areas struct contains response from API GetAreas.
type Areas struct {
	Status string `json:"status"`
	Data   struct {
		Title   string     `json:"title"`
		Content string     `json:"content"`
		Rows    []AreaRows `json:"rows"`
	} `json:"data"`
}

type AreasV3 struct {
	Metadata struct {
		Path           string `json:"path"`
		HttpStatusCode int    `json:"http_status_code"`
		HttpStatus     string `json:"http_status"`
		Timestamp      int    `json:"timestamp"`
	} `json:"metadata"`
	Data []struct {
		Suburb struct {
			Id   int     `json:"id"`
			Name string  `json:"name"`
			Lat  float64 `json:"lat"`
			Lng  float64 `json:"lng"`
		} `json:"suburb"`
		City struct {
			Id   int     `json:"id"`
			Name string  `json:"name"`
			Lat  float64 `json:"lat"`
			Lng  float64 `json:"lng"`
		} `json:"city"`
		Province struct {
			Id   int     `json:"id"`
			Name string  `json:"name"`
			Lat  float64 `json:"lat"`
			Lng  float64 `json:"lng"`
		} `json:"province"`
		Country struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"country"`
		Id       int     `json:"id"`
		Name     string  `json:"name"`
		Postcode string  `json:"postcode"`
		Lat      float64 `json:"lat"`
		Lng      float64 `json:"lng"`
	} `json:"data"`
	Pagination struct {
		CurrentPage     int `json:"current_page"`
		CurrentElements int `json:"current_elements"`
		TotalPages      int `json:"total_pages"`
		TotalElements   int `json:"total_elements"`
	} `json:"pagination"`
}

type AreaRows struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	PostCode string `json:"postcode"`
}

func (r AreasV3) ToAreas() Areas {
	var rows []AreaRows
	for _, val := range r.Data {
		rows = append(rows, AreaRows{
			ID:       val.Id,
			Name:     val.Name,
			Alias:    val.Name,
			PostCode: val.Postcode,
		})
	}
	return Areas{
		Status: r.Metadata.HttpStatus,
		Data: struct {
			Title   string     `json:"title"`
			Content string     `json:"content"`
			Rows    []AreaRows `json:"rows"`
		}{
			Title:   "Areas",
			Content: "",
			Rows:    rows,
		},
	}
}

// Cities struct contains response from API GetCities.

type Cities struct {
	Status string `json:"status"`
	Data   struct {
		Title   string     `json:"title"`
		Content string     `json:"content"`
		Rows    []CityRows `json:"rows"`
	} `json:"data"`
}

type CitiesV3 struct {
	Metadata struct {
		Path           string `json:"path"`
		HttpStatusCode int    `json:"http_status_code"`
		HttpStatus     string `json:"http_status"`
		Timestamp      int    `json:"timestamp"`
	} `json:"metadata"`
	Data []struct {
		Province struct {
			Id   int     `json:"id"`
			Name string  `json:"name"`
			Lat  float64 `json:"lat"`
			Lng  float64 `json:"lng"`
		} `json:"province"`
		Country struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"country"`
		Id   int     `json:"id"`
		Name string  `json:"name"`
		Lat  float64 `json:"lat"`
		Lng  float64 `json:"lng"`
	} `json:"data"`
	Pagination struct {
		CurrentPage     int `json:"current_page"`
		CurrentElements int `json:"current_elements"`
		TotalPages      int `json:"total_pages"`
		TotalElements   int `json:"total_elements"`
	} `json:"pagination"`
}

type CityRows struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ProvinceID   int    `json:"province_id"`
	ProvinceName string `json:"province_name"`
}

func (r CitiesV3) ToCities() Cities {
	var rows []CityRows
	for _, val := range r.Data {
		rows = append(rows, CityRows{
			ID:           val.Id,
			Name:         val.Name,
			ProvinceID:   val.Province.Id,
			ProvinceName: val.Province.Name,
		})
	}
	return Cities{
		Status: "",
		Data: struct {
			Title   string     `json:"title"`
			Content string     `json:"content"`
			Rows    []CityRows `json:"rows"`
		}{
			Title:   "Cities",
			Content: "",
			Rows:    rows,
		},
	}
}

// Countries struct contains response from API GetCountries.

type Countries struct {
	Status string `json:"status"`
	Data   struct {
		Title   string        `json:"title"`
		Content string        `json:"content"`
		Rows    []CountryRows `json:"rows"`
	} `json:"data"`
}

type CountryRows struct {
	CountryName string `json:"country_name"`
	CountryID   int    `json:"country_id"`
}

type CountriesV3 struct {
	Metadata struct {
		Path           string `json:"path"`
		HttpStatusCode int    `json:"http_status_code"`
		HttpStatus     string `json:"http_status"`
		Timestamp      int    `json:"timestamp"`
	} `json:"metadata"`
	Data []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"data"`
	Pagination struct {
		CurrentPage     int `json:"current_page"`
		CurrentElements int `json:"current_elements"`
		TotalPages      int `json:"total_pages"`
		TotalElements   int `json:"total_elements"`
	} `json:"pagination"`
}

func (r CountriesV3) ToCountries() Countries {
	var rows []CountryRows

	for _, val := range r.Data {
		rows = append(rows, CountryRows{
			CountryName: val.Name,
			CountryID:   val.Id,
		})
	}

	return Countries{
		Status: r.Metadata.HttpStatus,
		Data: struct {
			Title   string        `json:"title"`
			Content string        `json:"content"`
			Rows    []CountryRows `json:"rows"`
		}{
			Title:   "Countries",
			Content: "",
			Rows:    rows,
		},
	}
}

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Level struct {
	Id       int        `json:"id"`
	Level    int        `json:"level"`
	Type     string     `json:"type"`
	Name     string     `json:"name"`
	GeoCoord Coordinate `json:"geo_coord"`
	Postcode string     `json:"postcode"`
}

// Locations struct contains response from API SearchLocations.
type Locations struct {
	Status string `json:"status"`
	Data   struct {
		Title   string          `json:"title"`
		Content string          `json:"content"`
		Rows    []LocationsRows `json:"rows"`
	} `json:"data"`
}

type LocationsRows struct {
	AreaID       int    `json:"area_id"`
	AreaName     string `json:"area_name"`
	AreaAlias    string `json:"area_alias"`
	CityID       int    `json:"city_id"`
	CityName     string `json:"city_name"`
	Label        string `json:"label"`
	OrderList    int    `json:"order_list"`
	ProvinceID   int    `json:"province_id"`
	ProvinceName string `json:"province_name"`
	SuburbID     int    `json:"suburb_id"`
	SuburbName   string `json:"suburb_name"`
	SuburbAlias  string `json:"suburb_alias"`
	Value        string `json:"value"`
}

type LocationsV3 struct {
	Metadata structs.Metadata `json:"metadata"`
	Data     []struct {
		AdmLevel1   Level    `json:"adm_level_1"`
		AdmLevel2   Level    `json:"adm_level_2"`
		AdmLevel3   Level    `json:"adm_level_3"`
		AdmLevel4   Level    `json:"adm_level_4"`
		AdmLevel5   Level    `json:"adm_level_5"`
		AdmLevelCur Level    `json:"adm_level_cur"`
		DisplayTxt  string   `json:"display_txt"`
		Postcodes   []string `json:"postcodes"`
	} `json:"data"`
	Pagination struct {
		CurrentPage     int `json:"current_page"`
		CurrentElements int `json:"current_elements"`
		TotalPages      int `json:"total_pages"`
		TotalElements   int `json:"total_elements"`
	} `json:"pagination"`
}

func (r LocationsV3) ToLocations() Locations {
	var rows []LocationsRows

	for _, val := range r.Data {
		rows = append(rows, LocationsRows{
			AreaID:       val.AdmLevelCur.Id,
			AreaName:     val.AdmLevelCur.Name,
			AreaAlias:    val.AdmLevelCur.Name,
			CityID:       val.AdmLevel3.Id,
			CityName:     val.AdmLevel3.Name,
			Label:        val.DisplayTxt,
			OrderList:    0,
			ProvinceID:   val.AdmLevel2.Id,
			ProvinceName: val.AdmLevel2.Name,
			SuburbID:     val.AdmLevel4.Id,
			SuburbName:   val.AdmLevel4.Name,
			SuburbAlias:  val.AdmLevel4.Name,
			Value:        "",
		})
	}

	return Locations{
		Status: r.Metadata.HttpStatus,
		Data: struct {
			Title   string          `json:"title"`
			Content string          `json:"content"`
			Rows    []LocationsRows `json:"rows"`
		}{
			Title:   "Location",
			Content: "",
			Rows:    rows,
		},
	}
}

// Provinces struct contains response from API GetProvinces.
type Provinces struct {
	Status string `json:"status"`
	Data   struct {
		Title   string         `json:"title"`
		Content string         `json:"content"`
		Rows    []ProvinceRows `json:"rows"`
	} `json:"data"`
}

type ProvinceRows struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProvincesV3 struct {
	Metadata struct {
		Path           string `json:"path"`
		HttpStatusCode int    `json:"http_status_code"`
		HttpStatus     string `json:"http_status"`
		Timestamp      int    `json:"timestamp"`
	} `json:"metadata"`
	Data []struct {
		Id      int     `json:"id"`
		Name    string  `json:"name"`
		Lat     float64 `json:"lat"`
		Lng     float64 `json:"lng"`
		Country struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"country"`
	} `json:"data"`
	Pagination struct {
		CurrentPage     int `json:"current_page"`
		CurrentElements int `json:"current_elements"`
		TotalPages      int `json:"total_pages"`
		TotalElements   int `json:"total_elements"`
	} `json:"pagination"`
}

func (r ProvincesV3) ToProvince() Provinces {
	var rows []ProvinceRows
	for _, val := range r.Data {
		rows = append(rows, ProvinceRows{
			ID:   val.Id,
			Name: val.Name,
		})
	}
	return Provinces{
		Status: "",
		Data: struct {
			Title   string         `json:"title"`
			Content string         `json:"content"`
			Rows    []ProvinceRows `json:"rows"`
		}{
			Title:   "Provinces",
			Content: "",
			Rows:    rows,
		},
	}
}

// Suburbs struct contains response from API GetSuburbs.
type Suburbs struct {
	Status string `json:"status"`
	Data   struct {
		Title   string        `json:"title"`
		Content string        `json:"content"`
		Rows    []SuburbsRows `json:"rows"`
	} `json:"data"`
}

type SuburbsRows struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

type SuburbsV3 struct {
	Metadata struct {
		Path           string `json:"path"`
		HttpStatusCode int    `json:"http_status_code"`
		HttpStatus     string `json:"http_status"`
		Timestamp      int    `json:"timestamp"`
	} `json:"metadata"`
	Data []struct {
		Id   int     `json:"id"`
		Name string  `json:"name"`
		Lat  float64 `json:"lat"`
		Lng  float64 `json:"lng"`
		City struct {
			Id   int     `json:"id"`
			Name string  `json:"name"`
			Lat  float64 `json:"lat"`
			Lng  float64 `json:"lng"`
		} `json:"city"`
		Province struct {
			Id   int     `json:"id"`
			Name string  `json:"name"`
			Lat  float64 `json:"lat"`
			Lng  float64 `json:"lng"`
		} `json:"province"`
		Country struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"country"`
	} `json:"data"`
	Pagination struct {
		CurrentPage     int `json:"current_page"`
		CurrentElements int `json:"current_elements"`
		TotalPages      int `json:"total_pages"`
		TotalElements   int `json:"total_elements"`
	} `json:"pagination"`
}

func (r SuburbsV3) ToSuburbs() Suburbs {
	var rows []SuburbsRows

	for _, val := range r.Data {
		rows = append(rows, SuburbsRows{
			ID:   val.Id,
			Name: val.Name,
		})
	}
	return Suburbs{
		Status: "",
		Data: struct {
			Title   string        `json:"title"`
			Content string        `json:"content"`
			Rows    []SuburbsRows `json:"rows"`
		}{
			Title:   "Suburbs",
			Content: "",
			Rows:    rows,
		},
	}
}
