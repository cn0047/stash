from dataclasses import dataclass


@dataclass
class Foo:

    def __post_init__(self):
        print('postinit')


Foo()
