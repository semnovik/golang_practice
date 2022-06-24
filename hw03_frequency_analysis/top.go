package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	// Делим слова в строке, помещаем в список
	splitText := strings.Fields(text)
	wordAmountMap := make(map[string]int)
	var sortedList []sorted

	// Берем слово и считаем количество вхождений в список
	counter := 0
	for _, w := range splitText {
		for _, v := range splitText {
			if w == v {
				counter++
			}
		}
		wordAmountMap[w] = counter
		counter = 0
	}

	// Помещаем список солов и количество повторений в список структур
	for k, v := range wordAmountMap {
		sortedList = append(sortedList, sorted{k, v})
	}

	return lengthCorrection(sortMyList(sortedList))
}

type sorted struct {
	word   string
	amount int
}

func sortMyList(list []sorted) []sorted {
	sort.SliceStable(list, func(i, j int) bool { return list[i].word < list[j].word })
	sort.SliceStable(list, func(i, j int) bool { return list[i].amount > list[j].amount })
	return list
}

func lengthCorrection(sortedList []sorted) []string {
	var wordsAmount int
	var list []string

	if len(sortedList) < 10 {
		wordsAmount = len(sortedList)
	} else {
		wordsAmount = 10
	}

	for i := range sortedList[:wordsAmount] {
		list = append(list, sortedList[i].word)
	}

	return list
}
