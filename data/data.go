package data

import (
	"io/ioutil"
	"log"
)

// Data is an interface that is used in blockchain blocks.
type Data interface {
	Type() string
	Bytes() []byte
}
type stringData string

// String returns a string-like data object that can be used as data.
func String(data string) Data {
	return stringData(data)
}

// Document returns a data object read from the drive.
func Document(filepath string) Data {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Print("Could not load file.")
	}
	return stringData(content)
}

// Type returns the kind of data contained.
func (sd stringData) Type() string {
	return "string"
}

// Bytes returns the data as bytes.
func (sd stringData) Bytes() []byte {
	return []byte(sd)
}
