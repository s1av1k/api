package format

import (
	"io"
)

const HumanReadablePrefix = ""
const HumanReadableIndent = "  "

type Structure struct {
	MimeTypes            *[]string
	FileExtensions       *[]string
	InputFormatFileFunc  func(s *InputStructure) error
	InputFormatFunc      func(w io.Reader, s *InputStructure, hr bool) error
	OutputFormatFunc     func(w io.Writer, s *OutputStructure, hr bool) error
}

type InputStructure struct {
	File           string `yaml:"File"`
	Date           string `yaml:"Date"`
	Host           string `yaml:"Host"`
	Port           int    `yaml:"Port"`
	Content        []byte `yaml:"Content"`
	ContentSize    int64  `yaml:"ContentSize"`
	Error          string `yaml:"Error"`
	Status         int    `yaml:"Status"`
}

type OutputStructure struct {
	Time   struct {
		Stamp   int64  `yaml:"Stamp"`
		RFC3339 string `yaml:"RFC3339"`
	}
	Status struct{
		Code    int    `yaml:"Code"`
		Text    string `yaml:"Text"`
		Message string `yaml:"Message"`
	}
	Content     string `yaml:"Content"`
}
