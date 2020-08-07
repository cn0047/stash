kubeflow
-

[docs](https://www.kubeflow.org/docs/)
[pipelines](https://www.kubeflow.org/docs/pipelines/sdk/sdk-overview/)
[examples](https://github.com/kubeflow/pipelines/)

Kubeflow - ML toolkit for k8s.
Fairing - python package that makes it easy to train and deploy ML models on Kubeflow.
Component - self-contained set of code that performs one step in the ML workflow (pipeline).

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
%tensorboard --logdir './tb_logs' --bind_all
%tensorboard --logdir './tb_logs' --host 0.0.0.0 --port 6008 --reload_interval=60
````

````sh
%%writefile requirements.txt

tensorflow==2.1.0
````

````sh
!watch -n 1 nvidia-smi

docker run --rm -it nvidia/cuda sh -c 'watch -n 1 nvidia-smi'
````

````py
!kubectl describe pod

from kubernetes import client, config

pod_name = "gpu-main-0"
instance = client.CoreV1Api()
response = instance.read_namespaced_pod(name=pod_name, namespace='ns')
print(response)
print(response.status.pod_ip)
````
