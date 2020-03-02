Video
-

Video file is container for: video stream, audio & metadata.
Bit rate - amount of data encoded per second in a vide.
Resolution - pixel size of video (1920x1080 - HD).
DVD    - Bit rate = 4-8Mbps
Bluray - Bit rate = 24-40Mbps

Codec (Coder/Decoder) specified in metadata, like:
~~divx/xvid, ffmpeg, x264~~,
h.264, h.265(HEVC), mp3, aac, dolby, vp9.

Video format: standardized set of rules to store containers,
codecs, metadata, like: mp4, hls.

Best video for youtube: container: mp4, codec: h.264,
audio: AAC-LC, channels: stereo/ stereo 5.1, sample rate: 96khz,
aspect ratio: 19.6, frame rate: 60fps

Frame rate - count of images in 1 second (FPS).
24/30/60/120 - the bigger count - the higher quality and smoothly video.
For 60 frame rate it's possible to slow down video speed in video editor without losing quality.

#### H.264 or MPEG-4 Part 10 (AVC - Advanced Video Coding)

Used by 91% of video industry developers as of September 2019.
h.264 splits frames into blocks of 256 pixels each
and tries to predict what each subsequent block will look like
based on the rest of frames or previous ones.

#### GOP - Group of pictures

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
