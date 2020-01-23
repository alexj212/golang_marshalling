#Xml encoding example
This example will marshal a struct to a file on the disk and also unmarshalling back to a struct. The data will be encoded in an xml format.


~~~~

type Person struct {
	FirstName string `xml:"first,attr"`
	LastName  string `xml:"last,attr"`
	Email     string `xml:"email,attr"`
	City      string `xml:"city,attr"`
	Country   string `xml:"country,attr"`
	Age       int32  `xml:"age,attr"`
}


var Alex = &Person{
	FirstName: "Alex",
	LastName:  "Jeannopoulos",
	Email:     "alexj@backpocket.com",
	City:      "Plantation",
	Country:   "USA",
	Age:       48,
}

File Name:         person.xml
Size:              122
Permissions:       -rw-rw-r--
Last Modified:     2020-01-23 00:27:00.557998203 -0500 EST
Original Record:              Alex Jeannopoulos - email: alexj@backpocket.com city: Plantation country: USA age: 48
Copied Record  :              Alex Jeannopoulos - email: alexj@backpocket.com city: Plantation country: USA age: 48


Contents
  <Person first="Alex" last="Jeannopoulos" email="alexj@backpocket.com" city="Plantation" country="USA" age="48"></Person>
	
~~~~



### Writing struct to file    
     
    func writeXml(filePath string, object interface{}) error {
        file, err := os.Create(filePath)
        if err != nil {
            return err
        }
    
        xmlWriter := io.Writer(file)
        enc := xml.NewEncoder(xmlWriter)
        enc.Indent("  ", "    ")
        if err := enc.Encode(object); err != nil {
            return err
        }
    
        err = file.Close()
        return err
    }

    
### Reading file and marshalling to struct    
       
    func readXml(filePath string, object interface{}) error {
        file, err := os.Open(filePath)
        if err != nil {
            return err
        }
    
        rawXmlBytes, _ := ioutil.ReadAll(file)
        xml.Unmarshal(rawXmlBytes, object)
    
        err = file.Close()
        return err
    }