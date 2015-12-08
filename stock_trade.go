package main 


import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
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


func DecodeFromTrade(info []byte) {

	fmt.Println("DecodeFromTrade called")

	//fmt.Println(string(info))

	var data map[string]interface{}

	if err := json.Unmarshal(info, &data); err == nil {
		fmt.Println("From ", data["begin"], "To ", data["end"])
		//fmt.Println(data["zhubi_list"])
		data_array := data["zhubi_list"].([]interface{})
		fmt.Println("Total", len(data_array), "enties")
		//for i, v := range data_array {
			//entry := v.(map[string]interface{})
			//fmt.Println(i, ": ", entry["PRICE"])
			//TODO:insert into db
		//}
	}

}


func main() {

	fmt.Println("hello,welcome")

	target_site := TradeSiteGet(stock_code, EndTimeSlotKey[0], EndTimeSlotVal[EndTimeSlotKey[0]][1])

	fmt.Println(target_site)
	
	trade_info := SiteContentGet(target_site)

	DecodeFromTrade(trade_info)

}