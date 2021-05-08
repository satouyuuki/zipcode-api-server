package api

import (
	"fmt"
	"server2/pkg/db"
)

type Zip struct {
	Zipcode int `json:"zipcode"`
	PrefKana string `json:"pref_kana"`
	CityKana string `json:"city_kana"`
	TownKana string `json:"town_kana"`
	Prefectures string `json:"prefectures"`
	City string `json:"city"`
	Town string `json:"town"`
}

func FetchIndex() []Zip {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT zipcode, pref_kana, city_kana, town_kana, prefectures, city, town FROM test1 limit 3")
	if err != nil {
		panic(err.Error())
	}

	scanArgs := make([]Zip, 0)
	for rows.Next() {
		var value Zip
		err = rows.Scan(&value.Zipcode, &value.PrefKana, &value.CityKana, &value.TownKana, &value.Prefectures, &value.City, &value.Town)
		if err != nil {
			panic(err.Error())
		}
		scanArgs = append(scanArgs, value)
	}
	return scanArgs
}

// 郵便番号リソース /{zipcode}
func FetchByKey(zipcode string) []Zip {
	db := db.Connect()
	defer db.Close()
	rows, err := db.Query("SELECT zipcode, pref_kana, city_kana, town_kana, prefectures, city, town FROM test1 where zipcode = ?", zipcode)
	if err != nil {
		panic(err.Error())
	}

	scanArgs := make([]Zip, 0)
	for rows.Next() {
		var value Zip
		err = rows.Scan(&value.Zipcode, &value.PrefKana, &value.CityKana, &value.TownKana, &value.Prefectures, &value.City, &value.Town)
		if err != nil {
			panic(err.Error())
		}
		scanArgs = append(scanArgs, value)
	}
	return scanArgs
}

// 検索結果リソース /search/?q={query}
func Search(query string) []Zip {
	db := db.Connect()
	defer db.Close()
	// queryParam := query + '%'
	
	rows, err := db.Query("SELECT zipcode, pref_kana, city_kana, town_kana, prefectures, city, town FROM test1 where zipcode like ?", "'" + query + "%'")
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	// fmt.Println(rows)

	scanArgs := make([]Zip, 0)
	for rows.Next() {
		var value Zip
		err = rows.Scan(&value.Zipcode, &value.PrefKana, &value.CityKana, &value.TownKana, &value.Prefectures, &value.City, &value.Town)
		if err != nil {
			panic(err.Error())
		}
		scanArgs = append(scanArgs, value)
		fmt.Println(scanArgs)
	}
	return scanArgs
}
