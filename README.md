# BEXS [![GoDoc](https://godoc.org/github.com/diogoagra/bexs?status.svg)](https://godoc.org/github.com/diogoagra/bexs)
Go client for Bexs

## Installation
```
go get github.com/diogoagra/bexs
```

## Usage
```
package main

import (
	"fmt"

	"github.com/diogoagra/bexs"
)

func main() {
	client := bexs.New("bleutrade", "api-key", "api-secret", true)
	balances, err := client.GetBalances()

	if err != nil {
		return
	}

	fmt.Println(balances)
}

```
