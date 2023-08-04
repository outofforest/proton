package methods

import "reflect"

// Config stores config for generators
type Config struct {
	Type               reflect.Type
	NumOfBooleanFields uint64
}

// BitMapLength returns the number of required bytes to fit the number of bits representing boolean values
func BitMapLength(numOfBits uint64) uint64 {
	return (numOfBits + 7) / 8
}

// BitMapPosition returns the index of byte and the bit inside that byte where the value for indexed boolean value exists
func BitMapPosition(index uint64) (uint64, uint64) {
	return index / 8, index % 8
}
