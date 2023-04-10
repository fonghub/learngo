package internal

import (
	"fmt"
	"reflect"
)

//反射三定律

// FirstLaw 反射第一定律：反射可以将“接口类型变量”转换为“反射类型对象”
func FirstLaw() {
	//接口类型
	var x float64 = 3.4
	//反射类型
	tpe := reflect.TypeOf(x)
	val := reflect.ValueOf(x)
	fmt.Println("type:", tpe)              // type: float64
	fmt.Println("value:", val)             // value: 3.4
	fmt.Println("type kind:", tpe.Kind())  // type kind: float64
	fmt.Println("value kind:", val.Kind()) // value kind: float64
}

// SecondLaw 反射第二定律：反射可以将“反射类型对象”转换为“接口类型变量”
func SecondLaw() {
	//接口类型
	var x uint = 22
	//反射类型
	val := reflect.ValueOf(x)
	//接口类型
	y := val.Interface().(uint)
	fmt.Println(val)
	fmt.Println(val.Interface())
	fmt.Println(y)
}

// ThirdLaw 反射第三定律：如果要修改“反射类型对象”其值必须是“可写的”
func ThirdLaw() {
	//接口类型
	var x float64 = 3.4
	//反射类型
	val := reflect.ValueOf(x)
	fmt.Println("settability of val:", val.CanSet()) //settability of val: false
	// 执行以下语句，会报错，因为 val 变量不可写
	//val.SetFloat(3.55) // panic: reflect: reflect.Value.SetFloat using unaddressable value
	// 如果这行代码能够成功执行，它不会更新 val，
	// 虽然看起来变量 val 是根据 x 创建的，相反它会更新 x 存在于反射对象 val 内部的一个拷贝，
	// 而变量 x 本身完全不受影响。
	// 这会造成迷惑，并且没有任何意义，所以是不合法的。“可写性”就是为了避免这个问题而设计的。

	//反射指针变量，获取指针指向的数据
	val2 := reflect.ValueOf(&x).Elem()
	fmt.Println("settability of val2:", val2.CanSet()) //settability of val2: true
	//通过反射修改变量的值
	val2.SetFloat(3.55)
	//反射类型转接口类型，多种取值方式
	fmt.Println("val2.Interface().(float64)", val2.Interface().(float64))
	fmt.Println("val2.Interface()", val2.Interface())
	fmt.Println("val2", val2)
	fmt.Println("x", x)
}

func ReadAndWriteStruct() {
	type T struct {
		A int
		B string
	}
	t := T{23, "James"}
	//获取”可设置“指针指向的数据
	s := reflect.ValueOf(&t).Elem()
	// 获取数据的类型
	typeOfT := s.Type()
	//遍历结构体字段
	for i := 0; i < s.NumField(); i++ {
		//根据i取结构体字段
		f := s.Field(i)
		//输出结构体字段序号，字段名，值类型，值
		fmt.Printf("%d:%s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	//修改结构体字段
	s.Field(0).SetInt(24)
	s.Field(1).SetString("Kobe")
	fmt.Println("t is now ", t)
}
