TensorFlow
-

[docs](https://www.tensorflow.org/tutorials)

````sh
pip3 install tensorflow


virtualenv --system-site-packages -p python3 ./venv
source ./venv/bin/activate
python -c 'import tensorflow as tf; print(tf.reduce_sum(tf.random.normal([1000, 1000])))'
deactivate
````

computation graph - static
define before run

Constant - immutable value.
Placeholder - holds place for tensor that will be available at runtime.
Variable - mutable tensor values.
