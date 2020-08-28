import kfp
import kfp.dsl as dsl
from kubernetes.client.models import V1EnvVar


@dsl.component
def run_component():
  op = dsl.ContainerOp(
    name='run',
    image='cn007b/pi:ai.tf',
    command=['sh', '-c'],
    arguments=['python3 /app/regression.py'],
  )
  op.container.set_image_pull_policy('Always')
  return op


@dsl.pipeline(name='regression pipeline')
def r_pipeline():
  foo = V1EnvVar(name='FOO', value='bar')
  step_1 = run_component().add_env_variable(foo)


pipeline = kfp.Client().create_run_from_pipeline_func(r_pipeline, arguments={})
