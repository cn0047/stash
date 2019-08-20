Stackoverflow
-

[Top UA SO Users](http://data.stackexchange.com/stackoverflow/query/763779/top-ukrainian-stackoverflow-users)
[Top UA SO Users by Tag](https://data.stackexchange.com/stackoverflow/query/1004945/top-ukrainian-stackoverflow-users-by-tag)
[Trends](https://insights.stackoverflow.com/trends?tags=javascript%2Cphp)

[js](http://jsbin.com/)
[php](http://rextester.com)
[sql](http://sqlfiddle.com)
[regex](https://regex101.com/r/zO0kO8/1)
[bash](https://ideone.com/A8momR)

````
<kbd>Ctrl</kbd><kbd>Shift</kbd><kbd>P</kbd>

# SQL Fiddle:
<kbd>[SQL Fiddle demo 5.6](http://sqlfiddle.com/#!9/0a259/23)</kbd>
````

Hotkeys:
````
CTRL + K # Code
````

BigQuery:
````sql
https://console.cloud.google.com/marketplace/details/stack-exchange/stack-overflow

select display_name, reputation
from `bigquery-public-data.stackoverflow.users`
where location like "%ukraine%" or  location like "%Ukraine%"
order by reputation desc
limit 70
;

select tags
from `bigquery-public-data.stackoverflow.stackoverflow_posts` p
join `bigquery-public-data.stackoverflow.tags` t on t.name in (p.tags)
limit 5;
````

#### Code

````
[tag:xdebug]

<!-- language: lang-html -->
<!-- end snippet -->
````

JS:
````
<!-- begin snippet: js hide: false console: true babel: false -->
<!-- language: lang-js -->
console.log(200);
<!-- end snippet -->
````
