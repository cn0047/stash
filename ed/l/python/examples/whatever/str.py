def f1():
  x = 'foo'
  s = 'one line;' \
      f' line 2 with {x}'
  print(s) # one line; line 2 with foo


def f2():
  s = ('line 1 '
    'line 2 ')
  print(s)


# f1()
f2()
