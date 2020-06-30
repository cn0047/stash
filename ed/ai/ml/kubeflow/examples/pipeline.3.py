import kfp
import kfp.dsl as dsl
import kubernetes
from kubernetes import client
from kubernetes.client.models import V1EnvVar


def get_vol():
  return dsl.VolumeOp(
    name='kp1-vol',
    resource_name='kp1-pvc',
    size='1M',
    modes=dsl.VOLUME_MODE_RWO,
    # annotations={'metadata': {'finalizers': None}}
    # annotations={'finalizers': ''}
  )


def init_component(pv):
  return dsl.ContainerOp(
    name='init',
    image='cn007b/alpine',
    command=['sh', '-c'],
    arguments=['echo ok2 >> /tmp/x.txt'],
    pvolumes=pv
  )


def inspect_component(pv):
  return dsl.ContainerOp(
    name='inspect',
    image='cn007b/alpine',
    command=['sh', '-c'],
    arguments=['cat /tmp/x.txt'],
    pvolumes=pv
  )


def exit_component():
  return dsl.ContainerOp(
    is_exit_handler=True,
    name='exit',
    image='cn007b/alpine',
    arguments=["sh", "-c", "curl -XPOST 'https://realtimelog.herokuapp.com:443/m40p44jdi8l' -H 'Content-Type: application/json' -d '{\"code\": 200}'"]
  )


@dsl.pipeline(name='k_pipeline_3', description='')
def k_pipeline_3():
    v = get_vol()
    pv = {'/tmp': v.volume}
    step_exit = exit_component()
    with dsl.ExitHandler(step_exit):
        step_1 = init_component(pv)
        step_2 = inspect_component(pv).after(step_1)
        v.delete().after(step_2)


pipeline_conf = dsl.PipelineConf().set_image_pull_secrets([client.V1LocalObjectReference(name="regcred")])
pipeline_conf=kfp.dsl.PipelineConf().set_image_pull_secrets([client.V1LocalObjectReference(name="regcred")])
kfp.compiler.Compiler().compile(k_pipeline_3, 'pipeline.zip', pipeline_conf=pipeline_conf)
