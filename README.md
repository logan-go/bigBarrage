# bigBarrage
一个高并发弹幕系统

```
主要系统功能：
    ConnectKeeper
        持有和管理所有的用户连接
        管理连接过来的所有房间，房间里面都是连接
        连接Redis进行订阅
    Register
        管理和监控所有ConnectKeeper和MsgSender的活跃状态和负载状态
    ~~MsgSender~~
        ~~从队列里面获取要发送的内容，发送给对应的ConnectKeeper~~
```

===

###主要角色和设定讲解

```
ROOM：
    1.RoomID：每个消息互通的用户集合为一个Room，同一个Room内的消息是互相传递的，不同Room之间的消息是不会相互传递的。
    2.SubRoom：子房间，子房间的RoomID跟主RoomID是完全一直的，只是一个Room是由一个或者多个SubRoom组成的，不同的SubRoom分别散列在不同的ConnectKeeper里面。

Register：
    1.记录所有ConnectKeeper的地址，端口，状态
    2.记录所有MsgKeeper的地址，端口，状态

ConnectKeeper：
    1.维护所有连接
    2.对所有连接进行房间维护
    3.定期对Register发送本机的信息，包括：本机所有房间，房间人数，

```

###系统基本逻辑
```
1.Register只负责记录汇总信息，并且根据触发，把最新信息发送到Register的发布Key去
2.ConnectKeeper订阅Redis的key，获得消息之后发送key到Client端~
```
