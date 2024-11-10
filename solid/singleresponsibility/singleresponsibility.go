package main

import (
	"fmt"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++

	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	//
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// separateion of concerns

func (j *Journal) Save(fileName string) {
	_ = os.WriteFile(fileName, []byte(j.String()), 0644)
}

func (j *Journal) Load(fileName string) {
}

func (j *Journal) LoadFromWeb(fileName string) {
}

var LineSeparator = "\n"

func SaveToFile(j *Journal, fileName string) {
	_ = os.WriteFile(fileName, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, fileName string) {
	_ = os.WriteFile(fileName, []byte(strings.Join(j.entries, p.LineSeparator)), 0644)
}

func main() {
	j := &Journal{}
	j.AddEntry("i cried")
	j.AddEntry("i ate")

	fmt.Println(j.String())

	SaveToFile(j, "journal.txt")

	p := Persistence{"\r\n"}
	p.SaveToFile(j, "journal1.txt")
}
