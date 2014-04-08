package angier

import (
	"fmt"
	"reflect"
)

// Takes in two pointers to structs.  For any fields with the same name and type
// in the two structs, the values of the exported fields in the src struct are
// copied over to the corresponding field in the dest struct.
func Transfer(src interface{}, dest interface{}) error {

	// first, check that both are pointers
	if reflect.ValueOf(src).Type().Kind() != reflect.Ptr ||
		reflect.ValueOf(dest).Type().Kind() != reflect.Ptr {
		return fmt.Errorf("both the src and dest must be pointers")
	}

	// now, make sure that both are pointers to structs
	srcVal := reflect.Indirect(reflect.ValueOf(src))
	destVal := reflect.Indirect(reflect.ValueOf(dest))
	if srcVal.Type().Kind() != reflect.Struct ||
		destVal.Type().Kind() != reflect.Struct {
		return fmt.Errorf("both the src and dest must be pointers to structs")
	}

	// iterate over the fields of the src struct
	fieldsInSrc := srcVal.NumField()
	for i := 0; i < fieldsInSrc; i++ {
		fieldInSrc := srcVal.Type().Field(i)
		fieldName := fieldInSrc.Name

		// check to see if the dest struct has a field of the same value
		fieldInDest, exists := destVal.Type().FieldByName(fieldName)
		if !exists || fieldInDest.Type.Kind() != fieldInSrc.Type.Kind() {
			continue
		}

		// convert to reflect.Value
		srcFieldVal := srcVal.FieldByName(fieldName)
		destFieldVal := destVal.FieldByName(fieldName)

		// make sure the field is exported
		if !destFieldVal.CanSet() {
			continue
		}

		// set the field
		destFieldVal.Set(srcFieldVal)

	}

	return nil

}
