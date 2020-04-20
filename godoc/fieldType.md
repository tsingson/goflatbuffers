# field type

from idl.h

```
#define FLATBUFFERS_GEN_TYPES_SCALAR(TD) \
  TD(NONE,   "",       uint8_t,  byte,   byte,    byte,   uint8,   u8,   UByte, UInt8) \
  TD(UTYPE,  "",       uint8_t,  byte,   byte,    byte,   uint8,   u8,   UByte, UInt8) /* begin scalar/int */ \
  TD(BOOL,   "bool",   uint8_t,  boolean,bool,    bool,   bool,    bool, Boolean, Bool) \
  TD(CHAR,   "byte",   int8_t,   byte,   int8,    sbyte,  int8,    i8,   Byte, Int8) \
  TD(UCHAR,  "ubyte",  uint8_t,  byte,   byte,    byte,   uint8,   u8,   UByte, UInt8) \
  TD(SHORT,  "short",  int16_t,  short,  int16,   short,  int16,   i16,  Short, Int16) \
  TD(USHORT, "ushort", uint16_t, short,  uint16,  ushort, uint16,  u16,  UShort, UInt16) \
  TD(INT,    "int",    int32_t,  int,    int32,   int,    int32,   i32,  Int, Int32) \
  TD(UINT,   "uint",   uint32_t, int,    uint32,  uint,   uint32,  u32,  UInt, UInt32) \
  TD(LONG,   "long",   int64_t,  long,   int64,   long,   int64,   i64,  Long, Int64) \
  TD(ULONG,  "ulong",  uint64_t, long,   uint64,  ulong,  uint64,  u64,  ULong, UInt64) /* end int */ \
  TD(FLOAT,  "float",  float,    float,  float32, float,  float32, f32,  Float, Float32) /* begin float */ \
  TD(DOUBLE, "double", double,   double, float64, double, float64, f64,  Double, Double) /* end float/scalar */
#define FLATBUFFERS_GEN_TYPES_POINTER(TD) \
  TD(STRING, "string", Offset<void>, int, int, StringOffset, int, unused, Int, Offset<String>) \
  TD(VECTOR, "",       Offset<void>, int, int, VectorOffset, int, unused, Int, Offset<UOffset>) \
  TD(STRUCT, "",       Offset<void>, int, int, int,          int, unused, Int, Offset<UOffset>) \
  TD(UNION,  "",       Offset<void>, int, int, int,          int, unused, Int, Offset<UOffset>)
#define FLATBUFFERS_GEN_TYPE_ARRAY(TD) \
  TD(ARRAY,  "",       int,          int, int, int,          int, unused, Int, Offset<UOffset>)
  ```

  // The fields are:
  // - enum
  // - FlatBuffers schema type.
  // - C++ type.
  // - Java type.
  // - Go type.
  // - C# / .Net type.
  // - Python type.
  // - Rust type.
  // - Kotlin type.