docker run -ti --rm -v $PWD:/gh -w /gh tensorflow/tensorflow sh -c '
  python3 ed/ai/ml/tensorflow/examples/one.py
'

# regression
cd ed/ai/ml/tensorflow/examples
virtualenv ./venv
source ./venv/bin/activate
pip3 install numpy matplotlib tensorflow tensorboard
python3 regression.py
tensorboard --logdir=/tmp/tf_logs --bind_all
deactivate
