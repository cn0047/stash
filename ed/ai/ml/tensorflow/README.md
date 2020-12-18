TensorFlow
-

[docs](https://www.tensorflow.org/tutorials)

computation graph - static
define before run

````sh
pip3 install tensorflow
pip3 install https://storage.googleapis.com/tensorflow/mac/cpu/tensorflow-1.14.0-py3-none-any.whl
````
````py
import tensorflow as tf; tf.__version__
import tensorflow as tf; tf.test.is_gpu_available()
import tensorflow as tf; tf.config.experimental.list_physical_devices()
from tensorflow.python.client import device_lib; device_lib.list_local_devices()

print([
    # tf.config.list_physical_devices('GPU'),
    tf.config.experimental.list_physical_devices('GPU'),
])

# Create model using the TensorFlow Keras library
model = tf.keras.Sequential()
````

Constant - immutable value.
Placeholder - holds place for tensor that will be available at runtime.
Variable - mutable tensor values.

# TensorBoard

````sh
tensorboard --logdir './tb_logs' --bind_all
tensorboard --logdir './tf_logs' --bind_all
tensorboard --logdir 's3://path/to/logs/'
tensorboard --host 0.0.0.0 --port 6008 --logdir './tf_logs'
  --reload_interval=60             # refresh UI
  --bind_all                       # in v2
  --logdir_spec b1:tf_logs,b2:tfl2 # in v2

AWS_REGION=us-west-1 tensorboard --logdir s3://bkt/path

find / -name "*.tfevents.*" -type f
````
