kubeflow
-

[docs](https://www.kubeflow.org/docs/)
[pipelines](https://www.kubeflow.org/docs/pipelines/sdk/sdk-overview/)
[examples](https://github.com/kubeflow/pipelines/)

Kubeflow - ML toolkit for k8s.
Fairing - python package that makes it easy to train and deploy ML models on Kubeflow.

````sh
%%capture

!whoami
!rm -rf logs/

!pip3 install numpy --upgrade --user
!pip3 install -r requirements.txt --user

!kubectl get pods -n kubeflow

# !tensorboard --logdir=/home/jovyan/logs/ --bind_all
%load_ext tensorboard
import os; os.makedirs('./tb_logs', exist_ok=True)
%tensorboard --logdir './tb_logs'
````

````sh
%%writefile requirements.txt

tensorflow==2.1.0
````

````sh
!watch -n 1 nvidia-smi

docker run --rm -it nvidia/cuda nvidia-smi
````
