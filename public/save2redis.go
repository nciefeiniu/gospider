package public

import (
	"encoding/json"
	"fmt"
)

func clearRedisKey(key string) {
	cli := Get()
	defer cli.Close()
	_, _ = cli.Do("DEL", key)
}

func data2Json(data []map[string]string) string {
	if data == nil {
		return ""
	}
	mjson, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err, "json.Marshal error")
		return ""
	}
	return string(mjson)
}

func Savedata2RD(datas [] map[string]string, keyname string) {

	dataLength := 0
	for _, val := range datas {
		if val != nil {
			dataLength++
		}
	}

	finallyData := make([] map[string]string, 0, dataLength)
	for _, val := range datas {
		if val != nil {
			finallyData = append(finallyData, val)
		}
	}

	jsonData := data2Json(finallyData)

	if jsonData != "" {
		cli := Get()
		defer cli.Close()

		_, err :=cli.Do("SET", keyname, jsonData)
		fmt.Println(jsonData)

		if err != nil {
			fmt.Println(err)
		}
	}

}
