# Websocket

目前掌握的讯息已经很明确了
- 全双工
- 较小控制开销
- 心跳机制
- wss

!!!! 推荐看这个仓库的Readme. [link1st/GoWebsocket](https://github.com/link1st/gowebsocket)

[煎鱼的WebSocket系列](https://golang2.eddycjy.com/posts/ch4/02-protocol/)，了解握手协议和细节，例如心跳 onclose

跟着实现聊天室

在看看群友的 [gws的example](https://github.com/lxzan/gws/tree/master/examples)，了解chatroom和client、wss的实现细节。复现一遍，可以的话理解后用gorilla实现。