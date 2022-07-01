def func1():
    d = {'foo': 'bar'}
    print(d['foo'])
    try:
      print(d['bar'])
    except KeyError as e:
      print('bar is not in dict')


def func2():
    d1 = dict(foo='bar')
    d3 = {'foo': 'bar'}
    d3['code'] = 200
    print('[func2]', d1, d3, d3.get('code'), d3.get('x'))


# func1()
func2()
