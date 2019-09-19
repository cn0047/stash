<?php

return new \Phalcon\Config(array(

    'version' => '1.0',

    'database' => array(
        'adapter'    => 'Mysql',
        'host'       => 'localhost',
        'username'   => 'root',
        'password'   => '',
        'dbname'     => 'dbRelationships',
        'charset'    => 'utf8',

        'username'   => 'zii',
        'password'   => 'NGtveIF6EEprvY',
    ),

    'application' => array(
        'modelsDir' => 'app/models/',
        'baseUri' => '/dbRelationships/',
    ),

//    'dir' => [
//        'models' => __DIR__ . '/../models/',
//    ],

    /**
     * if true, then we print a new line at the end of each execution
     *
     * If we dont print a new line,
     * then the next command prompt will be placed directly on the left of the output
     * and it is less readable.
     *
     * You can disable this behaviour if the output of your application needs to don't have a new line at end
     */
    'printNewLine' => true
));
