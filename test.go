package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {

	for {
		data, err := fecthData()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data)

		time.Sleep(time.Second)
	}

}

func fecthData() (string, error) {
	body := strings.NewReader("dil_kodu=tr")
	req, err := http.NewRequest("POST", "https://www.haremaltin.com/dashboard/ajax/doviz", body)
	if err != nil {
		// handle err
		return "", err
	}
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		return "", err
	}
	defer resp.Body.Close()
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		return "", err
	}
	// var data map[string]interface{}
	// err = json.Unmarshal(jsonData, &data)
	// if err != nil {
	// 	return nil, err
	// }
	return string(jsonData), nil
}
