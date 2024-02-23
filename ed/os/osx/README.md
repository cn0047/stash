osx (darwin)
-

Since 1984.

The default shell in Mac OS X is the Bourne-again shell.
Chip Intel -> x64.
Chip M1, M2 -> ARM.

`jamf` [tool to manage apple device](https://www.jamf.com/).
`quicksilver` tool for quick app's launch ([default config](https://monosnap.com/file/UH48Ulyzm6IzKxhlWqBuozkAOgciUQ)).
`monosnap` tool for screenshots.
`GIPHY` tool to convert video into gif.

````sh
# .zshrc

. /Users/k/.bash_profile
````

````sh
# .bash_profile

setopt INTERACTIVE_COMMENTS

export PATH=/Users/k/go/bin:$PATH

export SSLKEYLOGFILE=/tmp/https.log
````

````sh
man grep

# open fonts directory
open /Applications/Utilities/Terminal.app/Contents/Resources/Fonts/

/usr/bin/open -a "/Applications/Google Chrome.app" 'http://google.com/'
# or
open 'http://google.com/'

# cpu cores count
sysctl -n hw.ncpu

$HOME/Library/Caches # cache dir

ls /usr/local/bin
````

Colorize CLI:

````sh
brew install grc

# .bash_profile
export CLICOLOR=xterm-color
export CLICOLOR=1
alias grep='grep --color=always'
````

## HotKeys

````sh
Option = Alt

Command-Shift-G         # goto file in finder (open input for file path)

Command-Control-O       # insert mode
Command-Control-Space   # emoji

Control-Shift-Power     # Lock screen
Command-Option-Power    # Put MacBook to sleep
Command-Option-Esc      # Force Quit

Command-Shift-.         # Show hidden files
Command-H:              # Hide
Command-M:              # Minimize
Command-Option-W:       # Close all windows at once

Option-Del              # Del word
Control-K               # Del row
Command-Del             # Del to home
Control-O               # Insert a new line after

Fn–U|D                  # Move page U/D
Fn–L|R                  # Scroll B/E document.
Command–U|D             # Move B/en document.
Command–L|R             # Move B/E current line.
Option–L|R              # Move B/E previous word.

Shift–Option–U|D|L|R    # Extended text selection
Command-Option          # Select text column in terminal

Command-Shift-3         # Screenshot ALL SCREEN in file on desktop
Command-Control-Shift-3 # Screenshot ALL SCREEN in buffer
Command-Shift-4         # Screenshot (SELECTED SHAPE) in file on desktop ✅
Command-Control-Shift-4 # Screenshot (SELECTED SHAPE) in buffer 💡

F1  # Bright down
F2  # Bright up
F11 # Sound down
F12 # Sound up

/ # go to path in finder
````

## Brew

````bash
brew ls
brew services list

brew info xdebug
brew --prefix xdebug # show install path
brew search xdebug
brew reinstall xdebug
brew list --versions xdebug # shows installed versions
brew remove xdebug

brew install grc # colors in terminal
brew install bash-completion
brew install colordiff
brew install jq
brew install yq
brew install md5sha1sum # md5sum
brew install qcachegrind # KCacheGrind
brew install inetutils # ftp telnet

brew install nmap
brew install sshuttle
brew install awscli
brew install mysql
brew install nginx
brew install ffmpeg
brew install jmeter
brew install mercurial
brew install postgresql
brew install redis
brew install terraform
brew install mongodb
brew install openssl@1.1 # install specific version

brew update && brew cleanup
brew upgrade jq

brew services start postgresql
````
