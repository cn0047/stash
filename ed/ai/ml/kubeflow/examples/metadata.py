import kfp
import kfp.dsl as dsl
from kubernetes import client
from kubernetes.client.models import V1EnvVar


def render_ui():
  op = dsl.ContainerOp(
    name='render ui',
    image='cn007b/alpine',
    command=['sh', '-c'],
    arguments=['echo \'{"outputs":[{"type": "markdown", "storage": "inline", "source": "\u2705 all ok."}]}\' > /m.json'],
    output_artifact_paths={'mlpipeline-ui-metadata': '/m.json'},
  )
  op.container.set_image_pull_policy('Always')
  r = {'cpu': '1', 'memory': '10M', 'nvidia.com/gpu': '0', 'ephemeral-storage': '1G'}
  op.container.resources = {'requests': r, 'limits': r}
  op.add_pod_label('tensorboard', 'true')
  return op


@kfp.dsl.pipeline(name='Render UI', description='')
def j_classifier_pipeline_sanity():
  render_ui()


p = j_classifier_pipeline_sanity
pod_node_selector = 'kubeflow'
pipeline_conf = dsl.PipelineConf()
pipeline_conf.set_image_pull_secrets([client.V1LocalObjectReference(name="x")])
pipeline_conf.set_default_pod_node_selector(label_name="kubernetes.io/instancegroup", value=pod_node_selector)
pipeline = kfp.Client().create_run_from_pipeline_func(p, arguments={}, pipeline_conf=pipeline_conf)
