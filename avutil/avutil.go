package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"

type (
	Dictionary C.struct_AVDictionary
	MediaType C.enum_AVMediaType
)

const (
	MediaTypeUnknown MediaType = C.AVMEDIA_TYPE_UNKNOWN
	MediaTypeVideo MediaType = C.AVMEDIA_TYPE_VIDEO
	MediaTypeAudio MediaType = C.AVMEDIA_TYPE_AUDIO
	MediaTypeData MediaType = C.AVMEDIA_TYPE_DATA
	MediaTypeSubtitle MediaType = C.AVMEDIA_TYPE_SUBTITLE
	MediaTypeAttachment MediaType = C.AVMEDIA_TYPE_ATTACHMENT
)
