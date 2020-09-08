from multiprocessing import Pool


def add(a, b):
  return a + b


def write(a, b):
    print(a, ':', b)


w = lambda a, b: print(a, ':', b)

a = ['1', '2', '3']
b = ['4', '5', '6']

pool = Pool(2)
# pool.starmap(write, zip(a,b))
# pool.starmap(w, zip(a,b)) #  Can't pickle <function <lambda> at 0x109fd8c80>
r = pool.starmap(add, zip(a,b))
pool.close()
pool.join()
print('pool.starmap result:', r)
