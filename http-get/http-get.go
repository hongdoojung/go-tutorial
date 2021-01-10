package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	jsonFile, err := os.Open("keys.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened keys.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	appid := result["appid"].(string)

	url := "https://api.openweathermap.org/data/2.5/weather?id=1835848&units=metric&appid="
	final_url := url + appid
	req, err := http.NewRequest("GET", final_url, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))

	var weather map[string]interface{}
	json.Unmarshal([]byte(body), &weather)
	weather_main := weather["main"].(map[string]interface{})
	temp := weather_main["temp"].(float64)
	fmt.Println(temp)
}
