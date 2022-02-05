package main

import (
	"fmt"
	"log"
	"time"

	sdk "github.com/jdxj/kiwivm-sdk-go"
)

func main() {
	client := sdk.NewClient("", "")
	stats, err := client.GetRawUsageStats()
	if err != nil {
		log.Fatalln(err)
	}

	var (
		in  int64
		out int64
	)
	for _, v := range stats.Data {
		in += v.NetworkInBytes
		out += v.NetworkOutBytes
		fmt.Printf("ts: %s, in: %d, out: %d\n",
			time.Unix(v.Timestamp, 0).Format(time.RFC3339),
			v.NetworkInBytes,
			v.NetworkOutBytes,
		)
	}
	fmt.Printf("in: %d, out: %d, total: %d\n", in, out, in+out)
}

func today0oClock() int64 {
	return (time.Now().Unix() / 86400) * 86400
}
