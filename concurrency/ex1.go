package main
import (
	"fmt"
)
func sayWord(word string){
	var i int=1
	for i<5{
		fmt.Println(word)
		i++
	}
}
func main(){
go sayWord("hi")
go sayWord("how you doin'")
sayWord("123")
}
