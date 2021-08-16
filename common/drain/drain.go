package drain

import "io"

//go:generate go run v2ray.com/core/common/errors/errorgen

type Drainer interface {
	AcknowledgeReceive(size int)
	Drain(reader io.Reader) error
}
