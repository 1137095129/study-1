package main

import (
	"fmt"
	"os"
	"strings"
)

func TestPanic() {
	defer func(){
		i := recover()
		fmt.Println(i)
	}()
	panic("aaaaa")
}

func getEnviron() {
	environ := os.Environ()
	for i := range environ {
		s := environ[i]
		split := strings.Split(s, "=")
		fmt.Println(fmt.Sprintf("%s=%s",split[0],split[1]))
	}
}

type MyTypeSlice []string

func newTypeSlice(strs []string, p *[]string) *MyTypeSlice {
	*p =strs
	return (*MyTypeSlice)(p)
}

func (ts *MyTypeSlice) Set(string2 string) error {
	*ts=MyTypeSlice(strings.Split(string2,","))
	return nil
}
func (ts *MyTypeSlice) String() string {
	return ""
}

func main() {
	//TestPanic()
	//arr := study.CreatePersonList(0)
	//for i := 0; i < 20; i++ {
	//	person := study.CreatePerson(fmt.Sprintf("aaa%d", i))
	//	person.SetAge(rand.Int())
	//	arr = append(arr, person)
	//}
	//for i := range arr {
	//	fmt.Println(arr[i])
	//}
	//sort.Sort(arr)
	//fmt.Println("--------------------------------")
	//for i := range arr {
	//	fmt.Println(arr[i])
	//}
	//getEnviron()
	//flag.Var()
}
