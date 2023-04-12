package goroutines

func CallSay() {
	//协程序，协程是一种用户态的轻量级线程
	//go Say("hello")
	//主程序，主程序结束时不会等待协程执行完毕，结果时输出“hello”可能比“world”少
	//Say("world")

	//协程的调度不保证执行的顺序
	//go log.Println("goroutine start")
	//Seq("goroutine")
	//go func() { log.Println("使用协程调用匿名函数") }()
	//go log.Println("goroutine end")
	//time.Sleep(1000 * time.Millisecond)

	//使用WaitGroup可以保证协程全部执行
	//WaitGp()

	//并发不安全，会出现脏读数据；使用锁可以保证并发安全
	//CallSecKill()

	//协程中的map并发不安全
	//Maps()
	//SyncMaps()

	//加减法生成工具
	//MakeHomeWork(100, 20)

	//模拟创建订单，异步处理订单
	MakeOrder()
}
