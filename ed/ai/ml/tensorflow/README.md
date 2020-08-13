TensorFlow
-

[docs](https://www.tensorflow.org/tutorials)

computation graph - static
define before run

````sh
pip3 install tensorflow

tensorboard --logdir './tb_logs' --bind_all
tensorboard --logdir './tf_logs' --bind_all
tensorboard --logdir './tb_logs' --host 0.0.0.0 --port 6008 --reload_interval=60
--bind_all                       # in v2
--logdir_spec b1:tf_logs,b2:tfl2 # in v2
````

````sh
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
