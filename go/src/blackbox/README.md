# Thrift Server/Client

running a simple thrift server:

```
 $ go run cmd/main.go --server
2021/09/08 14:41:45 main.go:124: hello
 => Running server on:  localhost:9003
2021/09/08 14:41:47 main.go:30: received ping()
2021/09/08 14:41:47 main.go:35: received GetVersion()
2021/09/08 14:41:47 main.go:40: received GetName()
2021/09/08 14:41:47 main.go:45: received LogLocation -  {
  "timestamp_unix_sec": 1631137307,
  "latitude_degrees": 33,
  "longitude_degrees": -122
 }
```

running a simple thrift client against the above:

```
 $ go run cmd/main.go
2021/09/08 14:41:47 main.go:124: hello
2021/09/08 14:41:47 main.go:97: ping() - start
2021/09/08 14:41:47 main.go:99: ping() - done
2021/09/08 14:41:47 main.go:104: get_version() - start
2021/09/08 14:41:47 main.go:107: get_version() -  000.001.001 blackbox
2021/09/08 14:41:47 main.go:115: LogLocation()
2021/09/08 14:41:47 main.go:157: => done:
```
