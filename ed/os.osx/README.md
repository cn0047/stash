osx
-

darwin

The default shell in Mac OS X is the Bourne-again shell.

`quicksilver` tool for quick app's launch ([default config](https://monosnap.com/file/UH48Ulyzm6IzKxhlWqBuozkAOgciUQ)).
`monosnap` tool for screenshots.
`GIPHY` tool to convert video into gif.

````sh
man grep

# open fonts directory
open /Applications/Utilities/Terminal.app/Contents/Resources/Fonts/

/usr/bin/open -a "/Applications/Google Chrome.app" 'http://google.com/'
# or
open 'http://google.com/'

# cpu cores count
sysctl -n hw.ncpu
````

## HotKeys

````sh
Option = Alt

Control-Shift-Power  # Lock screen
Command-Option-Power # Put MacBook to sleep
Command-Option-Esc   # Force Quit

Command-Shift-.   # Show hidden files
Command-h:        # Hide
Command-m:        # Minimize
Command-Option-w: # Close all windows at once

Option-Del  # Del word
Control-K   # Del row
Command-Del # Del to home
Control-O   # Insert a new line after

Fn–u|d      # Move page U/D
Fn–l|r      # Scroll B/E document.
Command–u|d # Move B/en document.
Command–l|r # Move B/E current line.
Option–l|r  # Move B/E previous word.

Shift–Option–U|D|L|R # Extended text selection

Command-Shift-3         # Screenshot ALL SCREEN in file on desktop
Command-Control-Shift-3 # Screenshot ALL SCREEN in buffer
Command-Shift-4         # Screenshot (SELECTED SHAPE) in file on desktop ✅
Command-Control-Shift-4 # Screenshot (SELECTED SHAPE) in buffer 💡

F1 # Bright down
F2 # Bright up
F11 # Sound down
F12 # Sound up
````

## Brew

````bash
brew ls
brew services list

brew search xdebug

brew install awscli
brew install bash-completion
brew install grc # colors in terminal
brew install jmeter
brew install jq
brew install md5sha1sum # md5sum
brew install mongodb
brew install mysql
brew install nginx
brew install nmap
brew install postgresql
brew install qcachegrind
brew install sshuttle
brew install telnet
brew install terraform

brew remove nginx
brew update && brew cleanup

brew services start postgresql
````
