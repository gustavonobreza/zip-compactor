package main

import (
	"archive/zip"
	"flag"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/ncruces/zenity"
)

var fromFlag *string
var toFlag *string
var quietFlag *bool
var selected []string
var target string
var Version string // Dinamic, generated in build time.

func init() {
	// Flags
	fromFlag = flag.String("from", "", "path of the file to be ziped")
	toFlag = flag.String("to", "", "path to create the ziped file")
	quietFlag = flag.Bool("q", false, "quit, to not open the explorer after finished")
	vFlag := flag.Bool("v", false, "get version of software")
	versionFlag := flag.Bool("version", false, "get version of software")

	flag.Parse()

	if (*vFlag) || (*versionFlag) {
		println(Version)
		os.Exit(0)
	}
}

func main() {
	hasFF := len(*fromFlag) != 0
	hasTF := len(*toFlag) != 0

	if hasFF || hasTF {
		// Use CLI
		cliUsage(hasFF, hasTF)
	} else {
		// Use GUI
		guiUsage()
	}

	// Validate the given paths
	validateInputs()

	ZipItems(target, selected)

	openExplorer()
}

// Handle with CLI usage to deal with user inputs and determinate what is the target and the selecteds paths.
func cliUsage(hasFF, hasTF bool) {
	// If just one of two flags is given
	if hasFF && !hasTF || !hasFF && hasTF {
		println("The 'from' and 'to' flags are required!")
		os.Exit(1)
	}
	selected = []string{*fromFlag}
	target = *toFlag
}

// Handle with GUI usage to deal with user inputs and determinate what is the target and the selecteds paths.
func guiUsage() {
	selected, _ = zenity.SelectFileMutiple(zenity.Title("Selecione um item para compactar"))
	target, _ = zenity.SelectFileSave(
		zenity.Title("Selecione o destino"),
		zenity.FileFilter{Patterns: []string{"*.zip", "*.ZIP"}},
		zenity.ConfirmOverwrite(),
	)
}

// Handle with target input. If the target not have the suffix ".zip" it'll applied.
func validateInputs() {
	if len(selected) == 0 || len(target) == 0 {
		println("The file to be zipped and path to result are required!")
		os.Exit(1)
	}

	if !strings.HasSuffix(strings.ToLower(target), ".zip") {
		target += ".zip"
	}
}

// Handle with the files to be zipped.
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

// Handle with the action of zip a file.
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

// Handle with the action of open the explorer (windows).
func openExplorer() {
	if *quietFlag {
		return
	}

	if runtime.GOOS == "windows" {
		splitedTargetPath := strings.Split(target, string(os.PathSeparator))
		parentOfTarget := strings.Join(splitedTargetPath[0:len(splitedTargetPath)-1], string(os.PathSeparator))
		err := exec.Command("explorer", parentOfTarget).Run()

		if err != nil {
			zenity.CancelLabel("error: " + err.Error())
			os.Exit(0)
		}
	}
}

// he Handle with errors to evit the repetion of handling with the same code
func he(err error, msg ...interface{}) {
	if err != nil {
		panic(err.Error())
	}
}
