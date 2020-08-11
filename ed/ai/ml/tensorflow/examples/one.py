import tensorflow as tf
import tensorflow.compat.v1 as tf1


def p(v):
  print("\n==========\n", v, "\n==========\n")


def f1():
  c = tf.constant('100')
  p(c)
  tf.print(c)


def f2():
  a = tf.Variable(2.0, name='a')
  p(a)


def f3():
  x1 = tf.constant([1,2,3,4])
  x2 = tf.constant([5,6,7,8])
  r = tf.multiply(x1, x2)
  p(r)


def f4():
  tf1.disable_v2_behavior()
  b = tf1.placeholder(tf1.float32, [None, 1], name='b')
  print("\n==========\n", b, "\n==========\n")


# f1()
# f2()
f3()
