package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	xlsx "github.com/tealeg/xlsx"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: git-xlsx-textconv file.xlsx")
	}
	excelFileName := os.Args[1]

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal(err)
	}

	for _, sheet := range xlFile.Sheets {
		for k, row := range sheet.Rows {
			cels := make([]string, len(row.Cells))
			for i, cell := range row.Cells {
				s := cell.String()
				s = strings.Replace(s, "\\", "\\\\", -1)
				s = strings.Replace(s, "\n", "\\n", -1)
				s = strings.Replace(s, "\r", "\\r", -1)
				s = strings.Replace(s, "\t", "\\t", -1)
				cels[i] = s
			}
			fmt.Printf("[%s:%d] %s\n", sheet.Name, k+1, strings.Join(cels, "\t"))
		}
	}
}
