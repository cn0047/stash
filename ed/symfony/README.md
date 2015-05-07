Symfony
-
2.6.6


####Running the Symfony Application
````
cd my_project_name/
cd ed/symfony/examples/bulletinBoard/

php app/console server:run
php app/console server:stop

then http://localhost:8000
````

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

// is it an Ajax request?
$request->isXmlHttpRequest();
// get a $_GET parameter
$request->query->get('page');
// get a $_POST parameter
$request->request->get('page');

$request->getPreferredLanguage(array('en', 'fr'));

$request->headers->get("User-Agent");
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

// create a JSON-response with a 200 status code
$response = new Response(json_encode(array('name' => $name)));
$response->headers->set('Content-Type', 'application/json');
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
app/                         - application configuration.
app/cache/{environment}/twig - twig template caching.
src/                         - all the project PHP code.
vendor/                      - any vendor libraries.
web/                         - web root directory, contains publicly accessible files.

Bundle Directory Structure:
Controller/          - controllers.
DependencyInjection/ - dependency injection extension classes.
Resources/config/    - configuration, routing.
Resources/views/     - templates.
Resources/public/    - images, stylesheets, etc.
web/                 - assets.
Tests/               - tests for the bundle.
````

####Controller
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
````

####Routing
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

####Templates
````php
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

// Global Template Variables.
app.security
app.user
app.request
app.session
app.environment
app.debug

// Symfony actually looks in two different locations for the template:
1. app/Resources/AcmeBlogBundle/views/Blog/index.html.twig
2. src/Acme/BlogBundle/Resources/views/Blog/index.html.twig

// Output Escaping
// IF YOU'RE USING TWIG TEMPLATES, THEN OUTPUT ESCAPING IS ON BY DEFAULT.
{{ article.body|raw }} - render normally
Hello <?php echo $view->escape($name) ?>
var myMsg = 'Hello <?php echo $view->escape($name, 'js') ?>';

// Debug.
{{ dump(articles) }}

// Syntax Checking.
php app/console twig:lint app/Resources/views/article/recent_list.html.twig
php app/console twig:lint app/Resources/views

<ul>
    {% for user in users if user.active %}
        <li>{{ user.username }}</li>
    {% else %}
        <li>No users found</li>
    {% endfor %}
</ul>

<div id="sidebar">
    {{ render(controller(
        'AcmeArticleBundle:Article:recentArticles',
        { 'max': 3 }
    )) }}
</div>

<a href="{{ path('article_show', {'id': 123, '_format': 'pdf'}) }}"> PDF Version </a>
````

Template Suffix:

| Filename | Format | Engine |
|--------|------|------|
|blog/index.html.twig|HTML|Twig|
|blog/index.html.php|HTML|PHP|
|blog/index.css.twig|CSS|Twig|

####Doctrine
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

````

####Testing
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

// Browsing
$client->back();
$client->forward();
$client->reload();
// Clears all cookies and the history
$client->restart();

$history = $client->getHistory();
$cookieJar = $client->getCookieJar();

// the HttpKernel request instance
$request = $client->getRequest();
// the BrowserKit request instance
$request = $client->getInternalRequest();
// the HttpKernel response instance
$response = $client->getResponse();
// the BrowserKit response instance
$response = $client->getInternalResponse();
$crawler = $client->getCrawler();

$container = $client->getContainer();
$kernel = $client->getKernel();

// enable the profiler for the very next request
$client->enableProfiler();
$crawler = $client->request('GET', '/profiler');
// get the profile
$profile = $client->getProfile();

// Redirecting
$crawler = $client->followRedirect();

$client->followRedirects();

$client->followRedirects(false);

$newCrawler = $crawler->filter('input[type=submit]')
    ->last()
    ->parents()
    ->first()
    ;
// Many other methods are also available:
filter('h1.title') - Nodes that match the CSS selector.
filterXpath('h1')  - Nodes that match the XPath expression.
eq(1)              - Node for the specified index.
first()            - First node.
last()             - Last node.
siblings()         - Siblings.
nextAll()          - All following siblings.
previousAll()      - All preceding siblings.
parents()          - Returns the parent nodes.
children()         - Returns children nodes.
reduce($lambda)    - Nodes for which the callable does not return false.

$crawler
    ->filter('h1')
    ->reduce(function ($node, $i) {
        if (!$node->getAttribute('class')) {
            return false;
        }
    })
    ->first()
;

// Extracting Information
// Returns the attribute value for the first node
$crawler->attr('class');
// Returns the node value for the first node
$crawler->text();
// Extracts an array of attributes for all nodes
// (_text returns the node value)
// returns an array for each element in crawler,
// each with the value and href
$info = $crawler->extract(array('_text', 'href'));
// Executes a lambda for each node and return an array of results
$data = $crawler->each(function ($node, $i) {
    return $node->attr('href');
});

$crawler->selectLink('Click here');
$link = $crawler->selectLink('Click here')->link();
$client->click($link);

$buttonCrawlerNode = $crawler->selectButton('submit');
$form = $buttonCrawlerNode->form();
$form = $buttonCrawlerNode->form(array(
    'name' => 'Fabien',
    'my_form[subject]' => 'Symfony rocks!',
));
$form = $buttonCrawlerNode->form(array(), 'DELETE');
$client->submit($form);
$client->submit($form, array(
    'name' => 'Fabien',
    'my_form[subject]' => 'Symfony rocks!',
));

// Change the value of a field
$form['name'] = 'Fabien';
$form['my_form[subject]'] = 'Symfony rocks!';
// Select an option or a radio
$form['country']->select('France');
// Tick a checkbox
$form['like_symfony']->tick();
// Upload a file
$form['photo']->upload('/path/to/lucas.jpg');

$client = static::createClient(array(
    'environment' => 'my_test_env',
    'debug' => false,
));
$client = static::createClient(array(), array(
    'HTTP_HOST' => 'en.example.com',
    'HTTP_USER_AGENT' => 'MySuperBrowser/1.0',
));
$client->request('GET', '/', array(), array(), array(
    'HTTP_HOST' => 'en.example.com',
    'HTTP_USER_AGENT' => 'MySuperBrowser/1.0',
));
````

####Validation
````php
use Symfony\Component\Validator\Constraints as Assert;
/**
 * @Assert\NotBlank()
 * @Assert\Length(min=3)
 */
public $name;

$validator = $this->get('validator');
$errors = $validator->validate($author);

$author = new Author();
$form = $this->createForm(new AuthorType(), $author);
$form->handleRequest($request);
if ($form->isValid()) {}

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

####Forms
````
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
page:181
