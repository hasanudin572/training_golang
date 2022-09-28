package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Age      int    `json:"Age"`
}

func main() {
	var jsonString = `{
	"Fullname":"airil",
	"Email":"airin@gmail.com",
	"Age": 29

}
`
	var result employee
	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("fullname :", result.Fullname)
	fmt.Println("email:", result.Email)
	fmt.Println("Age:", result.Age)
}
