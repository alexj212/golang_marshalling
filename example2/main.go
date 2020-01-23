package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func init() {

}

func writeXml(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	xmlWriter := io.Writer(file)
	enc := xml.NewEncoder(xmlWriter)
	//enc.Indent("  ", "    ")
	if err := enc.Encode(object); err != nil {
		return err
	}

	err = file.Close()
	return err
}

func readXml(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	rawXmlBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(rawXmlBytes, object)
	if err != nil {
		return err
	}

	err = file.Close()
	return err
}

func main() {
	run("/tmp/person.xml", Homer, new(Person))
	run("/tmp/people.xml", Family, new(People))
	run("/tmp/5000people.xml", Family5000, new(People))
}

func run(filename string, orig interface{}, copiedRecord interface{}) {

	var err error
	err = writeXml(filename, orig)
	if err != nil {
		log.Fatalf("Error writing person struct: %v", err)
	}

	fileStat, err := os.Stat(filename)
	if err != nil {
		log.Fatalf("Error reading file details: %v", err)
	}

	err = readXml(filename, copiedRecord)
	if err != nil {
		log.Fatalf("Error reading person struct: %v", err)
	}

	fmt.Printf("File Name      :         %s\n", fileStat.Name()) // Base name of the file
	fmt.Printf("Size           :         %d\n", fileStat.Size()) // Length in bytes for regular files
	fmt.Printf("Original Record:         %v\n", orig)
	fmt.Printf("Copied Record  :         %v\n", copiedRecord)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file contents: %v", err)
	}
	// Convert []byte to string and print to screen
	text := string(content)
	if len(content) < 255 {
		fmt.Printf("\n\nContents\n%s\n", text)
	}

	fmt.Printf("\n\n")
}
