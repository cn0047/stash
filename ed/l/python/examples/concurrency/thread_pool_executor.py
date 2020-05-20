from concurrent.futures import ThreadPoolExecutor
import threading
import random


def task(id: int):
    print(f"Executing task {id}")
    print("Task {} executed {}".format(id, threading.current_thread()))


def main():
    executor = ThreadPoolExecutor(max_workers=2)
    task1 = executor.submit(task, 1)
    task2 = executor.submit(task, 2)
    task3 = executor.submit(task, 3)


if __name__ == '__main__':
    main()
