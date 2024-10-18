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
-k                                   # Skip TLS verification
-o, --output FILE                    # Write output to <file> instead of stdout
-O, --remote-name                    # Write output to file wiht origin name
-L                                   # Follow location response header
-s, --silent
-S, --show-error
-u, --user USER[:PASSWORD]           # Server user and password
-v, --verbose
-x, --proxy [PROTOCOL://]HOST[:PORT] # Use proxy on given port

jh='Content-Type: application/json'
jh='Content-Type: application/json; charset=UTF-8 application/json; charset=UTF-8'
-H $jh
ua='Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36'
ua='User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36'
-H $ua

# (REST) JSON at POST.
# More examples available here: php.symfony/generateRESTapi.md
curl -XPOST http://localhost:3000/dishes -H 'Content-Type: application/json' -d '
  {"name": "newDish", "description": "newDesc"}
'

curl -X POST -H 'application/json' -d '{"key":"val"}' http://url.com

curl -X POST -H 'Content-Type: text/plain' -d @/tmp/foo.txt http://url.com

echo '{"fromFile":"yes"}' > /tmp/x.json \
  && curl -XPOST 'https://realtimelog.herokuapp.com:443/pvq6l1g0m8d' \
  -H 'Content-Type: application/json' -d @/tmp/x.json

curl http://url.com -d x=1 -d y=2

# post
c='c1'
s='s2'
urlencode() {
  php -r 'print(urlencode($argv[1]));' $1
}
curl -X POST "https://$h" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d 'clientId='`urlencode $c`'&secretKey='`urlencode $s`

curl -XPOST 'https://realtimelog.herokuapp.com:443/ka01kxtxsh' \
  -H 'Content-Type: application/json' -d '{"code":"200", "status": "OK"}'
# and
curl -XPOST 'https://realtimelog.herokuapp.com:443/ka01kxtxsh' \
  -H 'Content-Type: application/json' -d @- << EOF
{
  "code":"200",
  "status": "OK"
}
EOF

curl http://login:pass@base-http-auth.com/
curl -u login:pass http://base-http-auth.com/

# user == 'admin' && pass == 'password'
curl http://localhost:3000 -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQ='

# download file
curl $url -o $out

# download file and use remote name
curl -O $url

# download file with origin name into target dir
curl -O --create-dirs --output-dir /tmp/x $url

# upload file
curl http://localhost:8000 -F "file=@/home/kovpak/Downloads/download.jpg"
curl http://localhost:8000 -H "Content-Type: multipart/form-data" -F "file=@/Users/k/f.txt" -F "msg=MyFile"

# simple file upload to: https://gofile.io
curl -F email=cnfxlr@gmail.com -F file=@x.txt https://srv-file6.gofile.io/uploadFile

curl -T firmware.bin http://0.0.0.48/cgi-bin/package_install?hash=017

curl -A "User-Agent: $ua" http://localhost:8000

# ignore invalid SSL certificate
curl -k https://localhost:4433

# shows spent time (⏱)
cat << EOF >> /tmp/writeout.fmt.txt
\n\n
time_namelookup:    \t%{time_namelookup}s\n
time_redirect:      \t%{time_redirect}s\n
time_connect:       \t%{time_connect}s\n
time_appconnect:    \t%{time_appconnect}s\n
time_pretransfer:   \t%{time_pretransfer}s\n
time_starttransfer: \t%{time_starttransfer}s\n
Total time:         \t%{time_total}s\n
EOF
time curl -si https://realtimelog.herokuapp.com/rkc8q6llprn
curl -si -w "@/tmp/writeout.fmt.txt" https://realtimelog.herokuapp.com/rkc8q6llprn

curl -s 'https://github.com/cn007b' -o /dev/null -w '%{http_code}'
curl -s 'https://github.com/cn007b' -o /dev/null -w '%{time_total}'
````
