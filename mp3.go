package mp3

//#cgo pkg-config: libavcodec libavformat
/*
#include <libavcodec/avcodec.h>
#include <libavformat/avformat.h>
#include <stdio.h>
#include <stdlib.h>

int mp3_check(const char * url) {
	av_register_all();
	avcodec_register_all();

	AVFormatContext *ctx = NULL;
	int err = avformat_open_input(&ctx,url,NULL,NULL);
	if (err < 0) {
		return -1;
	}

	err = avformat_find_stream_info(ctx,NULL);
    if (err < 0) {
    	return -1;
    }

    AVCodec * codec = NULL;
	int strm = av_find_best_stream(ctx, AVMEDIA_TYPE_AUDIO, -1, -1, &codec, 0);
	if (strm < 0) {
		return -1;
	}

	AVCodecContext * codecCtx = ctx->streams[strm]->codec;

	err = avcodec_open2(codecCtx, codec, NULL);
	if (err < 0) {
		return -1;
	}

	if (strcmp(codec->name , "mp3")==0) {
		//Either image data or we are some kind of multimedia codec with mp3 audio
		if (ctx->nb_streams > 1) {
			int strm = av_find_best_stream(ctx, AVMEDIA_TYPE_VIDEO, -1, -1, &codec, 0);
			if (strm < 0) {
				return -1;
			}

			AVCodecContext * codecCtx = ctx->streams[strm]->codec;

			err = avcodec_open2(codecCtx, codec, NULL);
			if (err < 0) {
				return -1;
			}

			//Lets assume this is our picture!
			//TODO(sjon): Decode this picture.
			if (strcmp(codec->name , "mjpeg")==0 || strcmp(codec->name , "png")==0) {
				return 0;
			}

			return -1;
		}
		return 0;
	}

	return -1;
}
*/
import "C"
import "unsafe"

func IsMP3(filename string) bool {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	if r := C.mp3_check(cs); r >= 0 {
		return true
	}
	return false
}
