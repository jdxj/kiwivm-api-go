# KiwiVM-SDK-Go

This repository encapsulates kiwi api with go http client, which can be used to automate some tasks.

Note: This is not an official repository.

![logo](https://kiwivm.64clouds.com/1670298/img/kiwivm_logo_100x33px.png)

## Usage

[examples](./examples)

```go
package main

import (
	"fmt"
	"log"
	"time"

	sdk "github.com/jdxj/kiwivm-sdk-go"
	"github.com/jdxj/kiwivm-sdk-go/conf"
)

func main() {
	client := sdk.NewClient(conf.VeID, conf.APIKey)
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
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
