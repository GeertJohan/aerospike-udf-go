package udf

import "fmt"

func applyRecord(name string) *udfError {
	fmt.Printf("applying UDF %s to record\n", name)
	a, exists := functions[name]
	if !exists {
		return ErrFunctionNotAvailable
	}
	a.v.Call(nil)
	return nil
}

func applyStream(name string) *udfError {
	fmt.Printf("applying UDF %s to stream\n", name)
	panic("not implemented yet")
	return nil
}
