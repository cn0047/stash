RegEx
-

[tool](https://regex101.com/)
[email example](https://github.com/go-playground/validator/blob/0a9b75fbfdb5730dfff911981b428046167f33e6/regexes.go#L18)

IPv6: https://regex101.com/r/jpVRLB/1
IPv4: https://regex101.com/r/d88q1P/1
IPv4 address range: https://regex101.com/r/O4t93o/2/

| syntax          | description |
|-----------------|-------------|
|^                | start of string |
|\A               | absolute start of string |
|$                | end of string |
|\Z               | absolute end of string |
|.                | any single character |
|\                | escape |
|0                | 0 or more |
|+                | repetitive match, 1 or more |
|?                | optional match, 0 or 1 |
|*                | optional match, 0 or many |
|`*?, +?, ??`     | the *, +, and ? qualifiers are all greedy, `?` helps avoid it |
|+?               | repetitive match, 1 or more, prefer fewer |
|??               | optional match, 0 or 1, prefer zero |
|*?               | optional match, 0 or many, prefer fewer |
|{3}              | exactly 3 |
|{3,}             | 3 or more |
|{3,5}            | 3 or 4 or 5 |
|[abc]            | range (a or b or c) |
|[^abc]           | not a not b not c |
|[a-q]            | letter between a and q |
|[A-Q]            | upper case letter between A and Q |
|[0-7]            | digit between 0 and 7 |
|(...)            | group |
|\0               | null |
|\t               | tab |
|\n               | new line |
|\v               | vertical tab |
|\f               | new page |
|\r               | return |
|\c               | control character |
|\s               | white space [\s\t\r\n\f] |
|\S               | not white space |
|\d               | digit |
|\D               | not digit |
|\w               | word (a-z, A-Z, 0-9, _) |
|\W               | not word |
|\x               | hexadecimal digit |
|\O               | octal digit |
|\b               | word boundary |
|\B               | not word boundary |
|\Q               | begin quote. Any characters between \Q and \E including metacharacters (e.g. '+' or '.') will be treated as literals |
|\E               | end quote |
|\<               | start of word |
|\>               | end of word |
|?=               | positive lookahead assertion # `/foo(?=bar)/` -> foobar, but `bar` not included in result |
|?!               | negative lookahead # `/Java(?!Script)/` -> Java, not JavaScript |
|?<=              | positive lookbehind assetion # `/(?<=td)ms/` -> `tdms`, but `td` not included in result |
|?<!              | negative lookbehind # `/(?<!bar)foo/` -> not `barfoo`, but `bazfoo`, and `baz` not included in result |
|?>               | once-only sebexpression |
|?#               | comment |
|(?:...)          | pasive group, just group, cannot be linked by \1 # /(?:h.*)(f.*)/ -> htfm, group 1 = fm |
|\1\2             | back reference to matches that was in (...) |
|$1               | |
|$2               | |
|$`               | before matched string |
|$'               | after matched string |
|$+               | last matched string |
|((?<=^)\|(?<=,)) | Begin match with start of string (^) or comma. The `?<=` makes sure this is will not be replaced in preg_replace. |

|   |   |
|---|---|
| i | case insensitive |
| g | global search |
| m | multiple lines |
| s | treat string as single line |
| x | allow comments and white space in pattern |
| e | evaluate replacement |
| u | upgready|unicode |

````
/(['"])[^'"]*\1/

(aa){1,2}?                   # exactly 2 aa, not 4 aa (2 aa & 2 aa)
<em>(.+?)</em>               # exactly what in em tag
(['"])\w+\1                  # 's' or "d" but not 'mess". (back reference)
(?P<quote>['"])\w+(?P=quote) # named back reference

\*\w+\*          # *keep* or *secret* https://regex101.com/r/yO6wI6/1
(?<=\*)\w+(?=\*) # same as prev pattern
(?<!\*)\w+(?!\*) # https://regex101.com/r/hC8pG6/1

^(?!.*badword).*string.* # contains string but doesn't contain badword

/?(condition)true-pattern\|false-pattern/               # condition
/?(condition)true-pattern/                              # condition
(?{GROUP_MATCHED?}REPLACEMENT_IF_YES:REPLACEMENT_IF_NO} # conditional replacement
(?|(regex1)|(regex2))                                   # branch reset group
(\2two|(one))+                                          # forward references # /(\2two|(one))+/ -> oneonetwo
````

````sh
grep '$DC'      # environment var
grep -Pz 'r\nw' # match new line pattern

echo "Nate or nate"    | grep -P '(?<!N)a'
echo '{"token":"123"}' | grep -Po '(?<="token":")[^"]*'
echo '"total":127,'    | grep -P '(?!"total":)[0-9]*'
````

````sh
# Latitude and Longitude
^\([+-]?([1-8]?\d(\.\d+)?|90(\.0+)?), [+-]?((1[0-7]|[1-9])?\d(\.\d+)?|180(\.0+)?)\)$

# ok
(75, 180)
(+90.0, -147.45)
(77.11112223331, 149.99999999)
(+90, +180)
(90, 180)
(-90.00000, -180.0000)

# error
(75, 280)
(+190.0, -147.45)
(77.11112223331, 249.99999999)
(+90, +180.2)
(90., 180.)
(-090.00000, -180.0000)
````

````php
$s = '19 20 22 -1 -3-4-10 -7 2 10';
var_export(preg_split('/\s+|(?<!\s)(?=-)/', $s));
/*
array (
  0 => '19',
  1 => '20',
  2 => '22',
  3 => '-1',
  4 => '-3',
  5 => '-4',
  6 => '-10',
  7 => '-7',
  8 => '2',
  9 => '10',
)
*/
````

Named back reference.

````php
preg_match('/(?P<one>f..)/', 'foo bar', $matches);
var_export($matches);
/*
array (
  0 => 'foo',
  'one' => 'foo',
  1 => 'foo',
)
*/
````
