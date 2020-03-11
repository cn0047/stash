Tools for video
-

#### ffmpeg

[docs](https://www.ffmpeg.org/ffmpeg.html)
[source code](https://www.ffmpeg.org/download.html#repositories)

````sh
ffmpeg \
-formats  # show available formats
-codecs   # show available codecs
-decoders # show available decoders
-encoders # show available encoders

ffmpeg \
-c:v libx264                               # codec name
-f segment                                 # force format
-g 30                                      # GOP
-i video                                   # input
-loglevel 16                               # fatal logging level
-map                                       # for manual control of stream selection
-preset                                    # set the encoding preset
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
-y                                         # overwrite output
-crf 23                                    # Constant Rate Factor,
                                           # For x264 and x265 - values between 0 and 51 (0 - best quality)
-crf 23                                    # Default for x264
-crf 28                                    # Default for x265
-crf 31                                    # Default for libvpx (range between 0 and 63)

# change audio volume (256=normal)
ffmpeg -vol volume

ffmpeg -i f.avi f.mp4

ffmpeg -i v1.mov -q:v 0 v1.mp4 # mov -> mp4

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
````

#### ffprobe

[docs](https://ffmpeg.org/ffprobe.html)

````sh
ffprobe \
-i                                # read specified file
-of csv=p=0                       # alias for -print_format
-select_streams V:0               # select the specified streams
-show_entries stream=width,height # show a set of specified entries
-show_format                      # show format/container info
-show_frames                      # show frames info
-v quiet                          # logging level

ffprobe -show_streams -print_format json v1.mp4 2>/tmp/ffprobe.tmp

ffprobe v1.mp4
````
