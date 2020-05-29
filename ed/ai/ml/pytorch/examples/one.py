import torch


def f1():
  print([
    '[f1]',
    torch.cuda.is_available(),
    torch.get_default_dtype(),
  ], end='\n\n')


def f2():
  t = torch.tensor([10., 20.])
  print(t)
  t = torch.rand(2, 3)
  print(t)
  print(t.shape)
  t = torch.rand(2, 3, requires_grad=True)
  print(t)


f1()
f2()
