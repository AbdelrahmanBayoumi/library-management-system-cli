package readers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	readerModel "library-management-system-cli/model/readers"
	"net/http"
)

// ServeHTTP ...
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllReaders(w, r)
	case "POST":
		addReader(w, r)
	case "PUT":
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Not supported"}`))
	case "DELETE":
		w.Header().Set("Content-Type", "application/json")
		keys := r.URL.Query()
		id := keys.Get("id")
		if id == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "Not supported"}`))
			return
		}
		readers := readerModel.GetAllReaders()
		reader, flag := readers.GetReaderByID(id)
		if flag != 1 {
			// not found
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "No Reader Found With this id"}`))
		} else {
			reader.RemoveReader(&readers)
			json, err := json.Marshal(reader)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"message": "SERVER ERROR"}`))
				panic(err)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(string(json)))
			}
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

// SearchHandle ...
func SearchHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		keys := r.URL.Query()
		id := keys.Get("id")
		name := keys.Get("name")
		if id == "" && name == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "not found"}`))
			return
		} else if id != "" {
			// search by id and return reader
			w.Header().Set("Content-Type", "application/json")
			reader, flag := readerModel.GetAllReaders().GetReaderByID(id)
			if flag != 1 {
				// no id found
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"message": "no reader with this id"}`))
			} else {
				// id found
				json, _ := json.Marshal(reader)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(json))
			}
		} else if name != "" {
			// search by name and return reader
			w.Header().Set("Content-Type", "application/json")
			reader, flag := readerModel.GetAllReaders().GetReaderByName(name)
			if flag != 1 {
				// no id found
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"message": "no reader with this name"}`))
			} else {
				// id found
				json, _ := json.Marshal(reader)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(json))
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "not found"}`))
		}
	case "POST":
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	case "PUT":
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	case "DELETE":
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func addReader(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	if len(body) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "no body found"}`))
		return
	}
	var reader readerModel.Reader
	err := json.Unmarshal(body, &reader)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "SERVER ERROR"}`))
		panic(err)
	}
	readers := readerModel.GetAllReaders()
	reader.InsertReader(&readers)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Reader Added Successfully!"}`))
}

func getAllReaders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getAllReaders() ...")
	w.Header().Set("Content-Type", "application/json")
	readers := readerModel.GetAllReaders()
	json, err := json.Marshal(readers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "SERVER ERROR"}`))
		panic(err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(json))
	}
}
