GIT
-
<br>git version 2.5.0
<br>git version 2.15.0

RCS - Revision control system.

Git is a free and open source distributed version control system.

````sh
<<<<<<< HEAD
=======
>>>>>>> somehash

git config user.name "V.K."
git config user.email "cn007b@gmail.com"
git commit --amend --author="V.K. <cn007b@gmail.com>"



export GIT_TRACE_PACKET=1
export GIT_TRACE_PACK_ACCESS=1
export GIT_TRACE_PERFORMANCE=1
export GIT_TRACE_SETUP=1
export GIT_TRACE=1
export GIT_CURL_VERBOSE=1

export GIT_SSL_NO_VERIFY=1;
git config http.sslVerify 0

# in case of issue with push
git config  http.postBuffer 1048576000
git config https.postBuffer 1048576000
git config --global core.compression 0

# get config value
git config user.name
# set config value
git config user.name "V.K."
git config user.email "cn007b@gmail.com"

# from jenkins
git ls-remote -h https://github.com/W3Ltd/zii_core HEAD

git init

git clone ssh://gitolite@repo.com:1011/repoName.git
# lightweight clone certain branch
git clone --depth=1 --branch=master git://github.com/cn007b/vo
# clone usink ssh key
GIT_SSH_COMMAND='ssh -i ~/.ssh_my_key/id_rsa -o IdentitiesOnly=yes' git clone git@gitlab.com:pth/to/repo.git

# fetch - download objects and refs (branches and/or tags)
git fetch

# pull - incorporates changes from remote repo into the current branch.
git pull upstream master
git pull --ff-only
git pull --rebase
git pull --no-rebase

# pull recursevelly all subdirectories
find $(pwd) -type d -name ".git" | while read -r d; do
  repo="$(dirname "$d")"
  echo "Updating repo: $repo"
  cd "$repo"; git pull; cd -
done

git remote -v                      # shows remote repository
git remote show remoteRepoName     # shows all about remote repo (remote show origin)
git remote add origin ssh://gitolite@repo.com:1011/repoName.git
git remote update origin
git remote prune origin # when blocking reference, error: cannot lock ref 'refs/remotes/origin/brhcn'
git fetch remoteRepoName

git show commit-hash
git show HEAD^       # head parent
git show HEAD^2      # head second parent
git show HEAD~2      # first parent of first commit
HEAD^^ == HEAD~2

git commit -m 'Message'
git commit -m "PROJ-123 my comment..." # for JIRA
git commit --amend # update commit message
git commit --amend --author="V.K. <cn007b@gmail.com>"
git commit --date "2024-03-08" -m 'Message'
git commit --amend --date="2024-03-08" --no-edit

git log -2                # last 2 commits
git log -p                # shows commits & code in commit
git log -G $pattern       # find by regex pattern
git log --stat            # statiistic about changes
git log --no-meges        # log without merges
git log --follow file.txt # viewing history of deleted files !!!
git log --author=Jack
git log --date=short --no-merges --shortstat

git reflog # local repository log

git rev_parse branch
git rev-parse HEAD              # current commit hash
git rev-parse --abbrev-ref HEAD # branch name only

git diff --cached or diff --staged # after git add shows diff
git diff --stat # file names only
git diff branch..subBrach
git diff branch:file file # !!!

# remove untracked files
git clean -fd

git fetch --tags
git tag 1.1.1
git tag -d 1.1.1
git push origin --tags
git push origin v1.7

git blame -L 11,12 file

# avoid git push prompt
git remote set-url origin git+ssh://git@github.com/cn007b/my.git

# grab commit from another branch into current
git cherry-pick hash-to-commit

# develop - merge with squash
# master  - merge with fastforward
git merge
# in case of conflicts during merge:
git merge --abort

# update forked repo from source repo:
git remote add upstream https://github.com/cn007b/stash.git
git fetch upstream
git pull upstream master

# cleanup unnecessary files and optimize the local repository
git gc --aggressive

git reset x.go         # is the opposite of git add x.go
git reset --soft HEAD^ # undo a commit (delete accidentally committed file)
````

#### branch

````sh
git checkout --track -b master origin/master
git checkout --track -b develop origin/develop

git branch develop --set-upstream-to origin/develop

git show branchName
git brnach
git branch --no-merged
git branch --merged                                # branches merged with current
git branch -v                                      # all branches Y last branch commit
git branch -vv                                     # + tracking remote branch
git push remoteRepoName pushedBranch
git push remoteRepoName pushedBranch:newBranchName
git push remoteRepoName :branch                    # delete branch from remote repo
````

#### rebase

Rebasing is changing the base of your branch from one commit to another
making it appear as if you'd created your branch from a different commit.
Reason - maintain a linear project history.

Rebase is bad for bisect (it's destructing - commit is deleted)
<br>and rebase hide conflicts inside rebased commit
<br>and rebase corrupts history of comments in github and bitbucket.

And you have to: `git ph -f`

The Golden Rule of Rebasing - never use it on public branches.

````sh
git rebase -m # use merging strategies
git rebase -r # rebase merges

git rebase --abort

# squash - rebase 2 commits into 1:
git rebase -i HEAD~2
# for commit to squash write `s` as most left char
# :wq
git ph -f

# rebase to master (not merge master)
git rebase master
git ph -f

# undo rebase
git reset --hard ORIG_HEAD
# show ORIG_HEAD value
git log -1 ORIG_HEAD
````

If some feature has > 1 commit,
it hard to revert whole feature, have to revert each commit,
so it make sense to squash all commits related to feature into 1 commit.
Drawback here: will be only 1 commit, with last implementation,
and without all intermediate implementations.

#### bisect

````sh
git bisect start
git bisect bad       # tell that current situation is bad
git bisect good v1.0 # good commit tag
                     # check whether script working
git bisect bad       # if scripn not working
                     # check whether script working again
git bisect good      # if scripn is working
                     # ...
git bisect good
                     # ...
git bisect reset
````

#### submodule

````sh
cat .gitmodules            # get submodules
git submodule              # get submodules list
git submodule--helper list # get submodules list

git submodule init
git submodule update
git submodule status

git submodule sync --recursive
git submodule update --init --recursive
git submodule update --remote --recursive
````

#### config

.gitignore
````sh
.idea/
.DS_Store
````

````sh
/etc/gitconfig # system
~/.gitconfig   # user
.git/config    # project
````
````sh
[color]
    ui     = auto
    diff   = true
    status = auto
    branch = auto
[alias]
    b   = branch
    br  = branch
    ch  = checkout
    cb  = !sh -c 'git checkout --track -b $1 origin/$1' -
    s   = status -sb
    st  = status
    ssh = stash show
    sl  = stash list
    ss  = stash save --include-untracked
    sa  = stash apply
    sp  = stash pop
    su  = submodule update --remote --recursive
    rh  = reset HEAD
    rb  = !sh -c 'git rebase -i HEAD~$1' -
    ca  = commit -am
    cm  = commit -m
    l   = log --pretty=format:\"%h %ad | [%an] %s%d\" --graph --date=short --no-merges
    lg  = log --pretty=format:\"%h %ad | [%an] %s%d\" --graph --date=short
    f   = fetch origin
    p   = pull --ff-only
    pr  = pull --rebase
    po  = pull origin
    ph  = push origin HEAD
    pt  = push origin --tags
    m   = merge --no-ff
    ml  = merge origin/live
    mm  = merge origin/master
    d   = diff --word-diff
    df  = diff
    dfc = diff --cached
    dff = diff --name-only
    dfl = diff origin/live
    dfm = diff origin/master
[core]
    excludesfile = ~/.gitignore
    editor       = vim
    ignoreCase   = false
[user]
    name = V.K.
    email = cn007b@gmail.com
# [url "git@github.com:"]
#     insteadOf = https://github.com/
# [url "git@bitbucket.org:"]
#     insteadOf = https://bitbucket.org/

````
