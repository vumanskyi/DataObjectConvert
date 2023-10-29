package export

import (
	"errors"
	"fmt"
	"os"
)

type Exporter interface {
	Export(data interface{}) error
}

type DataExport struct {
	t string
}

func NewDataExport(t string) *DataExport {
	return &DataExport{t: t}
}

func (d *DataExport) Export(data interface{}) error {
	switch d.t {
	case "stdout":
		return stdout(data)
	}

	return errors.New("export type does not exist")
}

// stdout - export to console
func stdout(data interface{}) error {
	_, err := fmt.Fprintln(os.Stdout, data)

	return err
}

// file - export data to file
func file(data interface{}) error {
	panic("file export in progress")
}
