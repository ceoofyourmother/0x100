package main

import (
	"fmt"
	"io/fs"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	md "github.com/nao1215/markdown"
	"gopkg.in/yaml.v3"
)

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		panic(err)
	}
	if d.IsDir() {
		fmt.Printf("is a dir %s \n", d.Name())
	} else {
		fmt.Printf("is a file %s \n ", d.Name())
	}
	return nil
}

type markdown struct {
	Title string `yaml:"title"`
}

func escapeMarkdown(file string) string {

	content, err := os.ReadFile(fmt.Sprintf("%s/index.md", file))

	if err != nil {
		panic(err)
	}

	strings.ReplaceAll(string(content), "---", "")
	var contentStruct markdown
	err = yaml.Unmarshal(content, &contentStruct)
	if err != nil {
		fmt.Errorf("Erro ao inferir nome do post %s", err)
		panic(err)
	}
	return contentStruct.Title

}

func main() {
	f, _ := os.Create("../content/_index.md")
	defer f.Close()
	content := md.NewMarkdown(f)

	content.PlainText("---")
	content.PlainText("title: teste")
	content.PlainText("toc: false")
	content.PlainText("---")
	for i := 0; i < 2; i++ {
		content.PlainText("")
	}

	folders, _ := os.ReadDir("../content/")

	slices.Reverse(folders)
	for _, year := range folders {
		pathMontsFromYear, _ := os.ReadDir(fmt.Sprintf("../content/%s", year.Name()))
		for _, month := range pathMontsFromYear {
			var postsUrl []string
			urlPathBase := fmt.Sprintf("../content/%s/%s", year.Name(), month.Name())
			monthInt, _ := strconv.Atoi(month.Name())
			monthName := time.Month.String(time.Month(monthInt))
			content.H2f("%s - %s", year.Name(), monthName)
			days, _ := os.ReadDir(urlPathBase)
			for _, day := range days {
				posts, _ := os.ReadDir(fmt.Sprintf("%s/%s", urlPathBase, day.Name()))
				for _, post := range posts {
					urlPost := fmt.Sprintf("%s/%s/%s", urlPathBase, day.Name(), post.Name())
					postsUrl = append(postsUrl, urlPost)
				}
			}

			for _, post := range postsUrl {
				titlePost := escapeMarkdown(post)
				//TODO: dps arrumo esse split, o importante e funcionar!
				content.PlainText(fmt.Sprintf("- [%s](%s)", titlePost, strings.Split(post, "content")[1]))
			}
		}

	}

	content.Build()

}
