import kfp
import kfp.dsl as dsl
from kubernetes import client
from kubernetes.client.models import V1EnvVar


@kfp.dsl.pipeline(name='check x', description='')
def check_x():
  op1 = dsl.ContainerOp(
    name='write log',
    image='cn007b/alpine',
    command=['sh', '-c'],
    arguments=['echo ok > /tmp/x.log'],
    file_outputs={'x': '/tmp/x.log'}
  )
  op2 = dsl.ContainerOp(
    name='read log',
    image='cn007b/alpine',
    command=['sh', '-c'],
    arguments=['echo %s ' %op1.output]
  )


p = check_x
pod_node_selector = 'kubeflow'
pipeline_conf = dsl.PipelineConf()
pipeline_conf.set_image_pull_secrets([client.V1LocalObjectReference(name="x")])
pipeline_conf.set_default_pod_node_selector(label_name="kubernetes.io/instancegroup", value=pod_node_selector)
pipeline = kfp.Client().create_run_from_pipeline_func(p, arguments={}, pipeline_conf=pipeline_conf)
