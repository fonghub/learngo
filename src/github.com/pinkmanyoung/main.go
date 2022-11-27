package main

import "github.com/pinkmanyoung/effectivego/generics"

func main() {
	//var res any
	//-------------if----------
	//res := iftest.AbsTest(-12)
	//res = iftest.PowTest(3, 3, 10)
	//fmt.Println(res)
	//-------------if----------

	//-------------for----------
	//res := fortest.SumFor(10)
	//res := fortest.SumWhile()
	//fmt.Println(res)
	//fortest.TestFor()
	//f := []float64{1, -2, 3, -4, 5, -6, -7, -8}
	//ff := fortest.AbsIf(f)
	//fmt.Println(ff)
	//-------------for----------

	//-------------switch----------
	//res = switchtest.OsTest()
	//res = switchtest.AdjAge(20)
	//fmt.Println(res)
	//-------------switch----------

	//-------------defer----------
	//res = defertest.DeferTest()
	//res = defertest.DeferMore()
	//fmt.Println(res)
	//-------------defer----------

	//-------------structs----------
	//res = structs.Value()
	//res = structs.Point()
	//fmt.Println(res)
	//ss := structs.Construct("Tom", 23, "ST")
	//fmt.Println(ss)
	//structs.FourYear()
	//-------------structs----------

	//-------------array----------
	//res := arrays.Aget()
	//arrays.Aprint(res)
	//res := arrays.Sget(10)
	//arrays.Sprint(res)
	//arrays.PLenCap(res)
	//arrays.DoSlice(res)
	//arrays.Append(res, 23)
	//fmt.Println(res)
	//res := arrays.AppendSlice(128)
	//arrays.PLenCap(res)
	//-------------array----------

	//-------------map----------
	//res := maps.MapTest()
	//res := maps.MapValue()
	//maps.MapRange(res)
	//maps.MapCURD()
	//s := "Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning an array of substrings of s or an empty list if s contains only white space."
	//res := maps.MapExercise(s)
	//fmt.Println(res)
	//-------------map----------

	//-------------function----------
	//functions.FunctionCall()
	//-------------function----------

	//-------------methods----------
	//methods.Callabs()
	//methods.CallabsFunction()
	//methods.CallabsFloat(-52)
	//methods.Callscale(10)
	//methods.CallabsBoth()
	//methods.CallscaleBoth(5)
	//methods.CallabsFunctionBoth()
	//-------------methods----------

	//-------------interfaces----------
	//interfaces.Test1()
	//interfaces.Test2()
	//interfaces.Test3()
	//interfaces.Test4()
	//interfaces.Test5()
	//interfaces.Test6()
	//interfaces.Exercise()
	//interfaces.Stringer()
	//-------------interfaces----------

	//-------------types----------
	//types.TestType()
	//types.Testdo()
	//-------------types----------

	//-------------generics----------
	generics.NonGeneric()

	//-------------generics----------
}
