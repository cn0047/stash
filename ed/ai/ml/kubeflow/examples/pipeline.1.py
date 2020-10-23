# !pip3 install https://storage.googleapis.com/ml-pipeline/release/latest/kfp.tar.gz --upgrade --user

import json
import kfp
import kfp.dsl as dsl
from kubernetes import client


@dsl.component
def component_1():
  return dsl.ContainerOp(name='c1', image='cn007b/pi:ping')


@dsl.pipeline(name='k_pipeline_1', description='k_pipeline_1')
def k_pipeline_1():
  component_1()


pipeline_conf = dsl.PipelineConf().set_image_pull_secrets([client.V1LocalObjectReference(name="x")])
kfp.compiler.Compiler().compile(k_pipeline_1, 'pipeline.zip', pipeline_conf=pipeline_conf)
