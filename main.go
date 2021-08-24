package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
  "math"
)

// Different types for work with JSON data

type LastBlockInfo struct {
	Jsonrpc string
	Id      int
	Result  string
}

type ResultInfo struct {
	BaseFeePerGas    string
	Difficulty       string
	ExtraData        string
	GasLimit         string
	GasUsed          string
	Hash             string
	LogsBloom        string
	Miner            string
	MixHash          string
	Nonce            string
	Number           string
	ParentHash       string
	ReceiptsRoot     string
	Sha3Uncles       string
	Size             string
	StateRoot        string
	Timestamp        string
	TotalDifficulty  string
	Transactions     []map[string]string
	TransactionsRoot string
	Uncles           []string
}

type BlockInfo struct {
	Jsonrpc string
	Id      int
	Result  ResultInfo
}



// Handling transactions

func getLastBlockNumber() int64 {
	// Call API method to get last block number and return it
	var lastAddressURL = "https://api.etherscan.io/api?module=proxy&action=eth_blockNumber"
	var info LastBlockInfo

	err := GetJson(lastAddressURL, &info)
	if err != nil {
		fmt.Println(err)
	}

	number := strings.Split(info.Result, "x")[1]
	numint := HexToDec(number)

	return numint
}

func getTagInfo(number int64) ResultInfo {
	// Get info from transaction JSON and return it
	var resInfo BlockInfo
	numhex := DecToHex(number)
	numberTag := "0x" + string(numhex)
	infoURL := "https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=" + numberTag + "&boolean=true"
	GetJson(infoURL, &resInfo)
	return resInfo.Result
}

func modifyTagInfo(input ResultInfo, users *map[string]string) {
	// Iterate over transactions and apply changes to users map
	result := map[string]string{}

	for _, info := range input.Transactions {
		result["fromUser"] = info["from"]
		result["toUser"] = info["to"]
		result["Value"] = info["value"]
		modifyUsersInfo(*users, result)
	}
}

func modifyUsersInfo(users, info map[string]string) {
	// Work with Users map and change adresses' values
	_, sender_exists := users[info["fromUser"]]
	_, reciever_exists := users[info["toUser"]]

	switch {
	case sender_exists && reciever_exists:
		transaction(users, info)
	case sender_exists && !reciever_exists:
		users[info["toUser"]] = string(DecToHex(0))
		transaction(users, info)
	case !sender_exists && reciever_exists:
		users[info["fromUser"]] = string(DecToHex(0))
	  transaction(users, info)
	case !sender_exists && !reciever_exists:
		users[info["fromUser"]] = string(DecToHex(0))
		users[info["toUser"]] = string(DecToHex(0))
		transaction(users, info)
	}
}

func transaction(users, info map[string]string) {
	// Separate transaction between users in it's own function
	senderValue := HexToDec(strings.Split(users[info["fromUser"]], "x")[0])
	value := HexToDec(strings.Split(info["Value"], "x")[1])
	recieverValue := HexToDec(strings.Split(users[info["toUser"]], "x")[0])
	senderValue -= value
	recieverValue += value
	users[info["fromUser"]] = string(DecToHex(senderValue))
	users[info["toUser"]] = string(DecToHex(recieverValue))
}



// Main function

var myClient = http.Client{Timeout: 10 * time.Second}

func main() {
	usersStrings := map[string]string{}
  usersInts := map[string]float64{}
  var maxValue float64 = 0
  maxUserTag := ""

	lastBlockNumber := getLastBlockNumber()
	endBlockNumber := lastBlockNumber - 100


	for i := lastBlockNumber; i > endBlockNumber; i-- {
    // Get info about transactions and store it in map
		tagInfo := getTagInfo(i)
		modifyTagInfo(tagInfo, &usersStrings)
	}

	for user, value := range usersStrings {
		// Get modules of users's values
      usersInts[user] = math.Abs(float64(HexToDec(value)))
	}

  for user, value := range usersInts {
    // Get max value
    if value > maxValue {
      maxValue = value
      maxUserTag = user
    }
  }
  fmt.Println("Address with the value that changed the most: ", maxUserTag)
  fmt.Println("Address value changed by: ", maxValue)
}
