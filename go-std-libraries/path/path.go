package path

import (
	"path"
	"path/filepath"
	"fmt"
)

func DemoPathPackage() {
	const (
		url string = "https://maarumar.cloud/products/detail.go?id=12345"
		filePath string = "/usr/local/bin/go/src/path/path.go"
	)
	fmt.Println("(Path Package Demo)") // used for URL paths
	fmt.Println("Directory path:", path.Dir(url))
	fmt.Println("Base path:", path.Base(url))
	fmt.Println("Ext path:", path.Ext(url))

	fmt.Println("(File Path Package Demo)") // used for OS file paths
	fmt.Println("Directory path:", filepath.Dir(filePath))
	fmt.Println("Base path:", filepath.Base(filePath))
	fmt.Println("Ext path:", filepath.Ext(filePath))
	fmt.Println("Join path:", filepath.Join("usr", "local", "bin", "go"))
	fmt.Println("Is Abs:", filepath.IsAbs(filePath))
	fmt.Println("Is Local:", filepath.IsLocal(filePath))
}