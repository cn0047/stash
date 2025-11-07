Jira
-

[status](https://status.atlassian.com/)

* roadmap planner

````sh
# to link GitHub PR with issue in JIRA:
git commit -m "PRJ-1273: Some comment."
````

````sql
-- issues filter query
assignee WAS currentUser() DURING (-26w, now())
ORDER BY updated DESC
````

````
{panel:title=Tech details about branches}
Some content...
{panel}

{expand: Click to see more}
    Some info...
{expand}

{code:go}
{code}
````

# Confluence

Use "Insert markup" to add table to confluence page:
````
|| name || code ||
| foo | 1 |
| bar | 2 |
````
