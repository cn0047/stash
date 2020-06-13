import json
import requests


class Client():
  """REAL-TIME log client."""


  def __init__(self, id):
    h = 'https://realtimelog.herokuapp.com:443/'
    if id:
      self.url = h + id
      return
    res = requests.get(h)
    self.url = res.url


  def get_url(self):
    return self.url


  def msg(self, data):
    requests.post(
      self.url,
      headers={'Content-Type': 'application/json'},
      data=json.dumps(data),
    )


c = Client('x')
print('Open: ', c.get_url())
c.msg({'code': 200, 'status': 'OK'})
