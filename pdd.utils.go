package pddsdk

import (
	"fmt"
	"strconv"
	"time"
)

var TimeLocal, _ = time.LoadLocation("Asia/Shanghai")

func GetNow() time.Time {
	return time.Now().In(TimeLocal)
}

// GetString convert interface to string.
func GetString(v interface{}) string {
	switch result := v.(type) {
	case string:
		return result
	case []byte:
		return string(result)
	default:
		if v != nil {
			return fmt.Sprint(result)
		}
	}
	return ""
}

// GetInt convert interface to int.
func GetInt(v interface{}) int {
	switch result := v.(type) {
	case int:
		return result
	case int32:
		return int(result)
	case int64:
		return int(result)
	default:
		if d := GetString(v); d != "" {
			value, err := strconv.Atoi(d)
			if err != nil {
				panic(err.Error())
			}
			return value
		}
	}
	return 0
}

// GetInt64 convert interface to int64.
func GetInt64(v interface{}) int64 {
	switch result := v.(type) {
	case int:
		return int64(result)
	case int32:
		return int64(result)
	case int64:
		return result
	default:

		if d := GetString(v); d != "" {
			value, err := strconv.ParseInt(d, 10, 64)
			if err != nil {
				panic(err.Error())
			}
			return value
		}
	}
	return 0
}

// GetFloat64 convert interface to float64.
func GetFloat64(v interface{}) float64 {
	switch result := v.(type) {
	case float64:
		return result
	default:
		if d := GetString(v); d != "" {
			value, err := strconv.ParseFloat(d, 64)
			if err != nil {
				panic(err.Error())
			}
			return value
		}
	}
	return 0
}

//四舍五入，保留2位
func Round(number float64) float64 {
	numberStr := fmt.Sprintf("%.2f", number) //四舍五入，保留2位
	floatNum, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		panic(err.Error())
	}
	return floatNum
}

func StrToTime(dateText, timeLayout string) (timestamp int64) {
	//时间模板用 2006-01-02 15:04:05 ，据说是golang的诞生时间。
	var timeFormat string
	if timeLayout == "date" {
		timeFormat = "2006-01-02"
	} else if timeLayout == "datetime" {
		timeFormat = "2006-01-02 15:04:05"
	} else {
		timeFormat = timeLayout
	}
	loc, err := time.LoadLocation("Local")
	if err != nil {
		panic(err.Error())
	} //重要：获取时区
	theTime, err := time.ParseInLocation(timeFormat, dateText, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		panic(err.Error())
	}
	timestamp = theTime.Unix() //转化为时间戳 类型是int64
	return
}
