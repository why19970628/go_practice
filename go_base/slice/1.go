package main

import (
	"fmt"
	"sync"
)

func slice_cut_append() {
	x := []int{1, 2, 3}
	y := x[:2]
	fmt.Println(x, y, &x[0], &y[0])
	y = append(y, 50)
	fmt.Println(x, y, &x[0], &y[0])
	y = append(y, 60)
	fmt.Println(x, y, &x[0], &y[0])
	y[0] = 101
	fmt.Println(x)
}

type LetterFreq map[rune]int

func LetterCount(strs []string) LetterFreq {
	resp := LetterFreq{}
	wg := &sync.WaitGroup{}
	lock := &sync.Mutex{}
	wg.Add(len(strs))
	for i := 0; i < len(strs); i++ {
		go func(i int) {
			for j := 0; j < len(strs[i]); j++ {
				lock.Lock()
				resp[rune(strs[i][j])]++
				lock.Unlock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return resp
}

func LetterCount2(strs []string) LetterFreq {
	resp := LetterFreq{}
	wg := &sync.WaitGroup{}
	lock := &sync.Mutex{}
	wg.Add(len(strs))
	for i := 0; i < len(strs); i++ {
		go func(i int) {
			for j := 0; j < len(strs[i]); j++ {
				lock.Lock()
				resp[rune(strs[i][j])]++
				lock.Unlock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return resp
}

func main() {
	fmt.Println(LetterCount([]string{"123", "456", "123"}))
}
