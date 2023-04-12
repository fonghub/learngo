package goroutines

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync/atomic"
	"time"
)

type order struct {
	ID     uint64
	Status int
}

var (
	startNum  uint64
	chanOrder = make(chan uint64, 100)
)

// getUniqId 获取唯一的ID
func getUniqId() uint64 {
	return atomic.AddUint64(&startNum, 1)
}

func getFileName(id uint64) string {
	return fmt.Sprintf("%s%d", "goroutines/order/", id)
}

// WriteToFile 参数v为map或者struct
func WriteToFile(v any, fileName string) error {
	content, _ := json.Marshal(v)
	err := os.WriteFile(fileName, content, 0666)
	if err != nil {
		return err
	}
	return nil
}

func ReadFromFile(fileName string) (error, map[string]interface{}) {
	byteContent, err := os.ReadFile(fileName)
	if err != nil {
		return err, nil
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(byteContent, &res)
	if err != nil {
		return err, nil
	}
	return nil, res
}

// setOrder 生成订单，保存到文件
func setOrder() (error, uint64) {
	id := getUniqId()
	log.Println("setOrder=", id)
	order := &order{
		ID:     id,
		Status: 0,
	}
	fileName := getFileName(id)
	err := WriteToFile(order, fileName)
	if err != nil {
		return err, 0
	}
	return err, id
}

// updateOrder 每隔1000ms处理一个订单
func updateOrder() {
	//range 会反复从通道接收值，直到它关闭.
	//通道关闭，实际上时发送端关闭，接收端依然可以正常接收数据
	for id := range chanOrder {
		log.Println("updateOrder =", id)
		time.Sleep(1000 * time.Millisecond)
		fileName := getFileName(id)
		err, content := ReadFromFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		content["Status"] = 1
		err = WriteToFile(content, fileName)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("chanOrder has closed.")
}

// setStartNum 获取起始唯一ID
func setStartNum() {
	index := fmt.Sprintf("%s", "goroutines/index")
	err, content := ReadFromFile(index)
	if err != nil {
		log.Println(err)
	} else {
		startNum = uint64(content["index"].(float64))
	}
	log.Println("setStartNum= ", startNum)
}

// saveStartNum 保存唯一ID
func saveStartNum() {
	var content = make(map[string]uint64)
	content["index"] = startNum
	index := fmt.Sprintf("%s", "goroutines/index")
	err := WriteToFile(content, index)
	if err != nil {
		log.Fatal(err)
	}
}

// createOrder 每隔100ms创建一个订单，模拟每次20个订单
func createOrder() {
	setStartNum()
	endNum := startNum + 20
	for {
		err, orderId := setOrder()
		if err != nil {
			log.Fatal(err)
		}
		//订单ID发到chan
		chanOrder <- orderId
		time.Sleep(100 * time.Millisecond)
		if orderId >= endNum {
			close(chanOrder)
			log.Println("close the chanOrder.")
			break
		}
	}
	saveStartNum()
}

// MakeOrder 模拟产生订单数据，通过chan发送订单ID，使用goroutine从chan取订单ID，修改订单的状态
func MakeOrder() {
	go updateOrder()
	createOrder()
	for {

	}
	fmt.Println("make order success")
}
