package main

import "fmt"

var ch = make(chan int)
func main() {
	words := []string{"I", "am", "stupid", "and", "weak"}
	fmt.Printf("words:\t%v,\n", words)
	go useIf(words)
	go useSwitch(words)
	go useMap(words)
	for i:=0; i<3; i++ {
		<- ch
	}
}

func useIf(words []string) (new_words []string) {
	for _, word := range words {
		if word == "stupid" {
			word = "smart"
		} else if word == "weak" {
			word = "strong"
		}
		new_words = append(new_words, word)
	}
	fmt.Printf("useIfWords:\t%v\n",new_words)
	ch <- 1
	return
}

func useSwitch(words []string) (new_words []string) {
	for _, word := range words {
		switch word {
		case "stupid":
			word = "smart"
		case "weak":
			word = "strong"
		default:
		}
		new_words = append(new_words, word)
	}
	fmt.Printf("useSwitchWords:\t%v\n", new_words)
	ch <- 1
	return
}

func useMap(words []string) (new_words []string) {
	dict := map[string]string{"stupid": "smart", "weak": "strong"}
	for _, word := range words {
		if value, ok := dict[word]; ok {
			word = value
		}
		new_words = append(new_words, word)
	}
	fmt.Printf("uesMapWords:\t%v\n", new_words)
	ch <- 1
	return
}
