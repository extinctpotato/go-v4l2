// +build linux

package v4l2

func ioctl_ioc(dir uint, typ uintptr, nr uint, size uintptr) uint {
	return (dir << IOC_DIRSHIFT) |
		(uint(typ) << IOC_TYPESHIFT) |
		(nr << IOC_NRSHIFT) |
		(uint(size) << IOC_SIZESHIFT)
}

func ioctl_io(typ uintptr, nr uint, size uintptr) uint {
	return ioctl_ioc(IOC_NONE, typ, nr, 0)
}

func ioctl_ior(typ uintptr, nr uint, size uintptr) uint {
	return ioctl_ioc(IOC_READ, typ, nr, size)
}

func ioctl_iow(typ uintptr, nr uint, size uintptr) uint {
	return ioctl_ioc(IOC_WRITE, typ, nr, size)
}

func ioctl_iowr(typ uintptr, nr uint, size uintptr) uint {
	return ioctl_ioc(IOC_READ|IOC_WRITE, typ, nr, size)
}
