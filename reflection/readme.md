
[reflection](http://c.biancheng.net/view/4407.html)

Type 	reflect.Type	reflect.TypeOf
Value	reflect.Value	reflect.ValueOf

使用reflect.TypeOf函数可以获得值的 类型变量 reflect.Type
通过 类型变量 可以访问值的 类型信息

类型 Type 与 种类 Kind
类型 Type 有： int、string、bool、float32 等类型，以及使用 type 关键字定义的类型
种类（Kind）指的是对象归属的品种，是比类型更本质更基础的分类



反射三定律
- 反射可以将“接口类型变量”转换为“反射类型对象”
- 反射可以将“反射类型对象”转换为“接口类型变量”
- 如果要修改“反射类型对象”其值必须是“可写的”