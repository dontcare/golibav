package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"golibav/avutil"
	"unsafe"
)

type (
	Codec C.struct_AVCodec
	CodecContext C.struct_AVCodecContext
	CodecID C.enum_AVCodecID
	Packet C.struct_AVPacket
	ThreadType int
)

const (
	ThreadTypeFrame ThreadType = C.FF_THREAD_FRAME
	ThreadTypeSlice ThreadType = C.FF_THREAD_SLICE
)

func RegisterAll() {
	C.avcodec_register_all()
}

func (c *Codec) AllocContext3() *CodecContext {
	return (*CodecContext)(C.avcodec_alloc_context3((*C.struct_AVCodec)(c)))
}

func (ctx *CodecContext) FreeContext() {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(ctx)))
}

func (ctx *CodecContext) Close() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(ctx)))
}

func (ctx *CodecContext) CopyContext(ctx2 *CodecContext) int {
	return int(C.avcodec_copy_context((*C.struct_AVCodecContext)(ctx2), (*C.struct_AVCodecContext)(ctx)))
}

func (ctx *CodecContext) ActiveThreadType() int {
	return int(ctx.active_thread_type)
}

func (ctx *CodecContext) SetActiveThreadType(threadType ThreadType) {
	ctx.active_thread_type = (C.int)(threadType)
}

func (ctx *CodecContext) ThreadSafeCallbacks() int {
	return int(ctx.thread_safe_callbacks)
}

func (ctx *CodecContext) SetThreadSafeCallbacks(count int) {
	ctx.thread_safe_callbacks = (C.int)(count)
}

func (ctx *CodecContext) ThreadType() int {
	return int(ctx.thread_type)
}

func (ctx *CodecContext) ThreadCount() int {
	return int(ctx.thread_count)
}

func (ctx *CodecContext) SetThreadCount(count int) {
	ctx.thread_count = (C.int)(count)
}

func (ctx *CodecContext) CodecID() CodecID {
	return (CodecID)(ctx.codec_id)
}

func (ctx *CodecContext) CodecType() avutil.MediaType {
	return (avutil.MediaType)(ctx.codec_type)
}

func FindDecoder(id CodecID) *Codec {
	return (*Codec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

func FindDecoderByName(name string) *Codec {
	return (*Codec)(C.avcodec_find_decoder_by_name(C.CString(name)))
}

func FindEncoder(id CodecID) *Codec {
	return (*Codec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

func FindEncoderByName(name string) *Codec {
	return (*Codec)(C.avcodec_find_encoder_by_name(C.CString(name)))
}

func (ctx *CodecContext) Open2(codec *Codec, dict **avutil.Dictionary) int {
	return int(C.avcodec_open2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVCodec)(codec), (**C.struct_AVDictionary)(unsafe.Pointer(dict))))
}

func (ctx *CodecContext) DecodeVideo2(p *avutil.Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_video2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVFrame)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

func (ctx *CodecContext) Height() int {
	return int(ctx.height)
}

func (ctx *CodecContext) Width() int {
	return int(ctx.width)
}
