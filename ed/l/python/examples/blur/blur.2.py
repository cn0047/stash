import cv2

import PIL.Image
import numpy as np

from typing import Tuple


def blur_image(img: np.array, ksize: Tuple[int, int] = (10, 10)) -> np.array:
    """
    Median pooling blur.

    Parameters:
        img: image to blur.
        ksize: (width, height) pooling kernel, define size of each blurring region.
    """

    if len(img.shape) == 2:
        img = np.expand_dims(img, axis=-1)

    h, w, c = img.shape

    assert h % ksize[1] == 0 and w % ksize[0] == 0, 'Image sizes must be dividable by blur kernel sizes.'

    ny = h // ksize[1]
    nx = w // ksize[0]

    new_shape = (ny, ksize[1], nx, ksize[0]) + (c,)
    img = img.copy().reshape(new_shape)
    img[...] = np.median(img, axis=(1, 3), keepdims=True)

    if c > 1:
        img = img.reshape(h, w, c)
    else:
        img = img.reshape(h, w)

    return img


def b1():
    rec_img = PIL.Image.open('a.jpeg')
    img = np.array(rec_img)
    res_img = blur_image(img, ksize=(8, 8))
    cv2.imwrite('a.after.jpeg', res_img)


b1()
