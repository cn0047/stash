def f1(l):
  print('[f1]:', l[1:]) # [2, 3, 4, 5, 6, 7]


def f2(l):
  print('[f2]:', l[::2]) # [1, 3, 5, 7]


def queue():
  q = []
  q.append(1)
  q.append(2)
  q.append(3)
  q.append(4)
  v1 = q.pop(0)
  v2 = q.pop(0)
  print('[queue]:', v1, v2, q) # 1 2 [3, 4]


def stack():
  s = []
  s.insert(0, 1)
  s.insert(0, 2)
  s.insert(0, 3)
  s.insert(0, 4)
  v1 = s.pop(0)
  v2 = s.pop(0)
  print('[stack]:', v1, v2, s)


l = [1, 2, 3, 4, 5, 6, 7]
# f1(l)
# f2(l)
# queue()
stack()
