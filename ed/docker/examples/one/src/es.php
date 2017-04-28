<?php

$response = `curl -s -XGET localhost:9200/megacorp/employee/_search`;
$resultSet = json_decode($response, true);
var_dump($resultSet);
