Sumologic
-

[docs](https://help.sumologic.com/)

````
(_sourceName=my_svc)
| where msg = "log message"
| where msg matches "*regex*"
| where !(msg matches "http*")
````
