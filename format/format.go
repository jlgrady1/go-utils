package format

import "fmt"

// Sizes for formatting
const (
	BYTE     = 1.0
	KILOBYTE = 1024 * BYTE
	MEGABYTE = 1024 * KILOBYTE
	GIGABYTE = 1024 * MEGABYTE
	TERABYTE = 1024 * GIGABYTE
	PETABYTE = 1024 * TERABYTE
)

// ReadableByteSize takes the number of bytes as an unsigned int and returns
// a string containing the human readable format.
func ReadableByteSize(bytes uint64) string {
	var extension string
	value := float64(bytes)
	switch {
	case bytes >= PETABYTE:
		extension = "PB"
		value = value / PETABYTE
	case bytes >= TERABYTE:
		extension = "TB"
		value = value / TERABYTE
	case bytes >= GIGABYTE:
		extension = "GB"
		value = value / GIGABYTE
	case bytes >= MEGABYTE:
		extension = "MB"
		value = value / MEGABYTE
	case bytes >= KILOBYTE:
		extension = "KB"
		value = value / KILOBYTE
	default:
		extension = "B"
	}

	return fmt.Sprintf("%.1f %s", value, extension)
}
