import kfp
import kfp.dsl as dsl
from kubernetes import client
from kubernetes.client.models import V1EnvVar


PVC = 'volume-name'
pvolumes = {"/app": dsl.PipelineVolume(pvc=PVC), "/tmp":  dsl.PipelineVolume(pvc=PVC)}


@dsl.component
def init_component():
  return dsl.ContainerOp(
    name='init',
    image='cn007b/alpine',
    command=['rm', '-rf', '/tmp/x'],
    pvolumes=pvolumes
  )


@dsl.component
def run_component():
  op = dsl.ContainerOp(
    name='run',
    image='cn007b/alpine',
    command=["sh", "-c"],
    arguments=["echo ok > /tmp/x"],
    pvolumes=pvolumes
  )
  op.container.set_image_pull_policy('Always')
  return op


@dsl.component
def inspect_component():
  return dsl.ContainerOp(
    name='inspect',
    image='cn007b/pi:pinger',
    pvolumes=pvolumes
  )


@kfp.dsl.pipeline(name='k pipeline 2',description='')
def k_pipeline_2():
  env_var = V1EnvVar(name='CUDA_VISIBLE_DEVICES', value='0,1')
  step_1 = init_component()
  step_2 = run_component().add_env_variable(env_var).after(step_1)
  step_3 = inspect_component().add_env_variable(env_var).after(step_2)


pipeline_conf = dsl.PipelineConf().set_image_pull_secrets([client.V1LocalObjectReference(name="x")])
pipeline_conf=kfp.dsl.PipelineConf().set_image_pull_secrets([client.V1LocalObjectReference(name="x")])
kfp.compiler.Compiler().compile(k_pipeline_2, 'pipeline.zip', pipeline_conf=pipeline_conf)
pipeline = kfp.Client().create_run_from_pipeline_func(k_pipeline_2, arguments={}, pipeline_conf=pipeline_conf)
