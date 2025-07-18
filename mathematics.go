package astiav

//#include <libavutil/mathematics.h>
import "C"

// https://ffmpeg.org/doxygen/7.0/group__lavu__math.html#ga3daf97178dd1b08b5e916be381cd33e4
func Rescale(a, b, c int64) int64 {
	return int64(C.av_rescale(C.int64_t(a), C.int64_t(b), C.int64_t(c)))
}

// https://ffmpeg.org/doxygen/7.0/group__lavu__math.html#gaf02994a8bbeaa91d4757df179cbe567f
func RescaleQ(a int64, b Rational, c Rational) int64 {
	return int64(C.av_rescale_q(C.int64_t(a), b.c, c.c))
}

// https://ffmpeg.org/doxygen/7.0/group__lavu__math.html#gab706bfec9bf56534e02ca9564cb968f6
func RescaleQRnd(a int64, b Rational, c Rational, r Rounding) int64 {
	return int64(C.av_rescale_q_rnd(C.int64_t(a), b.c, c.c, C.enum_AVRounding(r)))
}
