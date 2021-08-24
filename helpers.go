package main

import (
	"encoding/json"
	"strconv"
)

func GetJson(url string, target interface{}) error {
	// Get JSON responce from web
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&target)
}

func DecToHex(input int64) string {
	// Turn integer into hex string
	num := strconv.FormatInt(input, 16)
	return num
}

func HexToDec(input string) int64 {
	// Turn hex string into integer
	num, _ := strconv.ParseInt(input, 16, 64)
	return num
}
