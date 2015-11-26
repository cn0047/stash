Mercurial
-

````
hg clone url # clone

hg status
hg status . # only files in this directory

hg diff   # all files in the repository
hg diff . # only files in this directory

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