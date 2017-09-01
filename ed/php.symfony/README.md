Symfony
-
2.6.6

[book](http://symfony.com/pdf/Symfony_book_2.6.pdf?v=4)
|
[cookbook](http://symfony.com/pdf/Symfony_cookbook_2.6.pdf?v=4)
|
[best practices](http://symfony.com/pdf/Symfony_best_practices_2.7.pdf?v=4)

#### Running the Symfony Application

````
cd my_project_name/
cd ed/symfony/examples/bulletinBoard/

php bin/console server:run
php bin/console server:stop

then http://localhost:8000
````

#### Installing the Symfony Installer

````
sudo curl -LsS http://symfony.com/installer -o /usr/local/bin/symfony
sudo chmod a+x /usr/local/bin/symfony
````

#### Creating the Symfony Application

````
symfony new my_project_name
symfony new my_project_name 2.3.26
````

#### Creating a Symfony Application with Composer

````
composer create-project symfony/framework-standard-edition my_project_name
composer create-project symfony/framework-standard-edition my_project_name "2.3.*"
````

#### The Directory Structure

````
app/                         - application configuration.
app/cache/{environment}/twig - twig template caching.
src/                         - all the project PHP code.
vendor/                      - any vendor libraries.
web/                         - web root directory, contains publicly accessible files.

Bundle Directory Structure:
Controller/                 - controllers.
DependencyInjection/        - dependency injection extension classes.
Resources/config/           - configuration, routing.
Resources/doc/              - index.rst.
Resources/meta/             - LICENSE.
Resources/views/            - templates.
Resources/translations/     - translations.
Resources/public/           - images, stylesheets, etc.
web/                        - assets.
Tests/                      - tests for the bundle.

// Override default directory structure
class AppKernel extends Kernel
{
    // Override the cache Directory
    public function getCacheDir()
    {
        return $this->rootDir.'/'.$this->environment.'/cache';
    }

    // Override the logs Directory
    public function getLogDir()
    {
       return $this->rootDir.'/'.$this->environment.'/logs';
    }
}
````

#### Application

````php
// Checking Symfony Application Configuration and Setup.
http://localhost:8000/config.php

// Check whether your project's dependencies contain any know security vulnerability.
php bin/console security:check
// Clear your cache.
php bin/console cache:clear --env=prod --no-debug
php bin/console cache:clear -e prod

php bin/console list --no-debug

// interactive mode
php bin/console --shell
php bin/console -s

php bin/console --shell --process-isolation
php bin/console -s --process-isolation

//  Create the Bundle.
php bin/console generate:bundle --namespace=Acme/DemoBundle --format=yml
php bin/console generate:bundle --namespace=Acme/TestBundle

````

#### Deploy

````
php app/check.php
composer install --no-dev --optimize-autoloader
php bin/console cache:clear --env=prod --no-debug
php bin/console assetic:dump --env=prod --no-debug
````

#### Configuration

````php
// Configuration
# app/config/config.yml
parameters:
    acme_hello.email.from: fabien@example.com

$container->getParameter('acme_hello.email.from');

// Services & Configuration
# app/config/config.yml
parameters:
    translator.class: Acme\HelloBundle\Translation\Translator

// src/Acme/DemoBundle/DependencyInjection/Compiler/OverrideServiceCompilerPass.php
namespace Acme\DemoBundle\DependencyInjection\Compiler;
use Symfony\Component\DependencyInjection\Compiler\CompilerPassInterface;
use Symfony\Component\DependencyInjection\ContainerBuilder;
class OverrideServiceCompilerPass implements CompilerPassInterface
{
    public function process(ContainerBuilder $container)
    {
        $definition = $container->getDefinition('original-service-id');
        $definition->setClass('Acme\DemoBundle\YourService');
    }
}

// External parameters in the service container
export SYMFONY__DATABASE__USER=user
export SYMFONY__DATABASE__PASSWORD=secre
````

#### Service Container

Service Container works as fast with huge services.yml file as with tiny one.
Same about memory usage.
Proved by benchmarking.

````php
# app/config/config.yml
services:
    my_mailer:
        class: Acme\HelloBundle\Mailer
        arguments: [sendmail]
        arguments: ["@=service('mailer_configuration').getMailerMethod()"]
        arguments: ["@=container.hasParameter('some_param') ? parameter('some_param') : 'default_value'"]
        calls:
            - [setMailer, ["@my_mailer"]]

public function sendEmailAction()
{
    $mailer = $this->get('my_mailer');
    $mailer->send('ryan@foobar.net', ...);
    // or
    global $kernel;
    var_dump($kernel->getContainer()->get('aws.s3'));
}

# app/config/config.yml
parameters:
    my_mailer.transport: sendmail
    mailer_password: "@@securepass"
services:
    my_mailer:
        class: Acme\HelloBundle\Mailer
        arguments: ["%my_mailer.transport%"]

# app/config/config.yml
imports:
    - { resource: "@AcmeHelloBundle/Resources/config/services.yml" }
    - { resource: "%kernel.root_dir%/parameters.yml" }

# app/config/config.yml
framework:
    secret: xxxxxxxxxx
    form: true
    csrf_protection: true
    router: { resource: "%kernel.root_dir%/config/routing.yml" }

// Debugging Services
php bin/console debug:container
php bin/console debug:container --show-private
php bin/console debug:container my_mailer
````

#### Performance

````php
composer dump-autoload --optimize
````
