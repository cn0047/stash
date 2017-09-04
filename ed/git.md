GIT
-
git version 2.5.0

Git is a free and open source distributed version control system .

````
# from jenkins
git ls-remote -h https://github.com/W3Ltd/ziipr_core HEAD

git init
git clone ssh://gitolite@repo.com:1011/repoName.git

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

git tag 1.1.1
git push origin --tags

git blame -L 11,12 file

export GIT_SSL_NO_VERIFY=1;
git config http.sslVerify 0

# avoid git push prompt
git remote set-url origin git+ssh://git@github.com/cn007b/my.git

# grab commit from another branch into current
git cherry-pick hash-to-commit
````

####bisect
````
git bisect start
git bisect bad       -- tell that current situation is bad
git bisect good v1.0 -- bad commit hash
                     -- test, is script working
git bisect bad       -- because all don't works
                     -- test, is script working
git bisect good      -- because all works fine
                     -- test script working
git bisect good
                     -- ...
git bisect reset
````

####diff
````
git diff --cached or diff --staged # after git add shows diff
git diff branch..subBrach
git diff branch:file file
````

####log
````
git log -2                # last 2 commits
git log -p                # shows commits & code in commit
git log --stat            # statiistic about changes
git log --no-meges        # log without merges
git log --follow file.txt # Viewing GIT history of deleted files.
````

####branch
````
git checkout --track -b master origin/master

git show branchName
git rev_parse branch
git brnach --no-merged
git brnach --merged                               # branches merged with current
git branch -v                                     # all branches Y last branch commit
git push remoteRepoName pusedBranch
git push remoteRepoName pusedBranch:newBranchName
git push remoteRepoName :newBranch                # delete branch from remote repo
````

####config

.gitignore
````
.idea/
.vagrant/
node_modules/
vendor/
.DS_Store
````

````
/etc/gitconfig # system
~/.gitconfig   # user
.git/config    # project
````
````
[color]
    ui     = auto
    diff   = true
    status = auto
    branch = auto
[alias]
    b   = branch
    br  = branch
    ch  = checkout
    s   = status -sb
    st  = status
    sl  = stash list
    f   = fetch origin
    ml  = merge origin/live
    mm  = merge origin/master
    d   = diff --word-diff
    df  = diff
    dfc = diff --cached
    dfl = diff origin/live
    dfm = diff origin/master
    dff = diff origin/live --name-only
    ca  = commit -am
    cm  = commit -m
    ph  = push origin HEAD
    lg  = log --pretty=format:\"%h %ad | [%an] %s%d\" --graph --date=short --no-merges
[user]
    name  = Vladimir Kovpak
    email = cn007b@gmail.com
[core]
    excludesfile = ~/.gitignore
    editor       = vim
````
