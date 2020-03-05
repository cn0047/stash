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

# change audio volume (256=normal)
ffmpeg -vol volume

ffmpeg -i f.avi f.mp4

ffmpeg -i v1.mov -q:v 0 v1.mp4 # mov -> mp4
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
