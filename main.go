package main

import (
	"fmt"
	"reflect"
)

func main() {
	if err := ChessesInit("chesses.json"); err != nil {
		fmt.Println("error:", err)
		return
	}
	for i, ch := range chesses {
		fmt.Println("-----------------------", i)
		for _, name := range ch.name {
			fmt.Println("name:", name)
		}
		fmt.Println("career:", reflect.TypeOf(ch.career).Name())
		for _, race := range ch.race {
			fmt.Println("race:", reflect.TypeOf(race).Name())
		}
		fmt.Println("number:", ColorAmount[reflect.TypeOf(ch.color)])
	}
}
