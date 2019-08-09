shell
-

````sh
sudo bash -c 'echo "APP_ENVIRONMENT=prod" > /etc/environment'
sudo sh -c 'echo "APP_ENVIRONMENT=prod" > /etc/environment'
sh -c 'echo 200'

env # prints all ENV vars

last # to see recent activity

~/.bash_history

# history:
# ⬆, ⬇ # keys to navigat through history
# Ctrl-p, Ctrl-n
# Ctrl-r # search in history
!207 # run from history

sudo !! # redo last command but as root

alias # show all aliases
alias la='ls -A'

type la # show alias
type cp
type /Users/k/web/kovpak/gh/ed/bash/examples/hw.sh

which git # /usr/local/bin/git

ldd /bin/ls # shows used *.so libraries by ls command

sleep 6 &
jobs # will shows scrips in background

fg # send command to foreground
bg # send command to background

# overwrite file
echo OK >| f

# don\'t add command to history (note the leading space)
 ls -l

echo There are ${#BASH_ALIASES[*]} aliases defined.

prompt \u@\h [\d]>
````

````sh
cd -       # go to previous dir
pushd path # remember path (save it at stack)
popd       # got to pushed path (and delete it from stack)
````
