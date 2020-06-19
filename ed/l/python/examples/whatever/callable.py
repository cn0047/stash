def simple_print(s):
  print(s)


def wrapped_print(s):
  print("\n===\n", s, "\n===\n")


def call_print_cb(cb):
  cb('from func f')


def f1():
  # ok
  func_name = 'simple_print'
  func_name = 'wrapped_print'
  call_print_cb(locals()[func_name])
  # error
  func_name = 'print'
  call_print_cb(globals()[func_name])


f1()
