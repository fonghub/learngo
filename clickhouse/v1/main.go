package main

import (
	"database/sql/driver"
	"fmt"
	"time"
	"v1/internal"
)

func main() {
	Insert()
}

// Create 建表
func Create() {
	createTableSql := "CREATE TABLE default.example\n(\n    `id`            UInt64      COMMENT 'id',\n    `md5`           String      COMMENT '单局id',\n    `win`           UInt32      COMMENT '赢得',\n    `betting`       UInt32      COMMENT '投注',\n    `end_chips`     UInt32      COMMENT '结束筹码',\n    `start_chips`   UInt32      COMMENT '开始筹码',\n    `game_id`       UInt16      COMMENT '游戏ID 整型',\n    `client`        UInt8       COMMENT '设备类型 1.PC端 2.IOS 3安卓横 4.其它',\n    `time`          UInt64      COMMENT '游戏时间',\n    `dtime`         DateTime    COMMENT '游戏时间',\n    `ip`            String      COMMENT '玩家IP地址',\n    `player_uid`    UInt32      COMMENT '玩家ID用户信息里面获取',\n    `currency`      String      COMMENT '货币种类用户信息里面获取',\n    `agent_id`      UInt16      COMMENT '总代理商id',\n    `sub_agent_id`  UInt16      COMMENT '子代理商id',\n    `status`        UInt8       COMMENT '订单状态 1.已下注 2.已结算',\n    `real_bet`      UInt32      COMMENT '有效下注',\n    `commission`    UInt32      COMMENT '抽水',\n)\nENGINE = ReplacingMergeTree\nPARTITION BY toYYYYMM(dtime)\nORDER BY id\nSETTINGS index_granularity = 8192\nCOMMENT '游戏日志';"
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
	tableName := "default.example"
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
			"RU",
			10 + i,
			100 + i,
			[]int16{1, 2, 3},
			time.Now(),
			time.Now(),
		}
		args[i] = arg
	}
	sql := `INSERT INTO example (country_code, os_id, browser_id, categories, action_day, action_time) VALUES (?, ?, ?, ?, ?, ?)`
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

func optimize(table string) error {
	err := internal.Ch.DoExec(fmt.Sprintf("optimize %s replace_test final", table))
	if err != nil {
		return err
	}
	return nil
}

// Select 查询
func Select() {
	table := "example"
	err := internal.Ch.Init()
	if err != nil {
		fmt.Println("err", err)
	}
	err = optimize(table)
	if err != nil {
		fmt.Println("err", err)
	}
	tailSql := "limit 20"
	_, err = internal.Ch.DoQuery("*", table, tailSql)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("success")
}

// Update 更新
func Update() {
	sql := `ALTER table default.example update country_code='cn' where os_id='72'`
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
