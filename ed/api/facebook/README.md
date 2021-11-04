facebook
-

[Doc](https://developers.facebook.com/)
[API](https://developers.facebook.com/docs/reference/api/)
[API explorer](https://developers.facebook.com/tools/explorer/)
[Token](https://developers.facebook.com/tools/debug/accesstoken/)

[Instant Articles](http://www.facebook.com/instant_articles/signup/)

````sh
token="EAAHA4AK7USIBAEQomCDv0Hs1nfNwZBivZBEQ3q5ZC8KdmbkwOwsJEEsjFPiI4YRksWwX55glrEj4GDOR2RB3tjVpZCum6eK9HRKNs8wGLJBBZCXEgPMljU3AZCHRxXIOC327zfo3tqb1MhfkIqVqFLzrMr5ZCV9YfTD2l38ZAHzvcHiu9Jf2mMgImEKhIOOePXpvmSWPws2US4NaKTM3pn0U"
pageId=1719701001418493

# me
curl -i -X GET "https://graph.facebook.com/v3.0/me?fields=id%2Cname&access_token="$token

# page info
curl "https://graph.facebook.com/"$pageId \
  -F 'method=get' \
  -F 'development_mode=true' \
  -F 'fields=about,attire,bio,location,parking,hours,emails,website' \
  -F 'access_token='$token

# get posts
curl -i -X GET "https://graph.facebook.com/v3.0/"$pageId"/instant_articles?access_token="$token

# test post
curl -i -X POST "https://graph.facebook.com/v3.0/"$pageId"/feed?published=false&message=An_unpublished_post&access_token="$token
````

