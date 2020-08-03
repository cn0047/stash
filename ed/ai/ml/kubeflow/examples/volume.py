import kfp
import kfp.dsl as dsl
from kubernetes import client
from kubernetes.client.models import V1EnvVar


def get_vol(n: str, s: str, m):
  return dsl.VolumeOp(name=f'{n}-j-pv', resource_name=f'{n}-j-pvc', storage_class='aws-efs', size=s, modes=m)


def component_w(pv):
  op = dsl.ContainerOp(
    name='write',
    image='cn007b/alpine',
    command=['sh', '-c'],
    arguments=["echo 'ok at: '`date` > /tmp/x.log"],
    pvolumes=pv,
  )
  op.container.set_image_pull_policy('Always')
  r = {'cpu': '1','memory': '10M','nvidia.com/gpu': '0'}
  op.container.resources = {'requests': r, 'limits': r}
  return op


def component_r(pv):
  op = dsl.ContainerOp(
    name='read',
    image='cn007b/alpine',
    command=['sh', '-c'],
    arguments=['cat /tmp/x.log'],
    pvolumes=pv,
  )
  op.container.set_image_pull_policy('Always')
  r = {'cpu': '1','memory': '10M','nvidia.com/gpu': '0'}
  op.container.resources = {'requests': r, 'limits': r}
  return op


@kfp.dsl.pipeline(name='check vol', description='')
def pipeline_check_vol():
  v1 = get_vol('main', '1M', dsl.VOLUME_MODE_RWM)
  pv = {'/tmp': v1.volume}
  c1 = component_w(pv)
  c2 = component_r(pv).after(c1)
  v1.delete().after(c2)


p = pipeline_check_vol
pod_node_selector = 'kubeflow'
pipeline_conf = dsl.PipelineConf()
pipeline_conf.set_image_pull_secrets([client.V1LocalObjectReference(name="x")])
pipeline_conf.set_default_pod_node_selector(label_name="kubernetes.io/instancegroup", value=pod_node_selector)
pipeline = kfp.Client().create_run_from_pipeline_func(p, arguments={}, pipeline_conf=pipeline_conf)
