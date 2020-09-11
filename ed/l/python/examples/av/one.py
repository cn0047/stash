import av
import os
import numpy as np

from PIL import Image


def to_dir_with_frames(v):
    os.system('rm -rf frames && mkdir frames')
    container = av.open(v)
    for frame in container.decode(video=0):
        img_pil = frame.to_image()
        # width, height = img_pil.size
        # img_arr = np.asarray(img_pil)
        img_pil.save('frames/_frame_%05d.jpg' % frame.index)


def demux_to_video(v):
    os.system('rm -rf res.mp4')
    container = av.open(v)
    in_video = container.streams.video[0]
    output = av.open('res.mp4', 'w')
    output.add_stream(template=in_video)
    for i, packet in enumerate(container.demux(in_video)):
        if packet.dts is None:
            continue
        output.mux(packet)
    output.close()


def demux_to_numpy_array(v):
    container = av.open(v)
    in_video = container.streams.video[0]
    frames = []
    for i, packet in enumerate(container.demux(in_video)):
        if packet.dts is None:
            continue
        for frame in packet.decode():
            img_pil = frame.to_image()
            # width, height = img_pil.size
            img_arr = np.asarray(img_pil)
            frames.append(img_arr)
    return np.asarray(frames)


def decode_to_numpy_array(v):
    container = av.open(v)
    frames = []
    for frame in container.decode(video=0):
        img_arr = np.asarray(frame.to_image())
        frames.append(img_arr)
    return np.asarray(frames)


def decode_to_numpy_array_2(v):
    container = av.open(v)
    container.streams.video[0].thread_type = 'AUTO'  # AUTO|FRAME
    frames = []
    for frame in container.decode(video=0):
        img_arr = np.asarray(frame.to_image())
        frames.append(img_arr)
    return np.asarray(frames)


def after_numpy_array_to_dir_with_frames(v):
    os.system('rm -rf frames && mkdir frames')
    frames = decode_to_numpy_array(v)

    i = 0
    for img_frame in frames:
        i += 1
        img = Image.fromarray(img_frame)
        img.save('frames/_frame_%05d.jpg' % i)


def from_numpy_array_to_video_file(v):
    frames = decode_to_numpy_array(v)

    container = av.open('res.mp4', mode='w')
    stream = container.add_stream('mpeg4') # , rate=fps
    stream.width = 1920
    stream.height = 1080
    stream.pix_fmt = 'yuv420p' # yuvj422p

    # option 1
    # for img_frame in frames:
    #     frame = av.VideoFrame.from_ndarray(img_frame, format='rgb24')
    #     for packet in stream.encode(frame):
    #         container.mux(packet)
    # for packet in stream.encode():
    #     container.mux(packet)

    # option 2
    for img_frame in frames:
        frame = av.VideoFrame.from_ndarray(img_frame, format='rgb24')
        packet = stream.encode(frame)
        container.mux(packet)

    container.close()


def from_numpy_array_to_buf_with_video_file(v):
    frames = decode_to_numpy_array(v)

    output = io.BytesIO()
    output.name = 'processed_video.mp4'

    container = av.open(output, mode='w')
    stream = container.add_stream('mpeg4') # h264|mpeg4
    stream.width = w
    stream.height = h
    stream.pix_fmt = 'yuv420p'
    for img_frame in frames:
        frame = av.VideoFrame.from_ndarray(img_frame, format='rgb24')
        packet = stream.encode(frame)
        container.mux(packet)
    container.close()

    output.seek(0)
    return output


v = 'path_to_video'
decode_to_numpy_array_2(v)
