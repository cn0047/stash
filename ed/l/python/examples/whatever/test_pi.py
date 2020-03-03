import unittest

from testpi import pi

class PiTest(unittest.TestCase):

    def test_pi(self):
        v = pi()
        print(v)

    @unittest.skip("WIP")
    def test_pi_2(self):
        pass
