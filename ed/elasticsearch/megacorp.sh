#!/bin/sh

# mapping for megacorp
curl -XPUT http://localhost:9200/megacorp -d '{
    "mappings" : {
      "employee": {
        "properties": {
            "first_name": {"type": "string", "index": "not_analyzed"},
            "last_name": {"type": "string", "index": "not_analyzed"},
            "age": {"type": "integer"},
            "about": {"type": "string", "index": "not_analyzed"},
            "last_login_at": {"type": "date", "format": "yyy-MM-dd"},
            "city": {"type": "string", "index": "not_analyzed"},
            "location": {"type": "geo_point", "lat_lon": "true"},
            "interests": {"type": "string"},
            "fetish": {"type": "nested"}
        }
      },
      "car": {
        "_parent" : {
            "type": "employee"
        },
        "properties": {
            "name": {"type": "string", "index": "not_analyzed"},
            "brand": {"type": "string", "index": "not_analyzed"},
            "about": {"type": "string", "index": "not_analyzed"}
        }
      }
    }
}'

# Create new documents (employee)
curl -XPUT localhost:9200/megacorp/employee/1?routing=JohnSmith -d '{
    "first_name" : "John",
    "last_name" : "Smith",
    "age" : 25,
    "about" : "I love to go rock climbing",
    "last_login_at": "2016-01-21",
    "city": "London",
    "location": {"lat": 51.5072, "lon": 0.1275},
    "interests": [ "sports", "music" ],
    "fetish": {"name": "none", "count": "no count", "desc": "no desc"}
}'
curl -XPUT localhost:9200/megacorp/employee/2 -d '{
    "first_name" : "Jane",
    "last_name" : "Smith",
    "age" : 32,
    "about" : "I like to collect rock albums",
    "last_login_at": "2016-01-10",
    "city": "Manchester",
    "location": {"lat": 53.4667, "lon": 2.2333},
    "interests": [ "music" ]
}'
curl -XPUT localhost:9200/megacorp/employee/3 -d '{
    "first_name" : "Douglas",
    "last_name" : "Fir",
    "age" : 35,
    "about": "I like to build cabinets",
    "last_login_at": "2016-02-14",
    "city": "Kyiv",
    "location": {"lat": 50.4500, "lon": 30.5233},
    "interests": [ "forestry", "cars" ]
}'
curl -XPUT localhost:9200/megacorp/employee/4 -d '{
    "first_name" : "Louis",
    "last_name" : "de Funès",
    "age" : 70,
    "about": "Actor. I like movies.",
    "last_login_at": "2012-03-04",
    "city": "Paris",
    "location": {"lat": 48.8567, "lon": 2.3508},
    "interests": [ "fantomas", "theatre", "hollywood" ]
}'
curl -XPUT localhost:9200/megacorp/employee/5 -d '{
    "first_name" : "Cristiano",
    "last_name" : "Ronaldo",
    "age" : 31,
    "about": "Footballer. I like sport.",
    "last_login_at": "2016-03-11",
    "city": "Santo António",
    "location": {"lat": 37.1939, "lon": 7.4158},
    "interests": [ "football", "cars", "casino" ]
}'
curl -XPUT localhost:9200/megacorp/employee/6 -d '{
    "first_name" : "Gennady",
    "last_name" : "Golovkin",
    "age" : 33,
    "about": "Professional Boxer",
    "last_login_at": "2016-03-03",
    "city": "Karaganda",
    "location": {"lat": 49.8333, "lon": 73.1667},
    "interests": [ "boxing", "WBA", "IBO", "cars" ]
}'
curl -XPUT localhost:9200/megacorp/employee/7 -d '{
    "first_name" : "Jackie",
    "last_name" : "Chan",
    "age" : 61,
    "about": "Martial Artist",
    "last_login_at": "2016-03-12",
    "city": "Hong Kong",
    "location": {"lat": 22.2783, "lon": 114.1747},
    "interests": [ "movie", "hollywood", "kong foo" ]
}'
curl -XPUT localhost:9200/megacorp/employee/13 -d '{
    "first_name" : "Wladimir",
    "last_name" : "Klitschko",
    "age" : 40,
    "about": "REALLY Professional Boxer. Longtime World Heavyweight Champion.",
    "last_login_at": "2016-04-22",
    "city": "Kyiv",
    "location": {"lat": 50.4501, "lon": 30.5234},
    "interests": [ "boxing", "sport", "movie", "hollywood" ]
}'
curl -XPUT localhost:9200/megacorp/employee/14 -d '{
    "first_name" : "Paul",
    "last_name" : "McCartney",
    "age" : 74,
    "about": "Music. London. etc...",
    "last_login_at": "2016-04-22",
    "city": "London",
    "location": {"lat": 51.5074, "lon": 0.1278},
    "interests": [ "london", "music" ]
}'
curl -XPUT localhost:9200/megacorp/employee/15 -d '{
    "first_name" : "Wayne",
    "last_name" : "Rooney",
    "age" : 30,
    "about": "is an English professional footballer",
    "last_login_at": "2016-07-07",
    "city": "London",
    "location": {"lat": 51.5074, "lon": 0.1278},
    "interests": ["football", "sport", "cars"],
    "fetish": {"name": "RANGE_ROVER_SPORT"}
}'
curl -XPUT localhost:9200/megacorp/employee/16 -d '{
    "first_name" : "Jayce",
    "last_name" : "Chan",
    "age" : 33,
    "about": "Son of Martial Artist",
    "last_login_at": "2016-07-11",
    "city": "Hong Kong",
    "location": {"lat": 22.2783, "lon": 114.1747},
    "interests": [ "movie", "hollywood", "kong foo", "father" ]
}'
