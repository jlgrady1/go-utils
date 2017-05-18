package format

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type FormatSuite struct{}

var _ = Suite(&FormatSuite{})

func (s *FormatSuite) TestReadableByteSizeB(c *C) {
	bytes := uint64(50)
	c.Assert(ReadableByteSize(bytes), Equals, "50.0 B")
}

func (s *FormatSuite) TestReadableByteSizeKB(c *C) {
	bytes := uint64(1536)
	c.Assert(ReadableByteSize(bytes), Equals, "1.5 KB")
}

func (s *FormatSuite) TestReadableByteSizeMB(c *C) {
	bytes := uint64(21286093)
	c.Assert(ReadableByteSize(bytes), Equals, "20.3 MB")
}

func (s *FormatSuite) TestReadableByteSizeGB(c *C) {
	bytes := uint64(4294967296)
	c.Assert(ReadableByteSize(bytes), Equals, "4.0 GB")
}

func (s *FormatSuite) TestReadableByteSizeTB(c *C) {
	bytes := uint64(7916483719988)
	c.Assert(ReadableByteSize(bytes), Equals, "7.2 TB")
}

func (s *FormatSuite) TestReadableByteSizePB(c *C) {
	bytes := uint64(1238489897526887)
	c.Assert(ReadableByteSize(bytes), Equals, "1.1 PB")
}
