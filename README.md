## Golang Data Marshalling


###Building Examples
this project provides a Makefile, to perform many common tasks such as running tests, running the formatter, and building examples.   
    
~~~~
$ make
build_info                     Build the container
help                           This help.
clean                          clean all binaries in bin dir
clean_binary                   clean binary in bin dir
test                           run tests
fmt                            run fmt
all                            build all
~~~~
  
* [Gob example](./example1/README.md)
* [XML example](./example2/README.md)
* [Json example](./example3/README.md)
* [Protobuf example](./example4/README.md)



### Encoding Sizes

|           |Encoding |Single Record   |3 Records  |   |
|-----------|---------|----------------|-----------|---|
|example1   |gob      |152             |342        |   |   
|example2   |xml      |217             |776        |   |   
|example3   |json     |160             |699        |   |   
|example4   |protobuf |61              |196        |   |   