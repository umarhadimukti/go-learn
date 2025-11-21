package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Building struct {
	Name string `required:"true" max:"100"` // this is a struct tag, read using reflect package
	Floors int `required:"true" min:"1"`
	Location string `required:"false"`
}

func isValid(value any) (bool, error) {
	var valueType reflect.Type = reflect.TypeOf(value)
	for i := 0; i < valueType.NumField(); i++ {
		var field reflect.StructField = valueType.Field(i)
		isRequired, _ := strconv.ParseBool(field.Tag.Get("required"))
		if isRequired {
			var fieldValue reflect.Value = reflect.ValueOf(value).Field(i)
			if fieldValue.IsZero() {
				return false, errors.New(fmt.Sprintf("Field %s is required", field.Name))
			}
		}
	}
	return true, nil
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type:", valueType)

	result := make(map[string]any)
	for i := 0; i < valueType.NumField(); i++ {
		var field reflect.StructField = valueType.Field(i)
		result[field.Name] = map[string]any{
			"type": field.Type.String(),
			"is_required": field.Tag.Get("required"),
			"min": field.Tag.Get("min"),
			"max": field.Tag.Get("max"),
		}
	}
	fmt.Println("Fields:", result)
}

func DemoReflectPackage() {
	var jsiBuilding Building = Building{"JSI Tower", 50, "Benhil, Jakarta"}
	readField(jsiBuilding)
	isValid, err := isValid(jsiBuilding)
	if err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Is valid:", isValid)
	}
}