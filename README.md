# parsuri


[![GoDoc](https://godoc.org/github.com/yunginnanet/parsuri?status.svg)](http://godoc.org/github.com/yunginnanet/parsuri)

[parsuri](https://github.com/yunginnanet/parsuri) is a Go library to parse suricata eve.json files with proper marshaling.

### Example

```golang
package main

import (
	"github.com/yunginnanet/parsuri"
	"log"
)

func main() {
	loader := parsuri.NewLoader()

	// Load the eve.json file asynchronously
	if err := loader.LoadOneFile("eve.json"); err != nil {
		log.Fatal(err)
	}

	// Range over the events and print dns answers to stdout
	for loader.More() {
		if err := loader.Err(); err != nil {
			log.Fatal(err)
		}
		event := loader.Event()
		if event.DNS != nil && !event.DNS.Empty() && event.DNS.Type == "answer" {
			log.Println(event.DNS)
		}
	}

	if err := loader.Err(); err != nil {
		log.Fatal(err)
	}
}
```

### Credit

This is a rewrite of [surevego](https://github.com/rhaist/surevego).

### License

- BSD-3 Copyright (c) 2017 **Robert Haist**

- BSD-3 Copyright (c) 2025 [yunginnanet](https://github.com/yunginnanet)
