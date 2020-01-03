Less
-

````sh
r        # repaint the screen
v        # edit the current file with $VISUAL or $EDITOR
:f       # print current file name
!command # execute the shell command
=        # prints info about the file

z # scroll forward
b # scroll backward
d # scroll forward one half of the screen
u # scroll backward one half of the screen

g   # go to line n or to line 1
G   # go to line n or to last line
:ng # go to line with number n
{([ # find close bracket } ) ]
})] # find open bracket { ( [

RightArrow         # right one half screen width
LeftArrow          # left  one half screen width
Control-RightArrow # right to last column displayed
Control-LeftArrow  # left  to first column

/pattern # search forward
?pattern # search backward
&pattern # display only lines which match pattern
n        # repeat previous search
N        # repeat previous search in reverse direction
ESC-u    # toggle search highlighting

m<letter> # mark the current position with <letter>
'<letter> # go to a previously marked position
''        # go to the previous position
````

Options:

````sh
-F      # exit if the entire file can be displayed on the first screen
+F      # follow
-X      # like clearing the screen
-N      # line numbers
-n      # no line numbers
-R      # raw control chars
-r      # raw control chars
-S      # truncate long lines
-i      # ignore case
-z [N]  # set size of window; -z-4
-# [N]  # horizontal scroll amount (0 = one half screen width)
-xn,... # sets tab stops; -x9,17 will set tabs at positions 9, 17, 25, 33, etc.

less -FX +F /var/log/syslog
````
