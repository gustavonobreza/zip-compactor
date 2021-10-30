package main

import (
	"archive/zip"
	"io"
	"os"

	"github.com/ncruces/zenity"
)

func main() {
	selected, _ := zenity.SelectFileMutiple(zenity.Title("Selecione um item para compactar"))

	if len(selected) == 0 {
		os.Exit(0)
	}
	for _, val := range selected {
		print(val)
	}

	target, _ := zenity.SelectFileSave(zenity.Title("Selecione o destino"), zenity.FileFilter{Patterns: []string{"*.zip", "*.ZIP"}})

	if len(target) == 0 {
		os.Exit(0)
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
	// str := reflect.TypeOf(msg).Kind() != "string"
	if err != nil {
		// if str {
		// 	panic(msg)
		// }
		panic(err.Error())
	}
}
