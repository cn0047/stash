Thrift
-

Types:
* bool - Boolean, one byte
* i8 (byte) - Signed 8-bit integer
* i16 - Signed 16-bit integer
* i32 - Signed 32-bit integer
* i64 - Signed 64-bit integer
* double - 64-bit floating point value
* string - String
* binary - Blob (byte array)
* list<t1> - Ordered list of one type
* set<t1> - Set of unique elements of one type
* map<t1,t2> - Map from one type to another

enum

Thrift JSON:
````sh
Content-Type: application/x-thrift

[1,"methondName",1,0,{"1":{"str":"foo"},"2":{"str":"bar"}}]

{"tf":"0"} # bool false
{"tf":"1"} # bool true
{"i32":9}  # int32
{"i64":9}  # int64

{"str":"this is string"}

{"lst":["str",2,"tag1","tag2"]}                     # list of strings with length 2
{"set":["rec",1,{"1":{"str":"foo"},"2":{"i32":9}}]} # set with 2 items
{"map":["str","str",1,{"en":"title"}]}              # map
{"map":["str","rec",1,{"en":{"1":{"str":"msg"}}}]}  # map

# obj 1
{"rec":{"1":{"str":"foo"}}}

# obj 2
{"rec":{
  "1":{"str":"foo"},
  "2":{"str":"bar"}
}}

# array of objects (1 object)
"1":{"set":["rec",1, {"1":{"str":"ok"},"2":{"i32":200}} ]}
"2":{"lst":["rec",1, {"1":{"str":"ok"}} ]}
````
