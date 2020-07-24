import boto3
from boto3.session import Session


ACCESS_KEY=''
SECRET_KEY=''


def get_client_1():
  return boto3.client('s3')


def get_client_2():
  return boto3.client('s3', aws_access_key_id=ACCESS_KEY, aws_secret_access_key=SECRET_KEY)


def ls(s3, b, p):
  resp = s3.list_objects_v2(Bucket=b, Prefix=p)
  for obj in resp['Contents']:
    file = obj['Key']
    print(file)


def ls_cb(s3, b, p, cb):
  resp = s3.list_objects_v2(Bucket=b, Prefix=p)
  for obj in resp['Contents']:
    file = obj['Key']
    cb(file)


def cp(s3, b, p, t):
  s3.download_file(b, p, t)


c = get_client_2()
ls(c, 'bkt', 'path/to/obj')
cp(c, 'bkt', 'path/to/obj', 'file/on/host/machine')
