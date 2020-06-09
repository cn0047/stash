import numpy as np
import matplotlib.pyplot as plt
import tensorflow.compat.v1 as tf
tf.disable_v2_behavior()

np.random.seed(101)
tf.set_random_seed(101)

# Genrating random linear data
with tf.name_scope('genrating_data'):
  # There will be 50 data points ranging from 0 to 50
  x = np.linspace(0, 50, 50)
  y = np.linspace(0, 50, 50)
  # Adding noise to the random linear data
  x += np.random.uniform(-4, 4, 50)
  y += np.random.uniform(-4, 4, 50)
  n = len(x) # Number of data points

# Plot of Training Data
with tf.name_scope('training_data_plot'):
  plt.scatter(x, y)
  plt.xlabel('x')
  plt.xlabel('y')
  plt.title("Training Data")
  plt.show()

with tf.name_scope('xywb'):
  X = tf.placeholder("float")
  Y = tf.placeholder("float")
  W = tf.Variable(np.random.randn(), name = "W")
  b = tf.Variable(np.random.randn(), name = "b")
  learning_rate = 0.01
  training_epochs = 1000

# Hypothesis
with tf.name_scope('hypothesis'):
  y_pred = tf.add(tf.multiply(X, W), b)
  # Mean Squared Error Cost Function
  cost = tf.reduce_sum(tf.pow(y_pred-Y, 2)) / (2 * n)
  # Gradient Descent Optimizer
  optimizer = tf.train.GradientDescentOptimizer(learning_rate).minimize(cost)
  # Global Variables Initializer
  init = tf.global_variables_initializer()

with tf.name_scope('sess'):
  # Starting the Tensorflow Session
  with tf.Session() as sess:
    file_writer = tf.summary.FileWriter('/tmp/tf_logs', sess.graph)
    # Initializing the Variables
    sess.run(init)
    # Iterating through all the epochs
    for epoch in range(training_epochs):
      # Feeding each data point into the optimizer using Feed Dictionary
      for (_x, _y) in zip(x, y):
        sess.run(optimizer, feed_dict = {X : _x, Y : _y})
      # Displaying the result after every 50 epochs
      if (epoch + 1) % 50 == 0:
        # Calculating the cost a every epoch
        c = sess.run(cost, feed_dict = {X : x, Y : y})
        print("Epoch", (epoch + 1), ": cost =", c, "W =", sess.run(W), "b =", sess.run(b))
    # Storing necessary values to be used outside the Session
    training_cost = sess.run(cost, feed_dict ={X: x, Y: y})
    weight = sess.run(W)
    bias = sess.run(b)

with tf.name_scope('predictions'):
  # Calculating the predictions
  predictions = weight * x + bias
  print("Training cost =", training_cost, "Weight =", weight, "bias =", bias, '\n')

# Plotting the Results
with tf.name_scope('results'):
  plt.plot(x, y, 'ro', label ='Original data')
  plt.plot(x, predictions, label ='Fitted line')
  plt.title('Linear Regression Result')
  plt.legend()
  plt.show()
