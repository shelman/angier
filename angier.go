package angier

import (
	"fmt"
	"reflect"
)

// Takes in two pointers to structs.  For any fields with the same name and type
// in the two structs, the values of the fields in the src struct are copied
// over to the corresponding field in the dest struct.
func Transfer(src interface{}, dest interface{}) error {

	// first, check that both are pointers
	if reflect.TypeOf(src).Kind() != reflect.Ptr ||
		reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return fmt.Errorf("both the src and dest must be pointers")
	}

	// now, make sure that both are pointers to structs
	srcVal := reflect.Indirect(reflect.ValueOf(src))
	destVal := reflect.Indirect(reflect.ValueOf(dest))
	if reflect.TypeOf(srcVal).Kind() != reflect.Struct ||
		reflect.TypeOf(destVal).Kind() != reflect.Struct {
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

		switch srcFieldVal.Type().Kind() {
		case reflect.Bool:
			destFieldVal.SetBool(srcFieldVal.Bool())
		case reflect.Int:
			destFieldVal.SetInt(srcFieldVal.Int())
		case reflect.String:
			destFieldVal.SetString(srcFieldVal.String())
		default:
			fmt.Println(fmt.Sprintf("field type %v not yet supported",
				srcFieldVal.Type().Kind()))
		}

	}

	return nil

}
