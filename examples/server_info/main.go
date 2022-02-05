package main

import (
	"fmt"
	"log"

	sdk "github.com/jdxj/kiwivm-sdk-go"
)

func main() {
	client := sdk.NewClient("", "")
	info, err := client.GetServiceInfo()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("plan: %s\n", info.Plan)
	fmt.Printf("vm type: %s\n", info.VmType)
	fmt.Printf("ip: %v\n", info.IpAddresses)
	fmt.Printf("os: %s\n", info.OS)
}
