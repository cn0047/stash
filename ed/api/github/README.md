GitHub
-

[Most Followed](https://github.com/search?utf8=%E2%9C%93&q=followers%3A%3E%3D10000&type=Users)
[Top Repositories](https://github.com/search?q=stars:%3E1&s=stars&type=Repositories)
[Sponsors](https://help.github.com/en/github/supporting-the-open-source-community-with-github-sponsors/about-github-sponsors)
[status](https://www.githubstatus.com/)

Don't forget to hit "Squash and merge" instead of "Merge pull request" for pr with > 1 commits.

````sh
# ~/.ssh/config
Host github.com
  IdentityFile ~/.ssh/id_rsa_my

````

````sh
curl -i https://api.github.com/users

curl https://api.github.com/users/cn007b
curl https://api.github.com/users/cn007b/repos
curl https://api.github.com/users/cn007b/orgs
# url to tar archive to last release
curl https://api.github.com/repos/cn007b/monitoring/releases | jq '.[0].assets[].browser_download_url' | grep linux
````

To add image into gist:

Clone gist by using gist URL.
Add images into gist repo & commit.
Now images public available so it's possible to copy image url from gist web page,
and add into page like this:

````sh
# where:
# 8dbfd59749ab7a75189584ff8ef837ad         - gist id.
# db97ba8d4b326447560562e54737ad7c70b24170 - commit hash id.
![caption](https://gist.github.com/cn007b/8dbfd59749ab7a75189584ff8ef837ad/raw/db97ba8d4b326447560562e54737ad7c70b24170/1.png)
````
