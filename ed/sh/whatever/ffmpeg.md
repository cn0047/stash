FFmpeg
-

[docs ffmpeg](https://www.ffmpeg.org/ffmpeg.html)
[docs ffprobe](https://ffmpeg.org/ffprobe.html)
[source code](https://www.ffmpeg.org/download.html#repositories)
[examples ffmpeg](https://github.com/leandromoreira/digital_video_introduction/blob/master/encoding_pratical_examples.md#split-and-merge-smoothly)

````sh
ffmpeg \
-formats  # show available formats
-filters  #
-codecs   # show available codecs
-decoders # show available decoders
-encoders # show available encoders

-crf 23                                    # Constant Rate Factor,
                                           # For x264 and x265 - values between 0 and 51 (0 - best quality)
-crf 23                                    # Default for x264
-crf 28                                    # Default for x265
-crf 31                                    # Default for libvpx (range between 0 and 63)
-c:v libx264                               # codec name
-f segment|image2                          # force format
-g 30                                      # GOP
-i video                                   # input
-loglevel 16                               # fatal logging level
-map                                       # for manual control of stream selection
-preset ultrafast                          # set the encoding preset
-q:v 3                                     # use fixed quality scale
-qscale:v 3                                # use fixed quality scale
-r $fps                                    # set frame rate
-reset_timestamps 1                        # reset timestamps at the beginning of each segment (default false)
-s 1x1                                     # set frame size
-segment_format_options movflags=faststart # set list of options for the container format used for the segments
-segment_time 2                            # set segment duration
-ss 0                                      # set the start time offset
-threads 0                                 # use all threads
-update                                    # continuously overwrite one file (default false)
-vcodec codec                              # force video codec
-y                                         # overwrite output
-vf select='between(n\,14\,14)'            # set video filter, @see: http://ffmpeg.org/ffmpeg-filters.html#select_002c-aselect



# change audio volume (256=normal)
ffmpeg -i $fmov -vol 256 'res-'$fmov
ffmpeg -i $fmov -vcodec copy -af "volume=1dB" 'res-'$fmov
ffmpeg -i $fmov -vcodec copy -af "volume=-1dB" 'res-'$fmov

ffmpeg \
-y \                              # global options
-c:a libfdk_aac -c:v libx264 \    # input options
-i bunny_1080p_60fps.mp4 \        # input url
-c:v libvpx-vp9 -c:a libvorbis \  # output options
bunny_1080p_60fps_vp9.webm        # output url

# cut from timestamp 00:01:50 to duration 10.5
ffmpeg -ss 00:01:50 -i <input> -t 10.5 -c copy <output>

# add text on video
ffmpeg -i <input> -vf \
drawtext="text='Test Text':x=100:y=50: fontsize=24:fontcolor=yellow:box=1:boxcolor=red" \
<output>



docker run -ti --rm -v $HOME/Downloads:/d -w /d xubuntu /bin/bash

fmov=v.mov
fmp4=v.mp4
f2mp4=v2.mp4

# mov -> mp4
ffmpeg -i $fmov $fmp4
ffmpeg -i $fmov -q:v 0 $fmp4

# video duration
duration=`ffprobe -show_streams -print_format json $fmp4 2>/tmp/ffprobe.tmp | jq -r '.streams[0].duration'`

# create dir with png frames
mkdir png
ffmpeg -i $fmp4 png/frame_%5d.png

# create video from png files (not perfect)
ffmpeg -loop 1 -i png/frame_%5d.png -c:v libx264 -pix_fmt yuv420p -t $duration $f2mp4

# get 15nth frame from video
ffmpeg -i $f2mp4 -vf select='between(n\,14\,14)' -vsync 0 found_frame_%d.png
````

````sh

ffprobe \
-i                                # read specified file
-of csv=p=0                       # alias for -print_format
-select_streams V:0               # select the specified streams
-show_entries stream=width,height # show a set of specified entries
-show_format                      # show format/container info
-show_frames                      # show frames info
-v quiet                          # logging level

ffprobe v1.mp4

ffprobe -show_streams -print_format json $fmp4 2>/tmp/ffprobe.tmp | jq
````
