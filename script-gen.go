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
	commitCount = 50

	printCountLimit = 255

	messageLengthMin = 6
	messageLengthMax = 20

	commitIntervalMin = 12 * 60 * 60
	commitIntervalMax = 24 * 60 * 60

	sourcePrefix = "src"
	outputDir    = "output"
)

func main() {
	timeNow := time.Now()
	rand.Seed(time.Now().UnixNano())

	codeCache := make(map[string]string)

	mkdir()

	scriptOut, err := os.OpenFile(outputDir+"/commit.sh", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("Error opening script file: %s\n", err)
		panic(err)
	}
	defer scriptOut.Close()

	scriptCommands := make([]string, 0)

	// Generate 700 commits
	for i := 0; i < commitCount; i++ {
		// What is my language today?
		language, path := getLanguage()

		// Time travel back 12 ~ 24 hours
		timeNow = timeNow.Add(-time.Duration(rand.Intn(commitIntervalMax-commitIntervalMin)+commitIntervalMin) * time.Second)

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

		if language == "java" || language == "scala" || language == "csharp" {
			fullCode = strings.Replace(fullCode, "CLASS_NAME", message, 1)
		}

		commitMessage := fmt.Sprintf("Print %s for %d times, in %s", message, repeatCount, language)

		languageSuffix := path[strings.LastIndex(path, ".")+1:]

		// Write code to disk
		err = ioutil.WriteFile(fmt.Sprintf("%s/%s/%s.%s", outputDir, language, message, languageSuffix), []byte(fullCode), 0644)
		if err != nil {
			fmt.Printf("Failed to write %s/%s.%s\n", outputDir, message, languageSuffix)
			panic(err)
		}

		// store commit command
		var command string
		command += fmt.Sprintf("git add %s.%s\n", message, languageSuffix)
		command += fmt.Sprintf("git commit -m \"%s\" --date %d\n\n", commitMessage, timeNow.Unix())
		scriptCommands = append(scriptCommands, command)
	}

	// Write script to disk
	for i := len(scriptCommands) - 1; i >= 0; i-- {
		_, err = scriptOut.WriteString(scriptCommands[i])
		if err != nil {
			fmt.Printf("Failed to write script\n")
			panic(err)
		}
	}
	if err != nil {
		fmt.Printf("Failed to write script file: %s\n", err)
		panic(err)
	}
	fmt.Printf("Wrote %d commits to script file\n", len(scriptCommands))
}

func randomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	result[0] = upperChars[rand.Intn(len(upperChars))]
	for i := 1; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func mkdir() {
	var languages = getLanguageList()
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, 0755)
	}

	for _, language := range languages {
		if _, err := os.Stat(fmt.Sprintf("%s/%s", outputDir, language)); os.IsNotExist(err) {
			os.Mkdir(fmt.Sprintf("%s/%s", outputDir, language), 0755)
		}
	}
}

func getLanguage() (lang, path string) {
	var languages = getLanguageList()
	var languageSuffix = getSuffixList()
	idx := rand.Intn(len(languages))
	return languages[idx], fmt.Sprintf("simple%s", languageSuffix[idx])
}

func getLanguageList() []string {
	return []string{"c", "cpp", "csharp", "go", "java", "javascript", "python", "ruby", "swift", "typescript", "php", "kotlin", "scala", "haskell", "rust", "lisp", "lua", "bash"}
}

func getSuffixList() []string {
	return []string{".c", ".cpp", ".cs", ".go", ".java", ".js", ".py", ".rb", ".swift", ".ts", ".php", ".kt", ".scala", ".hs", ".rs", ".lisp", ".lua", ".sh"}
}
