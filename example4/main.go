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
	run("/tmp/person.proto", Homer, new(Person))
	run("/tmp/people.proto", Family, new(People))
	run("/tmp/5000people.proto", Family5000, new(People))
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

	err = readProto(filename, copiedRecord)
	if err != nil {
		log.Fatalf("Error reading person struct: %v", err)
	}

	fmt.Printf("File Name      :         %s\n", fileStat.Name()) // Base name of the file
	fmt.Printf("Size           :         %d\n", fileStat.Size()) // Length in bytes for regular files

	origPeople, ok := orig.(*People)
	if ok {
		copyPeople := copiedRecord.(*People)
		fmt.Printf("Original Record:         %d\n", len(origPeople.Members))
		fmt.Printf("Copied Record  :         %d\n", len(copyPeople.Members))
	} else {
		// single person
		fmt.Printf("Original Record:         1\n")
		fmt.Printf("Copied Record  :         1\n")
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file contents: %v", err)
	}

	if len(content) < 255 {
		fmt.Printf("\n\nContents (hex)\n%s\n", hex.Dump(content))
	}

	fmt.Printf("\n\n")
}
