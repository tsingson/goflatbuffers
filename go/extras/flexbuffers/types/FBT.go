// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package types

import "strconv"

type FBT int8

const (
	FBTNULL                     FBT = 0
	FBTINT                      FBT = 1
	FBTUINT                     FBT = 2
	FBTFLOAT                    FBT = 3
	FBTKEY                      FBT = 4
	FBTSTRING                   FBT = 5
	FBTINDIRECT_INT             FBT = 6
	FBTINDIRECT_UINT            FBT = 7
	FBTINDIRECT_FLOAT           FBT = 8
	FBTMAP                      FBT = 9
	FBTVECTOR                   FBT = 10
	FBTVECTOR_INT               FBT = 11
	FBTVECTOR_UINT              FBT = 12
	FBTVECTOR_FLOAT             FBT = 13
	FBTVECTOR_KEY               FBT = 14
	FBTVECTOR_STRING_DEPRECATED FBT = 15
	FBTVECTOR_INT2              FBT = 16
	FBTVECTOR_UINT2             FBT = 17
	FBTVECTOR_FLOAT2            FBT = 18
	FBTVECTOR_INT3              FBT = 19
	FBTVECTOR_UINT3             FBT = 20
	FBTVECTOR_FLOAT3            FBT = 21
	FBTVECTOR_INT4              FBT = 22
	FBTVECTOR_UINT4             FBT = 23
	FBTVECTOR_FLOAT4            FBT = 24
	FBTBLOB                     FBT = 25
	FBTBOOL                     FBT = 26
	FBTVECTOR_BOOL              FBT = 36

	FBTVerifyValueMin FBT = 0
	FBTVerifyValueMax FBT = 36
)

var EnumNamesFBT = map[FBT]string{
	FBTNULL:                     "NULL",
	FBTINT:                      "INT",
	FBTUINT:                     "UINT",
	FBTFLOAT:                    "FLOAT",
	FBTKEY:                      "KEY",
	FBTSTRING:                   "STRING",
	FBTINDIRECT_INT:             "INDIRECT_INT",
	FBTINDIRECT_UINT:            "INDIRECT_UINT",
	FBTINDIRECT_FLOAT:           "INDIRECT_FLOAT",
	FBTMAP:                      "MAP",
	FBTVECTOR:                   "VECTOR",
	FBTVECTOR_INT:               "VECTOR_INT",
	FBTVECTOR_UINT:              "VECTOR_UINT",
	FBTVECTOR_FLOAT:             "VECTOR_FLOAT",
	FBTVECTOR_KEY:               "VECTOR_KEY",
	FBTVECTOR_STRING_DEPRECATED: "VECTOR_STRING_DEPRECATED",
	FBTVECTOR_INT2:              "VECTOR_INT2",
	FBTVECTOR_UINT2:             "VECTOR_UINT2",
	FBTVECTOR_FLOAT2:            "VECTOR_FLOAT2",
	FBTVECTOR_INT3:              "VECTOR_INT3",
	FBTVECTOR_UINT3:             "VECTOR_UINT3",
	FBTVECTOR_FLOAT3:            "VECTOR_FLOAT3",
	FBTVECTOR_INT4:              "VECTOR_INT4",
	FBTVECTOR_UINT4:             "VECTOR_UINT4",
	FBTVECTOR_FLOAT4:            "VECTOR_FLOAT4",
	FBTBLOB:                     "BLOB",
	FBTBOOL:                     "BOOL",
	FBTVECTOR_BOOL:              "VECTOR_BOOL",
}

var EnumValuesFBT = map[string]FBT{
	"NULL":                     FBTNULL,
	"INT":                      FBTINT,
	"UINT":                     FBTUINT,
	"FLOAT":                    FBTFLOAT,
	"KEY":                      FBTKEY,
	"STRING":                   FBTSTRING,
	"INDIRECT_INT":             FBTINDIRECT_INT,
	"INDIRECT_UINT":            FBTINDIRECT_UINT,
	"INDIRECT_FLOAT":           FBTINDIRECT_FLOAT,
	"MAP":                      FBTMAP,
	"VECTOR":                   FBTVECTOR,
	"VECTOR_INT":               FBTVECTOR_INT,
	"VECTOR_UINT":              FBTVECTOR_UINT,
	"VECTOR_FLOAT":             FBTVECTOR_FLOAT,
	"VECTOR_KEY":               FBTVECTOR_KEY,
	"VECTOR_STRING_DEPRECATED": FBTVECTOR_STRING_DEPRECATED,
	"VECTOR_INT2":              FBTVECTOR_INT2,
	"VECTOR_UINT2":             FBTVECTOR_UINT2,
	"VECTOR_FLOAT2":            FBTVECTOR_FLOAT2,
	"VECTOR_INT3":              FBTVECTOR_INT3,
	"VECTOR_UINT3":             FBTVECTOR_UINT3,
	"VECTOR_FLOAT3":            FBTVECTOR_FLOAT3,
	"VECTOR_INT4":              FBTVECTOR_INT4,
	"VECTOR_UINT4":             FBTVECTOR_UINT4,
	"VECTOR_FLOAT4":            FBTVECTOR_FLOAT4,
	"BLOB":                     FBTBLOB,
	"BOOL":                     FBTBOOL,
	"VECTOR_BOOL":              FBTVECTOR_BOOL,
}

func (v FBT) String() string {
	if s, ok := EnumNamesFBT[v]; ok {
		return s
	}
	return "FBT(" + strconv.FormatInt(int64(v), 10) + ")"
}