package main

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)){
	// fn("WTF is happening")

	val := getValue(x)
	
	for i :=0 ; i < val.NumField(); i++{
		field := val.Field(i)
		fmt.Println(field)

		switch field.Kind(){
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn )
		}
	}
}

func getValue(x interface{} ) reflect.Value  {

	val :=  reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr{
		fmt.Println("element",val.Elem())
		val = val.Elem()
	}
	return val
}