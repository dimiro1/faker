package defaults

import (
	"fmt"
	"reflect"
	"strconv"
)

// Apply apply the defaults values defined in the field tags.
// Currently the supported types are int, int32, int64, float32, float64, string, bool
// If the value is different of one these types the function returns an error.
//
// The function expects a pointer to a struct.
//
// Eg.
//
// type testStruct struct {
//      AString      string  `default:"Claudemiro"`
//      AInt         int     `default:"27"`
//      AInt32       int32   `default:"28"`
//      AInt64       int64   `default:"29"`
//      AFloat32     float32 `default:"3.1415"`
//      AFloat64     float64 `default:"3.1415"`
//      ABoolean     bool    `default:"true"`
//      NotAnnotated float64
// }
func Apply(val interface{}) error {
	const defaultTagName = "default"

	// Invalid type of the val
	if t := reflect.TypeOf(val).Kind(); t != reflect.Ptr {
		return fmt.Errorf("Apply: The val must be a pointer to a struct: %+v received", t)
	}

	value := reflect.ValueOf(val).Elem()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		// The field already has a value
		if !reflect.DeepEqual(reflect.Zero(field.Type()).Interface(), field.Interface()) {
			continue
		}

		tag := value.Type().Field(i).Tag.Get(defaultTagName)

		// Continue if the current field is not annotated
		if tag == "" {
			continue
		}

		typeName := field.Type().Name()

		switch typeName {
		case "bool":
			field.SetBool(tag == "true")
		case "string":
			field.SetString(tag)
		case "int":
			fallthrough
		case "int32":
			fallthrough
		case "int64":
			v, err := strconv.ParseInt(tag, 10, 64)
			if err != nil {
				return err
			}
			field.SetInt(v)
		case "float32":
			fallthrough
		case "float64":
			v, err := strconv.ParseFloat(tag, 64)
			if err != nil {
				return err
			}
			field.SetFloat(v)
		default:
			return fmt.Errorf("Apply: Not supported for type %s", typeName)
		}
	}

	return nil
}
