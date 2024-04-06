# Context

```go
type Context interface {    
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}    
    Err() error
    Value(key interface{}) interface{}
}

```

## emptyCtx
`emptyCtx` 实现了Context的四个方法，并提供了实例化的两个变量 `Background` 和 `TODO`
```go
// An emptyCtx is never canceled, has no values, and has no deadline.
// It is the common base of backgroundCtx and todoCtx.
type emptyCtx struct{}

func (emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (emptyCtx) Done() <-chan struct{} {
	return nil
}

func (emptyCtx) Err() error {
	return nil
}

func (emptyCtx) Value(key any) any {
	return nil
}

type backgroundCtx struct{ emptyCtx }

func (backgroundCtx) String() string {
	return "context.Background"
}

type todoCtx struct{ emptyCtx }

func (todoCtx) String() string {
	return "context.TODO"
}

```

`context.Background()` 和 `TODO()` 都是调用 `emptyCtx`


## valueCtx

### ValueCtx struct
```go
type valueCtx struct {
	Context
	key, val interface{}
}

func (c *valueCtx) Value(key interface{}) interface{} {
	if c.key == key {
		return c.val
	}
	return c.Context.Value(key)
}

```
`valueCtx` 利用一个 Context 类型的变量来表示父节点 context，所以当前 Context 继承了**父Context的所有信息**。

### WithValue
`WithValue` 用于向Context添加 Key-Value-pair:
```go
func WithValue(parent Context,key,val interface{}) Context{
    if key == nil{
        panic("nil key")
        }
    if !reflect.TypeOf(key).Comparable(){
        panic("key is not comparable")
    }
    return &valueCtx{parent,key,val}
}
```
这里添加键值对不是在原 context 结构体上直接添加，而是以此 context 作为父节点，**重新创建一个新的 `valueCtx` 子节点**，将键值对添加在子节点上，由此形成一条**context链**。获取value的过程就是在这条 context 链上由尾部上前搜寻：
![](/static/img/ctx.valctx.png)

## cancelCtx
```go

// A cancelCtx can be canceled. When canceled, it also cancels any children
// that implement canceler.
type cancelCtx struct {
	Context

	mu       sync.Mutex            // protects following fields
	done     atomic.Value          // of chan struct{}, created lazily, closed by first cancel call
	children map[canceler]struct{} // set to nil by the first cancel call
	err      error                 // set to non-nil by the first cancel call
	cause    error                 // set to non-nil by the first cancel call
}

// A canceler is a context type that can be canceled directly. The
// implementations are *cancelCtx and *timerCtx.
type canceler interface {
	cancel(removeFromParent bool, err, cause error)
	Done() <-chan struct{}
}
```

### cancelCtx struct

## Summary

context 主要用于父子任务直接的同步取消信号，本质上是一种协程调度的方式。

上游任务仅仅使用 Context 通知下游任务不再需要，但不会直接干涉和终端下游任务的执行，  
由下游任务执行决定后续的处理操作，就是说 Context 的取消操作是**无侵入的**。  
Context 是线程安全的，因为 Context 本身是不可变的，所以可以放心在多协程中传递使用。

> More:[Context - Go面试宝典](https://golang.design/go-questions/stdlib/context/why/)