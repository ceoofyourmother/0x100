package main

import (
	"bufio"
	"fmt"
	md "github.com/nao1215/markdown"
	"os"
	"strings"
	"time"
)

func main() {
	args := bufio.NewReader(os.Stdin)

	fmt.Println("Qual o titulo do post?")
	titleWithSpace, _ := args.ReadString('\n')
	titleWithoutSpace := strings.TrimSpace(titleWithSpace)
	title := strings.ReplaceAll(titleWithoutSpace, " ", "_")

	timeNow := time.Now()
	year := timeNow.Year()
	month := timeNow.Month()
	day := timeNow.Day()

	pathPost := fmt.Sprintf("../content/%d/%d/%d/%s", year, month, day, title)

	err := os.MkdirAll(pathPost, os.ModePerm)

	if err != nil {
		panic(err)
	}

	f, err := os.Create(fmt.Sprintf("%s/index.md", pathPost))

	content := md.NewMarkdown(f)
	content.PlainText("---")
	content.PlainText(fmt.Sprintf("title: %s", title))
	content.PlainText(fmt.Sprintf("date: %s", timeNow.Format(time.RFC3339)))
	content.PlainText("---")

	content.Build()

	if err != nil {
		panic(err)
	}

}
