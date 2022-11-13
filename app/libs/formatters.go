package libs

import (
	"encoding/json"
	"encoding/xml"
	"io"

	"gopkg.in/yaml.v3"
)

const (
	JSON_FORMAT = "json"
	YAML_FORMAT = "yaml"
	XML_FORMAT  = "xml"
)

type Formatter interface {
	WriteFormatted(io.Writer, interface{}) error
	ReadFormatted(r io.Reader, data interface{}) error
	GetFormatName() string
}

type JSONFormatter struct {
}

func (j JSONFormatter) ReadFormatted(r io.Reader, data interface{}) error {
	return json.NewDecoder(r).Decode(data)
}

func (j JSONFormatter) WriteFormatted(w io.Writer, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}

func (j JSONFormatter) GetFormatName() string {
	return JSON_FORMAT
}

type XMLFormatter struct {
}

func (x XMLFormatter) WriteFormatted(w io.Writer, data interface{}) error {
	w.Write([]byte(xml.Header))
	return xml.NewEncoder(w).Encode(data)
}

func (x XMLFormatter) ReadFormatted(r io.Reader, data interface{}) error {
	return xml.NewDecoder(r).Decode(data)
}

func (x XMLFormatter) GetFormatName() string {
	return XML_FORMAT
}

type YAMLFormatter struct {
}

func (y YAMLFormatter) ReadFormatted(r io.Reader, data interface{}) error {
	return yaml.NewDecoder(r).Decode(data)
}

func (y YAMLFormatter) WriteFormatted(w io.Writer, data interface{}) error {
	return yaml.NewEncoder(w).Encode(data)
}

func (y YAMLFormatter) GetFormatName() string {
	return YAML_FORMAT
}
