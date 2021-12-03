package main

import (
	"fmt"

	"golang.org/x/text/encoding/unicode"

	"github.com/hashicorp/vault/api"
)

func main() {
	e := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	foo := api.Auth{}
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", foo)
}
