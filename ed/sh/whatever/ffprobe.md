FFprobe
-

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
