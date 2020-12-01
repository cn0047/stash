import kfp
import kfp.dsl as dsl
from kubernetes import client


def hw():
  op = dsl.ContainerOp(
    name='hw',
    image='cn007b/alpine',
    command=['sh', '-c'],
    arguments=['echo hello world'],
  )
  op.container.set_image_pull_policy('Always')
  r = {'cpu': '1', 'memory': '10M', 'nvidia.com/gpu': '0', 'ephemeral-storage': '1G'}
  op.container.resources = {'requests': r, 'limits': r}
  return op


@kfp.dsl.pipeline(name='Render UI', description='')
def j_pipeline():
  hw()


p = j_pipeline
pipeline_conf = dsl.PipelineConf()
pipeline_conf.set_image_pull_secrets([client.V1LocalObjectReference(name="x")])
pipeline = kfp.Client().create_run_from_pipeline_func(p, arguments={}, pipeline_conf=pipeline_conf)
