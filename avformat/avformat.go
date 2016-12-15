package avformat

//#include <libavformat/avformat.h>
//#cgo pkg-config: libavformat
//static const AVStream *go_av_streams_get(const AVStream **streams, unsigned int n)
//{
//  return streams[n];
//}
import "C"

import (
	"golibav/avcodec"
	"golibav/avutil"
	"unsafe"
)

type (
	FormatContext C.struct_AVFormatContext
	InputFormat C.struct_AVInputFormat
	Stream C.struct_AVStream
)

func RegisterAll() {
	C.av_register_all()
}

func AllocContext() *FormatContext {
	return (*FormatContext)(C.avformat_alloc_context())
}

func (ctx *FormatContext) FreeContext() {
	defer C.avformat_free_context((*C.struct_AVFormatContext)(unsafe.Pointer(ctx)))
}

func (ctx *FormatContext) OpenInput(filename string, input *InputFormat, dict **avutil.Dictionary) int {
	return int(C.avformat_open_input((**C.struct_AVFormatContext)(unsafe.Pointer(&ctx)), C.CString(filename), (*C.struct_AVInputFormat)(input), (**C.struct_AVDictionary)(unsafe.Pointer(dict))))
}

func (ctx *FormatContext) CloseInput() {
	C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(&ctx)))
}

func (ctx *FormatContext) FormatFindStreamInfo(d **avutil.Dictionary) int {
	return int(C.avformat_find_stream_info((*C.struct_AVFormatContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

func (ctx *FormatContext) NumberOfStreams() uint {
	return uint(ctx.nb_streams)
}

func (ctx *FormatContext) Streams() []*Stream {
	count := ctx.NumberOfStreams()
	if count <= 0 {
		return nil
	}
	streams := make([]*Stream, 0, count)
	for i := uint(0); i < count; i++ {
		cStream := C.go_av_streams_get(ctx.streams, C.uint(i))
		stream := (*Stream)(unsafe.Pointer(cStream))
		streams = append(streams, stream)
	}
	return streams
}

func (s *Stream) Codec() *avcodec.CodecContext {
	return (*avcodec.CodecContext)(unsafe.Pointer(s.codec))
}

func (s *Stream) Metadata() *avutil.Dictionary {
	return (*avutil.Dictionary)(unsafe.Pointer(s.metadata))
}

func (ctx *FormatContext) ReadFrame(pkt *avcodec.Packet) int {
	return int(C.av_read_frame((*C.struct_AVFormatContext)(unsafe.Pointer(ctx)), (*C.struct_AVPacket)(unsafe.Pointer(pkt))))
}
