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

// TypeValueArray is the GType for a GstValueArray
var TypeValueArray = glib.Type(C.gst_value_array_get_type())

// ValueArrayValue represets a GstValueArray.
type ValueArrayValue glib.Value

// ValueArray converts the given slice of Go types into a ValueArrayValue.
// This function can return nil on any conversion or memory allocation errors.
func ValueArray(ss []interface{}) *ValueArrayValue {
	v, err := glib.ValueAlloc()
	if err != nil {
		return nil
	}
	C.gst_value_array_init(
		(*C.GValue)(unsafe.Pointer(v.GValue)),
		C.guint(len(ss)),
	)
	for _, s := range ss {
		val, err := glib.GValue(s)
		if err != nil {
			return nil
		}
		C.gst_value_array_append_value(
			(*C.GValue)(unsafe.Pointer(v.GValue)),
			(*C.GValue)(unsafe.Pointer(val.GValue)),
		)
	}
	out := ValueArrayValue(*v)
	return &out
}

// Size returns the size of the array.
func (v *ValueArrayValue) Size() uint {
	return uint(C.gst_value_array_get_size((*C.GValue)(unsafe.Pointer(v.GValue))))
}

// ValueAt returns the value at the index in the array, or nil on any error.
func (v *ValueArrayValue) ValueAt(idx uint) interface{} {
	gval := C.gst_value_array_get_value(
		(*C.GValue)(unsafe.Pointer(v.GValue)),
		C.guint(idx),
	)
	if gval == nil {
		return nil
	}
	out, err := glib.ValueFromNative(unsafe.Pointer(gval)).GoValue()
	if err != nil {
		return nil
	}
	return out
}

// ToGValue implements a glib.ValueTransformer.
func (v *ValueArrayValue) ToGValue() (*glib.Value, error) {
	out := glib.Value(*v)
	return &out, nil
}

// TypeValueList is the GType for a GstValueList
var TypeValueList = glib.Type(C.gst_value_list_get_type())

// ValueListValue represets a GstValueList.
type ValueListValue glib.Value

// ValueList converts the given slice of Go types into a ValueListValue.
// This function can return nil on any conversion or memory allocation errors.
func ValueList(ss []interface{}) *ValueListValue {
	v, err := glib.ValueAlloc()
	if err != nil {
		return nil
	}
	C.gst_value_list_init(
		(*C.GValue)(unsafe.Pointer(v.GValue)),
		C.guint(len(ss)),
	)
	for _, s := range ss {
		val, err := glib.GValue(s)
		if err != nil {
			return nil
		}
		C.gst_value_list_append_value(
			(*C.GValue)(unsafe.Pointer(v.GValue)),
			(*C.GValue)(unsafe.Pointer(val.GValue)),
		)
	}
	out := ValueListValue(*v)
	return &out
}

// Size returns the size of the list.
func (v *ValueListValue) Size() uint {
	return uint(C.gst_value_list_get_size((*C.GValue)(unsafe.Pointer(v.GValue))))
}

// ValueAt returns the value at the index in the lise, or nil on any error.
func (v *ValueListValue) ValueAt(idx uint) interface{} {
	gval := C.gst_value_list_get_value(
		(*C.GValue)(unsafe.Pointer(v.GValue)),
		C.guint(idx),
	)
	if gval == nil {
		return nil
	}
	out, err := glib.ValueFromNative(unsafe.Pointer(gval)).GoValue()
	if err != nil {
		return nil
	}
	return out
}

// Concat concatenates copies of this list and value into a new list. Values that are not of type
// TypeValueList are treated as if they were lists of length 1. dest will be initialized to the type
// TypeValueList.
func (v *ValueListValue) Concat(value *ValueListValue) *ValueListValue {
	out, err := glib.ValueAlloc()
	if err != nil {
		return nil
	}
	C.gst_value_list_concat(
		(*C.GValue)(unsafe.Pointer(out.GValue)),
		(*C.GValue)(unsafe.Pointer(v.GValue)),
		(*C.GValue)(unsafe.Pointer(value.GValue)),
	)
	o := ValueListValue(*out)
	return &o
}

// Merge merges copies of value into this list. Values that are not of type TypeValueList are treated as
// if they were lists of length 1.
//
// The result will be put into a new value and will either be a list that will not contain any duplicates,
// or a non-list type (if the lists were equal).
func (v *ValueListValue) Merge(value *ValueListValue) *ValueListValue {
	out, err := glib.ValueAlloc()
	if err != nil {
		return nil
	}
	C.gst_value_list_merge(
		(*C.GValue)(unsafe.Pointer(out.GValue)),
		(*C.GValue)(unsafe.Pointer(v.GValue)),
		(*C.GValue)(unsafe.Pointer(value.GValue)),
	)
	o := ValueListValue(*out)
	return &o
}

// ToGValue implements a glib.ValueTransformer.
func (v *ValueListValue) ToGValue() (*glib.Value, error) {
	out := glib.Value(*v)
	return &out, nil
}
