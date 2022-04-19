package v4l2

import "fmt"

type EmptyBufferError struct {
	Length uint32
	Offset uint32
}

type MmapError struct {
	Err error
	Fd  int
}

func (m *MmapError) Error() string {
	return fmt.Sprintf(
		"mmap on fd %d returned %s",
		m.Fd,
		m.Err.Error(),
	)
}

func (e *EmptyBufferError) Error() string {
	return fmt.Sprintf(
		"VIDIOC_QUERYBUF returned an empty buffer (length: %d, offset: %d)",
		e.Length,
		e.Offset,
	)
}
