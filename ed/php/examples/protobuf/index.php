<?php

/*
RUN in shell:

# Generate protobuf php files for "foo":
docker run -it --rm -v $PWD:/gh -w /gh/ed/php/examples/protobuf cn007b/php:7.1-protobuf-3 protoc --php_out=proto foo.proto

# Encode msg into protobuf:
docker run -it --rm -v $PWD:/gh -w /gh/ed/php/examples/protobuf cn007b/php:7.1-protobuf-3 bash -c '
    cat foo.msg | protoc --encode=Foo foo.proto > proto/foo.example
'
# Decoding msg from protobuf:
docker run -it --rm -v $PWD:/gh -w /gh/ed/php/examples/protobuf cn007b/php:7.1-protobuf-3 bash -c '
    cat foo.msg | protoc --encode=Foo foo.proto | protoc --decode Foo foo.proto
'

# Test:
docker run -it --rm -v $PWD:/gh -w /gh/ed/php/examples/protobuf cn007b/php:7.1-protobuf-3 composer install
docker run -it --rm -v $PWD:/gh -w /gh/ed/php/examples/protobuf cn007b/php:7.1-protobuf-3 php index.php

*/

require __DIR__ . '/vendor/autoload.php';

require __DIR__ . '/proto/GPBMetadata/Foo.php';
require __DIR__ . '/proto/Foo.php';

function f1($packed)
{
    $parsedFoo = new Foo();
    $parsedFoo->mergeFromString($packed);
    printf("RESULT: id=%s, bar=%s \n\n", $parsedFoo->getId(), $parsedFoo->getBar());
    var_dump($parsedFoo->getTags());
}

$foo = new Foo();
$foo->setId(911);
$foo->setBar('This is BAR.');
$foo->setTags(['php', 'protobuf']);
$packed = $foo->serializeToString();
var_dump($packed);

echo "\n\n-----------------------: 1\n\n";

f1($packed);

echo "\n\n-----------------------: 2\n\n";

$file = './proto/foo.example';
if (file_exists($file)) {
    $packed = file_get_contents($file);
    f1($packed);
}
