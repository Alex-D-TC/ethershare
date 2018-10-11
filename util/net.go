package util

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func MakeGobEncoder() (*gob.Encoder, *bytes.Buffer) {
	buffer := bytes.NewBuffer([]byte{})
	return gob.NewEncoder(buffer), buffer
}

func MakeGobDecoder() (*gob.Decoder, *bytes.Buffer) {
	buffer := bytes.NewBuffer([]byte{})
	return gob.NewDecoder(buffer), buffer
}

func GobEncode(data interface{}) ([]byte, error) {
	encoder, buffer := MakeGobEncoder()
	err := encoder.Encode(data)
	return buffer.Bytes(), err
}

func GobDecode(rawData []byte, destination interface{}) error {
	decoder, buffer := MakeGobDecoder()
	buffer.Write(rawData)
	return decoder.Decode(destination)
}

func MakeJSONEncoder() (*json.Encoder, *bytes.Buffer) {
	buff := bytes.NewBuffer([]byte{})
	return json.NewEncoder(buff), buff
}

func MakeJSONDecoder() (*json.Decoder, *bytes.Buffer) {
	buff := bytes.NewBuffer([]byte{})
	return json.NewDecoder(buff), buff
}

func JSONEncode(data interface{}) ([]byte, error) {
	encoder, buff := MakeJSONEncoder()
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func JSONDecode(raw []byte, dest interface{}) error {
	decoder, buff := MakeJSONDecoder()
	buff.Write(raw)
	return decoder.Decode(dest)
}
