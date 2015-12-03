Mercurial
-

````
hg version

hg init

hg clone url # clone

hg log
hg log -r 3
hg log -r 0272e0d5a517
hg log -r 1 -r 4
hg log -r 2:4
hg log -v -r 3 # --verbose
hg log -v -p -r 2 # --patch

hg parents

hg status
hg status . # only files in this directory

hg diff   # all files in the repository
hg diff . # only files in this directory
hg diff -r branchName:default

hg revert filename
hg up -C
hg up -C default # checkout branch default

hg pull -u

hg paths # path name and its url

hg branch               # show curr branch
hg branch newBranchName # new branch
hg branches -R

hg commit -m "Commit message"

hg push --new-branch
````

####config
````
~/.hgrc
````
````
[diff]
git = True
[extensions]
color =
[ui]
username = Vladimir Kovpak <cn007b@gmail.com>
````

page 33
