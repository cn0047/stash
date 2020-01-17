curl
-

````sh
-A, --user-agent
-b, --cookie                         # Read cookies from STRING/FILE
-b, --cookie STRING/FILE             # String or file to read cookies from (H)
-c, --cookie-jar                     # Write cookies to FILE
-d, --data DATA                      # HTTP POST data (H)
-e, --referer                        # Referer URL (H)
-F, --form CONTENT                   # Specify HTTP multipart POST data (H)
-I                                   # Only the HTTP-headers
-i                                   # Include the HTTP-headers
-o, --output FILE                    # Write output to <file> instead of stdout
-O, --remote-name                    # Write output to file wiht origin name
-s, --silent
-u, --user USER[:PASSWORD]           # Server user and password
-v, --verbose
-x, --proxy [PROTOCOL://]HOST[:PORT] # Use proxy on given port

-H 'Content-Type: application/json; charset=UTF-8 application/json; charset=UTF-8' \

# (REST) JSON at POST.
# More examples available here: https://github.com/cn007b/my/blob/master/ed/php.symfony/generateRESTapi.md
curl -XPOST http://localhost:3000/dishes \
    -H 'Content-Type: application/json' -d '{"name": "newDish", "description": "newDesc"}'
curl -X POST -H 'application/json' -d '{"key":"val"}' http://url.com
curl -X POST -H 'Content-Type: text/plain' -d @/tmp/foo.txt http://url.com
curl http://url.com -d x=1 -d y=2

curl http://login:pass@base-http-auth.com/
curl -u login:pass http://base-http-auth.com/

// user == 'admin' && pass == 'password'
curl http://localhost:3000 -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQ='

# download file
curl $host -o $out

# upload file
curl http://localhost:8000 -F "file=@/home/kovpak/Downloads/download.jpg"
curl http://localhost:8000 -H "Content-Type: multipart/form-data" -F "file=@/Users/k/f.txt" -F "msg=MyFile"

curl -T firmware.bin http://0.0.0.48/cgi-bin/package_install?hash=017

# shows spent time (⏱)
time curl -si https://realtimelog.herokuapp.com/test

curl -A 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36' \
  http://localhost:8000

# ignore invalid SSL Certificate
curl -k https://localhost:4433

curl -s 'https://github.com/cn007b' -o /dev/null -w '%{http_code}'
curl -s 'https://github.com/cn007b' -o /dev/null -w '%{time_total}'
````
