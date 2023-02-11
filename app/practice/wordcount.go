package practice

import (
	"fmt"
)

func WordCount(s string) map[string]int {
	var wordList []string
	var word string
	for i := 0; i < len(s); i++ {
		// 最後の単語は文字で終わり、リストに加える
		if i == int(len(s))-1 {
			word += s[i : i+1]
			wordList = append(wordList, word)
			// 確認している文章が文字だった場合、空文字に到達するまで文字を後ろから挿入する
		} else if s[i:i+1] != " " {
			word += s[i : i+1]
			// 確認している文章が空文字だった場合、その直前は単語としてまとめリストに入れる
		} else {
			wordList = append(wordList, word)
			word = ""
		}
	}
	wordMap := make(map[string]int)
	for i := 0; i < len(wordList); i++ {
		wordMap[wordList[i]] = 0
		// 自分自身の単語は1カウントし、別の場所に同じ単語があったらカウントを増やす
		for j := 0; j < len(wordList); j++ {
			if wordList[i] == wordList[j] {
				wordMap[wordList[i]] += 1
			}
		}
	}
	fmt.Println(wordMap)
	return wordMap
}

// example
// func main() {
// 	fmt.Println(WordCount("I am human"))
// }
