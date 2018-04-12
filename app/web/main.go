package main

import (
	"fmt"
	"github.com/tchalo/iselia/config"
	"github.com/tchalo/iselia/router"
)

func main() {
	is, err := config.NewIselia()

	if err != nil {
		fmt.Println(err)
		return
	}

	r := router.NewWebRouter(is)

	r.Run(":3000")
}
