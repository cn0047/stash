Jupyter Notebook
-

[docs](https://jupyter.org/documentation)

````sh
# run
jupyter notebook
````

````sh
!pwd

print? # prints help info

%lsmagic
%history
%%writefile? # prints help info

files = !ls
print(files)

!cat {files}

In[11]
Out[14]

_ # retruns last result or In/Out if clicked

return 'color: red'

Image(url='...')
````

````sh
%%writefile test.txt

Test file.
````

#### Keys

````sh
m            # markdown
y            # code

shift+l      # toggle line numbers
shift+tab    # info

ctrl+enter   # run cell
shift+enter  # run cell and select next cell

x            # cut
c            # copy
v            # paste
shift+v      # paste above
dd           # delete
a            # insert above
b            # insert belove

shift+m      # merge
ctrl+shift+- # split

command+[    # indentation
command+]
````
