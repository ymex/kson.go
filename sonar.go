package sonar

import (
	"encoding/json"
	"reflect"
	"strconv"
	"regexp"
	"errors"
	"strings"
	"log"
)

type Sonar struct {
	data  interface{}
	store map[string]interface{}
	keys  []string
	err   error
}

func NewSonar(data interface{}) *Sonar {
	kson := new(Sonar)
	if val, ok := data.(string); ok {
		kson.err = kson.UnmarshalJSON([]byte(val))
	} else if val, ok := data.([]byte); ok {
		kson.err = kson.UnmarshalJSON(val)
	} else {
		kson.err = errors.New("Sonar: cannot unmarshal data into Go value of type " + reflect.TypeOf(data).String())
	}
	//if sonar.HasError() {
	//	log.Panicln(sonar.err)
	//}
	return kson
}

func (k *Sonar) UnmarshalJSON(p []byte) error {
	return json.Unmarshal(p, &k.data)
}

func (k *Sonar) Interface() interface{} {
	return k.data
}

func (k *Sonar) Encode() ([]byte, error) {
	return k.MarshalJSON()
}

func (k *Sonar) EncodePretty() ([]byte, error) {
	return json.MarshalIndent(&k.data, "", "  ")
}

func (k *Sonar) MarshalJSON() ([]byte, error) {
	return json.Marshal(&k.data)
}

func (k *Sonar) Find(keys ... string) *Sonar {
	if len(keys) <= 0 {
		log.Panicln("Find() received have no arguments")
	}
	k.clear()
	k.keys = keys
	for _, key := range keys {
		err := k.parseLink(k.data, key)
		if err != nil {
			var text string
			if k.err !=nil {
				text = k.err.Error()+"、 "
			}
			k.err = errors.New(text+err.Error())
		}
	}
	if k.HasError() {
		log.Println(k.err)
	}
	return k
}

func (k *Sonar) HasError() bool {
	return k.err != nil
}

func (k *Sonar) clear() {
	k.store = make(map[string]interface{}, 0)
	k.keys = make([]string, 0)
}

//key->karr[1][2]
func (k *Sonar) parseLink(data interface{}, key string) error {

	aliasKey := k.aliasKey(key)
	realKey := k.realKey(key)
	subkeys := strings.Split(realKey, "->")

	var result interface{} = nil
	var err error = nil
	if len(subkeys) == 1 {
		result, err = k.parse(data, realKey)
	} else {
		for _, subkey := range subkeys {
			if result == nil {
				result, err = k.parse(data, subkey)
			} else {
				result, err = k.parse(result, subkey)
			}
		}
	}
	if err != nil {
		return err
	}

	k.store[aliasKey] = result
	return nil
}

//key[2][3]
func (k *Sonar) parse(data interface{}, key string) (interface{}, error) {

	indexs, err := getMutArrayIndexs(key)
	if err != nil {
		return nil, err
	}
	if isArrayKey(len(indexs), key) {
		var arrdata interface{} = nil
		arrayName := getMutArrayName(key)

		if len(arrayName) != 0 {
			if m, ok := data.(map[string]interface{}); ok {
				if val, ok := m[arrayName]; ok {
					arrdata = val
				}
			}
		} else {
			arrdata = data
		}

		if arrdata == nil {
			return nil, errors.New("can't find the valid key: '" + key+"'")
		}
		for i := 0; i < len(indexs); i++ {
			if lr, ok := arrdata.([]interface{}); ok {
				if i == len(indexs)-1 {
					if indexs[i] > len(lr)-1 {
						return nil, errors.New(key + " index out of range:" + strconv.Itoa(indexs[i]))
					}
					return lr[indexs[i]], nil
				} else {
					arrdata = lr[i]
				}
			}
		}
	}
	if m, ok := data.(map[string]interface{}); ok {
		if val, ok := m[key]; ok {
			return val, nil
		}
		return nil, errors.New("can't find the valid key: '" + key+"'")
	}
	return nil, errors.New("type assertion to map[string]interface{} failed")
}

func (k *Sonar) aliasKey(key string) string {
	index := strings.Index(key, ":")
	if index <= 0 || index == len(key)-1 {
		return key
	}
	return key[0:index]
}

func (k *Sonar) realKey(key string) string {
	index := strings.Index(key, ":")
	if index <= 0 || index == len(key)-1 {
		return key
	}
	return key[index+1:]
}

func getMutArrayName(key string) string {
	return key[0:strings.Index(key, "[")]
}

func isArrayKey(indexs int, key string) bool {
	return indexs >= 1 && strings.HasSuffix(key, "]")
}

func getMutArrayIndexs(arrayName string) ([]int, error) {
	indexs := make([]int, 0)
	reg := regexp.MustCompile("\\[(.+?)\\]")
	mt := reg.FindAllString(arrayName, -1)
	for _, val := range mt {
		if ival, err := strconv.Atoi(val[1:len(val)-1]); err == nil {
			indexs = append(indexs, ival)
		} else {
			return indexs, errors.New("key Grammatical errors")
		}
	}
	return indexs, nil
}

func (k *Sonar) Got(key string) *TypeTransform {
	val, ok := k.store[key]
	if ok {
		return &TypeTransform{data: val}
	}
	return &TypeTransform{data: nil}
}

func (k *Sonar) GotFirst() *TypeTransform {
	key := k.keys[0]
	return k.Got(key)
}

func (k *Sonar) GotLast() *TypeTransform {

	key := k.keys[len(k.keys)-1]
	return k.Got(key)
}

func (k *Sonar) GotPosition(index int) *TypeTransform {
	if index > len(k.keys)-1 || index <= 0 {
		log.Panicln(strconv.Itoa(index) + " index out of range")
	}
	key := k.keys[index]
	return k.Got(k.aliasKey(key))
}

func (k *Sonar) checkKeys() {
	if len(k.keys) <= 0 {
		log.Panicln("sonar have no key in it.")
	}
}

// Get returns a pointer to a new `Json` object
// for `key` in its `map` representation

func (k *Sonar) Get(key string) *Sonar {
	m, err := k._map()
	if err == nil {
		if val, ok := m[key]; ok {
			return &Sonar{data: val, err: nil}
		}
	}
	return &Sonar{data: nil, err: err}
}

func (k *Sonar) GetIndex(index int) *Sonar {
	a, err := k._array()
	if err == nil {
		if len(a) > index {
			return &Sonar{data: a[index], err: nil}
		}
	}
	return &Sonar{data: nil, err: err}
}

func (k *Sonar) Type() *TypeTransform {
	return &TypeTransform{data: k.data}
}

func (k *Sonar) _map() (map[string]interface{}, error) {
	if m, ok := (k.data).(map[string]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("type assertion to map[string]interface{} failed")
}

func (k *Sonar) _array() ([]interface{}, error) {
	if a, ok := (k.data).([]interface{}); ok {
		return a, nil
	}
	return nil, errors.New("type assertion to []interface{} failed")
}
