package internal

import (
	"database/sql/driver"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"sync"
)

type ch struct {
	mu   sync.Mutex
	conn driver.Conn
}

var (
	Ch  = new(ch)
	dsn = "tcp://192.168.99.105:9000?username=default&password=dm123456&database=default&compress=true&debug=true"
)

// Init 初始化 获得conn对象
func (c *ch) Init() error {
	if c.conn == nil {
		c.mu.Lock()
		defer c.mu.Unlock()
		conn, err := clickhouse.Open(dsn)
		if err != nil {
			return err
		}
		c.conn = conn
	}
	return nil
}

// Insert 插入
func (c *ch) Insert(sql string, args map[int][]driver.Value) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return err
	}
	stmt, err := c.conn.Prepare(sql)
	if err != nil {
		return err
	}
	for _, value := range args {
		if _, err = stmt.Exec(value); err != nil {
			return err
		}
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

// Update 更新（带参数）
func (c *ch) Update(sql string, args []driver.Value) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return err
	}
	stmt, err := c.conn.Prepare(sql)
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(args); err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

// DoExec 建表 删表 更新（不带参数） 删除
func (c *ch) DoExec(sql string) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return err
	}
	stmt, err := c.conn.Prepare(sql)
	if err != nil {
		return err
	}
	if _, err = stmt.Exec([]driver.Value{}); err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// DoQuery 查询
func (c *ch) DoQuery(fieldSql string, tableName string, tailSql string) ([]g.MapStrStr, error) {
	tx, err := c.conn.Begin()
	if err != nil {
		return nil, err
	}
	stmt, err := c.conn.Prepare(fmt.Sprintf("select %s from %s where %s", fieldSql, tableName, tailSql))
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query([]driver.Value{})
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	columns := rows.Columns()
	row := make([]driver.Value, len(columns))
	var ros []g.MapStrStr
	for rows.Next(row) == nil {
		ro := make(g.MapStrStr)
		for i, c := range columns {
			ro[c] = gconv.String(row[i])
		}
		ros = append(ros, ro)
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return ros, nil
}

// Optimize 执行optimize手动触发合并
func (c *ch) Optimize(table string) error {
	err := c.DoExec(fmt.Sprintf("optimize table %s final", table))
	if err != nil {
		return err
	}
	return nil
}

// DoCount 返回行数
func (c *ch) DoCount(tableName string, tailSql string) (int, error) {
	res, err := c.DoQuery("COUNT() as count", tableName, tailSql)
	if err != nil {
		return 0, err
	}
	return gconv.Int(res[0]["count"]), nil
}

// DoMaxId 返回最大id
func (c *ch) DoMaxId(tableName string, tailSql string) (uint64, error) {
	res, err := c.DoQuery("MAX(id) as max_id", tableName, tailSql)
	if err != nil {
		return 0, err
	}
	return gconv.Uint64(res[0]["max_id"]), nil
}

// DoGetField 返回字段的值
func (c *ch) DoGetField(field string, ret string, tableName string, tailSql string) (uint64, error) {
	res, err := c.DoQuery(field, tableName, tailSql)
	if err != nil {
		return 0, err
	}
	if len(res) < 1 {
		return 0, nil
	} else {
		return gconv.Uint64(res[0][ret]), nil
	}
}

// GetTailSql 返回拼接的sql
func (c *ch) GetTailSql(tailSql, orderBy string, pageSize, pageNum int) string {
	if orderBy == "" {
		orderBy = "id desc"
	}
	tailSql = fmt.Sprintf("%s order by %s", tailSql, orderBy)

	if pageSize == 0 {
		pageSize = 10
	}
	tailSql = fmt.Sprintf("%s limit %d", tailSql, pageSize)
	if pageNum == 0 {
		pageNum = 1
	}
	tailSql = fmt.Sprintf("%s offset %d", tailSql, (pageNum-1)*pageSize)
	return tailSql
}
