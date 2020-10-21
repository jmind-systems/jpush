# JPush

## Overview

## Usage

```go
package main

import (
    "log"
    "context"

    "github.com/jmind-systems/jpush"
    "github.com/jmind-systems/jpush/payload"
)

func main() {
    client, err := jpush.NewClient("key", "secret")
    if err != nil {
        log.Fatalf("err: %v", err)
    }

    notification := payload.NewNotification().
        Android()


    var req jpush.Request
    req.Notification = notification

    err = client.Push(context.Background(), nil)
    if err != nil {
        log.Fatalf("err: %v", err)
    }
}
```

## License

Project released under the terms of the MIT [license](./LICENSE).
