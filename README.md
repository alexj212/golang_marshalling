## Golang Data Marshalling


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