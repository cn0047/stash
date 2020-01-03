Environments
-

````php
export SYMFONY_ENV=prod

// How to Master and Create new Environments
class AppKernel extends Kernel
{
    public function registerContainerConfiguration(LoaderInterface $loader)
    {
        $loader->load($this->getRootDir().'/config/config_'.$this->getEnvironment().'.yml');
    }
}

$kernel = new AppKernel('prod', false);

# 'prod' environment (debug is always disabled for 'prod')
php bin/console command_name --env=prod
# 'test' environment and debug disabled
php bin/console command_name --env=test --no-debug

# app/config/config_benchmark.yml
imports:
    - { resource: config_prod.yml }
framework:
    profiler: { only_exceptions: false }

// change just this line
$kernel = new AppKernel('benchmark', false);

http://localhost/app_benchmark.php

// How to Optimize your Development Environment for Debugging
# Disabling the Bootstrap File and Class Caching
$loader = require_once __DIR__.'/../app/bootstrap.php.cache';
require_once __DIR__.'/../app/AppKernel.php';
$kernel = new AppKernel('dev', true);
$kernel->loadClassCache();
$request = Request::createFromGlobals();
// $loader = require_once __DIR__.'/../app/bootstrap.php.cache';
$loader = require_once __DIR__.'/../app/autoload.php';
require_once __DIR__.'/../app/AppKernel.php';
$kernel = new AppKernel('dev', true);
// $kernel->loadClassCache();
$request = Request::createFromGlobals();
````
