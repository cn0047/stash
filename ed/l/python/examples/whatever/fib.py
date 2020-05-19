import time


def fib(n: int) -> int:
    time.sleep(0.1)
    if n < 2:
        return 1
    return fib(n - 1) + fib(n - 2)


def main(n: int) -> int:
    return fib(n)


print(main(5))
