# Golang 的异常处理机制

## error

> `error`就是一个内置的接口

```go
type error interface{
    Error() string
}
```

也就是只要实现了`Error`方法的对象就是一个`error`

### 两种创建 error 的方法

```go
# 1
err1 := errors.New("a error")

# 2 就是对第一种的封装
err2 := fmt.ErrorOf("another error")
```

### 自定义 error

实现`error`接口，就是实现一个`Error`方法即可。

