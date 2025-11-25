# Text Analyzer (Beginner Go Project)

Simple Go program that:
- counts the number of words  
- counts the number of sentences  
- shows the most frequent words

## Usage

1. Open this folder in your terminal.
2. Run:
   ```sh
   go run main.go
   ```
3. Paste your text and press Enter!

---

## Explanation

- **countWords** counts words separated by spaces.
- **countSentences** counts sentences by searching for `.`, `!`, `?` as endings.
- **frequencyWords** gives you the most common words (case-insensitive) and shows the five most frequent.
