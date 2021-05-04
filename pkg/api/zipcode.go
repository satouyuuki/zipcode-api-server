package api

import (
	"github.com/satouyuuki/zipcode-api-server/pkg/db"
)

type Zipcode struct {
	local_code int
	old_zipcode int
	zipcode int
	pref_kana string
	city_kana string
	town_kana string
	prefectures string
	city string
	town string
	multiple_zipcode int
	koaza int
	tyome int
	multiple_town int
	update_flag int
	update_why int
}

// CREATE TABLE `test1` (
//   `local_code` int(11) NOT NULL,
//   `old_zipcode` int(5) NOT NULL,
//   `zipcode` int(7) NOT NULL,
//   `pref_kana` varchar(50) not null,
//   `city_kana` varchar(50) not null,
//   `town_kana` varchar(50) not null,
//   `prefectures` varchar(50) not null,
//   `city` varchar(50) not null,
//   `town` varchar(50) not null,
//   `multiple_zipcode` tinyint(1) not null,
//   `koaza` tinyint(1) not null,
//   `tyome` tinyint(1) not null,
//   `multiple_town` tinyint(1) not null,
//   `update_flag` tinyint(1) not null,
//   `update_why` tinyint(1) not null
// );


func FetchIndex() []Zipcode {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM test1 limit 3")
	if err !=å nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	} 

	values := make([]sql.RawBytes, len(columns))

	// rows.Scanには引数に`[]interface{}`が必要

	scanArgs := make([]interface{}, len(values))
	// fmt.Println("scanArgs = ", scanArgs)
	for i := range values {
		// scanArgs[i]にvalues[i]のポインタを設定
		scanArgs[i] = &values[i]
	}

	//Opening型のスライスに格納します
  returnValues := make([]Zipcode, len(values))
	// fmt.Println("values = ", values)
	// 行セットに対して繰り返し処理
	for rows.Next() {
		// columnを変数へ読み込む
		// fmt.Println("values2 = ", values)
		// 各行において、row.Scan()でcolumnを変数へ読み込んでいる
		err = rows.Scan(scanArgs...)
		fmt.Println("values3 = ", string(values[3]))
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			// fmt.Println(columns[i], ": ", value)
			returnValues = append(returnValues, value)
		}
		// fmt.Println("-------------------------");
		return returnValues
	}
}