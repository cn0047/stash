####Forms
````php
{{ form(form, {'attr': {'novalidate': 'novalidate'}}) }}

$form = $this->createFormBuilder($users, array(
    'validation_groups' => array('registration'),
))->add(...);

// Validation Groups
use Symfony\Component\OptionsResolver\OptionsResolverInterface;
public function setDefaultOptions(OptionsResolverInterface $resolver)
{
    $resolver->setDefaults(array(
       'validation_groups' => array('registration'),
    ));
    // Disabling Validation
    $resolver->setDefaults(array(
        'validation_groups' => false,
    ));
    // Groups based on the Submitted Data
    $resolver->setDefaults(array(
        'validation_groups' => array(
            'AppBundle\Entity\Client',
            'determineValidationGroups',
        )
    ));
    $resolver->setDefaults(array(
        'validation_groups' => function(FormInterface $form) {
            $data = $form->getData();
            if (Client::TYPE_PERSON == $data->getType()) {
                return array('person');
            }
            return array('company');
        },
    ));
    $resolver->setDefaults(array(
        'validation_groups' => function(FormInterface $form) {
            $data = $form->getData();
            if (Client::TYPE_PERSON == $data->getType()) {
              return array('Default', 'person');
            }
            return array('Default', 'company');
        },
    ));
}

$form = $this->createFormBuilder($task)
    ->add('nextStep', 'submit')
    ->add('previousStep', 'submit')
    ->getForm();

$form = $this->createFormBuilder($task)
    ->add('previousStep', 'submit', array(
        'validation_groups' => false,
    ))
    ->getForm();

// Built-in Field Types:
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

->add('dueDate', 'date', array('widget' => 'single_text'))
->add('dueDate', 'date', array(
    'widget' => 'single_text',
    'label' => 'Due Date',
))

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

// Rendering a Form in a Template
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

if ($form->isValid()) {
    $em = $this->getDoctrine()->getManager();
    $em->persist($task);
    $em->flush();
    return $this->redirectToRoute('task_success');
}

// Global Form Theming
# app/config/config.yml
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

// To automatically include the customized templates from the app/Resources/views/Form
# app/config/config.yml
framework:
    templating:
        form:
            resources:
                - 'Form'

// CSRF Protection
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

// Using a Form without a Class
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
````

####Security
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

{% if is_granted('ROLE_ADMIN') %}
    <a href="...">Delete</a>
{% endif %}

{% if app.user and is_granted('ROLE_ADMIN') %}

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

{% if is_granted(expression('"ROLE_ADMIN" in roles or (user and user.isSuperAdmin())')) %}
    <a href="...">Delete</a>
{% endif %}

{% if is_granted('IS_AUTHENTICATED_FULLY') %}
    <p>Username: {{ app.user.username }}</p>
{% endif %}

// Dynamically Encoding a Password
$user = new AppBundle\Entity\User();
$plainPassword = 'ryanpass';
$encoder = $this->container->get('security.password_encoder');
$encoded = $encoder->encodePassword($user, $plainPassword);
$user->setPassword($encoded);

// Hierarchical Roles
security:
    role_hierarchy:
        ROLE_ADMIN: ROLE_USER
        ROLE_SUPER_ADMIN: [ROLE_ADMIN, ROLE_ALLOWED_TO_SWITCH]

// Checking for Known Security Vulnerabilities in Dependencies
 php app/console security:check
````

####HTTP Cache
````php
// Symfony Reverse Proxy
require_once __DIR__.'/../app/bootstrap.php.cache';
require_once __DIR__.'/../app/AppKernel.php';
require_once __DIR__.'/../app/AppCache.php';
use Symfony\Component\HttpFoundation\Request;
$kernel = new AppKernel('prod', false);
$kernel->loadClassCache();
// wrap the default AppKernel with the AppCache one
$kernel = new AppCache($kernel);
$request = Request::createFromGlobals();
$response = $kernel->handle($request);
$response->send();
$kernel->terminate($request, $response);

// app/AppCache.php
use Symfony\Bundle\FrameworkBundle\HttpCache\HttpCache;
class AppCache extends HttpCache
{
    protected function getOptions()
    {
        return array(
            'debug' => false,
            'default_ttl' => 0,
            'private_headers' => array('Authorization', 'Cookie'),
            'allow_reload' => false,
            'allow_revalidate' => false,
            'stale_while_revalidate' => 2,
            'stale_if_error' => 60,
        );
    }
}

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
// ...
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
// ...
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

// Varying the Response
// set one vary header
$response->setVary('Accept-Encoding');
// set multiple vary headers
$response->setVary(array('Accept-Encoding', 'User-Agent'));

// Marks the Response stale
$response->expire();

// Force the response to return a proper 304 response with no content
$response->setNotModified();

// Set cache settings in one call
$response->setCache(array(
    'etag' => $etag,
    'last_modified' => $date,
    'max_age' => 10,
    's_maxage' => 10,
    'public' => true,
    // 'private' => true,
));

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

####Translations
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
####Service Container
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
    // ...
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

####Performance
````php
composer dump-autoload --optimize
````

####Internals
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

####Environments
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

page:210
