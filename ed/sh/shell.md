shell
-

````sh
sudo bash -c 'echo "APP_ENVIRONMENT=prod" > /etc/environment'
sudo sh -c 'echo "APP_ENVIRONMENT=prod" > /etc/environment'
sh -c 'echo 200'

last # to see recent activity

~/.bash_history

# history:
# ⬆, ⬇ # keys to navigat through history
# Ctrl-p, Ctrl-n
# Ctrl-r # search in history
!207 # run from history

sudo !! # redo last command but as root

# don\'t add command to history (note the leading space)
 ls -l

prompt \u@\h [\d]>
````

````sh
cd -       # go to previous dir
pushd path # remember path (save it at stack)
popd       # got to pushed path (and delete it from stack)
````

#### hotkeys

````
Ctrl-k # delete rest of line
Ctrl-u # delete to start of line
Ctrl-w # delete word
````
