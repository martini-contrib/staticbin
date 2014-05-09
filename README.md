# StaticBin [![wercker status](https://app.wercker.com/status/f4abe5d9213d02e00fc76f14c493048c/s/ "wercker status")](https://app.wercker.com/project/bykey/f4abe5d9213d02e00fc76f14c493048c) [![GoDoc](https://godoc.org/github.com/yosssi/staticbin?status.png)](https://godoc.org/github.com/yosssi/staticbin)

Martini middleware/handler for serving static files from binary data

## Usage

```go
package main

import "github.com/yosssi/staticbin"

func main() {
	// You have to pass the "Asset" function that go-bindata
	// (https://github.com/jteeuwen/go-bindata) generates.
	// staticbin.Classic(Asset) instance automatically serves static files
	// from the "public" directory in the root of your server.
	// You can serve from more directories by adding more staticbin.Static handlers.
	//   m.Use(staticbin.Static("assets", Asset)) // serve from the "assets" directory as well
	m := staticbin.Classic(Asset)

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Run()
}
```

## Doc

* [GoDoc](https://godoc.org/github.com/yosssi/staticbin)
