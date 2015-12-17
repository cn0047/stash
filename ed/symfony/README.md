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
//
$Request->query->has('foo');
// retrieve GET and POST variables respectively
$request->query->get('foo');
$request->request->get('bar', 'default value if bar does not exist');
$request->get('boo');
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

#### Forms
````php
{{ form(form, {'attr': {'novalidate': 'novalidate'}}) }}

$form = $this->createFormBuilder($users, array(
    'validation_groups' => array('registration'),
))->add(...);

$form = $this->createFormBuilder($task)
    ->add('nextStep', 'submit')
    ->add('previousStep', 'submit')
    ->add('dueDate', 'date', array('widget' => 'single_text'))
    ->add('dueDate', 'date', array(
        'widget' => 'single_text',
        'label' => 'Due Date',
    ))
    ->getForm();

$form = $this->createFormBuilder($task)
    ->setAction($this->generateUrl('target_route'))
    ->setMethod('GET')
    ->add('task', 'text')
    ->add('dueDate', 'date')
    ->add('save', 'submit')
    ->getForm();

$form = $this->createForm(new TaskType(), $task, array(
    'action' => $this->generateUrl('target_route'),
    'method' => 'GET',
));

if ($form->isValid()) {
    $em = $this->getDoctrine()->getManager();
    $em->persist($task);
    $em->flush();
    return $this->redirectToRoute('task_success');
}

# CSRF Protection
use Symfony\Component\OptionsResolver\OptionsResolverInterface;
class TaskType extends AbstractType
{
    public function setDefaultOptions(OptionsResolverInterface $resolver)
    {
        $resolver->setDefaults(array(
            'data_class' => 'AppBundle\Entity\Task',
            'csrf_protection' => true,
            'csrf_field_name' => '_token',
            // a unique key to help generate the secret token
            'intention' => 'task_item',
        ));
    }
}

# Using a Form without a Class
use Symfony\Component\HttpFoundation\Request;
// make sure you've imported the Request namespace above the class
public function contactAction(Request $request)
{
    $defaultData = array('message' => 'Type your message here');
    $form = $this->createFormBuilder($defaultData)
        ->add('name', 'text')
        ->add('email', 'email')
        ->add('message', 'textarea')
        ->add('send', 'submit')
        ->getForm();
    $form->handleRequest($request);
    if ($form->isValid()) {
        // data is an array with "name", "email", and "message" keys
        $data = $form->getData();
    }
}

# Built-in Field Types:
// Text Fields
• text
• textarea
• email
• integer
• money
• number
• password
• percent
• search
• url

// Choice Fields
• choice
• entity
• country
• language
• locale
• timezone
• currency

// Date and Time Fields
• date
• datetime
• time
• birthday

// Other Fields
• checkbox
• file
• radio

// Field Groups
• collection
• repeated

// Hidden Fields
• hidden

// Buttons
• button
• reset
• submit

// Base Fields
• form

# Rendering a Form in a Template
{{ form_start(form) }}
    {{ form_errors(form) }}
    {{ form_row(form.task) }}
    {{ form_row(form.dueDate) }}
    {{ form.vars.value.task }}
{{ form_end(form) }}

{{ form_start(form) }}
    {{ form_errors(form) }}
    <div>
        {{ form_label(form.task) }}
        {{ form_errors(form.task) }}
        {{ form_widget(form.task) }}
    </div>
    <div>
        {{ form_label(form.dueDate) }}
        {{ form_errors(form.dueDate) }}
        {{ form_widget(form.dueDate) }}
    </div>
    <div>
        {{ form_widget(form.save) }}
    </div>
{{ form_end(form) }}

{{ form_label(form.task, 'Task Description') }}

{{ form_widget(form.task, {'attr': {'class': 'task_field'}}) }}

{{ form.task.vars.id }}
{{ form.task.vars.full_name }}

{{ form_start(form, {'action': path('target_route'), 'method': 'GET'}) }}

$form->get('dueDate')->getData();
$form->get('dueDate')->setData(new \DateTime());

# Global Form Theming
// app/config/config.yml
twig:
    form_themes:
        - 'form/fields.html.twig'

// Customizing Form Output all in a Single File with Twig
{% extends 'base.html.twig' %}

{# import "_self" as the form theme #}
{% form_theme form _self %}

{# make the form fragment customization #}
{% block form_row %}
    {# custom field row output #}
{% endblock form_row %}

{% block content %}
    {{ form_row(form.task) }}
{% endblock %}
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

{% if items is not empty %}
{% for item in items  %}
    {{ item.tariffNum }}
{% endfor %}
{% endif %}

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
{% set seriesSendTimeUtc = (series.send_time_utc is defined) ? series.send_time_utc :  null %}
````
````
// Assets
app/console assets:install web/

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

#### HTTP Cache
````php
// Set cache settings in one call
$response->setCache(array(
    'etag' => $etag,
    'last_modified' => $date,
    'max_age' => 10,
    's_maxage' => 10,
    'public' => true,
    // 'private' => true,
));

// The Cache-Control Header
use Symfony\Component\HttpFoundation\Response;
$response = new Response();
// mark the response as either public or private
$response->setPublic();
$response->setPrivate();
// set the private or shared max age
$response->setMaxAge(600);
$response->setSharedMaxAge(600);
// set a custom Cache-Control directive
$response->headers->addCacheControlDirective('must-revalidate', true);

// Expiration with the Expires Header
$date = new DateTime();
$date->modify('+600 seconds');
$response->setExpires($date);

// Expiration with the Cache-Control Header
$response->setMaxAge(600);
// Same as above but only for shared caches
$response->setSharedMaxAge(600);

// Varying the Response
// set one vary header
$response->setVary('Accept-Encoding');
// set multiple vary headers
$response->setVary(array('Accept-Encoding', 'User-Agent'));

// Marks the Response stale
$response->expire();

// Force the response to return a proper 304 response with no content
$response->setNotModified();

// Validation with the ETag Header
// src/AppBundle/Controller/DefaultController.php
namespace AppBundle\Controller;
use Symfony\Component\HttpFoundation\Request;
class DefaultController extends Controller
{
    public function homepageAction(Request $request)
    {
        $response = $this->render('static/homepage.html.twig');
        $response->setETag(md5($response->getContent()));
        $response->setPublic(); // make sure the response is public/cacheable
        $response->isNotModified($request);
        return $response;
    }
}

// Validation with the Last-Modified Header
// src/AppBundle/Controller/ArticleController.php
namespace AppBundle\Controller;
use Symfony\Component\HttpFoundation\Request;
use AppBundle\Entity\Article;
class ArticleController extends Controller
{
    public function showAction(Article $article, Request $request)
    {
        $author = $article->getAuthor();
        $articleDate = new \DateTime($article->getUpdatedAt());
        $authorDate = new \DateTime($author->getUpdatedAt());
        $date = $authorDate > $articleDate ? $authorDate : $articleDate;
        $response->setLastModified($date);
        // Set response as public. Otherwise it will be private by default.
        $response->setPublic();
        if ($response->isNotModified($request)) {
            return $response;
        }
        // ... do more work to populate the response with the full content
        return $response;
    }
}

// Optimizing your Code with Validation
// src/AppBundle/Controller/ArticleController.php
namespace AppBundle\Controller;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\HttpFoundation\Request;
class ArticleController extends Controller
{
    public function showAction($articleSlug, Request $request)
    {
        // Get the minimum information to compute
        // the ETag or the Last-Modified value
        // (based on the Request, data is retrieved from
        // a database or a key-value store for instance)
        $article = ...;
        // create a Response with an ETag and/or a Last-Modified header
        $response = new Response();
        $response->setETag($article->computeETag());
        $response->setLastModified($article->getPublishedAt());
        // Set response as public. Otherwise it will be private by default.
        $response->setPublic();
        // Check that the Response is not modified for the given Request
        if ($response->isNotModified($request)) {
            // return the 304 Response immediately
            return $response;
        }
        // do more work here - like retrieving more data
        $comments = ...;
        // or render a template with the $response you've already started
        return $this->render('article/show.html.twig', array(
            'article' => $article,
            'comments' => $comments
        ), $response);
    }
}

// Using ESI in Symfony
# app/config/config.yml
framework:
    esi: { enabled: true }

public function aboutAction()
{
    $response = $this->render('static/about.html.twig');
    // set the shared max age - which also marks the response as public
    $response->setSharedMaxAge(600);
    return $response;
}

{# app/Resources/views/static/about.html.twig #}
{# you can use a controller reference #}
{{ render_esi(controller('AppBundle:News:latest', { 'maxPerPage': 5 })) }}
{# ... or a URL #}
{{ render_esi(url('latest_news', { 'maxPerPage': 5 })) }}
````

#### Translations
````php
echo $translator->trans('Hello World');

# app/config/config.yml
framework:
    translator: { fallbacks: [en] }
    default_locale: en

# app/config/routing.yml
contact:
    path: /{_locale}/contact
        defaults: { _controller: AppBundle:Contact:index }
        requirements:
            _locale: en|fr|de

public function indexAction()
{
    $locale = $request->getLocale();
    $request->getSession()->set('_locale', $locale);
    $translated = $this->get('translator')->trans('Symfony is great');
    return new Response($translated);
}

<!-- messages.fr.xlf -->
<?xml version="1.0"?>
<xliff version="1.2" xmlns="urn:oasis:names:tc:xliff:document:1.2">
    <file source-language="en" datatype="plaintext" original="file.ext">
        <body>
            <trans-unit id="1">
                <source>Symfony is great</source>
                <target>J'aime Symfony</target>
            </trans-unit>
        </body>
    </file>
</xliff>

{% trans %}Hello %name%{% endtrans %}
{% transchoice count %}
    {0} There are no apples|{1} There is one apple|]1,Inf] There are %count% apples
{% endtranschoice %}

{% trans with {'%name%': 'Fabien'} from "app" %}Hello %name%{% endtrans %}
{% trans with {'%name%': 'Fabien'} from "app" into "fr" %}Hello %name%{% endtrans %}
{% transchoice count with {'%name%': 'Fabien'} from "app" %}
    {0} %name%, there are no apples|{1} %name%, there is one apple|]1,Inf] %name%,
    there are %count% apples
{% endtranschoice %}

{{ message|trans }}
{{ message|transchoice(5) }}
{{ message|trans({'%name%': 'Fabien'}, "app") }}
{{ message|transchoice(5, {'%name%': 'Fabien'}, 'app') }}

{# text translated between tags is never escaped #}
{% trans %}
    <h3>foo</h3>
{% endtrans %}
{% set message = '<h3>foo</h3>' %}
{# strings and variables translated via a filter are escaped by default #}
{{ message|trans|raw }}
{{ '<h3>bar</h3>'|trans|raw }}

{% trans_default_domain "app" %}

<?php echo $view['translator']->trans('Symfony is great') ?>
<?php echo $view['translator']->transChoice(
    '{0} There are no apples|{1} There is one apple|]1,Inf[ There are %count% apples',
    10,
    array('%count%' => 10)
) ?>

// Each time you create a new translation resource
php app/console cache:clear

class Author
{
    /**
    * @Assert\NotBlank(message = "author.name.not_blank")
    */
    public $name;
}

{% trans %}Symfony2 is great{% endtrans %}
{{ 'Symfony2 is great'|trans }}
{{ 'Symfony2 is great'|transchoice(1) }}
{% transchoice 1 %}Symfony2 is great{% endtranschoice %}

$view['translator']->trans("Symfony2 is great");
$view['translator']->transChoice('Symfony2 is great', 1);

{% set message = 'Symfony2 is great' %}
{{ message|trans }}

// To inspect all messages in the fr locale for the AcmeDemoBundle, run:
php app/console debug:translation fr AcmeDemoBundle
php app/console debug:translation en AcmeDemoBundle --domain=messages

php app/console debug:translation en AcmeDemoBundle --only-unused
php app/console debug:translation en AcmeDemoBundle --only-missing
````

#### Doctrine
````php
// Generate Entities from an Existing Database !!!
php app/console doctrine:mapping:import --force AppBundle xml
php app/console doctrine:mapping:convert annotation ./src
php app/console doctrine:generate:entities AppBundle

php app/console doctrine:database:create
php app/console doctrine:database:drop --force
php app/console doctrine:generate:entity
php app/console doctrine:generate:entities AppBundle/Entity/Product
// update db schema from entities
php app/console doctrine:schema:update --force

// generates all entities in the AppBundle
php app/console doctrine:generate:entities AppBundle
// generates all entities of bundles in the Acme namespace
php app/console doctrine:generate:entities Acme

php app/console list doctrine
php app/console help doctrine:database:create
php app/console doctrine:ensure-production-settings --env=prod

// Generate CRUD !!!
php app/console generate:doctrine:crud --entity=AppBundle:EntityName

// in controller
$post=$this->get('doctrine')->getManager()->getRepository('AppBundle:Post')->find($id);
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

$errors = $this->get('validator')->validate(
    $request->get('id'),
    new \Symfony\Component\Validator\Constraints\Type(['type' => 'digit'])
);
````

#### Security
````php
# app/config/security.yml
security:
    providers:
        in_memory:
            memory: ~
    firewalls:
        dev:
            pattern: ^/(_(profiler|wdt)|css|images|js)/
            security: false
        default:
            anonymous: ~
            http_basic: ~

public function helloAction($name)
{
    // The second parameter is used to specify on what object the role is tested.
    $this->denyAccessUnlessGranted('ROLE_ADMIN', null, 'Unable to access this page!');
}

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Security;
/**
 * @Security("has_role('ROLE_ADMIN')")
 */
public function helloAction($name)
{
}

public function helloAction($name)
{
    if (!$this->get('security.authorization_checker')->isGranted('IS_AUTHENTICATED_FULLY')) {
        throw $this->createAccessDeniedException();
    }
    $user = $this->getUser();
    // the above is a shortcut for this
    $user = $this->get('security.token_storage')->getToken()->getUser();
    return new Response('Well hi there '.$user->getFirstName());
}

{% if is_granted('ROLE_ADMIN') %}
    <a href="...">Delete</a>
{% endif %}

{% if app.user and is_granted('ROLE_ADMIN') %}

{% if is_granted(expression('"ROLE_ADMIN" in roles or (user and user.isSuperAdmin())')) %}
    <a href="...">Delete</a>
{% endif %}

{% if is_granted('IS_AUTHENTICATED_FULLY') %}
    <p>Username: {{ app.user.username }}</p>
{% endif %}

# Hierarchical Roles
security:
    role_hierarchy:
        ROLE_ADMIN: ROLE_USER
        ROLE_SUPER_ADMIN: [ROLE_ADMIN, ROLE_ALLOWED_TO_SWITCH]
````

#### Service Container
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
php app/console debug:container
php app/console debug:container --show-private
php app/console debug:container my_mailer
````

#### Environments
````php
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
php app/console command_name --env=prod
# 'test' environment and debug disabled
php app/console command_name --env=test --no-debug

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

#### Performance
````php
composer dump-autoload --optimize
````

#### Internals
````php
$profile = $container->get('profiler')->loadProfileFromResponse($response);
$profile = $container->get('profiler')->loadProfile($token);
// get the latest 10 tokens
$tokens = $container->get('profiler')->find('', '', 10, '', '');
// get the latest 10 tokens for all URL containing /admin/
$tokens = $container->get('profiler')->find('', '/admin/', 10, '', '');
// get the latest 10 tokens for local requests
$tokens = $container->get('profiler')->find('127.0.0.1', '', 10, '', '');
// get the latest 10 tokens for requests that happened between 2 and 4 days ago
$tokens = $container->get('profiler')->find('', '', 10, '4 days ago', '2 days ago');

// on the production machine
$profile = $container->get('profiler')->loadProfile($token);
$data = $profiler->export($profile);
// on the development machine
$profiler->import($data);

// Configuration
# load the profiler
framework:
    profiler: { only_exceptions: false }
# enable the web profiler
web_profiler:
    toolbar: true
    intercept_redirects: true

_profiler:
    resource: "@WebProfilerBundle/Resources/config/routing/profiler.xml"
    prefix: /_profiler
````
