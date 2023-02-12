# Testing 库

Testing 是一个 go 提供自动测试标准库，使用 go test 指令会自动执行任意格式为 TestXxx 的测试方法，Xxx 是被测试的方法名

测试文件推荐与被测试文件位于同一个包文件下，并使用与被测试文件名字相同，在后面加上_test尾缀

[testing](https://pkg.go.dev/testing)

## 执行指令 go test

### 单元测试

`go test -run 测试方法名`

### 基准测试

`go test -bench 测试方法名`

[go test](https://pkg.go.dev/cmd/go#hdr-Testing_flags)