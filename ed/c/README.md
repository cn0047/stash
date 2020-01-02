c
-

````sh
# osx
cc -o one ed/c/examples/whatever/one.c
````

Data types:
* char
* int
* float
* double

````c
if (exp)
if (exp) st else st
if (exp) st else if (exp) st

while (exp) {}

do {} while(exp)

for (ifexp; condexp; loopexp) {}

switch (exp) {
  case cond: st break;
  default: st
}
````

````c
struct hello
{
  int index;
  float code;
};
hello h = {5, 200};
h.index
````
