twitter
-

[api](https://apps.twitter.com)
[dev](dev.twitter.com)

````bash
# to see tweet on twitter
open https://twitter.com/statuses/$tweetID
````

## twurl

````bash
gem install twurl

twurl authorize --consumer-key $ck --consumer-secret $cs

# check
twurl /1.1/statuses/home_timeline.json | jq
twurl /1.1/statuses/user_timeline.json | jq
````

````bash
# not working now
# curl -X GET 'http://search.twitter.com/search.json?q=007'

# tweet text msg
twurl -X POST -H api.twitter.com "/1.1/statuses/update.json?status=Hello! This Tweet was sent via the Twitter API." | jq

# upload image
twurl -X POST -H upload.twitter.com "/1.1/media/upload.json" --file "/Users/k/Downloads/j.jpg" --file-field "media"
# get
# twurl -X GET -H upload.twitter.com "/1.1/media/upload.json?command=STATUS&media_id="$mediaID
````
