#Xml encoding example
This example will marshal a struct to a file on the disk and also unmarshalling back to a struct. The data will be encoded in an xml format.


~~~~


message Person {
    string FirstName = 1;
    string LastName = 2;
    string Email = 3;
    string City = 4;
    string Country = 5;
    int32 Age = 6;
}


type Person struct {
	FirstName string `protobuf:"bytes,1,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=LastName" json:"LastName,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=Email" json:"Email,omitempty"`
	City      string `protobuf:"bytes,4,opt,name=City" json:"City,omitempty"`
	Country   string `protobuf:"bytes,5,opt,name=Country" json:"Country,omitempty"`
	Age       int32  `protobuf:"varint,6,opt,name=Age" json:"Age,omitempty"`
}




var Homer = &Person{
	FirstName: "Homer",
	LastName:  "Simpson",
	Email:     "homer@simpsons.com",
	City:      "Springfield",
	Country:   "USA",
	Age:       36,
}




$ ./bin/protobuf_example     
File Name      :         person.proto
Size           :         56
Original Record:         1
Copied Record  :         1


Contents (hex)
00000000  0a 05 48 6f 6d 65 72 12  07 53 69 6d 70 73 6f 6e  |..Homer..Simpson|
00000010  1a 12 68 6f 6d 65 72 40  73 69 6d 70 73 6f 6e 73  |..homer@simpsons|
00000020  2e 63 6f 6d 22 0b 53 70  72 69 6e 67 66 69 65 6c  |.com".Springfiel|
00000030  64 2a 03 55 53 41 30 24                           |d*.USA0$|



File Name      :         people.proto
Size           :         288
Original Record:         5
Copied Record  :         5


File Name      :         5000people.proto
Size           :         288000
Original Record:         5000
Copied Record  :         5000
	
~~~~



### Writing struct to file    
       
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



    
### Reading file and marshalling to struct    
          
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
