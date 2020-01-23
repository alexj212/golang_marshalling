package main

import (
	"encoding/hex"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"os"
)

func writeProto(filePath string, object proto.Message) error {

	serializedBytes, err := proto.Marshal(object)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filePath, serializedBytes, 0644); err != nil {
		return err
	}

	return nil
}

func readProto(filePath string, object proto.Message) error {
	// Read the existing address book.
	in, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(in, object); err != nil {
		return err
	}

	return nil
}

func main() {
	run("/tmp/person.proto", Alex, new(Person))
	run("/tmp/people.proto", Family, new(People))
}

func run(filename string, orig proto.Message, copiedRecord proto.Message) {

	var err error
	err = writeProto(filename, orig)
	if err != nil {
		log.Fatalf("Error writing person struct: %v", err)
	}

	fileStat, err := os.Stat(filename)
	if err != nil {
		log.Fatalf("Error reading file details: %v", err)
	}

	fmt.Printf("File Name:         %s\n", fileStat.Name())    // Base name of the file
	fmt.Printf("Size:              %d\n", fileStat.Size())    // Length in bytes for regular files
	fmt.Printf("Permissions:       %s\n", fileStat.Mode())    // File mode bits
	fmt.Printf("Last Modified:     %v\n", fileStat.ModTime()) // Last modification time

	err = readProto(filename, copiedRecord)
	if err != nil {
		log.Fatalf("Error reading person struct: %v", err)
	}

	fmt.Printf("Original Record:              %v\n", orig)
	fmt.Printf("Copied Record  :              %v\n", copiedRecord)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file contents: %v", err)
	}
	fmt.Printf("\n\nContents (hex)\n%s\n\n", hex.Dump(content))
}
