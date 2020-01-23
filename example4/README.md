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




var Alex = &Person{
	FirstName: "Alex",
	LastName:  "Jeannopoulos",
	Email:     "alexj@backpocket.com",
	City:      "Plantation",
	Country:   "USA",
	Age:       48,
}

File Name:         person.proto
Size:              61
Permissions:       -rw-r--r--
Last Modified:     2020-01-23 00:30:24.480701141 -0500 EST
Original Record:              FirstName:"Alex" LastName:"Jeannopoulos" Email:"alexj@backpocket.com" City:"Plantation" Country:"USA" Age:48 
Copied Record  :              FirstName:"Alex" LastName:"Jeannopoulos" Email:"alexj@backpocket.com" City:"Plantation" Country:"USA" Age:48 


Contents (hex)
00000000  0a 04 41 6c 65 78 12 0c  4a 65 61 6e 6e 6f 70 6f  |..Alex..Jeannopo|
00000010  75 6c 6f 73 1a 14 61 6c  65 78 6a 40 62 61 63 6b  |ulos..alexj@back|
00000020  70 6f 63 6b 65 74 2e 63  6f 6d 22 0a 50 6c 61 6e  |pocket.com".Plan|
00000030  74 61 74 69 6f 6e 2a 03  55 53 41 30 30           |tation*.USA00|

	
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
