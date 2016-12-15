package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/frame.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

type (
	Frame C.struct_AVFrame
	FrameSideData C.struct_AVFrameSideData
	FrameSideDataType C.enum_AVFrameSideDataType
	Buffer C.struct_AVBuffer
	BufferRef C.struct_AVBufferRef
	BufferPool C.struct_AVBufferPool
)

func FrameAlloc() *Frame {
	return (*Frame)(unsafe.Pointer(C.av_frame_alloc()))
}

func (f *Frame) Free() {
	C.av_frame_free((**C.struct_AVFrame)(unsafe.Pointer(&f)))
}

func (f *Frame) GetBuffer(a int) int {
	return int(C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(a)))
}

func (f *Frame) Ref(d *Frame) int {
	return int(C.av_frame_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(f))))
}

func (f *Frame) Clone() *Frame {
	return (*Frame)(C.av_frame_clone((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

func (f *Frame) Unref() {
	cf := (*C.struct_AVFrame)(unsafe.Pointer(f))
	C.av_frame_unref(cf)
}

func (f *Frame)MoveRef(d *Frame) {
	C.av_frame_move_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(f)))
}

func (f *Frame) IsWritable() int {
	return int(C.av_frame_is_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

func (f *Frame) MakeWritable() int {
	return int(C.av_frame_make_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

func (f *Frame) CopyProps(d *Frame) int {
	return int(C.av_frame_copy_props((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(f))))
}

func (f *Frame) GetPlaneBuffer(p int) *BufferRef {
	return (*BufferRef)(C.av_frame_get_plane_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(p)))
}

func (f *Frame) NewSideData(d FrameSideDataType, s int) *FrameSideData {
	return (*FrameSideData)(C.av_frame_new_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(d), C.int(s)))
}

func (f *Frame) GetSideData(t FrameSideDataType) *FrameSideData {
	return (*FrameSideData)(C.av_frame_get_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(t)))
}

func (f *Frame) Data() *uint8 {
	return (*uint8)(unsafe.Pointer((*C.uint8_t)(unsafe.Pointer(&f.data))))
}

func (f *Frame) Linesize() int {
	return int(*(*C.int)(unsafe.Pointer(&f.linesize)))
}
