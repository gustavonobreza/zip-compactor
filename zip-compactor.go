package main

import (
	"archive/zip"
	"flag"
	"io"
	"os"
	"strings"

	"github.com/ncruces/zenity"
)

func main() {
	fromFlag := flag.String("from", "", "path of the file to be ziped")
	toFlag := flag.String("to", "", "path to create the ziped file")

	flag.Parse()

	hasFF := len(*fromFlag) != 0
	hasTF := len(*toFlag) != 0

	if hasFF && !hasTF || !hasFF && hasTF {
		println("The 'from' and 'to' flags are required!")
		os.Exit(1)
	}

	var selected []string
	var target string

	if hasTF && hasFF {
		// Use CLI (just can select one file)
		selected = []string{*fromFlag}
		target = *toFlag
	} else {
		// Use GUI
		selected, _ = zenity.SelectFileMutiple(zenity.Title("Selecione um item para compactar"))
		target, _ = zenity.SelectFileSave(
			zenity.Title("Selecione o destino"),
			zenity.FileFilter{Patterns: []string{"*.zip", "*.ZIP"}},
			zenity.ConfirmOverwrite(),
		)
	}

	if len(selected) == 0 || len(target) == 0 {
		println("The file to be zipped and path to result are required!")
		os.Exit(1)
	}

	if !strings.HasSuffix(strings.ToLower(target), ".zip") {
		target += ".zip"
	}

	ZipItems(target, selected)
}

func ZipItems(filename string, items []string) {
	newZipFile, err := os.Create(filename)
	he(err)

	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, item := range items {
		if err := AddItemToZip(zipWriter, item); err != nil {
			he(err)
		}
	}
}

func AddItemToZip(zipWriter *zip.Writer, filename string) error {
	itemToZip, err := os.Open(filename)
	he(err)
	defer itemToZip.Close()

	info, err := itemToZip.Stat()
	he(err)

	header, err := zip.FileInfoHeader(info)
	he(err)

	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	he(err)
	_, err = io.Copy(writer, itemToZip)
	return err
}

// he Handle with errors to evit the repetion of handling with the same code
func he(err error, msg ...interface{}) {
	if err != nil {
		panic(err.Error())
	}
}
