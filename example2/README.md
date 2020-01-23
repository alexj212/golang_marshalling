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


var Homer = &Person{
	FirstName: "Homer",
	LastName:  "Simpson",
	Email:     "homer@simpsons.com",
	City:      "Springfield",
	Country:   "USA",
	Age:       36,
}



$ ./bin/xml_example    
File Name      :         person.xml
Size           :         115
Original Record:         Homer Simpson - email: homer@simpsons.com city: Springfield country: USA age: 36
Copied Record  :         Homer Simpson - email: homer@simpsons.com city: Springfield country: USA age: 36


Contents
<Person first="Homer" last="Simpson" email="homer@simpsons.com" city="Springfield" country="USA" age="36"></Person>


File Name      :         people.xml
Size           :         598
Original Record:         members: 5
Copied Record  :         members: 5


File Name      :         5000people.xml
Size           :         581017
Original Record:         members: 5000
Copied Record  :         members: 5000	
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