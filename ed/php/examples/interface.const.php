<?php

/**
 * It's possible for interfaces to have constants.
 * Interface constants works exactly like class constants
 * except they cannot be overridden by a class/interface that inherits them.
 */
interface Foo
{
    const A = 'Foo A';
}

class Bar implements Foo
{
    const B = 'Bar B';
}

echo Bar::A.PHP_EOL;
echo Bar::B.PHP_EOL;
/*
Foo A
Bar B
*/
?>

<?php

interface Foo2
{
    const A = 'Foo2 A';
}

class Bar2 implements Foo2
{
    const A = 'Bar2 A';
}

echo Bar::A.PHP_EOL;
/*
PHP Fatal error:  Cannot inherit previously-inherited or override constant A from interface Foo2
*/
