import multiprocessing

def f() -> None:
  print('I am in\n')

p = multiprocessing.Process(target=f)
p.start()
p.join()
