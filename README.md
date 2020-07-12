# go-panic-guide

🧠 **Go Panic 指南：收录常见触发 panic 的场景、原因解析与示例代码。**

本项目旨在帮助开发者深入理解 Go 程序在运行中可能出现 panic 的各种情况，以及如何正确地捕获、规避和测试这些 panic。

---

## 💡 为什么会 Panic？

`panic` 是 Go 语言在运行时遇到不可恢复错误时的机制，例如：
- 访问越界
- 类型断言失败
- 并发安全错误
- 通道阻塞
- 非法的内存访问等

---

## 📚 收录的 Panic 场景

| 序号 | 场景               | 是否可 recover | 示例函数                  |
|------|--------------------|----------------|---------------------------|
| 1    | 类型断言错误       | ✅ 可捕获      | `typeAssertionPanic()`       |
| 2    | map 并发写入       | ❌ 不可捕获    | `concurrentMapPanic()`    |
| 4    | 切片操作越界       | ✅ 可捕获      | `sliceOutOfRangePanic()`  |
| 5    | 通道阻塞（死锁）   | ❌ 不可捕获    | `channelDeadlockPanic()`  |
| 6    | 关闭已关闭的通道   | ✅ 可捕获      | `closeClosedChannel()`    |

---

## ⚠️ 关于通道死锁的说明

以下写法会导致程序直接 `fatal error`，而不是 `panic`，因此 `recover()` 无法捕获：

```go
func channelDeadlockPanic() {
    ch := make(chan int)
    ch <- 1 // 阻塞，无接收者，导致全部 goroutine 睡眠，引发 fatal error
}
```

**推荐的测试方式**：使用 `select` + `time.After` 来检测阻塞，而不是 `recover`。

---

## 🧪 单元测试说明

使用 [`testing`](https://pkg.go.dev/testing) 进行每个 panic 场景的验证，并尽量进行错误恢复以保证测试不中断。

---

## 🚀 如何运行

```bash
go test -v ./...
```

---

## 📂 目录结构

```
go-panic-guide/
├── go.mod
├── oparray # 切片操作
├── opchannel # 通道操作
├── opmap # 字典操作
├── panicking # panic 示例
├── README.md # 项目说明文档
└── typeassertion # 类型推断
```

---

## 📌 推荐阅读

- Go 官方文档：[https://go.dev/doc/](https://go.dev/doc/)
- Go blog - panic & recover：[https://blog.golang.org/defer-panic-and-recover](https://blog.golang.org/defer-panic-and-recover)

---

## 📜 License

MIT License.
