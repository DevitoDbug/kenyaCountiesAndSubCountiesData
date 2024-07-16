package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"os"
)

type CountyData struct {
	Code         string `json:"code"`
	County       string `json:"county"`
	HeadQuarters string `json:"headQuarters"`
	SubCounties  string `json:"subCounties"`
}

func main() {
	var CountiesData []CountyData
	var fileName = "countiesDataFile.json"
	c := colly.NewCollector()

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
			var rowData []string
			tr.ForEach("td", func(i int, td *colly.HTMLElement) {
				rowData = append(rowData, td.Text)
			})

			if len(rowData) >= 4 {
				county := CountyData{
					Code:         rowData[0],
					County:       rowData[1],
					HeadQuarters: rowData[2],
					SubCounties:  rowData[3],
				}
				CountiesData = append(CountiesData, county)
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://blog.afro.co.ke/list-of-counties/")

	//Creating file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error in writing file: ", err)
		return
	}
	defer func(file *os.File) {
		err2 := file.Close()
		if err2 != nil {
			fmt.Println("Error while closing file: ", err2)
		}
	}(file)

	//writing data to the file
	jsonData, err := json.MarshalIndent(CountiesData, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}
	err4 := os.WriteFile(fileName, jsonData, 0666)
	if err4 != nil {
		fmt.Println("Error when writing to file: ", err4)
	}

}
