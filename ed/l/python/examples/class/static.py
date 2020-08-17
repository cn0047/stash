class St:

  val = 5;


  def f1():
    St.val += 2;
    print("[f1] ", St.val)


  @staticmethod
  def f2():
    St.val += 2;
    print("[f2] ", St.val)


  @classmethod
  def f3(cls):
    cls.val += 2;
    print("[f3] ", cls.val)


St.f1()
St.f2()
St.f3()
