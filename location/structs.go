package location

// Areas struct contains response from API GetAreas.
type Areas struct {
	Status string `json:"status"`
	Data   struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Rows    []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Alias    string `json:"alias"`
			PostCode string `json:"postcode"`
		} `json:"rows"`
	} `json:"data"`
}

// Cities struct contains response from API GetCities.
type Cities struct {
	Status string `json:"status"`
	Data   struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Rows    []struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			ProvinceID   int    `json:"province_id"`
			ProvinceName string `json:"province_name"`
		} `json:"rows"`
	} `json:"data"`
}

// Countries struct contains response from API GetCountries.
type Countries struct {
	Status string `json:"status"`
	Data   struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Rows    []struct {
			CountryName string `json:"country_name"`
			CountryID   int    `json:"country_id"`
		} `json:"rows"`
	} `json:"data"`
}

// Locations struct contains response from API SearchLocations.
type Locations struct {
	Status string `json:"status"`
	Data   struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Rows    []struct {
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
		} `json:"rows"`
	} `json:"data"`
}

// Provinces struct contains response from API GetProvinces.
type Provinces struct {
	Status string `json:"status"`
	Data   struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Rows    []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"rows"`
	} `json:"data"`
}

// Suburbs struct contains response from API GetSuburbs.
type Suburbs struct {
	Status string `json:"status"`
	Data   struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Rows    []struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Alias string `json:"alias"`
		} `json:"rows"`
	} `json:"data"`
}
