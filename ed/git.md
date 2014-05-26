GIT
-

````
init
clone repo

diff --cached or diff --staged # after git add shows diff
diff branch..subBrach
diff branch:file file

log -2         # last 2 commits
log -p         # shows commits & code in commit
log --stat     # statiistic about changes
log --no-meges # log without merges

remote -v                      # shows remote repository
remote show remoteRepoName     # shows all about remote repo (remote show origin)
fetch remoteRepoName


brnach --merged    # branches merged with current
brnach --no-merged

show commit-hash
show HEAD^       # head parent
show HEAD^2      # head second parent
show HEAD~2      # first parent of first commit
HEAD^^ == HEAD~2

commit -m 'Message'

blame -L 11,12 file

/etc/gitconfig # system
~/.gitconfig   # user
````
.git/config    # project

####branch
````
show branch
rev_parse branch
branch -v                                     # all branches Y last branch commit
push remoteRepoName pusedBranch
push remoteRepoName pusedBranch:newBranchName
push remoteRepoName :newBranch                # delete branch from remote repo
````