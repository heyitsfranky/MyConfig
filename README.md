# MyConfig

**MyConfig** provides functionality for initializing and reading configuration settings from YAML files. It is designed to allow the initialization of nil pointers to structs efficiently and to also guarantee all parameters of a given struct exist in the referenced configuration file.

## Installation

To use **MyConfig** in your Go project, you can simply run:

```bash
go get github.com/heyitsfranky/MyConfig@latest
```

## Usage

Here's a basic example of how to use **MyConfig**:
```go
import (
    "fmt"
    "os"

    "github.com/heyitsfranky/myConfig"
)

// Define your configuration struct with YAML tags.
type MyConfigStruct struct {
    Key1 string `yaml:"key1"`
    Key2 string `yaml:"key2"`
}

func main() {
    var config *MyConfigStruct
    err := myConfig.Init("config.yaml", &config)
    if err != nil {
        fmt.Println("Error initializing config:", err)
        os.Exit(1)
    }

    // Your configuration is now populated in the 'config' variable.
    fmt.Println("Key1:", config.Key1)
    fmt.Println("Key2:", config.Key2)
}
```
## License

This package is distributed under the MIT License.
Feel free to contribute or report issues on GitHub.

Happy coding!