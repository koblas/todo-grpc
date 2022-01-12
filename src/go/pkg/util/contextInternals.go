package util

import (
	"fmt"
	"reflect"
	"unsafe"
)

func PrintContextInternals(ctx interface{}) {
	var printContextInternals func(ctx interface{}, indent string)

	printContextInternals = func(ctx interface{}, indent string) {
		rv := reflect.ValueOf(ctx)
		base := reflect.TypeOf(ctx)

		for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
			rv = rv.Elem()
			base = base.Elem()
		}

		if indent == "" {
			fmt.Printf("\n%sFields for %s.%s\n", indent, base.PkgPath(), base.Name())
		}

		if rv.Kind() == reflect.Struct {
			values := rv
			keys := base

			for i := 0; i < rv.NumField(); i++ {
				reflectValue := values.Field(i)
				reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()
				reflectField := keys.Field(i)

				if reflectField.Name == "Context" {
					printContextInternals(reflectValue.Interface(), indent+"  ")
				} else {
					fmt.Printf("%sfield name: %+v\n", indent, reflectField.Name)
					fmt.Printf("%svalue: %+v\n", indent, reflectValue.Interface())
				}
			}
		} else {
			fmt.Printf("%scontext is empty (int)\n", indent)
		}
	}

	printContextInternals(ctx, "")
}
