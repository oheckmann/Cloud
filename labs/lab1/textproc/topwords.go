// Find the top K most common words in a text document.
// Input path: location of the document, K top words
// Output: Slice of top K words
// For this excercise, word is defined as characters separated by a whitespace

// Note: You should use `checkError` to handle potential errors.

package textproc

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func TopWords(path string, K int) []WordCount {
	var log = fmt.Println
	result := []WordCount{}
	data, err := os.ReadFile(path)
	checkError(err)
	dataString := string(data)
	words := strings.Fields(dataString)
	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word] += 1
	}

	for key, value := range wordsMap {
		result = append(result, WordCount{Word: key, Count: value})
	}
	
	sortWordCounts(result)
	log(result[:K])
	return result[:K]
}

//--------------- DO NOT MODIFY----------------!

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

// Method to convert struct to string format
func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.

func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
