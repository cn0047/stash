<?php

$m = new Mongo();
$db = $m->selectDB('mydb');
$collection = new MongoCollection($db, 'users');

$collection->insert(['firstname' => 'Bob', 'lastname' => 'Jones']);

$collection->update(
    ['firstname' => 'Bob'],
    ['$set' => ['lastname' => 'Jones NEW', 'address' => '1 Smith Lane - '.uniqid()]]
);

$cursor = $collection->find();
$result = [];
foreach ($cursor as $doc) {
    $result[] = $doc;
}
var_export($result);

/*
array (
  0 =>
  array (
    '_id' =>
    MongoId::__set_state(array(
       '$id' => '561d7b3f113e9248168b4567',
    )),
    'firstname' => 'Bob',
    'lastname' => 'Jones NEW',
    'address' => '1 Smith Lane - 561d7b3f3dc98',
  ),
)
*/
