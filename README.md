# countOnce
just once? no, when appear many it run once, but it can run many times

### 灵感来源
  参考sync.once 只能发生一次,此包提供了并发时发生一次,多次调用仍有效
### code start
```
queueMutex = utils.NewQueueMutex()
queueMutex.Do(func() {
  do something...
}
```
  
