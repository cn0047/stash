<?php

namespace Foo;

class FooClass
{
    public function __construct() {
        echo 'Foo'.PHP_EOL;
    }
}
?>

<?php

namespace Bar;

class BarClass
{
    public function __construct() {
        echo 'Bar'.PHP_EOL;
    }
}
?>

<?php

namespace Foo;

use Bar\BarClass as FooClass2;
$class = new FooClass2;
/*
Bar
*/
?>

<?php

namespace Foo;

use Bar\BarClass as FooClass;
$class = new FooClass;
/*
PHP Fatal error:  Cannot use Bar\BarClass as FooClass because the name is already in use
*/
?>
