package gogridengine

import (
	"errors"
	"strconv"
)

/*
Because we have to serialize everything generically into strings, let's add some methods to return strongly typed values
(Such as converting *G to bytes and actual floats for load)
*/

//Load returns the formatted, type safe float value for the Short Load resource. Provide the window of length:
//load_short load_medium load_long
func (r ResourceList) Load(window string) (float64, error) {
	return r.getFloatValueFromList("load_" + window)
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
func (r ResourceList) FreeMemory() (StorageValue, error) {
	return r.getStorageValueFromList("mem_free")
}

//FreeSwap returns the type casted value for Free Swap Bytes
func (r ResourceList) FreeSwap() (StorageValue, error) {
	return r.getStorageValueFromList("swap_free")
}

//FreeVirtualMemory returns the type casted value for Free Virtual Memory Bytes
func (r ResourceList) FreeVirtualMemory() (StorageValue, error) {
	return r.getStorageValueFromList("virtual_free")
}

//TotalMemory returns the Type cast value for Total Memory in bytes
func (r ResourceList) TotalMemory() (StorageValue, error) {
	return r.getStorageValueFromList("mem_total")
}

//TotalSwap returns the Type cast value for Total Swap in bytes
func (r ResourceList) TotalSwap() (StorageValue, error) {
	return r.getStorageValueFromList("swap_total")
}

//TotalVirtual returns the Type cast value for the virtual memory total in bytes
func (r ResourceList) TotalVirtual() (StorageValue, error) {
	return r.getStorageValueFromList("virtual_total")
}

//MemoryUsed returns the type cast value for Memory used in bytes
func (r ResourceList) MemoryUsed() (StorageValue, error) {
	return r.getStorageValueFromList("mem_used")
}

//SwapUsed returns the type cast value for swap used in bytes
func (r ResourceList) SwapUsed() (StorageValue, error) {
	return r.getStorageValueFromList("swap_used")
}

//VirtualUsed returns the type cast value for swap used in bytes
func (r ResourceList) VirtualUsed() (StorageValue, error) {
	return r.getStorageValueFromList("virtual_used")
}

//CPU returns utilization type cast as a float
func (r ResourceList) CPU() (float64, error) {
	return r.getFloatValueFromList("cpu")
}

//ProcessorCount returns the number of processors type cast as integer64
func (r ResourceList) ProcessorCount() (int64, error) {
	return r.getIntegerValueFromList("num_proc")
}

//MSocketCount returns the socket count type cast as an int64
func (r ResourceList) MSocketCount() (int64, error) {
	return r.getIntegerValueFromList("m_socket")
}

//MCoreCount returns the core count type cast as an int64
func (r ResourceList) MCoreCount() (int64, error) {
	return r.getIntegerValueFromList("m_core")
}

//MThreadCount returns the thread count type cast as an int64
func (r ResourceList) MThreadCount() (int64, error) {
	return r.getIntegerValueFromList("m_thread")
}

//NPLoadAverage NP Load type cast as float 64
func (r ResourceList) NPLoadAverage() (float64, error) {
	return r.getFloatValueFromList("np_load_avg")
}

//NPLoadShort NP Load type cast as float 64
func (r ResourceList) NPLoadShort() (float64, error) {
	return r.getFloatValueFromList("np_load_short")
}

//NPLoadMedium NP Load type cast as float 64
func (r ResourceList) NPLoadMedium() (float64, error) {
	return r.getFloatValueFromList("np_load_medium")
}

//NPLoadLong NP Load type cast as float 64
func (r ResourceList) NPLoadLong() (float64, error) {
	return r.getFloatValueFromList("np_load_long")
}

//Used for extracting a storage value from the resource list to minimize function size
func (r ResourceList) getStorageValueFromList(KeyName string) (StorageValue, error) {
	resource, err := r.locateKey(KeyName)
	if err != nil {
		return StorageValue{}, err
	}

	return newStorageValue(resource.Value)
}

//Used for extracting a float from a resource list to minimize function size
func (r ResourceList) getFloatValueFromList(KeyName string) (float64, error) {
	resource, err := r.locateKey(KeyName)
	if err != nil {
		return 0, err
	}

	resconv, err := strconv.ParseFloat(resource.Value, 64)
	if err != nil {
		return 0, err
	}

	return resconv, nil
}

//Used for extracting an integer froma resource list to minimize function size
func (r ResourceList) getIntegerValueFromList(KeyName string) (int64, error) {
	resource, err := r.locateKey(KeyName)
	if err != nil {
		return 0, err
	}

	resconv, err := strconv.ParseInt(resource.Value, 10, 64)

	if err != nil {
		return 0, err
	}

	return resconv, nil
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
