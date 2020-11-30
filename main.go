package main

import (
	//"bufio"
	"fmt"
	//"os"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Letter Frquency Counter")
	fmt.Println("=======================")
	fmt.Println("Please write any text")
	//text, _ := reader.ReadString('\n')
	text := "heyy"
	var freq = GetLettersFrequency(text)
	fmt.Println("=======================")
	wg.Wait()
	PrintMap(freq);
}

func PrintMap(mapToPrint LetterMap) {
	for key, value := range mapToPrint {
		fmt.Printf("'%c' : %d\n", key, value)
	}
}