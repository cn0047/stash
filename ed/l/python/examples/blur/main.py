import imagesize
import cv2
import numpy


def tblr(f, height, width):
  if f == './z.png':
    top    = int(0.14 * height)
    bottom = int(0.25 * height)
    left   = int(0.60 * width)
    right  = int(0.71 * width)
  if f == './y.png':
    top    = int(0.16 * height)
    bottom = int(0.45 * height)
    left   = int(0.43 * width)
    right  = int(0.67 * width)
  if f == './t.png':
    top    = int(0.18 * height)
    bottom = int(0.37 * height)
    left   = int(0.35 * width)
    right  = int(0.63 * width)
  if f == './n.png':
    top    = int(0.065 * height)
    bottom = int(0.16 * height)
    left   = int(0.49 * width)
    right  = int(0.528 * width)
  return top, bottom, left, right


def p1():
  ksize = (11, 11)
  sigmaX = 10
  sigmaY = 10
  return ksize, sigmaX, sigmaY


def p2():
  ksize = (99, 99)
  sigmaX = 50
  sigmaY = 50
  return ksize, sigmaX, sigmaY


def p3():
  ksize = (99, 99)
  sigmaX = 99
  sigmaY = 99
  return ksize, sigmaX, sigmaY


def p4():
  ksize = (99, 99)
  sigmaX = 10
  sigmaY = 10
  return ksize, sigmaX, sigmaY


def p5():
  ksize = (99, 99)
  sigmaX = 10
  sigmaY= 10
  return ksize, sigmaX, sigmaY


def b1(f):
  frame = cv2.imread(f)
  height, width, *_ = frame.shape
  top, bottom, left, right = tblr(f, height, width)
  blur_bbox = frame[top:bottom, left:right]
  ksize, sigmax, sigmay = p5()
  blur_bbox = cv2.GaussianBlur(blur_bbox, ksize, sigmax, sigmaY=sigmay)
  # blur_bbox = cv2.blur(blur_bbox, (99, 99))
  frame[top:bottom, left:right] = blur_bbox
  cv2.imwrite(f+'.r.png', frame)


def b3(f):
  width, height = imagesize.get(f)
  top, bottom, left, right = tblr(f, height, width)
  blur_bbox = numpy.ones(shape = (top, bottom, left, right))
  frame[top:bottom, left:right] = blur_bbox
  cv2.imwrite(f+'.r.png', frame)


def b4(f):
  width, height = imagesize.get(f)
  frame = numpy.zeros((height, width, 3), numpy.uint8)
  cv2.imwrite(f+'.r.png', frame)


f=b4
# f('./t.png')
# f('./y.png')
# f('./z.png')
f('./n.png')
