package main

import (
	"fmt"
	"gatherup/sdk"
)

func main() {
	client := sdk.Client{
		Credentials: sdk.Credentials{
			ClientId: "1234",
			Bearer:   "abcd",
		},
		Url:       "https://app.gatherup.com/api",
		Aggregate: true,
	}

	response, err := client.Request(
		"/customers/get",
		map[string]string{},
	)

	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
	} else {
		if response.IsSuccess() {
			ok, data := response.Get("data")
			if ok {
				fmt.Println(data)
			} else {
				fmt.Println("Error: Invalid response!")
			}
		} else {
			fmt.Print("Error: ")
			fmt.Println(response.GetMessage())
		}
	}
}
