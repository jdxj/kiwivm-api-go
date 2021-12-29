package main

import (
	"log"

	sdk "github.com/jdxj/kiwivm-sdk-go"
)

func main() {
	var (
		veid   = ""
		apiKey = ""
	)
	client := sdk.NewClient(veid, apiKey)
	stats, err := client.GetRawUsageStats()
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range stats.Data {
		fmt.pr
	}
}
