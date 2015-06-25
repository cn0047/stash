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

php app/console server:run
php app/console server:stop

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
php app/console security:check
// Clear your cache.
php app/console cache:clear --env=prod --no-debug
php app/console cache:clear -e prod

php app/console list --no-debug

// interactive mode
php app/console --shell
php app/console -s

php app/console --shell --process-isolation
php app/console -s --process-isolation

//  Create the Bundle.
php app/console generate:bundle --namespace=Acme/DemoBundle --format=yml
php app/console generate:bundle --namespace=Acme/TestBundle

````

#### Deploy
````
php app/check.php
composer install --no-dev --optimize-autoloader
php app/console cache:clear --env=prod --no-debug
php app/console assetic:dump --env=prod --no-debug
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

doctrine:
    dbal:
        driver pdo_mysql
        dbname: symfony_project
        user: "%database.user%"
        password: "%database.password%"
````

#### Request
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

// is it an Ajax request?
$request->isXmlHttpRequest();
// get a $_GET parameter
$request->query->get('page');
// get a $_POST parameter
$request->request->get('page');

$request->getPreferredLanguage(array('en', 'fr'));

$request->headers->get("User-Agent");
````

#### Response
````php
use Symfony\Component\HttpFoundation\Response;
$response = new Response();
$response->setContent('<html><body><h1>Hello world!</h1></body></html>');
$response->setStatusCode(Response::HTTP_OK);
$response->headers->set('Content-Type', 'text/html');
// prints the HTTP headers followed by the content
$response->send();

// create a JSON-response with a 200 status code
$response = new Response(json_encode(array('name' => $name)));
$response->headers->set('Content-Type', 'application/json');
````

#### Routing
````
Special Routing Parameters:
_controller - which controller is executed
_format     - format
_locale     - locale

// Visualizing & Debugging Routes.
php app/console debug:router
php app/console debug:router article_show
php app/console router:match /blog/my-latest-post
````

#### Controller
````php
// Redirecting.
$this->redirectToRoute('homepage');
$this->redirectToRoute('homepage', array(), 301);
$this->redirect('http://symfony.com/doc');
new RedirectResponse($this->generateUrl('homepage'));

// Forwarding to Another Controller.
$response = $this->forward('AppBundle:Something:fancy', array('name' => $name, 'color' => 'green'));

// Accessing other Services.
$templating = $this->get('templating');
$router = $this->get('router');
$mailer = $this->get('mailer');

// Managing Errors and 404 Pages.
throw $this->createNotFoundException('The product does not exist');

// Managing the Session.
$session = $request->getSession();
$session->set('foo', 'bar');
$foobar = $session->get('foobar');

// Generating URLs.
$uri = $this->get('router')->generate('blog_show', array('slug' => 'my-blog-post'));
$url = $this->generateUrl('blog_show', array('slug' => 'my-blog-post'));
$url = $this->container->get('router')->generate('blog_show', array('slug' => 'my-blog-post'));
var url = Routing.generate('blog_show', {"slug": 'my-blog-post'});
// Generating Absolute URLs.
$this->generateUrl('blog_show', array('slug' => 'my-blog-post'), true);
// Generating URLs from a Template.
<a href="{{ path('blog_show', {'slug': 'my-blog-post'}) }}"> Read this blog post. </a>
<a href="{{ url('blog_show', {'slug': 'my-blog-post'}) }}"> Read this blog post. </a>

// Debug.
dump($articles);

// Overriding Controllers
# src/Acme/UserBundle/Controller/RegistrationController.php
namespace Acme\UserBundle\Controller;
use FOS\UserBundle\Controller\RegistrationController as BaseController;
class RegistrationController extends BaseController
{
    public function registerAction()
    {
        $response = parent::registerAction();
        // ... do custom stuff
        return $response;
    }
}
````

#### Templates
````php
// Symfony actually looks in two different locations for the template:
1. app/Resources/AcmeBlogBundle/views/Blog/index.html.twig
2. src/Acme/BlogBundle/Resources/views/Blog/index.html.twig

// Output Escaping
// If you're using twig templates, then output escaping is ON by default.
{{ article.body|raw }} - render normally
Hello <?php echo $view->escape($name) ?>
var myMsg = 'Hello <?php echo $view->escape($name, 'js') ?>';

// Syntax Checking.
php app/console twig:lint app/Resources/views/article/recent_list.html.twig
php app/console twig:lint app/Resources/views

// Debug.
{{ dump(articles) }}

// Global Template Variables.
app.security
app.user
app.request
app.session
app.environment
app.debug

{{ ... }}         - variable or the result of an expression
{% ... %}         - a tag that controls the logic of the template
{# ... #}         - comment
{{ title|upper }} - filters
{{ parent() }}    - block from the parent template

{{ include('article/article_details.html.twig', { 'article': article }) }}

// Links.
<a href="{{ path('_welcome') }}">Home</a>
<a href="{{ path('article_show', {'slug': article.slug}) }}"> {{ article.title }} </a>
// Absolute URL.
<a href="{{ url('_welcome') }}">Home</a>
<a href="<?php echo $view['router']->generate('_welcome', array(), true) ?>">Home</a>

// Assets.
<img src="{{ asset('images/logo.png') }}" alt="Symfony!" />
<link href="{{ asset('css/blog.css') }}" rel="stylesheet" type="text/css" />
<img src="{{ asset('images/logo.png', version='3.0') }}" alt="Symfony!" />
<img src="{{ asset('images/logo.png', absolute=true) }}" alt="Symfony!" />

// Stylesheets and JavaScripts
{% block stylesheets %}
    <link href="{{ asset('css/main.css') }}" rel="stylesheet" />
{% endblock %}
{% block javascripts %}
    <script src="{{ asset('js/main.js') }}"></script>
{% endblock %}
<link href="{{ asset('bundles/acmedemo/css/contact.css') }}" rel="stylesheet" />

<script src="{{ asset('js/script.js') }}"></script>

// Dumping Asset Files
php app/console assetic:dump --env=prod --no-debug
php app/console assetic:dump
php app/console assetic:watch

````

Template Suffix:

| Filename | Format | Engine |
|--------|------|------|
|blog/index.html.twig|HTML|Twig|
|blog/index.html.php|HTML|PHP|
|blog/index.css.twig|CSS|Twig|

#### Doctrine
````php
php app/console doctrine:database:create
php app/console doctrine:database:drop --force
php app/console doctrine:generate:entity
php app/console doctrine:generate:entities AppBundle/Entity/Product
php app/console doctrine:schema:update --force

// generates all entities in the AppBundle
php app/console doctrine:generate:entities AppBundle
// generates all entities of bundles in the Acme namespace
php app/console doctrine:generate:entities Acme

php app/console list doctrine
php app/console help doctrine:database:create
php app/console doctrine:ensure-production-settings --env=prod
````

#### Testing
````php
phpunit -c app/
phpunit -c app src/AppBundle/Tests/Util

$client = static::createClient();
$crawler = $client->request('GET', '/post/hello-world');
$client->request('POST', '/submit', array('name' => 'Fabien'));
$client->request(
    'POST',
    '/submit',
    array(),
    array(),
    array('CONTENT_TYPE' => 'application/json'),
    '{"name":"Fabien"}'
);

$client->request(
    $method,
    $uri,
    array $parameters = array(),
    array $files = array(),
    array $server = array(),
    $content = null,
    $changeHistory = true
);

$link = $crawler->selectLink('Go elsewhere...')->link();
$crawler = $client->click($link);

$form = $crawler->selectButton('validate')->form();
$crawler = $client->submit($form, array('name' => 'Fabien'));

$link = $crawler
    ->filter('a:contains("Greet")') // find all links with the text "Greet"
    ->eq(1) // select the second link in the list
    ->link() // and click it
    ;
$crawler = $client->click($link);

$form = $crawler->selectButton('submit')->form();
// set some values
$form['name'] = 'Lucas';
$form['form_name[subject]'] = 'Hey there!';
// submit the form
$crawler = $client->submit($form);

// Assert that the response matches a given CSS selector.
$this->assertGreaterThan(0, $crawler->filter('h1')->count());

$this->assertContains(
    'Hello World',
    $client->getResponse()->getContent()
);

// Assert that there is at least one h2 tag with the class "subtitle"
$this->assertGreaterThan(
    0,
    $crawler->filter('h2.subtitle')->count()
);

// Assert that there are exactly 4 h2 tags on the page
$this->assertCount(4, $crawler->filter('h2'));

// Assert that the "Content-Type" header is "application/json"
$this->assertTrue(
$client->getResponse()->headers->contains(
    'Content-Type',
    'application/json'
);

// Assert that the response content contains a string
$this->assertContains('foo', $client->getResponse()->getContent());
// ...or matches a regex
$this->assertRegExp('/foo(bar)?/', $client->getResponse()->getContent());

// Assert that the response status code is 2xx
$this->assertTrue($client->getResponse()->isSuccessful());
// Assert that the response status code is 404
$this->assertTrue($client->getResponse()->isNotFound());
// Assert a specific 200 status code
$this->assertEquals(
    200, // or Symfony\Component\HttpFoundation\Response::HTTP_OK
    $client->getResponse()->getStatusCode()
);

// Assert that the response is a redirect to /demo/contact
$this->assertTrue(
    $client->getResponse()->isRedirect('/demo/contact')
);
// ...or simply check that the response is a redirect to any URL
$this->assertTrue($client->getResponse()->isRedirect());

// Form submission with a file upload
use Symfony\Component\HttpFoundation\File\UploadedFile;
$photo = new UploadedFile(
    '/path/to/photo.jpg',
    'photo.jpg',
    'image/jpeg',
    123
);
$client->request(
    'POST',
    '/submit',
    array('name' => 'Fabien'),
    array('photo' => $photo)
);

// Perform a DELETE request and pass HTTP headers
$client->request(
    'DELETE',
    '/post/12',
    array(),
    array(),
    array('PHP_AUTH_USER' => 'username', 'PHP_AUTH_PW' => 'pa$$word')
);
````

#### Validation
````php
// Basic Constraints
• NotBlank
• Blank
• NotNull
• Null
• True
• False
• Type

// String Constraints
• Email
• Length
• Url
• Regex
• Ip
• Uuid

// Number Constraints
• Range

// Comparison Constraints
• EqualTo
• NotEqualTo
• IdenticalTo
• NotIdenticalTo
• LessThan
• LessThanOrEqual
• GreaterThan
• GreaterThanOrEqual

// Date Constraints
• Date
• DateTime
• Time

// Collection Constraints
• Choice
• Collection
• Count
• UniqueEntity
• Language
• Locale
• Country

// File Constraints
• File
• Image

// Financial and other Number Constraints
• CardScheme
• Currency
• Luhn
• Iban
• Isbn
• Issn

// Other Constraints
• Callback
• Expression
• All
• UserPassword
• Valid

/**
 * @Assert\Choice(
 * choices = { "male", "female" },
 * message = "Choose a valid gender."
 * )
 */
public $gender;

/**
 * @Assert\True(message = "The password cannot match your first name")
 */
public function isPasswordLegal() { return $this->firstName !== $this->password; }

// group
$errors = $validator->validate($author, null, array('registration'));
````