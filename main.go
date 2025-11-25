package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Welcome to the Text Analyzer!")
    fmt.Println("Paste your text here, then press Enter:")
    text, _ := reader.ReadString('\n')
    text = strings.TrimSpace(text)

    wordCount := countWords(text)
    sentenceCount := countSentences(text)
    wordFrequency := frequencyWords(text)

    fmt.Printf("\nWord count: %d\n", wordCount)
    fmt.Printf("Sentence count: %d\n", sentenceCount)
    fmt.Println("\nWord frequency (top 5):")
    for i, wf := range wordFrequency {
        if i >= 5 {
            break
        }
        fmt.Printf("%d. %s : %d\n", i+1, wf.word, wf.count)
    }
}

// Counts the number of words (separated by spaces)
func countWords(text string) int {
    words := strings.Fields(text)
    return len(words)
}

// Counts the number of sentences (ending with . ! or ?)
func countSentences(text string) int {
    re := regexp.MustCompile(`[.!?]`)
    matches := re.FindAllStringIndex(text, -1)
    return len(matches)
}

// Struct to sort word frequencies
type wordFrequencyStruct struct {
    word  string
    count int
}

// Calculates word frequency (case-insensitive)
func frequencyWords(text string) []wordFrequencyStruct {
    words := strings.FieldsFunc(text, func(c rune) bool {
        return c == ' ' || c == ',' || c == '.' || c == '!' || c == '?' || c == ';' || c == ':'
    })
    freq := make(map[string]int)
    for _, w := range words {
        w = strings.ToLower(w)
        freq[w]++
    }
    // Convert to slice and sort
    var sorted []wordFrequencyStruct
    for k, v := range freq {
        sorted = append(sorted, wordFrequencyStruct{word: k, count: v})
    }
    // Sort in descending order by count
    for i := 0; i < len(sorted); i++ {
        for j := i + 1; j < len(sorted); j++ {
            if sorted[j].count > sorted[i].count {
                sorted[i], sorted[j] = sorted[j], sorted[i]
            }
        }
    }
    return sorted
}
