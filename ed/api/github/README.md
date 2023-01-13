GitHub
-

[github API](https://docs.github.com/en/rest)
[status](https://www.githubstatus.com/)
[most followed](https://github.com/search?utf8=%E2%9C%93&q=followers%3A%3E%3D10000&type=Users)
[top repositories](https://github.com/search?q=stars:%3E1&s=stars&type=Repositories)
[sponsors](https://github.com/sponsors/cn007b)
[sponsors docs](https://help.github.com/en/github/supporting-the-open-source-community-with-github-sponsors/about-github-sponsors)
[diagram in markdown](https://github.com/mermaid-js/mermaid)
[watching page](https://github.com/watching)

````sh
# ~/.ssh/config
Host github.com
  IdentityFile ~/.ssh/id_rsa_my

# push gist config with http
url = https://gist.github.com/cn007b/b0136cc4db3ad91774e93dbaadcf0331
# push gist config with ssh
url = git@gist.github.com:b0136cc4db3ad91774e93dbaadcf0331.git
````

````sh
brew install gh # https://github.com/cli/cli

gh auth login

gh repo list $org -L 100
````

````sh
curl -i https://api.github.com/users

curl https://api.github.com/users/cn007b
curl https://api.github.com/users/cn007b/repos
curl https://api.github.com/users/cn007b/orgs
# url to tar archive to last release
curl https://api.github.com/repos/cn007b/monitoring/releases | jq '.[0].assets[].browser_download_url' | grep linux

https://api.github.com/orgs/thepkg
````

Don't forget to hit "Squash and merge" instead of "Merge pull request" for pr with > 1 commits.

To add image into gist:
Clone gist by using SSH gist URL.
Add images into gist repo & commit.
Now images public available so it's possible to copy image url from gist web page,
and add into page like this:

````sh
# where:
# 8dbfd59749ab7a75189584ff8ef837ad         - gist id.
# db97ba8d4b326447560562e54737ad7c70b24170 - commit hash id.
![caption](https://gist.github.com/cn007b/8dbfd59749ab7a75189584ff8ef837ad/raw/db97ba8d4b326447560562e54737ad7c70b24170/1.png)
````
