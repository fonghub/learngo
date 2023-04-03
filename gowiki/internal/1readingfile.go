package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

// Save 保存到文件
func (p *Page) Save() error {
	filename := p.Title + ".txt"
	filePath := p.GetPath(filename)
	return os.WriteFile(filePath, p.Body, 0600)
}

// GetPath 获取文件路径
func (p *Page) GetPath(filename string) string {
	return fmt.Sprintf("%s%s", "./data/", filename)
}

// LoadPage 从文件加载
func (p *Page) LoadPage(title string) (map[string]interface{}, error) {
	filename := title + ".txt"
	filePath := p.GetPath(filename)
	body, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var data = make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
