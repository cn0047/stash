import kfp
import kfp.dsl as dsl
from kubernetes import client


def c_print(msg):
  op = dsl.ContainerOp(
    name='one',
    image='cn007b/alpine',
    command=['sh', '-c'],
    arguments=[f'echo {msg}'],
  )
  op.container.set_image_pull_policy('Always')
  r = {'cpu': '1', 'memory': '10M', 'nvidia.com/gpu': '0', 'ephemeral-storage': '1G'}
  op.container.resources = {'requests': r, 'limits': r}
  return op


@kfp.dsl.pipeline(name='pipeline 1', description='')
def j_pipeline_1():
  c_print('p1')


@kfp.dsl.pipeline(name='pipeline 2', description='')
def j_pipeline_2():
  c_print('p2')


pipeline = kfp.Client().create_run_from_pipeline_func(j_pipeline_1, arguments={})
pipeline.wait_for_run_completion(timeout=30)
pipeline = kfp.Client().create_run_from_pipeline_func(j_pipeline_2, arguments={})
