package readers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"library-management-system-cli/utils"
	"log"
	"os"
	"strings"
)

const fileName string = "readers.json"

// Readers struct which contains
// an array of readers
type Readers struct {
	Readers []Reader `json:"readers"`
}

// Reader struct which contains the Reader data
type Reader struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Birthday   string `json:"birthday"`
	Height     string `json:"height"`
	Weight     string `json:"weight"`
	Employment string `json:"employment"`
}

// PrintAll ...
func (readers Readers) PrintAll() {
	for i := 0; i < len(readers.Readers); i++ {
		fmt.Println(readers.Readers[i].ToString())
	}
}

func getJSONstring() string {
	err := utils.CheckFile(fileName)
	if err != nil {
		log.Println(err)
	}

	// Open our jsonFile
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened", fileName)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return string(byteValue)
}

// GetAllReaders ...
func GetAllReaders() Readers {
	err := utils.CheckFile(fileName)
	if err != nil {
		log.Println(err)
	}

	// Open our jsonFile
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened", fileName)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// fmt.Println("byteValue:", string(byteValue))
	// we initialize our Readers array
	var readers Readers

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'readers' which we defined above
	json.Unmarshal(byteValue, &readers)

	return readers
}

// ToString ...
func (r Reader) ToString() string {
	return "ID: " + r.ID + ", " + "Name: " + r.Name + ", " + "Gender: " + r.Gender + ", " + "Birthday: " + r.Birthday + ", " + "Height: " + r.Height + ", " + "Weight: " + r.Weight + ", " + "Employment: " + r.Employment
}

// InsertReader function to insert a Reader readers.json
func (r Reader) InsertReader(readers *Readers) {
	(*readers).Readers = append((*readers).Readers, r)
	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(*readers)
	if err != nil {
		log.Println(err)
	}

	err = utils.CheckFile(fileName)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(fileName, dataBytes, 0644)
	if err != nil {
		log.Println(err)
	}
}

func removeIndex(s []Reader, index int) []Reader {
	return append(s[:index], s[index+1:]...)
}

// RemoveReader function to remove a Reader readers.json
func (r Reader) RemoveReader(readers *Readers) {
	var index int = 0
	for index = 0; index < len(readers.Readers); index++ {
		if readers.Readers[index].ID == r.ID {
			break
		}
	}
	(*readers).Readers = removeIndex((*readers).Readers, index)

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(*readers)
	if err != nil {
		log.Println(err)
	}

	err = utils.CheckFile(fileName)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(fileName, dataBytes, 0644)
	if err != nil {
		log.Println(err)
	}
}

// GetReaderByID ...
func (readers Readers) GetReaderByID(id string) (Reader, int) {
	for i := 0; i < len(readers.Readers); i++ {
		if readers.Readers[i].ID == id {
			return readers.Readers[i], 1
		}
	}
	return Reader{}, 0
}

// GetReaderByName ...
func (readers Readers) GetReaderByName(name string) (Reader, int) {
	for i := 0; i < len(readers.Readers); i++ {
		if strings.ToLower(readers.Readers[i].Name) == strings.ToLower(name) {
			return readers.Readers[i], 1
		}
	}
	return Reader{}, 0
}

func main() {
	readers := GetAllReaders()

	readers.PrintAll()
	fmt.Println("===================================")

	reader1, flag := readers.GetReaderByID("2")
	if flag != 1 {
		log.Println("ERROR in GetBookByID")
	} else {
		fmt.Println("Found:", reader1.ToString())
		reader1.RemoveReader(&readers)
		fmt.Println("===================================")
		readers.PrintAll()
	}
}
