GIT
-

````
init
clone repo

diff --cached or diff --staged # after git add shows diff
diff branch..subBrach

log -2         # last 2 commits
log -p         # shows commits & code in commit
log --stat     # statiistic about changes
log --no-meges # log without merges

remote -v                      # shows remote repository
remote show remoteReponame     # shows all about remote repo (remote show origin)
fetch remoteReponame
push remoteReponame pusedBranch
push remoteReponame pusedBranch:newBranch
push remoteReponame :newBranch # delete branch from remote repo

branch -v          # all branches Y last branch commit
brnach --merged    # branches merged with current
brnach --no-merged

show commit-hash
show branch
rev_parse branch
show HEAD^       # head parent
show HEAD^2      # head second parent
show HEAD~2      # first parent of first commit
HEAD^^ == HEAD~2

commit -m 'Message'

blame -L 11,12 file

/etc/gitconfig # system
~/.gitconfig   # user
.git/config    # project
````