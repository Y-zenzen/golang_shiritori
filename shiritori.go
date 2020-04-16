package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("GAME START")
	firstWord := "あ"
	start(firstWord)
}

func start(word string) {
	fmt.Printf("最初の文字: %s\n", word)
	text := input()
	judge(text, word)
}

func input() string {
	fmt.Println("文字を入力してください")

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	return text
}

func judge(text, word string) {
	//文字列を文字単位で扱う
	textRune := []rune(text)
	firstWord := string(textRune[0])
	size := len(textRune)

	if firstWord != word {
		fmt.Println("最初の文字が違います")
		fmt.Printf("入力した最初の文字: %s\n", firstWord)
		fmt.Println("もう一度入力してください")
		start(word)
	}

	// whiteList := []string{
	// 	"あ", "い", "う", "え", "お", "か", "き", "く", "け", "こ", "さ", "し", "す", "せ", "そ", "た", "ち", "つ", "て", "と", "な", "に", "ぬ", "ね", "の", "は", "ひ", "ふ", "へ", "ほ", "ま", "み", "む", "め", "も", "や", "ゆ", "よ", "ら", "り", "る", "れ", "ろ", "わ", "を", "ん",
	// 	"ア", "イ", "ウ", "エ", "オ", "カ", "キ", "ク", "ケ", "コ", "サ", "シ", "ス", "セ", "ソ", "タ", "チ", "ツ", "テ", "ト", "ナ", "ニ", "ヌ", "ネ", "ノ", "ハ", "ヒ", "フ", "ヘ", "ホ", "マ", "ミ", "ム", "メ", "モ", "ヤ", "ユ", "ヨ", "ラ", "リ", "ル", "レ", "ロ", "ワ", "ヲ", "ン",
	// }

	// for key, value := range text {
	// 	if !inArray(whiteList, value) {
	// 		fmt.Println("入力値が不正です")

	// 	}
	// }

	lastWord := string(textRune[size-1])
	fmt.Printf("最後の文字: %s", lastWord)

	if lastWord == "ん" {
		fmt.Println("最後の文字が「ん」です")
		end()
	}

	start(lastWord)
}

// func inArray(whiteList []string, word string) bool {
// 	for _, value := range whiteList {
// 		if word != value {
// 			return false
// 		}
// 	}
// 	return true
// }

func end() {
	fmt.Println("Game Over")
	os.Exit(0)
}
