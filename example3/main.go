package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func init() {

}

func writeJson(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	jsonWriter := io.Writer(file)
	enc := json.NewEncoder(jsonWriter)
	//enc.SetIndent("", "    ")
	if err := enc.Encode(object); err != nil {
		return err
	}

	err = file.Close()
	return err
}

func readJson(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	rawJsonBytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(rawJsonBytes, object)

	err = file.Close()
	return err
}

func main() {
	run("/tmp/person.json", Homer, new(Person))
	run("/tmp/people.json", Family, new(People))
	run("/tmp/5000people.json", Family5000, new(People))
}

func run(filename string, orig interface{}, copiedRecord interface{}) {

	var err error
	err = writeJson(filename, orig)
	if err != nil {
		log.Fatalf("Error writing person struct: %v", err)
	}

	fileStat, err := os.Stat(filename)
	if err != nil {
		log.Fatalf("Error reading file details: %v", err)
	}

	err = readJson(filename, copiedRecord)
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

	if len(content) < 255 {
		// Convert []byte to string and print to screen
		text := string(content)
		fmt.Printf("\n\nContents\n%s\n", text)
	}
	fmt.Printf("\n\n")
}
