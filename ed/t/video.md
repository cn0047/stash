Video
-

**Video** - a series of pictures / frames running at a given rate.

**Bit rate** - amount of data encoded per second in a video.
DVD - Bit rate = 4-8Mbps.
Bluray - Bit rate = 24-40Mbps.

**Resolution** - pixel size of video (1920x1080 - HD).

**Video file** is container for: video stream, audio & metadata.
Each stream is encoded by a different kind of codec.
Packets are pieces of data that contain bits of data that are decoded into raw frames.

**Codec** (Coder/Decoder) specified in metadata, like:
~~divx/xvid, ffmpeg, x264~~,
h.264, h.265(HEVC), mp3, aac, dolby, vp9.

Video format: standardized set of rules to store containers,
codecs, metadata, like: mp4, hls.

Best video for youtube: container:
  container: mp4,
  codec: h.264,
  audio: AAC-LC,
  channels: stereo/ stereo 5.1,
  sample rate: 96khz,
  aspect ratio: 19.6,
  frame rate: 60fps,

**Remuxing** - converting from one container to another one (avi -> mp4).

**Frame rate** - count of images in 1 second (FPS).
24/30/60/120 - the bigger count - the higher quality and smoothly video.
For 60 frame rate it's possible to slow down video speed in video editor without losing quality.

**GOP** - Group of pictures.

GOP - is a collection of successive pictures within a coded video stream.

GOP can contain next picture types:
* `I picture or I frame` (intra coded picture) – picture that coded independently of all other pictures.
* `P picture or P frame` (predictive coded picture) - contains motion-compensated difference information
relative to previously decoded pictures.
* `B picture or B frame` (bipredictive coded picture).
* `D picture or D frame` (DC direct coded picture).

`M=3, N=12` -> `IBBPBBPBBPBBI`:
3 tells the distance between two anchor frames (I or P).
12 one tells the distance between two full images (I-frames).

**Chrominance** is the signal used in video systems to convey the color information of the picture.
Chrominance is usually represented as two color-difference components:
U = B − Y (blue − luma) and V = R − Y (red − luma).

**Luma** represents the brightness in an image
(the "black-and-white" or achromatic portion of the image).
Luma represents the achromatic image, while the chroma components represent the color info.

**YCbCr** - is a family of color spaces used as a part of the color image pipeline
in video and digital photography systems.
Y is the luma component
and CB and CR are the blue-difference and red-difference chroma components.

**Chroma subsampling** is the practice of encoding images
by implementing less resolution for chroma information than for luma information.
Since the human visual system is much more sensitive to variations in brightness than color,
a video system can be optimized by devoting more bandwidth to the luma component,
than to the color difference components Cb and Cr.

**720p** (also called HD ready or standard HD)
is a progressive HDTV signal format with 720 horizontal lines
and an aspect ratio (AR) of 16:9, normally known as widescreen HDTV (1.78:1).

**WebM** is an audiovisual media file format.

**MPEG-4** - is a method of defining compression of audio and video data.

**H.264** or MPEG-4 Part 10 (**AVC** - Advanced Video Coding).

Used by 91% of video industry developers as of September 2019.
h.264 splits frames into blocks of 256 pixels each
and tries to predict what each subsequent block will look like
based on the rest of frames or previous ones.
(Uses decoded picture buffer to make predictions about next image in video).

**x264** is a free and open-source software library and CLI utility for encoding
video streams into the H.264/MPEG-4 AVC format.
