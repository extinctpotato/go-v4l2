// +build linux

package v4l2

func ioctl_ioc(dir int, typ int, nr int, size int) int {
	return (dir << IOC_DIRSHIFT) |
	(typ << IOC_TYPESHIFT) |
	(nr << IOC_NRSHIFT) |
	(size << IOC_SIZESHIFT)
}

func ioctl_io(typ int, nr int) int {
	return ioctl_ioc(IOC_NONE, typ, nr, 0)
}

func ioctl_ior(typ int, nr int, size int) int {
	return ioctl_ioc(IOC_READ, typ, nr, size)
}

func ioctl_iow(typ int, nr int, size int) int {
	return ioctl_ioc(IOC_WRITE, typ, nr, size)
}

func ioctl_iowr(typ int, nr int, size int) int {
	return ioctl_ioc(IOC_WRITE|IOC_READ, typ, nr, size)
}
