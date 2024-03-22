# channel

[Channel 实现广播](https://juejin.cn/post/6844903857395335182)

- 创建广播 Broadcast channel (Eventbus)
- 创建Observers channel，用于保存所有需要监听事件的 chan(observer), 是一个 observer数组
- 添加 Observer 请求的channel, 类型是chan 里保存(chan int) [chan chan int]
- 移除 Observer 请求的channel, 类型是chan 里保存(chan int) [chan chan int]
- 