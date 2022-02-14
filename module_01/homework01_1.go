package main

import "fmt"

func main() {
	words := []string{"I", "am", "stupid", "and", "weak"}
	useIfWords := useIf(words)
	useSwitchWords := useSwitch(words)
	uesMapWords := useMap(words)
	fmt.Printf("words:\t%v,\n", words)
	fmt.Printf("useIfWords:\t%v,\n", useIfWords)
	fmt.Printf("useSwitchWords:\t%v,\n", useSwitchWords)
	fmt.Printf("uesMapWords:\t%v\n", uesMapWords)
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
	return
}
