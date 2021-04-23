# shea golang framework

Shea is a useful framework enjoy :)

## Installation

### 1: install shea-go:
```bash
go get -u github.com/sheaxiang/shea-go
```
### 2: import it in your code:
```
import "github.com/sheaxiang/shea-go"
```

## Quick start

```
package main

import (
	"github.com/sheaxiang/shea-go"
	"net/http"
)

func main()  {
	r := shea.Default()

	r.GET("/index", func(c *shea.Context) {
		c.HTML(http.StatusOK, "<h1>Hello SheaGo!</h1>")
	})

	r.Run(":9999")
}
```
