package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

func PacketAlloc() *Packet {
	return (*Packet)(C.av_packet_alloc())
}

func (p *Packet) InitPacket() {
	C.av_init_packet((*C.struct_AVPacket)(p))
}

func (p *Packet) NewPacket(s int) int {
	return int(C.av_new_packet((*C.struct_AVPacket)(p), C.int(s)))
}

func (p *Packet) ShrinkPacket(s int) {
	C.av_shrink_packet((*C.struct_AVPacket)(p), C.int(s))
}

func (p *Packet) GrowPacket(s int) int {
	return int(C.av_grow_packet((*C.struct_AVPacket)(p), C.int(s)))
}

func (p *Packet) PacketFromData(d *uint8, s int) int {
	return int(C.av_packet_from_data((*C.struct_AVPacket)(p), (*C.uint8_t)(d), C.int(s)))
}

func (p *Packet) CopyPacket(r *Packet) int {
	return int(C.av_copy_packet((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(r)))
}

func (p *Packet) CopyPacketSideData(r *Packet) int {
	return int(C.av_copy_packet_side_data((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(r)))
}

func (p *Packet) Free() {
	C.av_packet_free((**C.struct_AVPacket)(unsafe.Pointer(p)))
}

func (p *Packet) StreamIndex() int {
	return int(p.stream_index)
}
