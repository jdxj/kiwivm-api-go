# KiwiVM-API-Go

This repository encapsulates [kiwi api](https://kiwivm.64clouds.com/) with
go http client, which can be used to automate some tasks.

Note: This is not an official repository.

![logo](https://kiwivm.64clouds.com/1670298/img/kiwivm_logo_100x33px.png)

## Usage

```shell
$ go get -u github.com/jdxj/kiwivm-api-go@v0.2.0
```

[examples](./examples)

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	kiwi "github.com/jdxj/kiwivm-api-go"
)

func main() {
	ctx := context.Background()
	client := kiwi.NewClient("your veid", "your api-key")
	stats, err := client.GetRawUsageStats(ctx)
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
