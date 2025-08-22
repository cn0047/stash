GitLab
-
<br>GitLab Enterprise Edition v16.4.1-ee (few codereviewers).
<br>GitLab Community Edition v16.4.1.

* Hosted GitLab.
* Self-managed GitLab.

````sh
# list all projects
i=0;
while [ $i -ne 100 ]; do;
  i=$(($i+1)); echo "$i";
  curl -H "Private-Token: $GITLAB_TOKEN" "https://gitlab.com/api/v4/projects?per_page=100&page=$i" >> /tmp/x.json
  echo "," >> /tmp/x.json
done

````
