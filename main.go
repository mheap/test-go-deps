package main

import (
	"fmt"
	tags "github.com/kong/go-apiops/tags"
	gdr "github.com/kong/go-database-reconciler/pkg/utils"
)

func main() {
	foo := tags.Tagger{}
	fmt.Printf("foo: %v\n", foo)
	fmt.Printf("%s", gdr.ImplementationTypeKongGateway)
}
