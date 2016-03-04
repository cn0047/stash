<?php

use Kahlan\Filter\Filter;
use Kahlan\Reporter\Coverage\Exporter\Coveralls;

require __DIR__ . '/vendor/autoload.php';

defined('APPLICATION_PATH')
|| define('APPLICATION_PATH', __DIR__);

require APPLICATION_PATH . '/app/controllers/IndexController.php';
require APPLICATION_PATH . '/app/services/ElasticSearch.php';

// It overrides some default option values.
// Note that the values passed in command line will overwrite the ones below.
$args = $this->args();
$args->argument('ff', 'default', 1);
$args->argument('coverage', 'default', 3);
$args->argument('coverage-scrutinizer', 'default', 'scrutinizer.xml');
$args->argument('coverage-coveralls', 'default', 'coveralls.json');
