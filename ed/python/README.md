Python
-
<br>2.7.10

[doc](https://www.python.org/)
[packages](https://pypi.org/)

Python - dynamic, strongly typed, general-purpose, high-level,
object-oriented, multipurpose programming language.

Philosophy:
* beautiful is better than ugly
* explicit is better than implicit
* simple is better than complex
* complex is better than complicated
* readability counts

````
python -V
python3 -V

pip -V
pip list
pip show requests
pip search requests
pip install requests
pip install -r requirements.txt

pip install requests virtualenv
````

````
global myvar
nonlocal myvar

name = "Bond"
"hi {0}".format(name)
f"hi {name}"

"""
multiline string.
"""
"also multi-" \
"line string"

# list
arr = ['a', 'b']
len(arr)
del arr[1]
el in arr == true

while c != 0:
  print(c)
  c -= 1

for el in arr:
  print(el)

range(5, 10, 2) == [5, 7, 9]

myDictionary = {
  "name": "Mark",
  "active": true
}

try:
  f()
except (MyError, NotMyError) as err:
  print(err)
  raise # raise error again
except KeyError as err:
  print(err)

yield val # works like in php

import myModule
from myModule import myFunc
from myModule import *
myModule.path # gets path to module

type(myvar) # type of var
isinstance(3, int)
myvar.mro() # info about class instance
getattr(var, attr)
hasattr(var, attr)
dir(var)

globals()
locals()

__init__()
__call__()
__str__()
__next__()
__enter__() # enter context (with keyword)
__end__() # exit context (with keyword)

@staticmethod
@property # getter
@p.setter # setter
````

````
# package:
my_package
└── __init__.py # package init file
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
* range
* set

Everything is an object.
