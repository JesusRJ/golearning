package main

import "fmt"

type Dog struct {
	name string
}

func main() {
    dogs := []Dog{Dog{"Stallone"}, Dog{"Murdock"}, Dog{"Bradock"}}
    var dogsPtr []*Dog

    for _, dog := range dogs {
        fmt.Printf("'dog' with name %s and pointer: <%p>\n", dog.name, &dog)
        dogsPtr = append(dogsPtr, &dog)
    }

    for _, dogPtr := range dogsPtr {
        fmt.Printf("'dogPtr' with name %s and pointer: <%p>\n", dogPtr.name, dogPtr)
    }
}
