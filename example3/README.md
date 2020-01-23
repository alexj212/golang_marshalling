#Xml encoding example
This example will marshal a struct to a file on the disk and also unmarshalling back to a struct. The data will be encoded in an xml format.


~~~~


type Person struct {
	FirstName     string `json:"first"`
	LastName      string `json:"last"`
	Email         string `json:"email"`
	City          string `json:"city"`
	Country       string `json:"country"`
	Age           int32  `json:"age"`
	OmittedInt    int32  `json:"OmittedInt,omitempty"`
	OmittedString string `json:"OmittedString,omitempty"`
	Excluded      string `json:"-"`
	private       string
}


var Homer = &Person{
	FirstName: "Homer",
	LastName:  "Simpson",
	Email:     "homer@simpsons.com",
	City:      "Springfield",
	Country:   "USA",
	Age:       36,
}


$ ./bin/json_example    
File Name      :         person.json
Size           :         110
Original Record:         Homer Simpson - email: homer@simpsons.com city: Springfield country: USA age: 36
Copied Record  :         Homer Simpson - email: homer@simpsons.com city: Springfield country: USA age: 36


Contents
{"first":"Homer","last":"Simpson","email":"homer@simpsons.com","city":"Springfield","country":"USA","age":36}



File Name      :         people.json
Size           :         560
Original Record:         members: 5
Copied Record  :         members: 5


File Name      :         5000people.json
Size           :         546014
Original Record:         members: 5000
Copied Record  :         members: 5000



	
~~~~

Struct values encode as JSON objects. Each exported struct field becomes a member of the object unless

the field's tag is "-", or
the field is empty and its tag specifies the "omitempty" option.
The empty values are false, 0, any nil pointer or interface value, and any array, slice, map, or string of length zero.



### Writing struct to file    
         
    func writeJson(filePath string, object interface{}) error {
        file, err := os.Create(filePath)
        if err != nil {
            return err
        }
    
        jsonWriter := io.Writer(file)
        enc := json.NewEncoder(jsonWriter)
        enc.SetIndent("", "    ")
        if err := enc.Encode(object); err != nil {
            return err
        }
    
        err = file.Close()
        return err
    }


    
### Reading file and marshalling to struct    
           
    func readJson(filePath string, object interface{}) error {
        file, err := os.Open(filePath)
        if err != nil {
            return err
        }
    
        rawJsonBytes, _ := ioutil.ReadAll(file)
        json.Unmarshal(rawJsonBytes, object)
    
        err = file.Close()
        return err
    }