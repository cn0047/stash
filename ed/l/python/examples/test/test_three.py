import unittest

from unittest.mock import MagicMock

class Foo:
  def foo() -> str:
    return 'foo'

class TestThree(unittest.TestCase):

    def test_a(self):
        self.assertEqual('foo'.upper(), 'FOO')

    def test_b(self):
        self.assertEqual('bar'.upper(), 'BAR')

    def test_c(self):
      m = MagicMock(Foo)
      m.foo.return_value = 'bar'

      self.assertEqual('bar', m.foo())


if __name__ == '__main__':
    unittest.main()
