Mercurial
-
version 3.4

````sh
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

#### stash

````sh
hg diff > ~/d.diff
hg revert -aC
hg import --no-commit ~/d.diff
````

#### log

````sh
hg log
hg log -r 3
hg log -r 0272e0d5a517
hg log -r 1 -r 4
hg log -r 2:4modified
hg log -v -r 3    # --verbose
hg log -v -p -r 2 # --patch
hg log -b .       # log for current branch
hg log -l 5       # limit 5

hg cat -r c6f5e97edfdf PATH_TO_FILE # see file from particular rev
````

#### status

````sh
hg status
hg status . # only files in this directory

# delete not tracked files
rm `hg st|awk '{if($1=="?") print $2}'`
````

#### diff

````sh
hg diff   # all files in the repository
hg diff . # only files in this directory
hg diff -r branchName:default
````

#### branch

````sh
hg branch               # show curr branch
hg branch newBranchName # new branch
hg branches             # all branches

hg identify --id --rev BRANCH_NAME # get revision by branch name
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

````sh
~/.hgrc
````
````
[diff]
git = True
[extensions]
color =
[ui]
username = Vladimir Kovpak <cn007b@gmail.com>
merge = internal:merge
ignore = ~/.hgignore
[alias]
lg = log -l 50 --template '{rev} {date|shortdate} | [{author}] {desc} \n'
b = branch
cm  = commit -m
r = revert -C
d = diff --color=always
df = diff
apply = import --no-commit
````
