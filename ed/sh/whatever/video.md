Tools for video
-

#### ffmpeg

[source code](https://www.ffmpeg.org/download.html#repositories)

````sh
ffmpeg \
-formats  # show available formats
-codecs   # show available codecs
-decoders # show available decoders
-encoders # show available encoders

# change audio volume (256=normal)
ffmpeg -vol volume

ffmpeg -i f.avi f.mp4
````
