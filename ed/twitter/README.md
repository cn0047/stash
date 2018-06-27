twitter
-

https://apps.twitter.com

## twurl

````sh
gem install twurl

twurl authorize --consumer-key kkkalH8MXr7iI7koqRgJ2KMWy --consumer-secret sssywtyALcKUvu04xKQmNnPmlu6un8D3SbsHUys2Z9HNfWBOy4

# check
twurl /1.1/statuses/home_timeline.json | jq
twurl /1.1/statuses/user_timeline.json | jq
````

````sh
# upload image
twurl -H upload.twitter.com -X POST "/1.1/media/upload.json" --file "/Users/k/Downloads/i.jpg" --file-field "media"
````
