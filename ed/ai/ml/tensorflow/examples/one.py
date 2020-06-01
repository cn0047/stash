"""
docker run -ti --rm -v $PWD:/gh -w /gh tensorflow/tensorflow sh -c '
  python3 ed/ai/ml/tensorflow/examples/one.py
'
"""

import tensorflow as tf


def p(v):
  print("\n==========\n", v, "\n==========\n")


def f1():
  c = tf.constant('100')
  p(c)


def f2():
  c = tf.constant('100')
  tf.print(c)


def f3():
  x1 = tf.constant([1,2,3,4])
  x2 = tf.constant([5,6,7,8])
  r = tf.multiply(x1, x2)
  p(r)



# f1()
# f2()
f3()
