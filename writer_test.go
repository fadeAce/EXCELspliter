package main

import (
	"testing"
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
)

func TestPathWrt(t *testing.T) {
	xlsx := excelize.NewFile()
	// Create a new sheet.
	index := xlsx.NewSheet("Sheet2")
	// Set value of a cell.
	xlsx.SetCellValue("Sheet2", "A2", "Hello -- world.")
	xlsx.SetCellValue("Sheet1", "B2", "100")
	xlsx.SetSheetRow("Sheet2", "", []string{"a", "b", "c"})
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
