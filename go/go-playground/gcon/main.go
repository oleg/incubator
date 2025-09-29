package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func main() {
	frequency := make([]int, len(allLetters))
	for i := 1000; i < 1010; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency)
	}
	time.Sleep(5 * time.Second)
	for i, c := range allLetters {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
}

// a-49344 b-9519 c-27499 d-24464 e-86476 f-15292 g-11676 h-22505 i-47698 j-821 k-3893 l-22233 m-19666 n-51797 o-51761 p-20141 q-1905 r-44794 s-51222 t-64298 u-16622 v-6029 w-7731 x-1938 y-7961 z-530
// a-46859 b-9400 c-26670 d-23758 e-79104 f-15039 g-11519 h-21949 i-45448 j-821 k-3877 l-21688 m-19230 n-49156 o-49169 p-19794 q-1897 r-42019 s-47732 t-59953 u-16265 v-5979 w-7660 x-1923 y-7906 z-530
func countLetters(url string, frequency []int) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("Server returned non-200 status code: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex]++
		}
	}
	fmt.Printf("Completed: %s\n", url)
}
