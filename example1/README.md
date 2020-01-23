#Gob encoding example
This example will marshal a struct to a file on the disk and also unmarshalling back to a struct. The data will be encoded in the golang gob binary format.


	
~~~~
var Alex = &Person{
	FirstName: "Alex",
	LastName:  "Jeannopoulos",
	Email:     "alexj@backpocket.com",
	City:      "Plantation",
	Country:   "USA",
	Age:       48,
}




  File Name:         person.gob
  Size:              152
  Permissions:       -rw-rw-r--
  Last Modified:     2020-01-23 00:22:05.139410546 -0500 EST
  Original Record:              Alex Jeannopoulos - email: alexj@backpocket.com city: Plantation country: USA age: 48
  Copied Record  :              Alex Jeannopoulos - email: alexj@backpocket.com city: Plantation country: USA age: 48


  Contents (hex)
  00000000  56 ff 81 03 01 01 06 50  65 72 73 6f 6e 01 ff 82  |V......Person...|
  00000010  00 01 06 01 09 46 69 72  73 74 4e 61 6d 65 01 0c  |.....FirstName..|
  00000020  00 01 08 4c 61 73 74 4e  61 6d 65 01 0c 00 01 05  |...LastName.....|
  00000030  45 6d 61 69 6c 01 0c 00  01 04 43 69 74 79 01 0c  |Email.....City..|
  00000040  00 01 07 43 6f 75 6e 74  72 79 01 0c 00 01 03 41  |...Country.....A|
  00000050  67 65 01 04 00 00 00 40  ff 82 01 04 41 6c 65 78  |ge.....@....Alex|
  00000060  01 0c 4a 65 61 6e 6e 6f  70 6f 75 6c 6f 73 01 14  |..Jeannopoulos..|
  00000070  61 6c 65 78 6a 40 62 61  63 6b 70 6f 63 6b 65 74  |alexj@backpocket|
  00000080  2e 63 6f 6d 01 0a 50 6c  61 6e 74 61 74 69 6f 6e  |.com..Plantation|
  00000090  01 03 55 53 41 01 60 00                           |..USA.`.|
	
~~~~


    
### Registering structs with gob encoder/decoder    
    func init() {
        gob.Register(Person{})
        gob.Register(People{})
    }
    



### Writing struct to file    
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
    
### Reading file and marshalling to struct    
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