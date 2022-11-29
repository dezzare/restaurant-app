package main

import (
	"fmt"
	"net/http"

	"github.com/dezzare/restaurant-app/menuservice/router"
)

func main() {
  r, err := router.SetupRouter()
    if err != nil {
        fmt.Println("Deu erro")
    }
  r.Run(":3000")
}
