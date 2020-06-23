v = 204
f = fr'smtng: {v} \n' # smtng: 204 \n
print(f)

# new formatting style
fs = '{} is not {} \n'
print(fs.format('foo', 'bar')) # foo is not bar

name = 'bond'
age = '20'
print(f'Hello, {name}. You are {age}.\n') # Hello, bond. You are 20.
