# %%capture
# !pip3 install https://storage.googleapis.com/ml-pipeline/release/latest/kfp.tar.gz --upgrade --user
# !pip3 install https://storage.googleapis.com/ml-pipeline/release/0.1.10/kfp.tar.gz --upgrade --user

# !pip3 install kfp kubernetes > /dev/null

import json
from string import Template
import kfp
import kfp.dsl as dsl
from kfp import components
from kfp.components import func_to_container_op
from kubernetes import client
from kubernetes.client.models import V1EnvVar


@dsl.component
def component_1():
  return dsl.ContainerOp(name='c3', image='cn007b/pi:ping')


@kfp.dsl.pipeline(name='k_pipeline_1', description='k_pipeline_1')
def k_pipeline_1():
  component_1()


pipeline_conf = dsl.PipelineConf().set_image_pull_secrets([client.V1LocalObjectReference(name="regcred")])
pipeline_conf=kfp.dsl.PipelineConf().set_image_pull_secrets([client.V1LocalObjectReference(name="regcred")])
kfp.compiler.Compiler().compile(k_pipeline_1, 'pipeline.zip', pipeline_conf=pipeline_conf)
