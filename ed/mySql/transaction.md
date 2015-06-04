Transaction
-

|terminal 1|terminal 2|
|-|-|
|select title from tree where id = 9; -- service 9||
|start transaction;||
|update tree set title = 'service 99' where id = 9;||
|Query OK||
||select title from tree where id = 9; -- service 9|
||start transaction;|
||update tree set title = 'service 999' where id = 9; -- hang|
|commit;||
||Query OK|
|select title from tree where id = 9; -- service 99||
||commit;|
||select title from tree where id = 9; -- service 99|
|select title from tree where id = 9; -- service 99||
