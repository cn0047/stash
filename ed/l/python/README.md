Python
-
<br>3.6
<br>2.7.10

[doc](https://www.python.org/)
[packages](https://pypi.org/)
[flake8](https://flake8.pycqa.org/)

Python - dynamic, ~~strongly typed~~, general-purpose, high-level,
object-oriented, multipurpose programming language.

Cython - superset of the Python, gives C-like performance
with code that is written mostly in Python.

Easier to ask for forgiveness than permission (EAFP) - `try` and `except` statements.
Look before you leap (LBYL) - tests for pre-conditions before making calls,
characterized by the presence of many `if` statements.

````sh
python -V
python3 -V

WORK_HOME=/path
PROJECT_HOME=/path
PYTHONUNBUFFERED=1 # output sends straight to terminal without buffering

PYTHONPATH=$HOME/dirWithScripts/:$PYTHONPATH
# PATH=$PATH:~/ed/l/python/examples/pkg python3 pkg

pylint x.py # lint file

python3 # REPL

cd ed/l/python/examples/whatever
python3 -m unittest
python -c 'print(200)'
````

````py
import platform; print(platform.python_version())

help("math")

global myvar
nonlocal myvar # don't use var from enclosing scope

print(s, end='', flush=True)

print(5 // 2) # floor division

# partition
o, _, t = "v1:v2".partition(':')
print(o, t) # v1 v2

# conditional statement
if cond:
  res = value1
else:
  res = value2
# conditional expression
res = value1 if cond else value2

while c != 0:
  print(c)
  c -= 1
while cond:
  # smth
else:
  # smth

for el in arr:
  print(el)

def trace(f, *args, **kwargs):
  return f(*args, **kwargs)

yield val # works like in php

globals()["foo"] = "bar" # global namespace
locals()["foo"] = "bar"
locals()
any([false, true]) # true
all([false, true]) # false
type(myvar) # type of var
isinstance(3, int)
issubclass()
hasattr(var, attr)
getattr(obj, attr)
setattr(obj, attr, val)
vars(self)['color'] = 'red'
delattr(obj, attr)
dir(var) # introspecting objects
repr(var) # representation, for DBG
str(var)

if __name__ == '__main__'
__doc__
__path__
__file__
__closure__
__bases__ # base classes
exception.__cause__
exception.__traceback__

obj.mro() # Method Resolution Order, info about class instance
obj.__mro__

obj.__dict__ # dictionary with obj's properties
vars(obj)

__new__(cls, *args, **kwargs)
__prepare__
__init__
__call__()
__get__()
__set__()
__delete__()
__str__(self)
__repr__(self)
__format__(self)
__iter__(self)
__next__()
__getitem__(self, index)
__getattr__(self, name)
__setattr__(self, name, val)
__eq__(self, rhs)
__hash__(self)
__ne__(self, rhs)
__len__(self)
__enter__() # enter context (with keyword)
__end__() # exit context (with keyword)
__exit__()

@classmethod # bound to a class not to instance, like staticmethod.
@property # getter
@propertyName.setter # setter

"""
Parameters
----------
conf: MyConfigType
    Main config object.
"""

"""
:param conf: Main config object.
"""
````

#### Data types

* int
* float (64-bit)
* bool
* none

* str
* bytes
* list
* dict
* tuple
* set
* range

Everything is an object.

````py
from decimal import Decimal
5/2  # 2.5
5//2 # 2
5%2  # 1
Decimal('0.8') - Decimal('0.7') # Decimal('0.1')

-7 % 3                       # 2
Decimal('-7') % Decimal('3') # Decimal('-1')

# round func for Decimal
round(2.675, 2)            # 2.67
round(Decimal('2.675'), 2) # Decimal('2.68')

# helpful stuff:
from collections import Counter, defaultdict, namedtuple, deque
````

**str** strings immutable.

````py
str = "multiline" \
      " string!"

r'ok' # raw string
f'hello, {name}\n' # format with vars
'{a}->{b}'.format(a='foo', b='bar') # foo->bar

name = "Bond"
print("hi {0}".format(name)) # hi Bond
print(f"hi {name}") # hi Bond
index = 7; print('frame_%05d.jpg' % index)
````

**bytes**.

````py
b'data' # bytes
````

**list**.

````py
arr = ['a', 'b', True] # list
len(arr)
del arr[1]
'b' in arr # True

data.append(7) # add element to list
data[1:]       # remove 1st el
data[:-1]      # remove last el
data[-1:]      # get last el
data[::-1]     # reverse list
data[:]        # copy
````

**dict**.

````py
md = {'foo': 1, "name": "Mark", "active": true} # dict - dictionary
````

**tuple** immutable sequence of arbitrary objects.

````py
t = ("f", 2, False) # tuple
t[0] # "f"

t = (1,2,3)
print(t) # (1, 2, 3)
print(*t) # 1 2 3
````

**set** unordered collection on unique elements.

````py
s = {1, 2, True} # set
s.pop()
s.add("orange")
2 in s # True
"banana" in s # False
s.remove("banana")
len(s)

s = set()
s |= {1} # add into set
s |= {2}
s -= {1} # delete from set
````

**range**.

````py
range(5, 10, 2) == [5, 7, 9]
````

#### Features

**comprehensions**.

**generator**.

**trampoline** using generators to substitute recursion.

**generator expression** like comprehensions but generator.

**module** - usually just a single file with code.

````py
import myModule             # module
from myModule import myFunc
from myModule import *
from .dir import *          # relative import
from . import common        # from common dir in same dir
from ..dir import *         # from parent dir
myModule.path               # gets path to module
````

**package** - module which contains other modules.

````py
# add pkg
sys.path.append('pkg_dir')
import pkg
import ..pkg # relative import
# package dir:
my_package
└── __init__.py # package init file

__main__.py # must be in executable dir
````

**eggs** - is logical structure embodying the release of a specific version
of project, comprising its code, resources, and metadata.

**wheel** - new standard of distribution (intended to replace eggs).

**exception** ([hierarchy](https://docs.python.org/3/library/exceptions.html#exception-hierarchy)).

````py
ValueError

try:
  f()
except (MyError, NotMyError) as err:
  print(err)
  raise # raise error again
except KeyError as err:
  print(err)
except KeyError2 as err:
  pass
except KeyError3 as err:
  raise
except KeyError4 as err:
  raise IndexError()
finally:
  # ...

try:
  # ...
except OSError:
  # ...
else:
  # ...
````

**lambda**.

````py
def first_name(name):
  return name.split()[0]

lambda name: name.split()[-1]
````

**metaclass**.

Default metaclass for all classes is `type`.

**abstract class**.

**unittest**.

stub - provides predefined answers.

fake - behaves like something real (has realistic implementation inside)
but with predefined results (file, db (real db but in memory one), webserver, etc.).

dummy - usually none.

mock - build object with predefined behavior.

spy - checks whether method was called.

##### Concurrency

**threading** - run threads in the same memory space.
`threading.Lock, threading.Semaphore`.

**multiprocessing** - run processes in separate memory space.
`multiprocessing.Pool`.

**queue** - `from queue import Queue, Pipe, JoinableQueue`.

**ThreadPoolExecutor** - `from concurrent.futures import ThreadPoolExecutor`.

**asyncio** `asyncio.get_event_loop().run_until_complete(mid.run_worker())`.

**coroutine** - special type of function that deliberately yield control to caller,
but does not end its context in the process, instead maintaining it in an idle state.
difference with generator - coroutine can accept arguments after it's been initially called,
whereas a generator can't.

**coroutine function** function with `async` prefix.

**coroutine object** - what function declared with an async keyword returned,
`CoroutineObject = CoroutineFunction()`.
