Mercurial
-
version 3.4

````
hg version

hg init

hg clone url # clone

hg paths # path name and its url

hg parents
hg tags

hg revert filename
hg up -C
hg up -C default # checkout branch default

hg pull -u

hg commit -m "Commit message"

hg push --new-branch
````

#### log
````
hg log
hg log -r 3
hg log -r 0272e0d5a517
hg log -r 1 -r 4
hg log -r 2:4modified
hg log -v -r 3    # --verbose
hg log -v -p -r 2 # --patch
hg log -b .       # log for current branch
hg log -l 5       # limit 5
````

#### status
````
hg status
hg status . # only files in this directory
````

#### diff
````
hg diff   # all files in the repository
hg diff . # only files in this directory
hg diff -r branchName:default
````

#### branch
````
hg branch               # show curr branch
hg branch newBranchName # new branch
hg branches             # all branches
````

#### conflict example
````sh
hg clone https://cn007b@bitbucket.org/cn007b/hg-hello hello

cd hello
echo 'Mercurial test named "Hello world"' > REDME.md
hg add .
hg commit -m 'Init.'
cd ../

hg clone hello my-new-hello
cd my-new-hello
echo '-' >> REDME.md
echo 'v2.' >> REDME.md
echo '' >> REDME.md
hg commit -m 'Added version to REDME.'
cd ../

cd hello
echo '' >> REDME.md
echo "It's just a test." >> REDME.md
hg commit -m 'Added description to REDME.'

hg pull ../my-new-hello/
hg merge
````

#### config
````
~/.hgrc
````
````
[diff]
git = True
[extensions]
color =
[ui]
merge = internal:merge
username = Vladimir Kovpak <cn007b@gmail.com>
[alias]
lg = log --template '{rev} {date|shortdate} | [{author}] {desc}\n'
b = branch
cm  = commit -m
````
