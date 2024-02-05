package cmd

import (
	"reflect"
	"time"
)

// mergeStructs merges fields from src into dst, excluding specified fields
func MergeStructs(dst, src interface{}) {
	dstVal := reflect.ValueOf(dst).Elem()
	srcVal := reflect.ValueOf(src).Elem()

	for i := 0; i < dstVal.NumField(); i++ {
		dstField := dstVal.Field(i)
		srcField := srcVal.Field(i)

		// Check if the field is settable and not in the list of fields to exclude
		if dstField.CanSet() && dstField.Type() == srcField.Type() {
			// Example exclusion check: skip "AccessToken" field
			if dstVal.Type().Field(i).Name != "AccessToken" {
				dstField.Set(srcField)
			}
		}
	}
}

func FormatTime(unixTime int64, layout string) (string, error) {
	t := time.Unix(unixTime, 0).Local()
	return t.Format(layout), nil
}
