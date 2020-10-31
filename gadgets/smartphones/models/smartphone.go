package models

//Smartphone model structure for smartphone
type Smartphone struct {
	Id            int64 // similar to long in Java - ID autoincremental in SQL
	Name          string
	Price         float32
	CountryOrigin string
	Os            string
}

// CreateSmartPhoneCMD command to create a new smartphone
type CreateSmartPhoneCMD struct {
	Name          string  `json:"name"`
	Price         float32 `json:"price"`
	CountryOrigin string  `json:"country_origin"`
	Os            string  `json:"os"`
}
