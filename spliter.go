package main

import (
	"fmt"
	"flag"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"./config"
	"./const"
	"./utils"
)

func main() {

	tar := config.Conf{}

	path := flag.String("p", "./sheet.xlsx", "path for given xlsx file to split from")
	length := flag.Int("l", 100, "length for a splitting")
	head := flag.Bool("head", false, "whether to keep head to per file split by length")
	sheet := flag.String("s", "sheet1", "what sheet you want to split at")
	target := flag.String("t", "./doc", "where the output files targeting at")
	config := flag.String("config", "./conf.yaml", "where the config file is at")

	flag.Parse()

	if *config != "" {
		d, err := ioutil.ReadFile(*config)
		if err != nil {
			fmt.Println(err)
		}
		err = yaml.Unmarshal([]byte(d), &tar)
		if err != nil {
			fmt.Println(err)
		}
		*head = tar.Head
		*path = tar.File
		*sheet = tar.SourceSheet
		*length = tar.Length
		*target = tar.TargetPath

		*head = tar.Head
		*head = tar.Head
	}

	fmt.Println("path:", *path)
	fmt.Println("length:", *length)
	fmt.Println("head:", *head)
	fmt.Println("sheet:", *sheet)
	fmt.Println("target:", *target)
	fmt.Println("config", *config)

	xlsx, err := excelize.OpenFile(*path)

	if err != nil {
		fmt.Println(err)
		return
	}

	rows := xlsx.GetRows(*sheet)
	header := rows[0]

	fmt.Println("")

	sum := (len(rows))/(*length) + 1
	if *head {
		sum = (len(rows)-1)/(*length) + 1
	}
	fmt.Println("sum is ", len(rows))
	fmt.Println("split to ", sum, "files")

	files := [][][]string{}

	if *head {
		for i, v := range header {
			if i == 0 {
				fmt.Print("head of this file is : ")
			}
			fmt.Print(v, "\t")
		}

		fmt.Println("")
		fmt.Println("begin splitting!")

		rowrows := rows[1:]
		// begin split
		for index, row := range rowrows {
			i := index / (*length)
			if len(files) == i {
				fmt.Println("pending cache page : ", i)
				files = append(files, [][]string{header})
			}
			single := files[i]
			single = append(single, row)
			files[i] = single
			if index == len(rowrows)-1 {
				fmt.Println("last one is ", row, " at page", i+1)
			}
		}
	} else {

	}

	val := 0
	for _, v := range files {
		val += len(v)
	}

	fmt.Println(len(files))
	fmt.Println("all lines are sum to : ", val)

	for i, v := range files {
		sheetname := tar.Output.Sheet
		fmt.Println("now file ", i+1)
		name := fmt.Sprintf("%s-%d", "split", i+1)
		xlsx := excelize.NewFile()
		index := xlsx.NewSheet(sheetname)
		if sheetname != _const.DEFAULT_SHEET_NAME {
			xlsx.DeleteSheet(_const.DEFAULT_SHEET_NAME)
		}
		for l, row := range v {
			for x, value := range row {
				xAxis := fmt.Sprintf("%s%d", _const.UnitMap[x], l+1)
				xlsx.SetCellValue(sheetname, xAxis, value)
			}
		}
		xlsx.SetActiveSheet(index)

		//find if dir exist
		ok, _ := utils.PathExists(*target)
		if !ok {
			err := os.Mkdir(*target, os.ModePerm)
			fmt.Println(err)
		}

		err := xlsx.SaveAs(*target + "/" + name + ".xlsx")
		if err != nil {
			fmt.Println(err)
		}
	}
}
