package main 


import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

var EndTimeSlotKey []string = []string{"09", "10"} 

var EndTimeSlotVal = map[string][]string {
	"09" : []string{"35", "40", "45", "50", "55"},
	"10" : []string{"00", "05", "10", "15", "20", "25", "30", "35", "40", "45", "50", "55"},
}



var stock_code string = "600704"


func TradeSiteGet(code string, hour string, min string) string {
	site := "http://quotes.money.163.com/service/zhubi_ajax.html?symbol=" + code + "&end=" + hour + "%3A" + min + "%3A00"
	return site
}

func SiteContentGet(site string) []byte {
	var b []byte
	resp, err := http.Get(site)
	defer resp.Body.Close()  
	if err != nil {
		fmt.Println(err)
	} else {
		b, _ = ioutil.ReadAll(resp.Body)
	}
	//fmt.Println(body)
	return b
}

func FloatToString(input_num float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'f', 6, 64)
}


func DecodeFromTrade(info []byte, db *sql.DB) {

	fmt.Println("DecodeFromTrade called")

	//fmt.Println(string(info))

	var data map[string]interface{}
	var sql string

	if err := json.Unmarshal(info, &data); err == nil {
		fmt.Println("From ", data["begin"], "To ", data["end"])
		//fmt.Println(data["zhubi_list"])
		data_array := data["zhubi_list"].([]interface{})
		fmt.Println("Total", len(data_array), "enties")
		for _, v := range data_array {
			entry := v.(map[string]interface{})
			//fmt.Println(i, ": ", entry["PRICE"])
			//TODO:insert into db
			sql = `insert into trade(TRADE_TYPE, PRICE, TURNOVER) values(` + FloatToString(entry["TRADE_TYPE"].(float64)) + `,` + FloatToString(entry["PRICE"].(float64)) + `,` + FloatToString(entry["TURNOVER_INC"].(float64)) + `);`
			//fmt.Println(sql)
			db.Exec(sql)
		}
	}

}


func main() {

	fmt.Println("hello,welcome")


	dbname := "./" + stock_code + ".db"
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		fmt.Println("test.db open failed")
		fmt.Println(err)
	}

	fmt.Println(dbname, "open sucessfully")

	sqlstring := `create table trade (TRADE_TYPE integer, PRICE text, TURNOVER integer);`
	db.Exec(sqlstring)

	target_site := TradeSiteGet(stock_code, EndTimeSlotKey[0], EndTimeSlotVal[EndTimeSlotKey[0]][0])

	fmt.Println(target_site)
	
	trade_info := SiteContentGet(target_site)

	DecodeFromTrade(trade_info, db)

	sqlstring = `select * from trade;`

	rows, err := db.Query(sqlstring)
	defer rows.Close()

	for rows.Next() {
		var trade_type int 
		var price float64
		var turnover int 
		err = rows.Scan(&trade_type, &price, &turnover)
		fmt.Printf("%d, %f, %d\n", trade_type, price, turnover)
	}

	db.Close()

}