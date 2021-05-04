package api

import (
	// "fmt"
	"server2/pkg/db"
)

type Zip struct {
	Id int `json:"id"`
	LocalCode int `json:"local_code"`
	OldZipcode int `json:"old_zipcode"`
	Zipcode int `json:"zipcode"`
	PrefKana string `json:"pref_kana"`
	CityKana string `json:"city_kana"`
	TownKana string `json:"town_kana"`
	Prefectures string `json:"prefectures"`
	City string `json:"city"`
	Town string `json:"town"`
	MultipleZipcode int `json:"multiple_zipcode"`
	Koaza int `json:"koaza"`
	Tyome int `json:"tyome"`
	MultipleTown int `json:"multiple_town"`
	UpdateFlag int `json:"update_flag"`
	UpdateWhy int `json:"update_why"`
}

func FetchIndex() []Zip {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM test1 limit 3")
	if err != nil {
		panic(err.Error())
	}

	scanArgs := make([]Zip, 0)
	for rows.Next() {
		var value Zip
		err = rows.Scan(&value.LocalCode, &value.OldZipcode, &value.Zipcode, &value.PrefKana, &value.CityKana, &value.TownKana, &value.Prefectures, &value.City, &value.Town, &value.MultipleZipcode, &value.Koaza, &value.Tyome, &value.MultipleTown, &value.UpdateFlag, &value.UpdateWhy, &value.Id)
		if err != nil {
			panic(err.Error())
		}
		scanArgs = append(scanArgs, value)
	}
	return scanArgs
}
