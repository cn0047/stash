<?php

echo sprintf('"%04d"', 1).PHP_EOL; // "0001"

var_dump(in_array('test', array(0))); // bool(true) - Because test converts to integer.
