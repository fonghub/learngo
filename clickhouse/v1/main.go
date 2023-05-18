package main

import (
	"database/sql/driver"
	"fmt"
	"time"
	"v1/internal"
)

func main() {
	//Create()
	//Drop()
	//Insert()
	//Select()
	Update()
}

// Create 建表
func Create() {
	createTableSql := "CREATE TABLE default.test_log" +
		"(" +
		"`index` UInt64," +
		"`id` String COMMENT '日志唯一id'," +
		"`string` String," +
		"`timestamp` Int64," +
		"`datetime` DateTime " +
		") " +
		"ENGINE = ReplacingMergeTree " +
		"PARTITION BY toYYYYMM(datetime) " +
		"ORDER BY id " +
		"SETTINGS index_granularity = 8192"
	err := internal.Ch.Init()
	if err != nil {
		fmt.Println(err)
	}
	err = internal.Ch.DoExec(createTableSql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success")
}

// Drop 删表
func Drop() {
	tableName := "default.test_log"
	err := internal.Ch.Init()
	if err != nil {
		fmt.Println(err)
	}
	dropTableSql := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
	err = internal.Ch.DoExec(dropTableSql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success")
}

// Insert 插入
func Insert() {
	var args = map[int][]driver.Value{}
	for i := 0; i < 100; i++ {
		arg := []driver.Value{
			10 + i,
			fmt.Sprintf("%s%d", "md5-", i),
			"test",
			time.Now().UnixMilli(),
			time.Now(),
		}
		args[i] = arg
	}
	sql := `INSERT INTO default.test_log (*) VALUES (?, ?, ?, ?, ?)`
	err := internal.Ch.Init()
	if err != nil {
		fmt.Println("err", err)
	}
	err = internal.Ch.Insert(sql, args)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("success")
}

// Select 查询
func Select() {
	table := "default.test_log"
	err := internal.Ch.Init()
	if err != nil {
		fmt.Println("err", err)
	}
	err = internal.Ch.Optimize(table)
	if err != nil {
		fmt.Println("err", err)
	}
	tailSql := "index>10 limit 20"
	_, err = internal.Ch.DoQuery("*", table, tailSql)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("success")
}

// Update 更新
func Update() {
	sql := "ALTER table default.test_log update `string`='testing' where `index`='10'"
	err := internal.Ch.Init()
	if err != nil {
		fmt.Println("Init", err)
	}
	err = internal.Ch.DoExec(sql)
	if err != nil {
		fmt.Println("DoExec", err)
	}
	fmt.Println("success")
}
