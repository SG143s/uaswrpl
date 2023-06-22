package sqlop

import (
	"database/sql"

	str "github.com/SG143s/uaswrpl/structsn"
	_ "github.com/go-sql-driver/mysql"
)

var db sql.DB

func Sqinit() {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/utsdb")

	if err != nil {
		panic(err)
	}
	db = *dbt
	dbt.Close()
}

func Sqclose() {
	db.Close()
}

func UpdName(name string, nname string) {
	_, err := db.Query("CALL changename(?, ?)", name, nname)
	if err != nil {
		panic(err)
	}
}

func GetSearch(key string, min string, max string) []str.Items {
	row, err := db.Query("CALL search(?, ?, ?)", key, min, max)
	if err != nil {
		panic(err)
	}
	var t1 str.Items
	var t2 []str.Items
	for row.Next() {
		err := row.Scan(&t1.Indexs, &t1.Urls, &t1.Name, &t1.Sku, &t1.SellP, &t1.OrP, &t1.Cur, &t1.Avail, &t1.Col, &t1.Cat, &t1.Sour, &t1.WebSou, &t1.Breadc, &t1.Desc, &t1.Brand, &t1.Images, &t1.Country, &t1.Lang, &t1.AvgRate, &t1.RevC, &t1.Crawledat)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, t1)
	}
	return t2
}
