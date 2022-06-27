// Http handler implementations
// ServeHTTP
package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()

	fmt.Println(time.Day())
}
