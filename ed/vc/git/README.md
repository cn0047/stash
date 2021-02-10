GIT
-

<br>git version 2.5.0
<br>git version 2.15.0

RCS - Revision control system.

Git is a free and open source distributed version control system.

````sh
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
git config user.name "V. Kovpak"
git config user.email "cn007b@gmail.com"

# from jenkins
git ls-remote -h https://github.com/W3Ltd/zii_core HEAD

git init

git clone ssh://gitolite@repo.com:1011/repoName.git
# lightweight clone certain branch
git clone --depth=1 --branch=master git://github.com/cn007b/vo

git remote -v                      # shows remote repository
git remote show remoteRepoName     # shows all about remote repo (remote show origin)
git remote add origin ssh://gitolite@repo.com:1011/repoName.git
git remote update origin
git fetch remoteRepoName

git show commit-hash
git show HEAD^       # head parent
git show HEAD^2      # head second parent
git show HEAD~2      # first parent of first commit
HEAD^^ == HEAD~2

git commit -m 'Message'
git commit -m "PROJ-123 my comment..." # for JIRA
git commit --amend # update commit message
git commit --amend --author="V.Kovpak <cn007b@gmail.com>"

# remove untracked files
git clean -fd

git tag 1.1.1
git push origin --tags

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
git remote add upstream https://github.com/golang/go.git
git fetch upstream
git pull upstream master

# cleanup unnecessary files and optimize the local repository
git gc --aggressive

git reset x.go         # is the opposite of git add x.go
git reset --soft HEAD^ # undo a commit (delete accidentally committed file)
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
git rebase --abort

# squash - rebase 2 commits into 1:
git rebase -i HEAD~2
# for commit to squash write `s` as most left char
# :wq
git ph -f

# don't merge master but rebase to master
git rebase master
git ph -f
````

If some feature has > 1 commit,
it hard to revert whole feature, have to revert each commit,
so it make sense to squash all commits related to feature into 1 commit.

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

#### diff

````sh
git diff --cached or diff --staged # after git add shows diff
git diff branch..subBrach
git diff branch:file file # !!!
````

#### log

````sh
git log -2                # last 2 commits
git log -p                # shows commits & code in commit
git log --stat            # statiistic about changes
git log --no-meges        # log without merges
git log --follow file.txt # viewing history of deleted files !!!
git log --author=Jack

git log --date=short --no-merges --shortstat
````

#### branch

````sh
git checkout --track -b master origin/master
git checkout --track -b develop origin/develop

git branch develop --set-upstream-to origin/develop

git show branchName
git rev-parse HEAD # current commit hash
git rev_parse branch
git brnach
git rev-parse --abbrev-ref HEAD # branch name only
git branch --no-merged
git branch --merged                               # branches merged with current
git branch -v                                     # all branches Y last branch commit
git branch -vv                                    # + tracking remote branch
git push remoteRepoName pushedBranch
git push remoteRepoName pushedBranch:newBranchName
git push remoteRepoName :branch                   # delete branch from remote repo
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
    rh  = reset HEAD
    rb  = rebase -i HEAD~2
    ca  = commit -am
    cm  = commit -m
    l   = log --pretty=format:\"%h %ad | [%an] %s%d\" --graph --date=short --no-merges
    lg  = log --pretty=format:\"%h %ad | [%an] %s%d\" --graph --date=short
    f   = fetch origin
    p   = pull
    po  = pull origin
    ph  = push origin HEAD
    pt  = push origin --tags
    m   = merge --no-ff
    ml  = merge origin/live
    mm  = merge origin/master
    d   = diff --word-diff
    df  = diff
    dfc = diff --cached
    dfl = diff origin/live
    dfm = diff origin/master
    dff = diff origin/live --name-only
[core]
    excludesfile = ~/.gitignore
    editor       = vim
    ignoreCase   = false
[user]
    name = V.Kovpak
[url "git@github.com:"]
    insteadOf = https://github.com/
[url "git@bitbucket.org:"]
    insteadOf = https://bitbucket.org/
````
