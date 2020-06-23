import requests

res = requests.get('https://api.github.com/users/cn007b')
print(res)
print(res.text)
