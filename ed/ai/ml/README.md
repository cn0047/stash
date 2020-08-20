Machine Learning
-

ML - perform task without using explicit instructions (relying on patterns).

Deep Learning - area of ML whose goal is to learn complex functions
using special neural network architectures that are "deep" (consist of many layers).

Classic Problems:
* Classification (supervised learning) (spam).
* Regression (supervised learning) (predicts a number).
* Clustering (unsupervised learning).
* Density.
* Sampling.
* Reinforcement Learning.
* Rule extraction.

Model - after training the system (detecting patterns in the data)
a model is created to make predictions.

Dataset must be divided into: train, validation, test.

Serving ML Model - make trained model available to serve prediction requests.

Precision and recall - example: given 8 dogs and 4 cats,
ML found 8 dogs (real: 5 dogs (true positives) & 3 cats (false positives)),
so precision = 5/8 while its recall is 5/12.

**SUPERVISED learning** - when we have info, like cat is on picture.
(email spam/not spam, cancer tumor).
Needs training data.
````
Supervised ML Algorithml -> Train -> Trained Model
prepare data -> select algorithml -> train model -> test model
data = 70% training, 30% testing
````

**UNSUPERVISED learning** - no info.
(database of custome data (market segments), news.google.com articles).
Identify clusters of like data.
````
UML Algorithml -> Classify -> Model
````

**reinforcement learning**.

## Neural network

Neural network - class of machine learning algorithm used to model complex patterns in datasets
using multiple hidden layers and non-linear activation functions.

Artificial Neural Networks (ANN) - systems "learn" to perform tasks
by considering examples, generally without being programmed with task-specific rules.

Neuron - takes a group of weighted inputs, applies an activation function (mathematical)
and returns an output.

Synapse - like road in a neural network, connect input to neuron, neuron to neuron, and neuron to output.

Bias - additional constant attached to neuron and added to the weighted input
before the activation function is applied.

Input Layer - holds the data your model will train on.
Hidden Layer - sits between the input and output layers and applies an activation function.
Output Layer - the final layer in a network.

Affine transformation - linear function.

Activation function - live inside neural network layers
and modify the data they receive before passing it to the next layer.
Common activation functions: ReLU, logit, tanh, step.

Loss (cost) function - wrapper around model’s predict function that tells us
"how good" the model is at making predictions for a given set of parameters.

Gradient accumulation - mechanism to split the batch of samples for training a neural network
into several mini-batches of samples that will be run sequentially.

## Tensor

Tensor - N-dimensional array of primitive values (data sent between nodes in data graph (edges)).
Scalar - 0-D Tensor.
Vector - 1-D tensor.
Matrix - 2-D tensor.
N-D Matrix - N-D tensor.
Rank - number of dimensions in tensor.
Shape - number of elements in each dimension.
