package gogridengine

import (
	"encoding/xml"
	"strconv"
)

//ResourceList is a slice of resources primarily used for sourcing internally and setup of receiver based functions
type ResourceList []Resource

//Resource is a general representation of a system resource attached to a host, identified purely by Key-value pairs. These can be best thought of as Unix Load, Memory resource allocations, etc
type Resource struct {
	XMLName xml.Name `xml:"resource" json:"-"`
	Name    string   `xml:"name,attr" json:"name"`
	Type    string   `xml:"type,attr" json:"type"`
	Value   string   `xml:",innerxml"`
}

//StorageValue breaks down string metrics from a computer storage standpoint (ie 10.2G) so that it can be calculated to bytes
type StorageValue struct {
	Size  float64 `json:"size"`
	Scale string  `json:"scale"`
	Bytes int64   `json:"bytes"`
}

func newStorageValue(input string) (StorageValue, error) {
	var sv StorageValue
	sv.Scale = string(input[len(input)-1])
	remainingBytes := input[:(len(input) - 1)]
	remainingSize, err := strconv.ParseFloat(string(remainingBytes), 64)
	if err != nil {
		return StorageValue{}, err
	}
	sv.Size = remainingSize

	//Now the case statement
	switch sv.Scale {
	case "G":
		sv.Bytes = int64(sv.Size * (1000 * 1000 * 1000))
	case "M":
		sv.Bytes = int64(sv.Size * (1000 * 1000))
	case "T":
		sv.Bytes = int64(sv.Size * (1000 * 1000 * 1000 * 1000))
	}

	return sv, nil
}
