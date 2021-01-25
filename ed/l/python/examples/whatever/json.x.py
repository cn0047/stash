import json


def from_str():
    s = '{"data": {"foo": "bar"}}'
    j = json.loads(s)
    data = j['data'] if 'data' in j else None
    info = j['info'] if 'info' in j else None
    print(data)
    print(info)


from_str()
