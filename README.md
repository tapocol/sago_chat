# Sago Chat

An extension to support chat in sago. Adds an action called "chat" that will
distribute a "chat" action back to the client for all connected sessions.

## Usage

Install and import

```sh
$ go get github.com/craigjackson/sago_chat
```

```go
import "github.com/craigjackson/sago_chat"
```

Before Run(), Init sago_chat:

```go
func main() {
  // ...
  sago_chat.Init()
  // ...
  sago.Run()
}
```

## License

The MIT License - Copyright (c) 2013 Craig Jackson

