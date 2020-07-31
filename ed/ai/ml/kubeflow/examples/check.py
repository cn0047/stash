import kfp
import kfp.dsl as dsl
from kubernetes import client
from kubernetes.client.models import V1EnvVar


@dsl.component
def check_cpu_component():
  op = dsl.ContainerOp(
    name='check CPU',
    image='x',
    command=['sh', '-c'],
    arguments=['y'],
  )
  op.container.set_image_pull_policy('Always')
  r = {'cpu': '1','memory': '2G','nvidia.com/gpu': '0'}
  op.container.resources = {'requests': r,'limits': r}
  return op


@dsl.component
def check_gpu_component():
  op = dsl.ContainerOp(
    name='check GPU',
    image='x',
    command=['sh', '-c'],
    arguments=['y'],
  )
  op.container.set_image_pull_policy('Always')
  r = {'cpu': '1','memory': '2G','nvidia.com/gpu': '1'}
  op.container.resources = {'requests': r,'limits': r}
  return op


@dsl.component
def check_s3_component():
  op = dsl.ContainerOp(
    name='check S3',
    image='cn007b/alpine',
    command=['sh', '-c'],
    arguments=['aws s3 ls s3://...'],
  )
  op.container.set_image_pull_policy('Always')
  r = {'cpu': '1','memory': '50M','nvidia.com/gpu': '0'}
  op.container.resources = {'requests': r, 'limits': r}
  return op


@dsl.component
def write_log_component():
  op = dsl.ContainerOp(
    name='write log',
    image='cn007b/alpine',
    command=['sh', '-c'],
    arguments=["echo 'sanity check performed at: '`date` > /tmp/object-classifier-sanity-check.log"],
    output_artifact_paths={'object-classifier-sanity-check': '/tmp/object-classifier-sanity-check.log'},
  )
  op.container.set_image_pull_policy('Always')
  r = {'cpu': '1','memory': '10M','nvidia.com/gpu': '0'}
  op.container.resources = {'requests': r, 'limits': r}
  return op


def render_ui():
  script = """
import json
import os
import sys

meta = {
    'outputs': [
        {
          'type': 'markdown',
          'storage': 'inline',
          'source': '✅ sanity check - passed.'
        }
    ]
}
with open('/mlpipeline-ui-metadata.json', 'w') as f:
    json.dump(meta, f)
    """
  op = dsl.ContainerOp(
    name='render ui',
    image='python:alpine3.7',
    command=['sh', '-c'],
    arguments=[f'python -c "{script}"'],
    output_artifact_paths={'mlpipeline-ui-metadata': '/mlpipeline-ui-metadata.json'},
  )
  op.container.set_image_pull_policy('Always')
  r = {'cpu': '1','memory': '10M','nvidia.com/gpu': '0'}
  op.container.resources = {'requests': r, 'limits': r}
  return op


@kfp.dsl.pipeline(name='sanity check', description='')
def pipeline_sanity():
  c1 = check_s3_component()
  c2 = check_cpu_component()
  c3 = check_gpu_component().add_env_variable(V1EnvVar(name='CUDA_VISIBLE_DEVICES', value='0,1'))
  write_log_component().after(c1, c2, c3)
  render_ui().after(c1, c2, c3)


p = pipeline_sanity
pod_node_selector = 'kubeflow'
pipeline_conf = dsl.PipelineConf()
pipeline_conf.set_image_pull_secrets([client.V1LocalObjectReference(name="xx")])
pipeline_conf.set_default_pod_node_selector(label_name="kubernetes.io/instancegroup", value=pod_node_selector)
pipeline = kfp.Client().create_run_from_pipeline_func(p, arguments={}, pipeline_conf=pipeline_conf)
