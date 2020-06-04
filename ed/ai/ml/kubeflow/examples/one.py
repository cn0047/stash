from kubernetes import client
from kubernetes.client.models import V1EnvVar

def f1():
  env_var = V1EnvVar(name='CUDA_VISIBLE_DEVICES', value='-1')
  print(env_var)


def f2():
    client.CoreV1Api()
    instance = client.CoreV1Api()
    print(instance.read_namespace(name='kubeflow'))
    response = instance.read_namespaced_pod(name='mypod', namespace='kubeflow')
    print(response)


# f1()
f2()
