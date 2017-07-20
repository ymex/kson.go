package sonar

import (
	"errors"
	"encoding/json"
	"reflect"
	"strconv"
	"log"
)

type TypeTransform struct {
	data interface{}
	err  error
}

func NewTypeTransform(data interface{}) *TypeTransform {
	return &TypeTransform{data: data}
}

func (t *TypeTransform)Interface() interface{}  {
	return t.data
}

//Bool guarantees the return of a `bool` (with optional default) and error
func (t *TypeTransform) Bool(values ... bool) (bool, error) {
	var def bool = false

	if len(values) > 1 {
		log.Panicf("Bool() received too many arguments %d", len(values))
	}
	if len(values) == 1 {
		def = values[0]

	}
	if s, ok := (t.data).(bool); ok {
		return s, nil
	}
	return def, errors.New("type assertion to bool failed")
}

//Float64 guarantees the return of a `float64` (with optional default) and error
func (t *TypeTransform) Float64(values ... float64) (float64, error) {
	var def float64

	if len(values) > 1 {
		log.Panicf("Float64() received too many arguments %d", len(values))
	}
	if len(values) == 1 {
		def = values[0]

	}

	switch t.data.(type) {
	case json.Number:
		return t.data.(json.Number).Float64()
	case float32, float64:
		return reflect.ValueOf(t.data).Float(), nil
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(t.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(t.data).Uint()), nil
	}
	return def, errors.New("invalid value type")
}

//Int guarantees the return of a `int` (with optional default) and error
func (t *TypeTransform) Int(values ... int) (int, error) {

	var def int

	if len(values) > 1 {
		log.Panicf("Int() received too many arguments %d", len(values))
	}
	if len(values) == 1 {
		def = values[0]

	}

	switch t.data.(type) {
	case json.Number:
		i, err := t.data.(json.Number).Int64()
		return int(i), err
	case float32, float64:
		return int(reflect.ValueOf(t.data).Float()), nil
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(t.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(t.data).Uint()), nil
	}
	return def, errors.New("invalid value type")
}

// Int64 guarantees the return of a `int64` (with optional default) and error
func (t *TypeTransform) Int64(values ... int64 ) (int64, error) {

	var def int64

	if len(values) > 1 {
		log.Panicf("MustInt64() received too many arguments %d", len(values))
	}
	if len(values) == 1 {
		def = values[0]

	}

	switch t.data.(type) {
	case json.Number:
		return t.data.(json.Number).Int64()
	case float32, float64:
		return int64(reflect.ValueOf(t.data).Float()), nil
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(t.data).Int(), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(t.data).Uint()), nil
	}
	return def, errors.New("invalid value type")
}

// Uint64 guarantees the return of a `uint64` (with optional default) and error
func (t *TypeTransform) Uint64(values ... uint64) (uint64, error) {

	var def uint64

	if len(values) > 1 {
		log.Panicf("MustUint64() received too many arguments %d", len(values))
	}
	if len(values) == 1 {
		def = values[0]

	}

	switch t.data.(type) {
	case json.Number:
		return strconv.ParseUint(t.data.(json.Number).String(), 10, 64)
	case float32, float64:
		return uint64(reflect.ValueOf(t.data).Float()), nil
	case int, int8, int16, int32, int64:
		return uint64(reflect.ValueOf(t.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(t.data).Uint(), nil
	}
	return def, errors.New("invalid value type")
}

// String  guarantees the return of a `string` (with optional default) and error
func (t *TypeTransform) String(values ... string) (string, error) {
	var def string = ""

	if len(values) > 1 {
		log.Panicf("String() received too many arguments %d", len(values))
	}
	if len(values) == 1 {
		def = values[0]

	}
	if s, ok := (t.data).(string); ok {
		return s, nil
	}
	return def, errors.New("type assertion to string failed")
}

// Bytes guarantees the return of a `[]byte` (with optional default) and error
func (t *TypeTransform) Bytes(values ... []byte) ([]byte, error) {

	var def []byte

	if len(values) > 1 {
		log.Panicf("Bytes() received too many arguments %d", len(values))
	}
	if len(values) == 1 {
		def = values[0]

	}

	if s, ok := (t.data).(string); ok {
		return []byte(s), nil
	}
	return def, errors.New("type assertion to []byte failed")
}

// Map guarantees the return of a `map[string]interface{}` (with optional default) and error
func (t *TypeTransform) Map(values ... map[string]interface{}) (map[string]interface{}, error) {

	var def map[string]interface{}

	if len(values) > 1 {
		log.Panicf("Map() received too many arguments %d", len(values))
	}
	if len(values) == 1 {
		def = values[0]

	}

	if m, ok := (t.data).(map[string]interface{}); ok {
		return m, nil
	}
	return def, errors.New("type assertion to map[string]interface{} failed")
}

// Array guarantees the return of a `array` (with optional default) and error
func (t *TypeTransform) Array(values ... []interface{}) ([]interface{}, error) {
	var def []interface{}
	if len(values) > 1 {
		log.Panicf("Array() received too many arguments %d", len(values))
	}
	if len(values) == 1 {
		def = values[0]

	}
	if a, ok := (t.data).([]interface{}); ok {
		return a, nil
	}
	return def, errors.New("type assertion to []interface{} failed")
}

// StringArray guarantees the return of a  `array` of `string` (with optional default) and error
func (t *TypeTransform) StringArray(values ... []string) ([]string, error) {

	var def []string

	if len(values) > 1 {
		log.Panicf("StringArray() received too many arguments %d", len(values))
	}else if len(values) == 1 {
		def = values[0]

	}

	arr, err := t.Array()
	if err != nil {
		return def, err
	}
	retArr := make([]string, 0, len(arr))
	for _, a := range arr {
		if a == nil {
			retArr = append(retArr, "")
			continue
		}
		s, ok := a.(string)
		if !ok {
			return nil, errors.New("type assertion to []string failed")
		}
		retArr = append(retArr, s)
	}
	return retArr, nil
}

func (t *TypeTransform)ToString(values ... string) string {
	var def string

	if len(values) > 1 {
		log.Panicf("String() received too many arguments %d", len(values))
	}else if len(values) == 1 {
		def = values[0]

	}
	if val , err := t.String();err == nil {
		return val
	}else if val , err := t.Float64();err == nil {
		return strconv.FormatFloat(val, 'f', -1, 64)
	}else if val ,err := t.Int64();err == nil {
		return strconv.FormatInt(val,10)
	}else if val, err := t.Bool(); err == nil {
		return strconv.FormatBool(val)
	}

	t.err = errors.New("type assertion to string failed")
	return def
}

func (t *TypeTransform)ToBool(values ... bool) bool {
	var def bool

	if len(values) > 1 {
		log.Panicf("String() received too many arguments %d", len(values))
	}else if len(values) == 1 {
		def = values[0]

	}
	if val , err := t.Bool(); err == nil {
		return val
	}else if val, err := t.String();err == nil{
		if bval , berr := strconv.ParseBool(val) ;berr == nil{
			return bval
		}
	}
	return def
}

func (t *TypeTransform)ToInt(values ... int) int {
	var def int

	if len(values) > 1 {
		log.Panicf("String() received too many arguments %d", len(values))
	}else if len(values) == 1 {
		def = values[0]

	}
	if val, err := t.Int();err == nil {
		return val
	}else if val,err :=t.String(); err == nil {
		if ival,ierr := strconv.ParseInt(val,10,64);ierr == nil {
			return int(ival)
		}
	}
	return def
}

func (t *TypeTransform)ToFloat(values ... float64) float64 {
	var def float64

	if len(values) > 1 {
		log.Panicf("String() received too many arguments %d", len(values))
	}else if len(values) == 1 {
		def = values[0]

	}
	if val, err := t.Float64();err == nil {
		return val
	}else if val,err :=t.String(); err == nil {
		if fval,ferr := strconv.ParseFloat(val,64);ferr == nil {
			return fval
		}
	}
	return def
}

func (t *TypeTransform)ToMap(values ... map[string]interface{})  map[string]interface{} {
	var def map[string]interface{}

	if len(values) > 1 {
		log.Panicf("String() received too many arguments %d", len(values))
	}else if len(values) == 1 {
		def = values[0]

	}
	if val,err := t.Map();err == nil {
		return val
	}
	return def
}