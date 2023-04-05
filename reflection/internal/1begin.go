package internal

import (
	"fmt"
	"reflect"
)

type Enum int

const Zero Enum = 0

type B byte

const b B = 0

func Demo1() {

	typeOfA := reflect.TypeOf(Zero)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	typeOfB := reflect.TypeOf(b)
	fmt.Println(typeOfB.Name(), typeOfB.Kind())

	type cat struct {
		Name string
		Type int `json:"type" id:"100"`
	}
	typeOfCat := reflect.TypeOf(cat{})
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	//对指针变量反射，无法通过Name和Kind方法获取变量的类型和种类，
	//而是要通过以下方法
	ins := &cat{}
	typeOfIns := reflect.TypeOf(ins)
	fmt.Println(typeOfIns.Name(), typeOfIns.Kind())

	//通过Elem方法获取指针指向的元素的类型，这个获取的过程称为取元素
	typeOfIns = typeOfIns.Elem()
	fmt.Println(typeOfIns.Name(), typeOfIns.Kind())

	//结构体属性
	inIns := cat{Name: "mini", Type: 1}
	typeOfinIns := reflect.TypeOf(inIns)
	for i := 0; i < typeOfinIns.NumField(); i++ {
		fieldType := typeOfinIns.Field(i)
		fmt.Printf("name:%v ,tag:'%v'\n", fieldType.Name, fieldType.Tag)
	}

	if catType, ok := typeOfinIns.FieldByName("Type"); ok {
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}
