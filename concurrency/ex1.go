package main
import (
	"fmt"
	"sync"
)
func sayWord(word string,grp *sync.WaitGroup){
	var i int=0
	for i<5{
		fmt.Println(word)
		i++
	}
	grp.Done()
}
func main(){
group1 :=new(sync.WaitGroup)
go sayWord("1",group1)
go sayWord("2",group1)
group1.Add(2)
group1.Wait()
}
