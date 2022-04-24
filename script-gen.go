package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	commitCount = 700

	printCountLimit = 255

	messageLengthMin = 6
	messageLengthMax = 20

	sourcePrefix = "src"
	outputDir    = "output"
)

func main() {
	timeNow := time.Now()

	codeCache := make(map[string]string)

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, 0755)
	}

	scriptOut, err := os.OpenFile(outputDir+"/commit.sh", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("Error opening script file: %s\n", err)
		panic(err)
	}
	defer scriptOut.Close()

	// Generate 700 commits
	for i := 0; i < commitCount; i++ {
		// What is my language today?
		language, path := getLanguage()

		// Time travel back 12 ~ 24 hours
		timeNow = timeNow.Add(-time.Duration(rand.Intn(12*60*60)+12*60*60) * time.Second)

		// Generate a random message
		message := randomString(rand.Intn(messageLengthMax-messageLengthMin) + messageLengthMin)

		// Set repeat count
		repeatCount := rand.Intn(printCountLimit)

		// Read source code from cache, or disk
		var fullCode string
		if code, ok := codeCache[path]; ok {
			fullCode = fmt.Sprintf(code, repeatCount, message)
		} else {
			fb, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", sourcePrefix, path))
			if err != nil {
				fmt.Printf("%s/%s does not exist\n", sourcePrefix, path)
				panic(err)
			}
			codeCache[path] = string(fb)
			fullCode = fmt.Sprintf(string(fb), repeatCount, message)
		}

		if language == "java" || language == "scala" {
			fullCode = strings.Replace(fullCode, "CLASS_NAME", message, 1)
		}

		commitMessage := fmt.Sprintf("Print %s for %d times, in %s", message, repeatCount, language)

		languageSuffix := path[strings.LastIndex(path, ".")+1:]

		// Write code to disk
		err = ioutil.WriteFile(fmt.Sprintf("%s/%s.%s", outputDir, message, languageSuffix), []byte(fullCode), 0644)
		if err != nil {
			fmt.Printf("Failed to write %s/%s.%s\n", outputDir, message, languageSuffix)
			panic(err)
		}

		// Write commit to script
		scriptOut.WriteString(fmt.Sprintf("git add %s.%s\n", message, languageSuffix))
		scriptOut.WriteString(fmt.Sprintf("git commit -m \"%s\" --date %d\n", commitMessage, timeNow.Unix()))
		scriptOut.WriteString("\n")
	}

}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func getLanguage() (lang, path string) {
	var languages = []string{
		"c",
		"c++",
		"c#",
		"go",
		"java",
		"javascript",
		"python",
		"ruby",
		"swift",
		"typescript",
		"php",
		"kotlin",
		"scala",
		"haskell",
		"rust",
		"lisp",
		"lua",
		"bash",
	}
	var languageSuffix = []string{
		".c",
		".cpp",
		".cs",
		".go",
		".java",
		".js",
		".py",
		".rb",
		".swift",
		".ts",
		".php",
		".kt",
		".scala",
		".hs",
		".rs",
		".lisp",
		".lua",
		".sh",
	}
	idx := rand.Intn(len(languages))
	return languages[idx], fmt.Sprintf("simple%s", languageSuffix[idx])
}
