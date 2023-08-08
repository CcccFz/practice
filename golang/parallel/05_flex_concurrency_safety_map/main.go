package main

import (
	"fmt"
	"reflect"
	"sync"
)

type ConcurrentMap struct {
	m       sync.Map
	keyType reflect.Type
	valType reflect.Type
}

func (cMap *ConcurrentMap) Load(key interface{}) (val interface{}, ok bool) {
	if cMap.keyType != reflect.TypeOf(key) {
		return
	}
	return cMap.m.Load(key)
}

func (cMap *ConcurrentMap) Store(key, val interface{}) error {
	keyType, valType := reflect.TypeOf(key), reflect.TypeOf(val)

	if cMap.keyType == nil || cMap.valType == nil {
		cMap.keyType, cMap.valType = keyType, valType
		cMap.m.Store(key, val)
		return nil
	}

	if cMap.keyType != keyType {
		return fmt.Errorf("存入了错误的key类型: %s", keyType.String())
	}
	if cMap.valType != valType {
		return fmt.Errorf("存入了错误的val类型: %s", valType.String())
	}

	cMap.m.Store(key, val)
	return nil
}

func main() {
	var m ConcurrentMap
	var err error

	if err = m.Store(1, "aa"); err != nil {
		panic(err)
	}

	if err = m.Store(2, "bb"); err != nil {
		panic(err)
	}

	if v, ok := m.Load(2); ok {
		fmt.Println(v.(string))
	}
}
