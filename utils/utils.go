package utils

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/gobeam/stringy"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pw string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("could not encrypt password")
	}
	return string(bytes)
}

func SetMapFieldToStruct(myStruct interface{}, fieldName string, fieldValue interface{}) error {
	structValue := reflect.ValueOf(myStruct).Elem()
	str := stringy.New(fieldName)
	pascalCase := str.PascalCase()
	structFieldValue := structValue.FieldByName(pascalCase.Get())

	if !structFieldValue.IsValid() {
		return fmt.Errorf("field: %s does not exist in struct", fieldName)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("cannot set %s field value", fieldName)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(fieldValue)
	if structFieldType != val.Type() {
		return errors.New("provided value type didn't match struct field type")
	}

	structFieldValue.Set(val)
	return nil
}

func FillStructFromMap(myStruct interface{}, myMap map[string]interface{}) error {
	for k, v := range myMap {
		err := SetMapFieldToStruct(myStruct, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func BoolToCheckboxValue(val bool) string {
	if val {
		return "on"
	}
	return ""
}

func CheckboxValueToBool(val string) bool {
	return val == "on"
}

func NullTime(pgTime pgtype.Timestamp) *time.Time {
	var t *time.Time
	if pgTime.Valid {
		t = &pgTime.Time
	}
	return t
}
