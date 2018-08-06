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
	log "github.com/sirupsen/logrus"
)

func main() {

	tar := config.Conf{}

	path := flag.String("p", "./sheet.xlsx", "path for given xlsx file to split from")
	length := flag.Int("l", 100, "length for a splitting")
	head := flag.Bool("head", false, "whether to keep head to per file split by length")
	sheet := flag.String("s", "sheet1", "what sheet you want to split at")
	target := flag.String("t", "./doc", "where the output files targeting at")
	config := flag.String("config", "./conf.yaml",
	`
			# path for taget file
			path: ./lzy.xlsx
			
			# whether to involved header for file in
			head: true
			
			# length for a splitting
			length: 250
			
			#  where the output files targeting at
			target: ./target
			
			# what sheet you want to split at
			sheet: 1
			
			# ---- advanced configs
			
			# output sheets
			output:
			  # output sheet after splitting
			  sheet: one
			
			default:
			  # when mastering blank space
			  blank: false
		  `	,
	)

	flag.Parse()

	if *config != "" {
		d, err := ioutil.ReadFile(*config)
		if err != nil {
			log.Info(err)
		}
		err = yaml.Unmarshal([]byte(d), &tar)
		if err != nil {
			log.Info(err)
		}
		*head = tar.Head
		*path = tar.File
		*sheet = tar.SourceSheet
		*length = tar.Length
		*target = tar.TargetPath

		*head = tar.Head
		*head = tar.Head
	}

	log.Info("path:", *path)
	log.Info("length:", *length)
	log.Info("head:", *head)
	log.Info("sheet:", *sheet)
	log.Info("target:", *target)
	log.Info("config", *config)

	xlsx, err := excelize.OpenFile(*path)

	if err != nil {
		log.Info(err)
		return
	}

	rows := xlsx.GetRows(*sheet)
	header := rows[0]

	log.Info("")

	sum := (len(rows))/(*length) + 1
	if *head {
		sum = (len(rows)-1)/(*length) + 1
	}
	log.Info("sum is ", len(rows))
	log.Info("split to ", sum, "files")

	files := [][][]string{}

	if *head {
		for i, v := range header {
			if i == 0 {
				fmt.Print("head of this file is : ")
			}
			fmt.Print(v, "\t")
		}

		log.Info("")
		log.Info("begin splitting!")

		rowrows := rows[1:]
		// begin split
		for index, row := range rowrows {
			i := index / (*length)
			if len(files) == i {
				log.Info("pending cache page : ", i)
				files = append(files, [][]string{header})
			}
			single := files[i]
			single = append(single, row)
			files[i] = single
			if index == len(rowrows)-1 {
				log.Info("last one is ", row, " at page", i+1)
			}
		}
	} else {

	}

	val := 0
	for _, v := range files {
		val += len(v)
	}

	log.Info(len(files))
	log.Info("all lines are sum to : ", val)

	for i, v := range files {
		sheetname := tar.Output.Sheet
		log.Info("now file ", i+1)
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
			log.Info(err)
		}

		err := xlsx.SaveAs(*target + "/" + name + ".xlsx")
		if err != nil {
			log.Info(err)
		}
	}
}
