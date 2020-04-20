# FlatBuffers(flatc) encode design (Go)

## 0. refence 

 1. [FlatBuffers Internals](https://google.github.io/flatbuffers/flatbuffers_internals.html)
 2. [FlatBuffers Binary Format](https://github.com/dvidelabs/flatcc/blob/master/doc/binary-format.md)
  3. [Writing a schema](https://google.github.io/flatbuffers/flatbuffers_guide_writing_schema.html)

## 1. components of flatc

### 1.1 flatc componts type 

1. in-line storage item, scalars / structs 
2. memory blocks, include header / table ( root table ) / vtable / vectors / string / 

table ( vtable ) / union / string / vector ------ never stored in-line.
scalar / struct --------- stored in-line 

### 1.2 memory blocks 

* header
* table
* vector
* string
* vtable
* ( structs ) 

> Space between the above blocks are zero padded and present in order to ensure proper alignment. Structs must be packed as densely as possible according the alignment rules that apply - this ensures that all implementations will agree on the layout. The blocks must not overlap in memory but two blocks may be shared if they represent the same data such as sharing a string.


### 1.3 scalars

all scalars golang support :
```
// bool
// uint8, uint16, uint32, uint64 (unsigned)
// int8, int16, int32, int64     (two's complement)
// float32, float64              (IEEE-754)
```

enum, use int8 in go 

## 2. flatc encode / decode 

