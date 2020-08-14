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
* map<t1,t2> - Map from one type to another
* list<t1> - Ordered list of one type
* set<t1> - Set of unique elements of one type

enum

Thrift JSON:
````sh
{"tf":"0"}                      # bool false
{"tf":"1"}                      # bool true

{"str":"this is string"}

{"lst":["str",2,"tag1","tag2"]} # list of strings with length 2
````
