# DiskMap: persistent key value storage in Go

DiskMap aims to be the most simple, not the most perforamt

It doesn't have
- cache
- index
- transactions

And it's not the fastest one, but
- it's simple <- the most important
- it's concurent safe
- It's very easy to use

## Usage

```go
import "github.com/yourusername/diskmap"

func main() {
    db := diskmap.NewDiskMap("/path/to/storage")

    // Set a key value pair
    err := db.Set("key", []byte("value"))
    if err != nil {
        // Handle error
    }

    // Get the value
    value, err := db.Get("key")
    if err != nil {
        // Handle error
    }
}
```

## Contributions

DiskMap is open for contributions. If you have any ideas how to make it more perforant keeping still simple, you are very welcome
