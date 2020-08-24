package main

import (
	"errorDemo/client"
	"fmt"
)

func main() {
	i, e := client.PostMetadata("Vipin")
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", e, e)
}
