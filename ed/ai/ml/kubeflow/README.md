kubeflow
-

[docs](https://www.kubeflow.org/docs/)
[pipelines](https://www.kubeflow.org/docs/pipelines/sdk/sdk-overview/)
[examples](https://github.com/kubeflow/pipelines/)
[examples](https://github.com/kubeflow/examples)

Kubeflow - ML toolkit for k8s.
Fairing - python package that makes it easy to train and deploy ML models on Kubeflow.
Component - self-contained set of code that performs one step in the ML workflow (pipeline).

Artifacts in UI tab will be available once ContainerOp finished.

## Jupyter notebooks:

````sh
%%capture

!whoami
!rm -rf logs/

!pip3 install numpy --upgrade --user
!pip3 install -r requirements.txt --user

!kubectl get pods -n kubeflow

pipeline.run_info

%load_ext tensorboard
%reload_ext
%load_ext tensorboard.notebook
import os; os.makedirs('./tb_logs', exist_ok=True)
%tensorboard --logdir './tb_logs'

!watch -n 1 nvidia-smi
docker run --rm -it nvidia/cuda sh -c 'watch -n 1 nvidia-smi'

!kubectl describe pod
````

````sh
%%writefile requirements.txt

tensorflow==2.1.0
````

````py
op = dsl.ContainerOp(
    name='run',
    image='cn007b/alpine'
)
op.add_pod_label('lbl', 'foo')
op.add_pod_annotation('anno', 'bar')

# Pod info
from kubernetes import client, config
pod_name = "gpu-main-0"
instance = client.CoreV1Api()
response = instance.read_namespaced_pod(name=pod_name, namespace='ns')
print(response)
print(response.status.pod_ip)

# Run pipeline
client = kfp.Client()
experiment = client.create_experiment('j_run')
run = client.run_pipeline(
    experiment_id=experiment.id,
    job_name='j_run_job',
    pipeline_id=PIPELINE_ID,
)

# Sidecar
sc1 = dsl.Sidecar(
    name='sc1',
    image='cn007b/alpine',
    args=[
        "/bin/sh",
        "-c",
        "curl -i -XPOST 'https://realtimelog.herokuapp.com:443/rkc8q6llprn' -H 'Content-Type: application/json' -d '{\"status\": \"204\"}'",
    ]
)
op = dsl.ContainerOp(
    name='op1',
    image='cn007b/alpine',
    command=['sh', '-c'],
    arguments=['while true; do echo `date`; sleep 0.5; done'],
    sidecars=[sc1],
)
````
