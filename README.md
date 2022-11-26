# aws-zerolog
Zerolog wrapper for AWS Logger

## Installation

```bash
go get github.com/ziflex/aws-zerolog
```

## Quick start

```go
package main

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/rs/zerolog"
	"github.com/ziflex/aws-zero"
)

func main() {
	logger := zerolog.New(os.Stdout)
	conf, err := config.LoadDefaultConfig(context.Background(), config.WithLogger(awszero.New(logger)))
	
	if err != nil {
		panic(err)
    }
	
	println(conf)
}
```