def f1():
  x = 'foo'
  s = 'one line;' \
      f' line 2 with {x}'
  print(s) # one line; line 2 with foo

  s = f"""
  1) x = {x}
  """
  print(s)


def f2():
  s = ('line 1 '
    'line 2 ')
  print(s)


def multi_line_format():
  s = 'works'
  msg = f"""
  well, it {s}
  """
  print(msg) # well, it works


def multi_line_concat():
  msg = """
    foo
  """ \
  """
  bar
  """
  print(msg)


# f1()
# f2()
# multi_line_format()
multi_line_concat()
