package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/theckman/yacspin"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

var header = "\t\t\tRepresentations Board Meeting"

type AgendaItem struct {
	Author     string
	Title      string
	LogInDate  string
	Reader     string
	Todo       string
	Circulate  bool
	Decision   string
	ActionDate string
	LastName   string
	FirstName  string
	Salutation string
	Email      string
	FileName   string
	Notes      string
}

var circulationAffirmation = "yes"
var documentTitle = "\t\t\t\tRepresentations Board Meeting\n"
var datePlaceholder = "\t\t\t\t\t- FILL IN DATE -\n\n\n\n"
var readForMeetingPlaceholder = "Read for Meeting\n\n"
var preliminaryReportPlaceholder = "Preliminary Report\n\n"

func formatData(data [][]string) []AgendaItem {
	var agenda []AgendaItem
	for i, line := range data {
		notHeaderRow := i >= 1
		if notHeaderRow {
			var item AgendaItem
			for index, value := range line {
				switch index {
				case 0:
					item.Author = value
				case 1:
					item.Title = value
				case 2:
					item.LogInDate = value
				case 3:
					item.Reader = value
				case 4:
					item.Todo = value
				case 5:
					circulate := true
					if strings.ToLower(value) != circulationAffirmation {
						circulate = false
					}
					item.Circulate = circulate
				case 6:
					item.Decision = value
				case 7:
					item.ActionDate = value
				case 8:
					item.LastName = value
				case 9:
					item.FirstName = value
				case 10:
					item.Salutation = value
				case 11:
					item.Email = value
				case 12:
					item.FileName = value
				case 13:
					item.Notes = value
				}
			}
			agenda = append(agenda, item)
		}
	}
	return agenda
}

func readFile(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func writeTitle(f *os.File) {
	f.WriteString(documentTitle)
}

func writeDatePlaceholder(f *os.File) {
	f.WriteString(datePlaceholder)
}

func writeAuthor(author string, f *os.File) {
	output := fmt.Sprintf("%s\n", author)
	f.WriteString(output)
}

func writeBookTitle(title string, f *os.File) {
	output := fmt.Sprintf("\t\"%s\"\n", title)
	f.WriteString(output)
}

func buildTabString(tabNumber int) string {
	tabs := ""
	for i := 0; i <= tabNumber; i++ {
		tabs += "\t"
	}
	return tabs
}

func calculateMaxTabs(len float64) int {
	// Max tab size per line that keeps Login date aligned is 7.
	tabSize := 7.0
	characterCount := math.Ceil(len / 5)
	return int(tabSize - characterCount)
}

func writeReaderAndLoginDate(reader string, loginDate string, f *os.File) {
	maxTabs := calculateMaxTabs(float64(len(reader)))
	tabs := buildTabString(maxTabs)
	output := fmt.Sprintf("%sLogin date: %s\n\n", tabs, loginDate)

	if reader != "" {
		output = fmt.Sprintf("\t%s%sLogin date: %s\n\n", reader, tabs, loginDate)
	}

	f.WriteString(output)
}

func writeAgendaItem(item AgendaItem, f *os.File) {
	writeAuthor(item.Author, f)
	writeBookTitle(item.Title, f)
	writeReaderAndLoginDate(item.Reader, item.LogInDate, f)
}

func writeSection(items []AgendaItem, f *os.File, title string, condition bool) {
	f.WriteString(title)
	for _, item := range items {
		if item.Circulate == condition {
			writeAgendaItem(item, f)
		}
	}
}

func writeAgendaFile(items []AgendaItem) {
	now := time.Now()
	agendaFileName := fmt.Sprintf("Agenda-%s.txt", now.Format("2006-01-02"))
	f, err := os.Create(agendaFileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	writeTitle(f)
	writeDatePlaceholder(f)
	writeSection(items, f, readForMeetingPlaceholder, true)
	writeSection(items, f, preliminaryReportPlaceholder, false)
}

func main() {
	csvName := flag.String("file", "data.csv", "Name of the CSV file to read. Default: data.csv")
	flag.Parse()
	cfg := yacspin.Config{
		Frequency:       100 * time.Millisecond,
		CharSet:         yacspin.CharSets[11],
		Suffix:          " Agenda Maker",
		SuffixAutoColon: true,
		Message:         fmt.Sprintf("reading %s data", *csvName),
		StopCharacter:   "✓",
		StopColors:      []string{"fgGreen"},
	}

	spinner, err := yacspin.New(cfg)
	if err != nil {
		return
	}

	spinner.Start()
	time.Sleep(2 * time.Second)
	fileData := readFile(*csvName)
	agendaItems := formatData(fileData)
	spinner.Message("writing agenda file")
	time.Sleep(2 * time.Second)
	writeAgendaFile(agendaItems)
	spinner.Stop()
}
