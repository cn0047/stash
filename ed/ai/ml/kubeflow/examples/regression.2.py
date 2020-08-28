import kfp
import kfp.dsl as dsl
from kubernetes.client.models import V1EnvVar


@dsl.pipeline(name='regression pipeline')
def r_pipeline():
  c1 = dsl.ContainerOp(
    name='run',
    image='cn007b/pi:ai.tf',
    command=['sh', '-c'],
    arguments=['python3 /app/regression.py'],
  )
  c1.container.set_image_pull_policy('Always')


pipeline = kfp.Client().create_run_from_pipeline_func(r_pipeline, arguments={})
