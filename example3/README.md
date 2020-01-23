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


var Alex = &Person{
	FirstName:     "Alex",
	LastName:      "Jeannopoulos",
	Email:         "alexj@backpocket.com",
	City:          "Plantation",
	Country:       "USA",
	Age:           48,
	OmittedInt:    156,
	OmittedString: "not omitted",
	Excluded:      "hide me",
}

File Name:         person.json
Size:              211
Permissions:       -rw-rw-r--
Last Modified:     2020-01-23 00:27:29.471034899 -0500 EST
Original Record:              Alex Jeannopoulos - email: alexj@backpocket.com city: Plantation country: USA age: 48
Copied Record  :              Alex Jeannopoulos - email: alexj@backpocket.com city: Plantation country: USA age: 48


Contents
{
    "first": "Alex",
    "last": "Jeannopoulos",
    "email": "alexj@backpocket.com",
    "city": "Plantation",
    "country": "USA",
    "age": 48,
    "OmittedInt": 156,
    "OmittedString": "not omitted"
}

	
~~~~



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