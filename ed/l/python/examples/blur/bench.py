import cv2


def b(f):
  img = cv2.imread(f)

  top, bottom, left, right = 286, 308, 111, 144
  rec = img[top:bottom, left:right]
  rec = cv2.GaussianBlur(rec, (3, 9), 5)
  img[top:bottom, left:right] = rec

  cv2.imwrite('img.after.py.png', img)


if __name__ == '__main__':
  b('./img.before.png')
