YouTube
-

See: https://github.com/youtube/api-samples/tree/master/go

````
# @see
/Users/kovpakvolodymyr/.credentials/youtube-go.json

# get playlists
go run playlists.go oauth2.go errors.go --channelId=UC4UxIHgjkBoXCJ3XdYXwaRg

# upload video
go run upload_video.go errors.go oauth2.go \
  --filename=$f.y.mp4 --title=$h1 --description=$desc --keywords=$tags
````
