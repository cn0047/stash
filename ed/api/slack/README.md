slack
-

[status](https://status.slack.com/)
[emoji](https://slackmojis.com/)
[status](https://status.slack.com/)

````
@here

:this:
````

````sh
curl https://hooks.slack.com/services/TTT/BBBB/HHH -X POST --data-urlencode 'payload={
"channel": "#deploy",
"username": "bot",
"text": "This is bot.",
"icon_emoji": ":ghost:"
}'

curl https://hooks.slack.com/services/TTT/BBBB/HHH -X POST -H 'Content-type: application/json' --data '{
  "text":"Hello, World!"
}'
````
