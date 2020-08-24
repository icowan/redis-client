# redis client

## 使用

```bash
go get github.com/icowan/redis-client
```

## 单点Redis

```golang
rds := NewRedisClient("127.0.0.1:6379", "admin", "", 1)
defer func() {
    _ = rds.Close()
}()

_ = rds.Set("hello", "world", time.Second*10)
v, err := rds.Get("hello")
if err != nil {
    log.Fatal()
}

log.Print(v)
```

## 集群Redis

```golang
rds := NewRedisClient("127.0.0.1:6379,127.0.0.1:7379,127.0.0.1:8379", "admin", "", 1)
defer func() {
    _ = rds.Close()
}()

_ = rds.Set("hello", "world", time.Second*10)
v, err := rds.Get("hello")
if err != nil {
    log.Fatal()
}

log.Print(v)
```
