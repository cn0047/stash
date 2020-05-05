d = {'foo': 'bar'}
print(d['foo'])
try:
  print(d['bar'])
except KeyError as e:
  print("bar is not in dict")
