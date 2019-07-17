GitHub
-

[Most Followed](https://github.com/search?utf8=%E2%9C%93&q=followers%3A%3E%3D10000&type=Users)
[Top Repositories](https://github.com/search?q=stars:%3E1&s=stars&type=Repositories)

````bash
curl -i https://api.github.com/users

curl https://api.github.com/users/cn007b
curl https://api.github.com/users/cn007b/repos
curl https://api.github.com/users/cn007b/orgs
````

````bash
# for JIRA:
git commit -m "PROJ-123 my comment..."
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
![code](https://gist.github.com/cn007b/8dbfd59749ab7a75189584ff8ef837ad/raw/db97ba8d4b326447560562e54737ad7c70b24170/1.png)
````