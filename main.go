package main

import "fmt"

func emit(wordChannel chan string, doneChannel chan bool) {
	i := 0
	words := []string{"feed", "the", "monkey"}
	for {
		select {
		case wordChannel <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		case <-doneChannel:
			close(doneChannel)
			return
		}
	}
}
func main() {
	wordChannel := make(chan string)
	doneChannel := make(chan bool)

	go emit(wordChannel, doneChannel)

	for i := 0; i <= 100; i++ {
		fmt.Println(<-wordChannel)
	}
	doneChannel <- true
}
