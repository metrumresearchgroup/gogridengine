package gogridengine

import (
	"encoding/xml"
	"errors"
	"strconv"
)

//ResourceList is a slice of resources primarily used for sourcing internally and setup of receiver based functions
type ResourceList []Resource

//Resource is a more general reported item from
type Resource struct {
	XMLName xml.Name `xml:"resource"`
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

/*
Because we have to serialize everything generically into strings, let's add some methods to return strongly typed values
(Such as converting *G to bytes and actual floats for load)
*/

//Load returns the formatted, type safe float value for the Short Load resource. Provide the window of length:
//load_short load_medium load_long
func (r ResourceList) Load(window string) (float64, error) {
	resource, err := r.locateKey("load_" + window)
	if err != nil {
		return 0, err
	}

	resconv, err := strconv.ParseFloat(resource.Value, 64)
	if err != nil {
		return 0, err
	}

	return resconv, nil
}

//NumberofProcessors is the formatted, type-safe value for the num_proc xml attribute
func (r ResourceList) NumberofProcessors() (int32, error) {
	resource, err := r.locateKey("num_proc")
	if err != nil {
		return 0, err
	}

	resconv, err := strconv.ParseInt(resource.Value, 10, 32)
	if err != nil {
		//Failure to convert to an integer for some reason
		return 0, errors.New("Failure to convert to an integer from string")
	}

	return int32(resconv), nil
}

//FreeMemory returns the type safe values for Memory free (in bytes)
func (r ResourceList) FreeMemory() (int64, error) {
	resource, err := r.locateKey("mem_free")
	if err != nil {
		return 0, err
	}
	identifier := resource.Value[len(resource.Value)-1]
	value := resource.Value[-(len(resource.Value) - 2)]
	size, err := strconv.ParseInt(string(value), 10, 64)
	if err != nil {
		//Failure to parse the integer component of the string
		return 0, errors.New("Failure to parse the integer component of the string")
	}

	switch identifier {
	case 'G':
		return size * (1000 * 1000), nil
	case 'M':
		return size * (1000), nil
	case 'T':
		return size * (1000 * 1000 * 1000), nil
	}

	//If no known size, break to 0
	return 0, errors.New("Unknown Size Identifier. Accepted are 'G', 'M', 'T'")
}

func (r ResourceList) locateKey(key string) (*Resource, error) {
	for _, c := range r {
		if c.Name == key {
			return &c, nil
		}
	}

	//If none are found:
	return &Resource{}, errors.New("Could not located the requested key")
}
