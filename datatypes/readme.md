Basic Types:
number, string, boolean
- number types
platform-dependent integers
- int, uint, uintptr
Explicitly sized signed integers
- int8, int16, int32, int64
Explicitly sized unsigned integers
- uint8, uint16, uint32, uint64
Specialized integer aliases
- byte, rune

Aggregates:
Array ([n]T) # explanation: a fixed-size array of n elements of type T
Struct (struct) # explanation: a struct is a collection of fields of different types

References:
slice ([]T) # explanation: a dynamic array that can grow and shrink
map (map[K]V) # explanation: a hash table that maps keys to values
channel (chan T) # explanation: a communication channel that allows goroutines to communicate with each other
pointer (*T) # explanation: a pointer to a value of type T

Interfaces:
interface (interface{}) # explanation: an interface is a set of method signatures
Empty Interface:
interface{} # explanation: an empty interface is a generic interface that can hold any type
