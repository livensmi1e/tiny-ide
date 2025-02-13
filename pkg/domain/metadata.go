package domain

import "github.com/livensmi1e/tiny-ide/pkg/constant"

type Metadata struct {
	Stdout string
	Stderr string
	Time   string
	Memory string
}

var DefaultMetadata = Metadata{
	Stdout: "",
	Stderr: "",
	Time:   constant.DefaultTime,
	Memory: constant.DefaultMemory,
}
