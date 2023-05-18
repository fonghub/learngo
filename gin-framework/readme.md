[gin framework](https://gin-gonic.com/zh-cn/docs/)


Engine 结构体

RouterGroup 结构体是 Engine 结构体的一个匿名字段

RouterGroup 结构体

实现 IRouter 和 IRoutes 接口

所以Engine结构体也实现了 IRouter 和 IRoutes 接口


IRouter 接口
包含 IRoutes 接口，比 IRoutes 多一个 Group 方法

IRoutes 接口