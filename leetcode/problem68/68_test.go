package problem68

import (
	"fmt"
	"strings"
	"testing"
)

func fullJustify(words []string, maxWidth int) []string {
	justify := func(line []string, wordLength int) string {
		spaceCount := maxWidth - wordLength
		if len(line) == 1 {
			return line[0] + strings.Repeat(" ", spaceCount)
		}
		eachSpaceSize := spaceCount / (len(line) - 1)
		remainder := spaceCount % (len(line) - 1)
		result := ""
		for i := 0; i < len(line)-1; i++ {
			result += line[i]
			result += strings.Repeat(" ", eachSpaceSize)
			if remainder > 0 {
				result += strings.Repeat(" ", 1)
				remainder--
			}
			spaceCount -= eachSpaceSize
		}
		result += line[len(line)-1]
		return result
	}

	result := make([]string, 0)
	iter := 0
	for iter < len(words) {
		currentLine := []string{words[iter]}
		totalLength := len(words[iter])
		wordLength := len(words[iter])

		iter++
		for iter < len(words) && totalLength+len(words[iter])+1 <= maxWidth {
			totalLength += len(words[iter]) + 1
			wordLength += len(words[iter])
			currentLine = append(currentLine, words[iter])
			iter++
		}
		if iter != len(words) {
			result = append(result, justify(currentLine, wordLength))
		} else {
			lastLine := ""
			spaceCount := maxWidth - wordLength
			for i := 0; i < len(currentLine)-1; i++ {
				lastLine += currentLine[i]
				lastLine += strings.Repeat(" ", 1)
				spaceCount -= 1
			}

			lastLine += currentLine[len(currentLine)-1]
			lastLine += strings.Repeat(" ", spaceCount)

			result = append(result, lastLine)
		}
	}

	return result
}

func Test68(t *testing.T) {
	
	justified:=fullJustify(strings.Split(strings.ReplaceAll(strings.ReplaceAll(`
	justify := func(line []string, wordLength int) string {
		spaceCount := maxWidth - wordLength
		if len(line) == 1 {
			return line[0] + strings.Repeat(" ", spaceCount)
		}
		eachSpaceSize := spaceCount / (len(line) - 1)
		remainder := spaceCount % (len(line) - 1)
		result := ""
		for i := 0; i < len(line)-1; i++ {
			result += line[i]
			result += strings.Repeat(" ", eachSpaceSize)
			if remainder > 0 {
				result += strings.Repeat(" ", 1)
				remainder--
			}
			spaceCount -= eachSpaceSize
		}
		result += line[len(line)-1]
		return result
	}

	result := make([]string, 0)
	iter := 0
	for iter < len(words) {
		currentLine := []string{words[iter]}
		totalLength := len(words[iter])
		wordLength := len(words[iter])

		iter++
		for iter < len(words) && totalLength+len(words[iter])+1 <= maxWidth {
			totalLength += len(words[iter]) + 1
			wordLength += len(words[iter])
			currentLine = append(currentLine, words[iter])
			iter++
		}
		if iter != len(words) {
			result = append(result, justify(currentLine, wordLength))
		} else {
			lastLine := ""
			spaceCount := maxWidth - wordLength
			for i := 0; i < len(currentLine)-1; i++ {
				lastLine += currentLine[i]
				lastLine += strings.Repeat(" ", 1)
				spaceCount -= 1
			}

			lastLine += currentLine[len(currentLine)-1]
			lastLine += strings.Repeat(" ", spaceCount)

			result = append(result, lastLine)
		}
	}

	return result
	`,"\n",""),"\t",""), " "), 45)
	for _,v:=range justified{
		fmt.Println(v)
	}
}
