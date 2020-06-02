import tensorflow.compat.v1 as tf
tf.disable_v2_behavior()


def p(v):
  print("\n==========\n", v, "\n==========\n")


def f1():
  b = tf.placeholder(tf.float32, [None, 1], name='b')
  p(b)


f1()
