package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/samplefmt.h>
import "C"
import (
	"unsafe"
)

const (
	AV_SAMPLE_FMT_NONE = int(C.AV_SAMPLE_FMT_NONE)
	AV_SAMPLE_FMT_U8   = int(C.AV_SAMPLE_FMT_U8)
	AV_SAMPLE_FMT_S16  = int(C.AV_SAMPLE_FMT_S16)
	AV_SAMPLE_FMT_S32  = int(C.AV_SAMPLE_FMT_S32)
	AV_SAMPLE_FMT_FLT  = int(C.AV_SAMPLE_FMT_FLT)
	AV_SAMPLE_FMT_DBL  = int(C.AV_SAMPLE_FMT_DBL)

	AV_SAMPLE_FMT_U8P  = int(C.AV_SAMPLE_FMT_U8P)
	AV_SAMPLE_FMT_S16P = int(C.AV_SAMPLE_FMT_S16P)
	AV_SAMPLE_FMT_S32P = int(C.AV_SAMPLE_FMT_S32P)
	AV_SAMPLE_FMT_FLTP = int(C.AV_SAMPLE_FMT_FLTP)
	AV_SAMPLE_FMT_DBLP = int(C.AV_SAMPLE_FMT_DBLP)
	AV_SAMPLE_FMT_S64  = int(C.AV_SAMPLE_FMT_S64)
	AV_SAMPLE_FMT_S64P = int(C.AV_SAMPLE_FMT_S64P)

	AV_SAMPLE_FMT_NB = int(C.AV_SAMPLE_FMT_NB)
)

func AvGetSampleFmtName(sampleFmt int) string {
	return C.GoString(C.av_get_sample_fmt_name((C.enum_AVSampleFormat)(sampleFmt)))
}

func AvSamplesAllocArrayAndSamples(nbChannels, nbSamples int, sampleFmt int, align int) (ret int, audioData **uint8, linesize int) {
	var linesize32 int32
	ret = (int)(C.av_samples_alloc_array_and_samples(
		(***C.uint8_t)(unsafe.Pointer(&audioData)),
		(*C.int)(&linesize32),
		C.int(nbChannels), C.int(nbSamples),
		C.enum_AVSampleFormat(sampleFmt), C.int(align)))
	return ret, audioData, int(linesize32)
}

func AvSamplesFreeArrayAndSamples(audioData **uint8) {
	C.av_freep(unsafe.Pointer(audioData))
	C.av_free(unsafe.Pointer(audioData))
}
