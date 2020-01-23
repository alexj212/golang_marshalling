package main

import (
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	gob.Register(Person{})
	gob.Register(People{})
}

func writeGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(object)
	if err != nil {
		return err
	}

	err = file.Close()
	return err
}

func readGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(object)
	if err != nil {
		return err
	}

	err = file.Close()
	return err
}

func main() {
	run("/tmp/person.gob", Homer, new(Person))
	run("/tmp/people.gob", Family, new(People))
	run("/tmp/5000people.gob", Family5000, new(People))
}

func run(filename string, orig interface{}, copiedRecord interface{}) {

	var err error
	err = writeGob(filename, orig)
	if err != nil {
		log.Fatalf("Error writing person struct: %v", err)
	}

	fileStat, err := os.Stat(filename)
	if err != nil {
		log.Fatalf("Error reading file details: %v", err)
	}

	err = readGob(filename, copiedRecord)
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
		fmt.Printf("\n\nContents (hex)\n%s\n", hex.Dump(content))
	}

	fmt.Printf("\n\n")
}
