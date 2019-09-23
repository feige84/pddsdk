package pddsdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	client := NewClient("xxx", "xxx")
	//client.GetCache = FileGetCache
	//client.SetCache = FileSetCache
	//client.WriteErrLog = InsertErrLog
	client.CacheLife = 0
	client.AccessToken = ""
	data3, apiErr3 := client.DdkGoodsSearch("纯棉男女宝宝裤子春秋冬季新生儿幼童装外出婴儿衣服大PP加绒长裤", nil, "", "", 1, 1, 10, 0, 0, 0, false, false)
	jsonData3, jsonErr3 := json.Marshal(data3)
	fmt.Println("jsonData:", string(jsonData3), jsonErr3, apiErr3)
	fmt.Println("return:", data3, apiErr3)
}
