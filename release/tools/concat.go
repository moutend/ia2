package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	header = "api_all_headers.idl"
	idls   = []string{
		"IA2CommonTypes.idl",
		"AccessibleRelation.idl",
		"AccessibleAction.idl",
		"AccessibleRole.idl",
		"AccessibleStates.idl",
		"Accessible2.idl",
		"Accessible2_2.idl",
		"Accessible2_3.idl",
		"AccessibleComponent.idl",
		"AccessibleValue.idl",
		"AccessibleText.idl",
		"AccessibleText2.idl",
		"AccessibleEditableText.idl",
		"AccessibleHyperlink.idl",
		"AccessibleHypertext.idl",
		"AccessibleHypertext2.idl",
		"AccessibleTable.idl",
		"AccessibleTable2.idl",
		"AccessibleTableCell.idl",
		"AccessibleImage.idl",
		"AccessibleEventID.idl",
		"AccessibleApplication.idl",
		"AccessibleDocument.idl",
		"IA2TypeLibrary.idl",
	}
)

func main() {
	if len(os.Args) < 3 {
		return
	}

	basePath := os.Args[1]
	outputPath := os.Args[2]

	file, err := os.Create(outputPath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data, err := ioutil.ReadFile(filepath.Join(basePath, header))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(file, "%s\n", data)

	for _, idl := range idls {
		data, err := ioutil.ReadFile(filepath.Join(basePath, idl))

		if err != nil {
			log.Fatal(err)
		}

		for _, line := range strings.Split(string(data), "\n") {
			if strings.HasPrefix(line, "import") {
				continue
			}

			fmt.Fprintf(file, "%s\n", line)
		}
	}

	log.Println("-- concat.go done")
}

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}
