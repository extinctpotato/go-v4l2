// +build linux

package v4l2

const (
	IOC_NRBITS   = 8
	IOC_TYPEBITS = 8

	// Check arch!
	IOC_SIZEBITS = 14
	IOC_DIRBITS  = 2

	// Check arch!
	IOC_NONE  = 0
	IOC_WRITE = 1
	IOC_READ  = 2

	IOC_NRMASK   = ((1 << IOC_NRBITS) - 1)
	IOC_TYPEMASK = ((1 << IOC_TYPEBITS) - 1)
	IOC_SIZEMASK = ((1 << IOC_SIZEBITS) - 1)
	IOC_DIRMASK  = (1 << IOC_DIRBITS)

	IOC_NRSHIFT   = 0
	IOC_TYPESHIFT = (IOC_NRSHIFT + IOC_NRBITS)
	IOC_SIZESHIFT = (IOC_TYPESHIFT + IOC_TYPEBITS)
	IOC_DIRSHIFT  = (IOC_SIZESHIFT + IOC_SIZEBITS)
)

const (
	V4L2_BUF_TYPE_VIDEO_CAPTURE = 1

	V4L2_FIELD_ANY  = 0
	V4L2_FIELD_NONE = 1

	V4L2_PIX_FMT_JPEG = 'J' | 'P'<<8 | 'E'<<16 | 'G'<<24
	V4L2_PIX_FMT_H264 = 'H' | '2'<<8 | '6'<<16 | '4'<<24

	V4L2_MEMORY_MMAP = 1

	VIDIOC_DQBUF       = 0xc0445611
	VIDIOC_QUERYCAP    = 0x80685600
	VIDIOC_REQBUFS     = 0xc0145608
	VIDIOC_G_EXT_CTRLS = 0xc0185647
	VIDIOC_STREAMON    = 0x40045612
	VIDIOC_STREAMOFF   = 0x40045613
	VIDIOC_S_CTRL      = 0xc008561c
)

// Controls (from linux/v4l2-controls.h)
const (
	// Control classes
	V4L2_CTRL_CLASS_USER = 0x00980000
	V4L2_CTRL_CLASS_MPEG = 0x00990000

	// User-class control IDs
	V4L2_CID_BASE               = V4L2_CTRL_CLASS_USER | 0x900
	V4L2_CID_USER_BASE          = V4L2_CID_BASE
	V4L2_CID_USER_CLASS         = V4L2_CTRL_CLASS_USER | 1
	V4L2_CID_BRIGHTNESS         = V4L2_CID_BASE + 0
	V4L2_CID_CONTRAST           = V4L2_CID_BASE + 1
	V4L2_CID_SATURATION         = V4L2_CID_BASE + 2
	V4L2_CID_HUE                = V4L2_CID_BASE + 3
	V4L2_CID_AUTO_WHITE_BALANCE = V4L2_CID_BASE + 12
	V4L2_CID_DO_WHITE_BALANCE   = V4L2_CID_BASE + 13
	V4L2_CID_RED_BALANCE        = V4L2_CID_BASE + 14
	V4L2_CID_BLUE_BALANCE       = V4L2_CID_BASE + 15
	V4L2_CID_GAMMA              = V4L2_CID_BASE + 16
	V4L2_CID_EXPOSURE           = V4L2_CID_BASE + 17
	V4L2_CID_AUTOGAIN           = V4L2_CID_BASE + 18
	V4L2_CID_GAIN               = V4L2_CID_BASE + 19
	V4L2_CID_HFLIP              = V4L2_CID_BASE + 20
	V4L2_CID_VFLIP              = V4L2_CID_BASE + 21

	// MPEG-class control IDs
	V4L2_CID_MPEG_BASE                    = V4L2_CTRL_CLASS_MPEG | 0x900
	V4L2_CID_MPEG_CLASS                   = V4L2_CTRL_CLASS_MPEG | 1
	V4L2_CID_MPEG_VIDEO_B_FRAMES          = V4L2_CID_MPEG_BASE + 202
	V4L2_CID_MPEG_VIDEO_GOP_SIZE          = V4L2_CID_MPEG_BASE + 203
	V4L2_CID_MPEG_VIDEO_BITRATE           = V4L2_CID_MPEG_BASE + 207
	V4L2_CID_MPEG_VIDEO_REPEAT_SEQ_HEADER = V4L2_CID_MPEG_BASE + 226
	V4L2_CID_MPEG_VIDEO_H264_I_PERIOD     = V4L2_CID_MPEG_BASE + 358
	V4L2_CID_MPEG_VIDEO_H264_LEVEL        = V4L2_CID_MPEG_BASE + 359
	V4L2_CID_MPEG_VIDEO_H264_PROFILE      = V4L2_CID_MPEG_BASE + 363

	// H.264 Levels
	V4L2_MPEG_VIDEO_H264_LEVEL_1_0 = 0
	V4L2_MPEG_VIDEO_H264_LEVEL_1B  = 1
	V4L2_MPEG_VIDEO_H264_LEVEL_1_1 = 2
	V4L2_MPEG_VIDEO_H264_LEVEL_1_2 = 3
	V4L2_MPEG_VIDEO_H264_LEVEL_1_3 = 4
	V4L2_MPEG_VIDEO_H264_LEVEL_2_0 = 5
	V4L2_MPEG_VIDEO_H264_LEVEL_2_1 = 6
	V4L2_MPEG_VIDEO_H264_LEVEL_2_2 = 7
	V4L2_MPEG_VIDEO_H264_LEVEL_3_0 = 8
	V4L2_MPEG_VIDEO_H264_LEVEL_3_1 = 9
	V4L2_MPEG_VIDEO_H264_LEVEL_3_2 = 10
	V4L2_MPEG_VIDEO_H264_LEVEL_4_0 = 11
	V4L2_MPEG_VIDEO_H264_LEVEL_4_1 = 12
	V4L2_MPEG_VIDEO_H264_LEVEL_4_2 = 13
	V4L2_MPEG_VIDEO_H264_LEVEL_5_0 = 14
	V4L2_MPEG_VIDEO_H264_LEVEL_5_1 = 15

	// H.264 Profiles
	V4L2_MPEG_VIDEO_H264_PROFILE_BASELINE             = 0
	V4L2_MPEG_VIDEO_H264_PROFILE_CONSTRAINED_BASELINE = 1
	V4L2_MPEG_VIDEO_H264_PROFILE_MAIN                 = 2
	V4L2_MPEG_VIDEO_H264_PROFILE_EXTENDED             = 3
	V4L2_MPEG_VIDEO_H264_PROFILE_HIGH                 = 4
	V4L2_MPEG_VIDEO_H264_PROFILE_HIGH_10              = 5
	V4L2_MPEG_VIDEO_H264_PROFILE_HIGH_422             = 6
	V4L2_MPEG_VIDEO_H264_PROFILE_HIGH_444_PREDICTIVE  = 7
	V4L2_MPEG_VIDEO_H264_PROFILE_HIGH_10_INTRA        = 8
	V4L2_MPEG_VIDEO_H264_PROFILE_HIGH_422_INTRA       = 9
	V4L2_MPEG_VIDEO_H264_PROFILE_HIGH_444_INTRA       = 10
	V4L2_MPEG_VIDEO_H264_PROFILE_CAVLC_444_INTRA      = 11
	V4L2_MPEG_VIDEO_H264_PROFILE_SCALABLE_BASELINE    = 12
	V4L2_MPEG_VIDEO_H264_PROFILE_SCALABLE_HIGH        = 13
	V4L2_MPEG_VIDEO_H264_PROFILE_SCALABLE_HIGH_INTRA  = 14
)
