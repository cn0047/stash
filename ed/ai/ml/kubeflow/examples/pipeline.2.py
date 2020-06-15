# %%capture
# !pip3 install https://storage.googleapis.com/ml-pipeline/release/latest/kfp.tar.gz --upgrade --user
# !pip3 install https://storage.googleapis.com/ml-pipeline/release/0.1.10/kfp.tar.gz --upgrade

# !pip3 install kfp kubernetes > /dev/null

import json
from string import Template
import kfp
import kfp.dsl as dsl
from kfp import components
from kfp.components import func_to_container_op
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
    command=['touch', '/tmp/x'],
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


@kfp.dsl.pipeline(name='pipeline 1',description='')
def pipeline_1():
  env_var = V1EnvVar(name='CUDA_VISIBLE_DEVICES', value='0,1')
  step_1 = init_component()
  step_2 = run_component().add_env_variable(env_var).after(step_1)
  step_3 = inspect_component().add_env_variable(env_var).after(step_2)


pipeline_conf = dsl.PipelineConf().set_image_pull_secrets([client.V1LocalObjectReference(name="regcred")])
pipeline_conf=kfp.dsl.PipelineConf().set_image_pull_secrets([client.V1LocalObjectReference(name="regcred")])
pipeline = kfp.Client().create_run_from_pipeline_func(pipeline_1, arguments={}, pipeline_conf=pipeline_conf)
kfp.compiler.Compiler().compile(pipeline_1, 'pipeline.zip', pipeline_conf=pipeline_conf)
