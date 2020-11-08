package tpprovider

import (
	"fmt"
	"time"
)

/*
service ->provider ->[]apis --> []usage
*/

// GenID to automatic ID generation where is required and must be unique
// ToDo: This function must be enhanced.
func GenID(typeid, name string) string {
	return fmt.Sprintf("%s_%s_%s", typeid, name, time.Now().Day())
}

// RemoveSliceIndex to remove an element from slice at index i
func RemoveSliceIndex(s []interface{}, index int) []interface{} {
	return append(s[:index], s[index+1:]...)
}
