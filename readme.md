# os signal wrapper

```go
import (
  "gopkg.in/karlseguin/signals.v1"
)

func main() {
  signals.Handle(signals.SIGUSR2, reload)
  RunServer()
}

func reload() bool {
  //do reload
  return true
}
```

Returning true causes the code to continue listening for further signals.
