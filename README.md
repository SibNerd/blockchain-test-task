# About
This program computes and outputs ethereum address, which balance changed the most for the last 100 block in format:

> Address with the value that changed the most: *address*

> Address value changed by: *absolute amount*

## Structure

### helpers.go
* _**GetJSON**_ - получение JSON по web-адресу get JSON
* _**HexToDec**_ - change hex number to decimal 
* _**DecToHex**_ - change decimal number to hex

### main.go
* _**getLastBlockNumber**_ - get last block address
* _**getTagInfo**_ - get block info with given address
* _**modifyTagInfo**_ - change blocks' data
* _**modifyUsersInfo**_ - change transaction participants' data
* _**makeTransaction**_ - make transaction between sender and reciever

  - **Why separate fucntions?**
    - modifyUsersInfo checks addresses of Sender and Reciever and makes transactions with different conditions. To prevent copypaste of huge amount of code, separate function for making transaction was created.
