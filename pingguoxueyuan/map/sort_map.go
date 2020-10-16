package main


import "fmt"
import "sort"

func main() {
	map1 := make(map[string]int)
	map1["Mon"]=1
	map1["Tue"]=2
	map1["Wed"]=3
	map1["Thu"]=4
	map1["Fri"]=5
	map1["Sat"]=6
	map1["Sun"]=7
	fmt.Println(map1)

	invMap := make(map[int]string,len(map1))
	for k,v := range map1{
		invMap[v]=k
	}
	fmt.Println("inverted:")
	for k,v := range invMap{
		fmt.Printf("The %d day is %s\n",k,v)
	}
	fmt.Println("inverted and sorted:")

	value := make([]int,len(invMap))
	i := 0
	for val,_ := range invMap{
		value[i]=val
		i++
	}
	sort.Ints(value)
	for _,j := range value{
		fmt.Printf("The %d day is %s\n",j,invMap[j])
	}
}