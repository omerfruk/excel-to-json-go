package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize"
	"io/ioutil"
)

type yurt struct {
	Ulke     string
	Il       string
	Ilce     string
	YurtAdi  string
	Cinsiyet string
}

func main() {
	f, err := excelize.OpenFile("4-İL-İLÇE BAZINDA YURTLARIN KAPASİTESİ.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("KAPASİTE TABLOSU")
	if err != nil {
		fmt.Println(err)
		return
	}
	var ulke string
	var yurtlar []yurt
	for _, row := range rows {
		if len(row) == 11 {
			if row[2] != "" {
				if row[0] != "KIBRIS" {
					ulke = "TÜRKİYE"
				} else {
					ulke = "KIBRIS"
				}
				yurtlar = append(yurtlar, yurt{
					Ulke:     ulke,
					Il:       row[0],
					Ilce:     row[1],
					YurtAdi:  row[3],
					Cinsiyet: row[4],
				})
			}
		}
	}

	yurtlarJson, _ := json.MarshalIndent(&yurtlar, "", " ")

	ioutil.WriteFile("Yurtlar.json", yurtlarJson, 0644)

	fmt.Println(rows)
}
