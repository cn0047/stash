TensorFlow
-

[docs](https://www.tensorflow.org/tutorials)

````sh
pip3 install tensorflow


virtualenv --system-site-packages -p python3 ./venv
source ./venv/bin/activate
python -c 'import tensorflow as tf; print(tf.reduce_sum(tf.random.normal([1000, 1000])))'
deactivate

from tensorflow.python.client import device_lib; device_lib.list_local_devices()
import tensorflow as tf; tf.test.is_gpu_available()
print([
    tf.__version__,
    tf.test.is_gpu_available(),
    # tf.config.list_physical_devices('GPU'),
    tf.config.experimental.list_physical_devices('GPU'),
])

# Create model using the TensorFlow Keras library
model = tf.keras.Sequential()
````

computation graph - static
define before run

Constant - immutable value.
Placeholder - holds place for tensor that will be available at runtime.
Variable - mutable tensor values.
