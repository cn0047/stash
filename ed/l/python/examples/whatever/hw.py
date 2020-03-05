"""
This is simple hw script.
"""


def f2():
    name = "Bond"
    print("hi {0}".format(name))
    print(f"hi {name}")


def f():
    """
    multiline string.
    """
    str = "Hello" \
          " World!"
    print(f"{str} \n")


f()
f2()
