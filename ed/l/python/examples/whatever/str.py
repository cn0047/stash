def f1():
  x = 'foo'
  s = 'one line;' \
      f' line 2 with {x}'
  print(s) # one line; line 2 with foo


f1()
