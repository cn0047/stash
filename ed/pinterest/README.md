pinterest
-

https://developers.pinterest.com/apps/
https://developers.pinterest.com/docs/

#### Get API Token

````sh
open https://developers.pinterest.com/apps/
# go to app page and add `https://localhost:3443/` on app web page into "Redirect URIs"

docker run -it --rm -v $PWD:/app -w /app -p 3443:443 \
  -v $PWD/docker/nginx/https/php-fpm.conf:/etc/nginx/conf.d/default.conf \
  -v $PWD/docker/nginx/https/localhost.crt:/ssl/localhost.crt \
  -v $PWD/docker/nginx/https/localhost.key:/ssl/localhost.key \
  -v $PWD:/gh \
  cn007b/php /bin/bash -c '
    service php7.1-fpm start;
    service nginx start;
    tail -f /dev/stdout
  '

appID={id}
secret={secret}

# Getting your authorization code
open 'https://api.pinterest.com/oauth/?response_type=code&redirect_uri=https://localhost:3443/&scope=read_public,write_public,read_relationships,write_relationships&state=randomString&client_id='$appID

# Grab code from previous ⬆ request
code={code}

#Getting your access token
curl -i -X POST 'https://api.pinterest.com/v1/oauth/token' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -d 'grant=authorization_code&grant_type=authorization_code&client_id='$appID'&client_secret='$secret'&code='$code

token="{token}"
````

#### API

````sh
# boards list
curl 'https://api.pinterest.com/v1/me/boards/?access_token='$token | jq

# create board
curl -X POST 'https://api.pinterest.com/v1/boards/?access_token='$token \
  -H 'Content-Type: application/json' -d '{"name": "test"}'

# create new pin
curl -X POST 'https://api.pinterest.com/v1/pins/?access_token='$token \
  -H 'Content-Type: application/json' -d '{
    "board":"2616723",
    "note":"my test",
    "link":"https://cn007b.tumblr.com/",
    "image_url":"https://i.kinja-img.com/gawker-media/image/upload/t_original/v66gt9dn0kxsmhpxmcmq.jpg"
  }'

# get pins from all boards
curl 'https://api.pinterest.com/v1/me/pins/?limit=5&fields=id,created_at,link,url,creator,board,note,color,counts,image,metadata&access_token='$token | jq

# get pins from board
board='335659047167093089'
curl 'https://api.pinterest.com/v1/boards/'$board'/pins/?limit=10&fields=id,created_at,link,url,creator,board,note&access_token='$token | jq

# get pin by id
curl -L 'https://api.pinterest.com/v1/pins/348817933631869565?access_token='$token | jq
````
