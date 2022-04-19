//go:build !gst116
// +build !gst116

package gst

/*
#include "gst.go.h"
*/
import "C"

import (
	"unsafe"

	"github.com/tinyzimmer/go-glib/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{
			T: TypeValueArray,
			F: marshalValueArray,
		},
		{
			T: TypeValueList,
			F: marshalValueList,
		},
	}

	glib.RegisterGValueMarshalers(tm)
}

func marshalValueArray(p unsafe.Pointer) (interface{}, error) {
	val := toGValue(p)
	out := ValueArrayValue(*glib.ValueFromNative(unsafe.Pointer(val)))
	return &out, nil
}

func marshalValueList(p unsafe.Pointer) (interface{}, error) {
	val := glib.ValueFromNative(unsafe.Pointer(toGValue(p)))
	out := ValueListValue(*glib.ValueFromNative(unsafe.Pointer(val)))
	return &out, nil
}
