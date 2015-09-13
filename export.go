package udf

import "C"
import (
	"fmt"
	"runtime"
)

const (
	udfAPIVersionMajor = 1
	udfAPIVersionMinor = 0
)

//export aerospike_udf_go_get_api_version_major
func aerospike_udf_go_get_api_version_major() C.int {
	return C.int(udfAPIVersionMajor)
}

//export aerospike_udf_go_get_api_version_minor
func aerospike_udf_go_get_api_version_minor() C.int {
	return C.int(udfAPIVersionMinor)
}

//export aerospike_udf_go_get_property
func aerospike_udf_go_get_property(cName *C.char) *C.char {
	name := C.GoString(cName)
	switch name {
	case "conn-name":
		return C.CString("github.com/GeertJohan/aerospike-udf-go")
	case "conn-version":
		return C.CString("0.1")
	case "go-version":
		return C.CString(runtime.Version())
	}
	return C.CString("")
}

//export aerospike_udf_go_setup
func aerospike_udf_go_setup() C.int {
	err := setupFunctions()
	if err != nil {
		fmt.Printf("error in aerospike_udf_go_setup during setupFunctions: %v\n", err)
		return C.int(err.code)
	}

	return C.int(0)
}

//export aerospike_udf_go_apply_record
func aerospike_udf_go_apply_record(udfName *C.char) (retval C.int) {
	defer recoverPanicForExported(&retval)
	name := C.GoString(udfName)

	err := applyRecord(name)
	if err != nil {
		fmt.Printf("error in aerospike_udf_go_apply_record: %v\n", err)
		return C.int(err.code)
	}

	return C.int(0)
}

//export aerospike_udf_go_apply_stream
func aerospike_udf_go_apply_stream(udfName *C.char) (retval C.int) {
	defer recoverPanicForExported(&retval)
	name := C.GoString(udfName)

	err := applyStream(name)
	if err != nil {
		fmt.Printf("error in aerospike_udf_go_apply_stream: %v\n", err)
		return C.int(err.code)
	}

	return C.int(0)
}

// We're calling third-party functions in apply_record and apply_stream. Third-party code could cause a panic.
// To protect aerospike, we should recover the panic inside this package, and return an errorcode to aerospike.
func recoverPanicForExported(retvalPtr *C.int) {
	if r := recover(); r != nil {
		fmt.Println(r)
		*retvalPtr = 255
	}
}
