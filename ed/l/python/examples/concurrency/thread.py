import threading

def f(val: int) -> None:
  print(f'got val: {val} \n')

t = threading.Thread(target=f, args=(204,))
t.start()
t.join()
