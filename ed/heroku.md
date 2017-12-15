Heroku
-

````
# start
heroku apps

heroku config:set GITHUB_USERNAME=bond -a realtimelog

# show ENVIRONMENT variables
heroku config -a realtimelog

heroku logs -a realtimelog --tail

heroku run bash -a realtimelog
````
