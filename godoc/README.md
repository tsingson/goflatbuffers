# dev note



## 0. reference 
[https://github.com/mzaks/FlatBuffersSwift/wiki/FlatBuffers-Explained](https://github.com/mzaks/FlatBuffersSwift/wiki/FlatBuffers-Explained)

## 1. flatbuffers internal

[http://google.github.io/flatbuffers/md__internals.html]( http://google.github.io/flatbuffers/md__internals.html)

### 1.1 components



> On purpose, the format leaves a lot of details about where exactly things live in memory undefined, e.g. fields in a table can have any order, and objects to some extend can be stored in many orders. This is because the format doesn't need this information to be efficient, and it leaves room for optimization and extension (for example, fields can be packed in a way that is most compact). Instead, the format is defined in terms of offsets and adjacency only. This may mean two different implementations may produce different binaries given the same input values, and this is perfectly valid.

fields in a table can have any order



> The most important and generic offset type (see `flatbuffers.h`) is `uoffset_t`, which is currently always a `uint32_t`, and is used to refer to all tables/unions/strings/vectors (these are never stored in-line). 

table / union / string / vector ------ never stored in-line.

scalar / struct --------- stored in-line 

### 1.2 struct

struct -------- They are always stored inline in their parent (a struct, table, or vector)

### 1.3 table  

> They start with an `soffset_t` to a vtable. This is a signed version of `uoffset_t`, since vtables may be stored anywhere relative to the object. This offset is substracted (not added) from the object start to arrive at the vtable start. This offset is followed by all the fields as aligned scalars (or offsets).