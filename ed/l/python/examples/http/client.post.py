import requests
import json

requests.post(
  'https://realtimelog.herokuapp.com:443/q405uin07bm',
  data=json.dumps({'code': 200, 'status': 'OK'}),
  headers={'Content-Type': 'application/json'},
)
