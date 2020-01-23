#Gob encoding example
This example will marshal a struct to a file on the disk and also unmarshalling back to a struct. The data will be encoded in the golang gob binary format.


	
~~~~

var Homer = &Person{
	FirstName: "Homer",
	LastName:  "Simpson",
	Email:     "homer@simpsons.com",
	City:      "Springfield",
	Country:   "USA",
	Age:       36,
}



$ ./bin/gob_example         
File Name      :         person.gob
Size           :         147
Original Record:         Homer Simpson - email: homer@simpsons.com city: Springfield country: USA age: 36
Copied Record  :         Homer Simpson - email: homer@simpsons.com city: Springfield country: USA age: 36


Contents (hex)
00000000  56 ff 81 03 01 01 06 50  65 72 73 6f 6e 01 ff 82  |V......Person...|
00000010  00 01 06 01 09 46 69 72  73 74 4e 61 6d 65 01 0c  |.....FirstName..|
00000020  00 01 08 4c 61 73 74 4e  61 6d 65 01 0c 00 01 05  |...LastName.....|
00000030  45 6d 61 69 6c 01 0c 00  01 04 43 69 74 79 01 0c  |Email.....City..|
00000040  00 01 07 43 6f 75 6e 74  72 79 01 0c 00 01 03 41  |...Country.....A|
00000050  67 65 01 04 00 00 00 3b  ff 82 01 05 48 6f 6d 65  |ge.....;....Home|
00000060  72 01 07 53 69 6d 70 73  6f 6e 01 12 68 6f 6d 65  |r..Simpson..home|
00000070  72 40 73 69 6d 70 73 6f  6e 73 2e 63 6f 6d 01 0b  |r@simpsons.com..|
00000080  53 70 72 69 6e 67 66 69  65 6c 64 01 03 55 53 41  |Springfield..USA|
00000090  01 48 00                                          |.H.|



File Name      :         people.gob
Size           :         442
Original Record:         members: 5
Copied Record  :         members: 5


File Name      :         5000people.gob
Size           :         283162
Original Record:         members: 5000
Copied Record  :         members: 5000


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