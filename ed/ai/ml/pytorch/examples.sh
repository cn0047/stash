docker run -ti --rm -v $PWD:/gh -w /gh cn007b/python sh -c '
  pip3 install torch numpy matplotlib
  python3 /gh/ed/ai/ml/pytorch/examples/demo.py
'
