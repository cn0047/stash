Symfony
-
2.6.6

####Request
````php
use Symfony\Component\HttpFoundation\Request;
$request = Request::createFromGlobals();
// the URI being requested (e.g. /about) minus any query parameters
$request->getPathInfo();
// retrieve GET and POST variables respectively
$request->query->get('foo');
$request->request->get('bar', 'default value if bar does not exist');
// retrieve SERVER variables
$request->server->get('HTTP_HOST');
// retrieves an instance of UploadedFile identified by foo
$request->files->get('foo');
// retrieve a COOKIE value
$request->cookies->get('PHPSESSID');
// retrieve an HTTP request header, with normalized, lowercase keys
$request->headers->get('host');
$request->headers->get('content_type');
$request->getMethod(); // GET, POST, PUT, DELETE, HEAD
$request->getLanguages(); // an array of languages the client accepts
````

####Response
````php
use Symfony\Component\HttpFoundation\Response;
$response = new Response();
$response->setContent('<html><body><h1>Hello world!</h1></body></html>');
$response->setStatusCode(Response::HTTP_OK);
$response->headers->set('Content-Type', 'text/html');
// prints the HTTP headers followed by the content
$response->send();
````

####Installing the Symfony Installer
````
sudo curl -LsS http://symfony.com/installer -o /usr/local/bin/symfony
sudo chmod a+x /usr/local/bin/symfony
````

####Creating the Symfony Application
````
symfony new my_project_name
symfony new my_project_name 2.3.26
````

####Creating a Symfony Application with Composer
````
composer create-project symfony/framework-standard-edition my_project_name
composer create-project symfony/framework-standard-edition my_project_name "2.3.*"
````

####Running the Symfony Application
````
cd my_project_name/
php app/console server:run
php app/console server:stop

then http://localhost:8000
````

####Application
````
// Checking Symfony Application Configuration and Setup.
http://localhost:8000/config.php

// Check whether your project's dependencies contain any know security vulnerability.
php app/console security:check

// Clear your cache.
php app/console cache:clear --env=prod --no-debug

//  Create the Bundle.
php app/console generate:bundle --namespace=Acme/DemoBundle --format=yml
php app/console generate:bundle --namespace=Acme/TestBundle
````

####The Directory Structure
````
app/    - application configuration.
src/    - all the project PHP code.
vendor/ - any vendor libraries.
web/    - web root directory, contains publicly accessible files.

Bundle Directory Structure:
Controller/          - controllers.
DependencyInjection/ - dependency injection extension classes.
Resources/config/    - configuration, routing.
Resources/views/     - templates.
Resources/public/    - images, stylesheets, etc.
web/                 - assets.
Tests/               - tests for the bundle.
````