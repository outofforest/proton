package main

import (
	"github.com/outofforest/proton"
	"github.com/outofforest/proton/test/pkg1"
	"github.com/outofforest/proton/test/pkg2"
)

//go:generate go run .

func main() {
	proton.Generate("../types.proton.go",
		pkg2.SubMsg{},
		pkg1.MsgUint64{},
		pkg1.MsgUint32{},
		pkg1.MsgUint16{},
		pkg1.MsgUint8{},
		pkg1.MsgInt64{},
		pkg1.MsgInt32{},
		pkg1.MsgInt16{},
		pkg1.MsgInt8{},
		pkg1.MsgBool3{},
		pkg1.MsgBool10{},
		pkg1.MsgFloat64{},
		pkg1.MsgFloat32{},
		pkg1.MsgString{},
		pkg1.MsgArray{},
		pkg1.MsgSlice{},
		pkg1.MsgMap{},
		pkg1.MsgMapString{},
		pkg1.MsgStruct{},
		pkg1.MsgStructAnonymous{},
		pkg1.MsgArrayUint8{},
		pkg1.MsgArrayInt8{},
		pkg1.MsgArrayFloat32{},
		pkg1.MsgArrayFloat64{},
		pkg1.MsgSliceUint8{},
		pkg1.MsgSliceInt8{},
		pkg1.MsgSliceFloat32{},
		pkg1.MsgSliceFloat64{},
		pkg1.MsgMixed{},
		pkg1.MsgUint64Custom{},
		pkg1.MsgUint32Custom{},
		pkg1.MsgUint16Custom{},
		pkg1.MsgUint8Custom{},
		pkg1.MsgInt64Custom{},
		pkg1.MsgInt32Custom{},
		pkg1.MsgInt16Custom{},
		pkg1.MsgInt8Custom{},
		pkg1.MsgBoolCustom{},
		pkg1.MsgFloat64Custom{},
		pkg1.MsgFloat32Custom{},
		pkg1.MsgStringCustom{},
		pkg1.MsgArrayCustom{},
		pkg1.MsgSliceCustom{},
		pkg1.MsgMapCustom{},
		pkg1.MsgArrayUint8Custom{},
		pkg1.MsgArrayUint8Custom2{},
		pkg1.MsgArrayInt8Custom{},
		pkg1.MsgArrayFloat32Custom{},
		pkg1.MsgArrayFloat64Custom{},
		pkg1.MsgSliceUint8Custom{},
		pkg1.MsgSliceUint8Custom2{},
		pkg1.MsgSliceInt8Custom{},
		pkg1.MsgSliceFloat32Custom{},
		pkg1.MsgSliceFloat64Custom{},
		pkg1.MsgMixedCustom{},
	)
}
