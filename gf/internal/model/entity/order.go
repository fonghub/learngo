// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2023-02-13 00:17:08
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Order is the golang structure for table order.
type Order struct {
	Id      int         `json:"id"      ` //
	Uid     int         `json:"uid"     ` //
	OrderId int         `json:"orderId" ` //
	Time    *gtime.Time `json:"time"    ` //
}
