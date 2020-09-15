from multiprocessing import Pool


POOL_SIZE = 4


def add(a, b):
  return a + b


def write(a, b):
    print(a, ':', b)


def with_write_func(a, b):
    pool = Pool(POOL_SIZE)
    pool.starmap(write, zip(a, b))
    pool.close()
    pool.join()


def with_lambda(a, b):
    w = lambda a, b: print(a, ':', b)
    pool = Pool(POOL_SIZE)
    pool.starmap(w, zip(a,b))  # Can't pickle <function <lambda> at 0x109fd8c80>
    pool.close()
    pool.join()


def with_add_func(a, b):
    pool = Pool(POOL_SIZE)
    r = pool.starmap(add, zip(a, b))
    pool.close()
    pool.join()
    print('[with_add_func] result:', r)  # [with_add_func] result: ['14', '25', '36']


a = ['1', '2', '3', 'a', 'b', 'c']
b = ['4', '5', '6', 'd', 'e', 'f']
# with_write_func(a, b)
# with_lambda(a, b)
with_add_func(a, b)
