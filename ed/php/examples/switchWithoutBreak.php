<?php

function foo($scenario = '')
{
    switch ($scenario) {
        case '':
        case 'one':
            echo 1;
        case '':
        case 'two':
            echo 2;
        default:
            echo 3;
    }
}

echo PHP_EOL.'Call one: ';
foo('one');
echo PHP_EOL.'Call two: ';
foo('two');
echo PHP_EOL.'Call "": ';
foo();
echo PHP_EOL;

/*
Call one: 123
Call two: 23
Call "": 123
*/
