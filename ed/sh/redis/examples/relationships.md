Hash vs Set in Relationships.
-

Relationships like on twitter:
````sh
+--------+----------+---------+
| type   | fromUser | toTuser |
|--------+----------+---------+
| follow | user10   | user1   |
| follow | user10   | user2   |
| follow | user10   | user3   |
+--------+----------+---------+
| follow | user20   | user1   |
| follow | user20   | user10  |
+--------+----------+---------+
| follow | user30   | user10  |
| follow | user30   | user20  |
+--------+----------+---------+
````

#### Set:

````sh
sadd srelationships:user10:following user1 user2 user3
sadd srelationships:user10:followers user20 user30

# get relationships for user10
smembers srelationships:user10:following
smembers srelationships:user10:followers

# add new follower for user10
sadd srelationships:user10:followers user40
# remove follower for user10
srem srelationships:user10:followers user40

# check is value present in set
sismember srelationships:user10:followers user40
````

#### Hash:

````sh
hmset hrelationships:user10:following user1 1 user2 1 user3 1
hmset hrelationships:user10:followers user20 1 user30 1

# get relationships for user10
hgetall hrelationships:user10:following
hgetall hrelationships:user10:followers

# add new follower for user10
hmset hrelationships:user10:followers user40 1
# remove follower for user10
hdel hrelationships:user10:followers user40

# check is value present in set
hexists hrelationships:user10:followers user40
````
