Generate RESTful api
-

````sql
create table vendor (
  id int not null auto_increment key,
  name varchar(50) not null default '' unique,
  country varchar(50) not null default ''
);
create table model (
  id int not null auto_increment key,
  vendor_id int not null,
  name varchar(50) not null default '',
  foreign key (vendor_id) references vendor(id) on delete restrict
);
````

#### Generate entities

````
php app/console doctrine:mapping:import --force AppBundle xml
php app/console doctrine:mapping:convert annotation ./src
php app/console doctrine:generate:entities AppBundle
````

https://github.com/voryx/restgeneratorbundle

composer require voryx/restgeneratorbundle dev-master

````
php app/console voryx:generate:rest --entity=AppBundle:Vendor
````

#### Examples

````
curl -i http://127.0.0.1:8000/app_dev.php/api/vendors.json
curl -i -H "Content-Type: application/json" -X POST -d '{"name" : "BMW", "country" : "Germany"}' http://127.0.0.1:8000/app_dev.php/api/vendors.json
curl -i -H "Content-Type: application/json" -X POST -d '{"name" : "Aston Martin", "country" : "UK"}' http://127.0.0.1:8000/app_dev.php/api/vendors.json
curl -i http://127.0.0.1:8000/app_dev.php/api/vendors/1.json
curl -i -H "Content-Type: application/json" -X POST -d '{"name" : "ZAZ", "country" : "UA"}' http://127.0.0.1:8000/app_dev.php/api/vendors.json
curl -i -H "Content-Type: application/json" -X PUT -d '{"name" : "ZaZ", "country" : "Ukraine"}' http://127.0.0.1:8000/app_dev.php/api/vendors/3.json
curl -i -H "Content-Type: application/json" -X PATCH -d '{"name" : "autoZaZ"}' http://127.0.0.1:8000/app_dev.php/api/vendors/3.json
curl -i -X DELETE http://127.0.0.1:8000/app_dev.php/api/vendors/3.json

curl -i http://127.0.0.1:8000/app_dev.php/api/models.json
curl -i -H "Content-Type: application/json" -X POST -d '{"vendor" : 1, "name" : "m6"}' http://127.0.0.1:8000/app_dev.php/api/models.json
curl -i -H "Content-Type: application/json" -X POST -d '{"vendor" : 2, "name" : "db9 v12"}' http://127.0.0.1:8000/app_dev.php/api/models.json
curl -i http://127.0.0.1:8000/app_dev.php/api/models/1.json
curl -i -H "Content-Type: application/json" -X POST -d '{"vendor" : 4, "name" : "lanos"}' http://127.0.0.1:8000/app_dev.php/api/models.json
curl -i -H "Content-Type: application/json" -X PUT -d '{"name" : "lanos xl"}' http://127.0.0.1:8000/app_dev.php/api/models/3.json
curl -i -H "Content-Type: application/json" -X PATCH -d '{"vendor" : 1, "name" : "x5"}' http://127.0.0.1:8000/app_dev.php/api/models/3.json
curl -i -X DELETE http://127.0.0.1:8000/app_dev.php/api/models/3.json
````
