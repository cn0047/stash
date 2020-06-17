import boto3
from boto3.session import Session


ACCESS_KEY=''
SECRET_KEY=''


def get_client_1():
  return boto3.client('s3')


def get_client_2():
  return boto3.client('s3', aws_access_key_id=ACCESS_KEY, aws_secret_access_key=SECRET_KEY)


def ls1(b, p):
  s3 = get_client_2()
  resp = s3.list_objects_v2(Bucket=b, Prefix=p)
  for obj in resp['Contents']:
    files = obj['Key']
    print(files)


ls1('bkt', 'path/to/obj')
