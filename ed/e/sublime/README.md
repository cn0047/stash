Sublime Text
-
<br>3

[docs](https://www.sublimetext.com/docs/index.html)

### OSX

````sh
u=k
u=kovpakvolodymyr

# subl in terminal
sudo ln -s "/Applications/Sublime\ Text.app/Contents/SharedSupport/bin/subl" /usr/local/bin/

# delete project from list
vim /Users/$USER/Library/Application\ Support/Sublime\ Text\ 3/Local/Auto\ Save\ Session.sublime_session
vim /Users/$USER/Library/Application\ Support/Sublime\ Text\ 3/Local/Backup\ Auto\ Save\ Session.sublime_session
vim /Users/$USER/Library/Application\ Support/Sublime\ Text\ 3/Local/Backup\ Session.sublime_session
vim /Users/$USER/Library/Application\ Support/Sublime\ Text\ 3/Local/Session.sublime_session

# packages list
subl /Users/$USER/Library/Application\ Support/Sublime\ Text\ 3/Packages/User/Package\ Control.sublime-settings

# packages dir
ls /Users/$USER/Library/Application\ Support/Sublime\ Text\ 3/Packages
subl /Users/$USER/Library/Application\ Support/Sublime\ Text\ 3/Packages/User/SublimeLinter/Solarized\ V\ \(Light\)\ \(SL\).tmTheme
subl /Users/$USER/Library/Application\ Support/Sublime\ Text\ 3/Packages/User/SublimeLinter/my.tmTheme

# keys
ALT+CMD+O # Insert key
````

Quick File Open settings:
````js
{
  "files": [
    "/Users/k/web/..."
  ]
}
````

Search:
````sh
ed/l/nodejs,-*/node_modules/*
ed/l/go,-*/vendor/*
````
