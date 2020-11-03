def func1():
    d = {'foo': 'bar'}
    print(d['foo'])
    try:
      print(d['bar'])
    except KeyError as e:
      print("bar is not in dict")


def func2():
    d1 = dict(foo='bar')
    d2 = {'foo': 'bar'}
    d3 = {"foo": "bar"}
    print(d1)
    print(d2)
    print(d3)


# func1()
func2()
