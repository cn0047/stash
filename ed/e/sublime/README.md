Sublime Text
-
3

### OSX

````sh
u=k
u=kovpakvolodymyr

# subl in terminal
sudo ln -s "/Applications/Sublime Text.app//Contents/SharedSupport/bin/subl" /usr/local/bin/

# packages list
subl /Users/$u/Library/Application\ Support/Sublime\ Text\ 3/Packages/User/Package\ Control.sublime-settings

# packages dir
ls /Users/$u/Library/Application\ Support/Sublime\ Text\ 3/Packages
subl /Users/$u/Library/Application\ Support/Sublime\ Text\ 3/Packages/User/SublimeLinter/Solarized\ V\ \(Light\)\ \(SL\).tmTheme
subl /Users/$u/Library/Application\ Support/Sublime\ Text\ 3/Packages/User/SublimeLinter/my.tmTheme

# keys
ALT+CMD+O # Insert key
````

Quick File Open settings:
````
{
  "files": [
    "/Users/k/web/..."
  ]
}
````
