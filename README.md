# autorun

A Go package for managing application auto-start configuration across different operating systems.

## Features

- Query current auto-run status for your application
- Enable/disable auto-run on system startup
- Cross-platform support (Windows/macOS/Linux)

## Installation

```bash
go get github.com/baohuiming/autorun
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/baohuiming/autorun"
)

func main() {
	config := &autorun.AutoRunConfig{
		AppName:        "MyApp",
		ExecutablePath: "/usr/local/bin/myapp",
		CompanyName:    "com.example",
	}

	// Query auto-run status
	isEnabled, err := autorun.QueryAutoRun(config)
	if err != nil {
		fmt.Printf("Error querying status: %v\n", err)
	}
	fmt.Printf("Auto-run enabled: %v\n", isEnabled)

	// Enable auto-run
	if err := autorun.EnableAutoRun(config); err != nil {
		fmt.Printf("Error enabling auto-run: %v\n", err)
	}

	// Disable auto-run
	if err := autorun.DisableAutoRun(config); err != nil {
		fmt.Printf("Error disabling auto-run: %v\n", err)
	}
}
```

## License
This project is licensed under the MIT License.