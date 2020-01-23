## Golang Data Marshalling
This repo is a collection of examples of marshalling golang structs to a file. The examples use several formats such as GOB, XML, Json and Protocol Buffers and will encode and decode a struct to a []byte, and write that to disk.   

[Slide Deck](https://docs.google.com/presentation/d/e/2PACX-1vQbl58wkTwbOwhlQXfxNgcSkMdmiazT-4MGkF-_x5B0fiBYzBUaFr7kNdbxNy7dgQDmrcNvJpg28e17/pub?start=false&loop=false&delayms=3000)

First presented to the Go Miami/FLL Meetup on Jan 23, 2020. 



###Building Examples
this project provides a Makefile, to perform many common tasks such as running tests, running the formatter, and building examples.   
    
~~~~
$ make
all                            build all examples into ./bin dir.
build_info                     Display environment variables
clean                          clean all binaries in bin dir
clean_binary                   clean binary in bin dir
fmt                            run fmt
gob_example                    Build gob_example binary in bin dir.
help                           This help.
json_example                   Build json_example binary in bin dir
protobuf_example               Build protobuf_example binary in bin dir
test                           run tests
xml_example                    Build xml_example binary in bin dir
~~~~
  
* [Gob example](./example1/README.md)
* [XML example](./example2/README.md)
* [Json example](./example3/README.md)
* [Protobuf example](./example4/README.md)



### Encoding Sizes

|           |Encoding |Single Record   |5 Records  | 5000 Records  |
|-----------|---------|----------------|-----------|---------------|
|example1   |gob      |147 bytes       |442 bytes  | 283162 bytes  |   
|example2   |xml      |115             |598        | 581017        |   
|example3   |json     |110             |560        | 546014        |   
|example4   |protobuf |56              |288        | 288000        |   