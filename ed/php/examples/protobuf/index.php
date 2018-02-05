<?php

/*
RUN in shell:

docker run -it --rm -v $PWD:/gh -w /gh/ed/php/examples/protobuf nphp composer install
docker run -it --rm -v $PWD:/gh -w /gh/ed/php/examples/protobuf nphp php index.php

*/

require __DIR__ . '/vendor/autoload.php';

require __DIR__ . '/proto/GPBMetadata/Foo.php';
require __DIR__ . '/proto/Foo.php';

$foo = new Foo();
$foo->setId(911);
$foo->setBar('This is BAR.');
$foo->setTags(['php', 'protobuf']);
$packed = $foo->serializeToString();
var_dump($packed);

echo "\n\n-----------------------\n\n";

$parsedFoo = new Foo();
$parsedFoo->mergeFromString($packed);
printf('RESULT: id=%s, bar=%s', $parsedFoo->getId(), $parsedFoo->getBar());
var_dump($parsedFoo->getTags());
