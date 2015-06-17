Symfony
-
2.6.6

[book](http://symfony.com/pdf/Symfony_book_2.6.pdf?v=4)
|
[cookbook](http://symfony.com/pdf/Symfony_cookbook_2.6.pdf?v=4)

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
````php
// Checking Symfony Application Configuration and Setup.
http://localhost:8000/config.php

// Check whether your project's dependencies contain any know security vulnerability.
php app/console security:check

// Clear your cache.
php app/console cache:clear --env=prod --no-debug

//  Create the Bundle.
php app/console generate:bundle --namespace=Acme/DemoBundle --format=yml
php app/console generate:bundle --namespace=Acme/TestBundle

// How to Install 3rd Party Bundles
• A) Add Composer Dependencies
    composer require friendsofsymfony/user-bundle
• B) Enable the Bundle
    // app/AppKernel.php
    class AppKernel extends Kernel
    {
        public function registerBundles()
        {
            $bundles = array(
                new FOS\UserBundle\FOSUserBundle(),
            );
        }
    }
• C) Configure the Bundle
    app/console config:dump-reference AsseticBundle
    app/console config:dump-reference assetic

// Bundle Name
Namespace                     | Bundle Class Name
------------------------------+----------------------
Acme\Bundle\BlogBundle        | AcmeBlogBundle
Acme\Bundle\Social\BlogBundle | AcmeSocialBlogBundle
Acme\BlogBundle               | AcmeBlogBundle

// Configuration
# app/config/config.yml
parameters:
    acme_hello.email.from: fabien@example.com

$container->getParameter('acme_hello.email.from');

// Custom Validation Constraints
use Symfony\Component\Validator\ConstraintValidator;
use Symfony\Component\Validator\Constraint;
use Symfony\Component\Validator\Context\ExecutionContextInterface;
class ContainsAlphanumericValidator extends ConstraintValidator
{
    public function validate($value, Constraint $constraint)
    {
        if ($this->context instanceof ExecutionContextInterface) {
            // the 2.5 API
            $this->context->buildViolation($constraint->message)
                ->setParameter('%string%', $value)
                ->addViolation()
            ;
        } else {
            // the 2.4 API
            $this->context->addViolation(
                $constraint->message,
                array('%string%' => $value)
            );
        }
    }
}

// How to Use Bundle Inheritance to Override Parts of a Bundle
# src/Acme/UserBundle/AcmeUserBundle.php
namespace Acme\UserBundle;
use Symfony\Component\HttpKernel\Bundle\Bundle;
class AcmeUserBundle extends Bundle
{
    public function getParent()
    {
        return 'FOSUserBundle';
    }
}

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

// How to Override any Part of a Bundle
// Services & Configuration
# app/config/config.yml
parameters:
    translator.class: Acme\HelloBundle\Translation\Translator

# src/Acme/DemoBundle/DependencyInjection/Compiler/OverrideServiceCompilerPass.php
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

// Validation Metadata
# src/Acme/UserBundle/Resources/config/validation.yml
FOS\UserBundle\Model\User:
    properties:
        plainPassword:
            - NotBlank:
                groups: [AcmeValidation]
            - Length:
                min: 6
                minMessage: fos_user.password.short
                groups: [AcmeValidation]

// How to Remove the AcmeDemoBundle
#  app/AppKernel.php
class AppKernel extends Kernel
{
    public function registerBundles()
    {
        $bundles = array(...);
            if (in_array($this->getEnvironment(), array('dev', 'test'))) {
            // comment or remove this line:
            // $bundles[] = new Acme\DemoBundle\AcmeDemoBundle();
        }
    }
}

// Remove Bundle Routing
in app/config/routing_dev.yml remove the _acme_demo

// Remove Bundle Configuration

// Remove the Bundle from the Filesystem
echo $this->container->get('kernel')->getBundle('AcmeDemoBundle')->getPath();

// How to Load Service Configuration inside a Bundle
# src/Acme/HelloBundle/DependencyInjection/AcmeHelloExtension.php
namespace Acme\HelloBundle\DependencyInjection;
use Symfony\Component\HttpKernel\DependencyInjection\Extension;
use Symfony\Component\DependencyInjection\ContainerBuilder;
class AcmeHelloExtension extends Extension
{
    public function load(array $configs, ContainerBuilder $container)
    {
        // ... you'll load the files here later
    }
}

// Manually Registering an Extension Class
use Acme\HelloBundle\DependencyInjection\UnconventionalExtensionClass;
class AcmeHelloBundle extends Bundle
{
    public function getContainerExtension()
    {
        return new UnconventionalExtensionClass();
    }
}

// Using the load() Method
use Symfony\Component\DependencyInjection\Loader\XmlFileLoader;
use Symfony\Component\Config\FileLocator;
// ...
public function load(array $configs, ContainerBuilder $container)
{
    $loader = new XmlFileLoader(
        $container,
        new FileLocator(__DIR__.'/../Resources/config')
    );
    $loader->load('services.xml');
}

// How to Create Friendly Configuration for a Bundle
# app/config/config.yml
acme_social:
    twitter:
        client_id: 123
        client_secret: $ecret

// Processing the $configs Array
array(
    // values from config.yml
    array(
        'twitter' => array(
            'client_id' => 123,
            'client_secret' => '$secret',
        ),
    ),
    // values from config_dev.yml
    array(
        'twitter' => array(
            'client_id' => 456,
        ),
    ),
)

// src/Acme/SocialBundle/DependencyInjection/Configuration.php
namespace Acme\SocialBundle\DependencyInjection;
use Symfony\Component\Config\Definition\Builder\TreeBuilder;
use Symfony\Component\Config\Definition\ConfigurationInterface;
class Configuration implements ConfigurationInterface
{
    public function getConfigTreeBuilder()
    {
        $treeBuilder = new TreeBuilder();
        $rootNode = $treeBuilder->root('acme_social');
        $rootNode
            ->children()
            ->arrayNode('twitter')
            ->children()
            ->integerNode('client_id')->end()
            ->scalarNode('client_secret')->end()
            ->end()
            ->end() // twitter
            ->end()
        ;
        return $treeBuilder;
    }
}

public function load(array $configs, ContainerBuilder $container)
{
    $configuration = new Configuration();
    $config = $this->processConfiguration($configuration, $configs);
}

#  src/Acme/HelloBundle/DependencyInjection/AcmeHelloExtension.php
namespace Acme\HelloBundle\DependencyInjection;
use Symfony\Component\DependencyInjection\ContainerBuilder;
use Symfony\Component\HttpKernel\DependencyInjection\ConfigurableExtension;
class AcmeHelloExtension extends ConfigurableExtension
{
    // note that this method is called loadInternal and not load
    protected function loadInternal(array $mergedConfig, ContainerBuilder $container)
    {
        // ...
    }
}

// Processing the Configuration yourself
public function load(array $configs, ContainerBuilder $container)
{
    $config = array();
    // let resources override the previous set value
    foreach ($configs as $subConfig) {
        $config = array_merge($config, $subConfig);
    }
    // ... now use the flat $config array
}

// Dump the Configuration
config:dump-reference

// How to Simplify Configuration of multiple Bundles
# src/Acme/HelloBundle/DependencyInjection/AcmeHelloExtension.php
namespace Acme\HelloBundle\DependencyInjection;
use Symfony\Component\HttpKernel\DependencyInjection\Extension;
use Symfony\Component\DependencyInjection\Extension\PrependExtensionInterface;
use Symfony\Component\DependencyInjection\ContainerBuilder;
class AcmeHelloExtension extends Extension implements PrependExtensionInterface
{
    public function prepend(ContainerBuilder $container)
    {
        // get all bundles
        $bundles = $container->getParameter('kernel.bundles');
        // determine if AcmeGoodbyeBundle is registered
        if (!isset($bundles['AcmeGoodbyeBundle'])) {
            // disable AcmeGoodbyeBundle in bundles
            $config = array('use_acme_goodbye' => false);
            foreach ($container->getExtensions() as $name => $extension) {
                switch ($name) {
                    case 'acme_something':
                    case 'acme_other':
                        // set use_acme_goodbye to false in the config of
                        // acme_something and acme_other note that if the user manually
                        // configured use_acme_goodbye to true in the app/config/config.yml
                        // then the setting would in the end be true and not false
                        $container->prependExtensionConfig($name, $config);
                        break;
                }
            }
        }
        // process the configuration of AcmeHelloExtension
        $configs = $container->getExtensionConfig($this->getAlias());
        // use the Configuration class to generate a config array with
        // the settings "acme_hello"
        $config = $this->processConfiguration(new Configuration(), $configs);
        // check if entity_manager_name is set in the "acme_hello" configuration
        if (isset($config['entity_manager_name'])) {
            // prepend the acme_something settings with the entity_manager_name
            $config = array('entity_manager_name' => $config['entity_manager_name']);
            $container->prependExtensionConfig('acme_something', $config);
        }
    }
}

// How to Set external Parameters in the Service Container
export SYMFONY__DATABASE__USER=user
export SYMFONY__DATABASE__PASSWORD=secre

doctrine:
    dbal:
        driver pdo_mysql
        dbname: symfony_project
        user: "%database.user%"
        password: "%database.password%"

# app/config/config.yml
imports:
- { resource: parameters.php }

# app/config/parameters.php
include_once('/path/to/drupal/sites/default/settings.php');
$container->setParameter('drupal.database.url', $db_url);
````

####Deploy
````
php app/check.php
composer install --no-dev --optimize-autoloader
php app/console cache:clear --env=prod --no-debug
php app/console assetic:dump --env=prod --no-debug
````

####Organize configuration files
````php
// app/AppKernel.php
use Symfony\Component\HttpKernel\Kernel;
use Symfony\Component\Config\Loader\LoaderInterface;
class AppKernel extends Kernel
{
    public function registerContainerConfiguration(LoaderInterface $loader)
    {
        $loader->load($this->getRootDir().'/config/config_'.$this->getEnvironment().'.yml');
        $loader->load($this->getRootDir().'/config/'.$this->getEnvironment().'/config.yml');
        $loader->load($this->getRootDir().'/config/environments/'.$this->getEnvironment().'.yml');
    }
}

# app/config/dev/config.yml
imports:
    - { resource: '../common/config.yml' }
    - { resource: 'parameters.yml' }
    - { resource: 'security.yml' }
    # silently discard errors when the loaded file doesn't exist
    # for confings not under the git.
    - { resource: '/etc/sites/mysite.com/parameters.yml', ignore_errors: true }
````

####The Directory Structure
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
````

####Override default directory structure
````php
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

// How to Customize Error Pages
app/
    └─ Resources/
        └─ TwigBundle/
            └─ views/
                └─ Exception/
                    ├─ error404.html.twig
                    ├─ error403.html.twig
                    ├─ error.html.twig # All other HTML errors (including 500)
                    ├─ error404.json.twig
                    ├─ error403.json.twig
                    ├─ error.json.twig # All other JSON errors (including 500)

// Example 404 Error Template
{# app/Resources/TwigBundle/views/Exception/error404.html.twig #}
{% extends 'base.html.twig' %}
{% block body %}
    <h1>Page not found</h1>
        {# example security usage, see below #}
    {% if app.user and is_granted('IS_AUTHENTICATED_FULLY') %}
    {% endif %}
        <p>
            The requested page couldn't be located. Checkout for any URL
            misspelling or <a href="{{ path('homepage') }}">return to the homepage</a>.
        </p>
{% endblock %}

// Overriding the Default ExceptionController
# app/config/config.yml
twig:
    exception_controller: AppBundle:Exception:showException

// How to Define Controllers as Services
# src/AppBundle/Controller/HelloController.php
namespace AppBundle\Controller;
use Symfony\Component\HttpFoundation\Response;
class HelloController
{
    public function indexAction($name)
    {
        return new Response('<html><body>Hello '.$name.'!</body></html>');
    }
}

# app/config/services.yml
services:
    app.hello_controller:
        class: AppBundle\Controller\HelloController

// Referring to the Service
$this->forward('app.hello_controller:indexAction', array('name' => $name));

# app/config/routing.yml
hello:
    path: /hello
    defaults: { _controller: app.hello_controller:indexAction }

// Alternatives to base Controller Methods
# src/AppBundle/Controller/HelloController.php
namespace AppBundle\Controller;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
class HelloController extends Controller
{
    public function indexAction($name)
    {
        return $this->render(
            'AppBundle:Hello:index.html.twig',
            array('name' => $name)
        );
    }
}

// Templating
# src/AppBundle/Controller/HelloController.php
namespace AppBundle\Controller;
use Symfony\Bundle\FrameworkBundle\Templating\EngineInterface;
use Symfony\Component\HttpFoundation\Response;
class HelloController
{
    private $templating;
    public function __construct(EngineInterface $templating)
    {
        $this->templating = $templating;
    }
    public function indexAction($name)
    {
        return $this->templating->renderResponse(
            'AppBundle:Hello:index.html.twig',
            array('name' => $name)
        );
    }
}

# app/config/services.yml
services:
    app.hello_controller:
        class: AppBundle\Controller\HelloController
        arguments: ["@templating"]
````

####Console Command
````php
// How to Create a Console Command
# src/AppBundle/Command/GreetCommand.php
namespace AppBundle\Command;
use Symfony\Bundle\FrameworkBundle\Command\ContainerAwareCommand;
use Symfony\Component\Console\Input\InputArgument;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;
class GreetCommand extends ContainerAwareCommand
{
    protected function configure()
    {
        $this
            ->setName('demo:greet')
            ->setDescription('Greet someone')
            ->addArgument(
                'name',
                InputArgument::OPTIONAL,
                'Who do you want to greet?'
            )
            ->addOption(
                'yell',
                null,
                InputOption::VALUE_NONE,
                'If set, the task will yell in uppercase letters'
            )
            ;
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $name = $input->getArgument('name');
        $logger = $this->getContainer()->get('logger');
        $logger->info('Executing command for '.$name);
        $locale = $input->getArgument('locale');
        $translator = $this->getContainer()->get('translator');
        $translator->setLocale($locale);
        if ($name) {
            $text = 'Hello '.$name;
            $logger->info('Greeted: '.$text);
        } else {
            $text = 'Hello';
            $logger->warning('Yelled: '.$text);
        }
        if ($input->getOption('yell')) {
            $text = strtoupper($text);
        }
        $output->writeln($text);
    }
}

php app/console demo:greet Fabien

// Testing Commands
use Symfony\Component\Console\Tester\CommandTester;
use Symfony\Bundle\FrameworkBundle\Console\Application;
use AppBundle\Command\GreetCommand;
class ListCommandTest extends \PHPUnit_Framework_TestCase
{
    public function testExecute()
    {
        // mock the Kernel or create one depending on your needs
        $application = new Application($kernel);
        $application->add(new GreetCommand());
        $command = $application->find('demo:greet');
        $commandTester = new CommandTester($command);
        $commandTester->execute(
            array(
               'name' => 'Fabien',
                '--yell' => true,
            )
        );
        $this->assertRegExp('/.../', $commandTester->getDisplay());
        // ...
    }
}

use Symfony\Component\Console\Tester\CommandTester;
use Symfony\Bundle\FrameworkBundle\Console\Application;
use Symfony\Bundle\FrameworkBundle\Test\KernelTestCase;
use AppBundle\Command\GreetCommand;
class ListCommandTest extends KernelTestCase
{
    public function testExecute()
    {
        $kernel = $this->createKernel();
        $kernel->boot();
        $application = new Application($kernel);
        $application->add(new GreetCommand());
        $command = $application->find('demo:greet');
        $commandTester = new CommandTester($command);
        $commandTester->execute(
            array(
                'name' => 'Fabien',
                '--yell' => true,
            )
        );
        $this->assertRegExp('/.../', $commandTester->getDisplay());
        // ...
    }
}

// How to Use the Console
php app/console cache:clear -e prod
php app/console list --no-debug

// interactive mode
php app/console --shell
php app/console -s

php app/console --shell --process-isolation
php app/console -s --process-isolation

// How to Generate URLs and Send Emails from the Console
# app/config/parameters.yml
parameters:
    router.request_context.host: example.org
    router.request_context.scheme: https
    router.request_context.base_url: my/path

class DemoCommand extends ContainerAwareCommand
{
    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $context = $this->getContainer()->get('router')->getContext();
        $context->setHost('example.com');
        $context->setScheme('https');
        $context->setBaseUrl('my/path');
        // ... your code here
    }
}

// Enabling automatic Exceptions Logging
# app/config/services.yml
services:
kernel.listener.command_dispatch:
    class: AppBundle\EventListener\ConsoleExceptionListener
    arguments:
        logger: "@logger"
    tags:
            - { name: kernel.event_listener, event: console.exception }

# app/config/services.yml
services:
kernel.listener.command_dispatch:
    class: AppBundle\EventListener\ConsoleExceptionListener
    arguments:
        logger: "@logger"
    tags:
        - { name: kernel.event_listener, event: console.exception }

// src/AppBundle/EventListener/ConsoleExceptionListener.php
namespace AppBundle\EventListener;
use Symfony\Component\Console\Event\ConsoleExceptionEvent;
use Psr\Log\LoggerInterface;
class ConsoleExceptionListener
{
    private $logger;
    public function __construct(LoggerInterface $logger)
    {
        $this->logger = $logger;
    }

    public function onConsoleException(ConsoleExceptionEvent $event)
    {
        $command = $event->getCommand();
        $exception = $event->getException();
        $message = sprintf(
            '%s: %s (uncaught exception) at %s line %s while running console command `%s`',
            get_class($exception),
            $exception->getMessage(),
            $exception->getFile(),
            $exception->getLine(),
            $command->getName()
        );
        $this->logger->error($message, array('exception' => $exception));
    }
}

// How to Define Commands as Services
# app/config/config.yml
services:
    acme_hello.command.my_command:
        class: Acme\HelloBundle\Command\MyCommand
        tags:
            - { name: console.command }

// src/Acme/DemoBundle/Command/GreetCommand.php
namespace Acme\DemoBundle\Command;
use Acme\DemoBundle\Entity\NameRepository;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;
class GreetCommand extends Command
{
    protected $nameRepository;
    public function __construct(NameRepository $nameRepository)
    {
        $this->nameRepository = $nameRepository;
        parent::__construct();
    }
    protected function configure()
    {
        $defaultName = $this->nameRepository->findLastOne();
        $this
            ->setName('demo:greet')
            ->setDescription('Greet someone')
            ->addOption(
                'name',
                '-n',
                InputOption::VALUE_REQUIRED,
                'Who do you want to greet?',
                $defaultName
            )
        ;
    }
    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $name = $input->getOption('name');
        $output->writeln($name);
    }
}
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

<script src="{{ asset('js/script.js') }}"></script>

{% block javascripts %}
    {% javascripts '@AppBundle/Resources/public/js/*' %}
        <script src="{{ asset_url }}"></script>
    {% endjavascripts %}
{% endblock %}
{% javascripts
    '@AppBundle/Resources/public/js/*'
    '@AcmeBarBundle/Resources/public/js/form.js'
    '@AcmeBarBundle/Resources/public/js/calendar.js' %}
    <script src="{{ asset_url }}"></script>
{% endjavascripts %}
{% javascripts
    '@AppBundle/Resources/public/js/thirdparty/jquery.js'
    '@AppBundle/Resources/public/js/*' %}
    <script src="{{ asset_url }}"></script>
{% endjavascripts %}
{% javascripts '@AppBundle/Resources/public/js/*' output='js/compiled/main.js' %}
    <script src="{{ asset_url }}"></script>
{% endjavascripts %}

{% block stylesheets %}
    {% stylesheets 'bundles/app/css/*' filter='cssrewrite' %}
        <link rel="stylesheet" href="{{ asset_url }}" />
    {% endstylesheets %}
{% endblock %}

{% image '@AppBundle/Resources/public/images/example.jpg' %}
    <img src="{{ asset_url }}" alt="Example" />
{% endimage %}

// Using Named Assets
# app/config/config.yml
assetic:
    assets:
        jquery_and_ui:
            inputs:
                - '@AppBundle/Resources/public/js/thirdparty/jquery.js'
                - '@AppBundle/Resources/public/js/thirdparty/jquery.ui.js'

{% javascripts
    '@jquery_and_ui'
    '@AppBundle/Resources/public/js/*' %}
    <script src="{{ asset_url }}"></script>
{% endjavascripts %}

// Filters
# app/config/config.yml
assetic:
    filters:
        uglifyjs2:
            bin: /usr/local/bin/uglifyjs

{% javascripts '@AppBundle/Resources/public/js/*' filter='uglifyjs2' %}
    <script src="{{ asset_url }}"></script>
{% endjavascripts %}

// Dumping Asset Files
php app/console assetic:dump --env=prod --no-debug

# app/config/config_dev.yml
assetic:
    use_controller: false

php app/console assetic:dump
php app/console assetic:watch

//  Minify CSS/JS Files (Using UglifyJS and UglifyCSS)
1)
npm install -g uglify-js
uglifyjs --help
2)
cd /path/to/your/symfony/project
npm install uglify-js --prefix app/Resources
"./app/Resources/node_modules/.bin/uglifyjs" --help

# app/config/config.yml
assetic:
    filters:
        uglifyjs2:
            # the path to the uglifyjs executable
            bin: /usr/local/bin/uglifyjs

# Configure the node Binary
# app/config/config.yml
assetic:
    # the path to the node executable
    node: /usr/bin/nodejs
    filters:
        uglifyjs2:
            # the path to the uglifyjs executable
            bin: /usr/local/bin/uglifyjs

{% javascripts '@AppBundle/Resources/public/js/*' filter='uglifyjs2' %}
    <script src="{{ asset_url }}"></script>
{% endjavascripts %}

// Disable Minification in Debug Mode
{% javascripts '@AppBundle/Resources/public/js/*' filter='?uglifyjs2' %}
    <script src="{{ asset_url }}"></script>
{% endjavascripts %}

// UglifyCSS
1)
npm install -g uglifycss
2)
cd /path/to/your/symfony/project
npm install uglifycss --prefix app/Resources

# app/config/config.yml
assetic:
    filters:
        uglifycss:
            bin: /usr/local/bin/uglifycss

{% stylesheets 'bundles/App/css/*' filter='uglifycss' filter='cssrewrite' %}
    <link rel="stylesheet" href="{{ asset_url }}" />
{% endstylesheets %}

// Image Optimization
# app/config/config.yml
assetic:
    filters:
        jpegoptim:
            bin: path/to/jpegoptim

{% image '@AppBundle/Resources/public/images/example.jpg'
    filter='jpegoptim' output='/images/example.jpg' %}
    <img src="{{ asset_url }}" alt="Example"/>
{% endimage %}

// Shorter Syntax
# app/config/config.yml
assetic:
    filters:
        jpegoptim:
            bin: path/to/jpegoptim
    twig:
        functions:
            # jpegoptim: ~
            jpegoptim: { output: images/*.jpg }

<img src="{{ jpegoptim('@AppBundle/Resources/public/images/example.jpg') }}" alt="Example"/>

// Filter to a specific File Extension
# app/config/config.yml
assetic:
    filters:
        coffee:
            bin: /usr/bin/coffee
            node: /usr/bin/node
            node_paths: [/usr/lib/node_modules/]
            # apply_to: "\.coffee$" # Filtering Based on a File Extension

{% javascripts '@AppBundle/Resources/public/js/example.coffee' filter='coffee' %}
    <script src="{{ asset_url }}"></script>
{% endjavascripts %}
{% javascripts '@AppBundle/Resources/public/js/example.coffee'
    '@AppBundle/Resources/public/js/another.coffee'
    filter='coffee' %}
    <script src="{{ asset_url }}"></script>
{% endjavascripts %}
{% javascripts '@AppBundle/Resources/public/js/example.coffee'
    '@AppBundle/Resources/public/js/another.coffee'
    '@AppBundle/Resources/public/js/regular.js' %}
    <script src="{{ asset_url }}"></script>
{% endjavascripts %}

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

####Handle File Uploads with Doctrine
````php
# First, create a simple Doctrine entity
// src/AppBundle/Entity/Document.php
namespace AppBundle\Entity;
use Doctrine\ORM\Mapping as ORM;
use Symfony\Component\Validator\Constraints as Assert;
/**
* @ORM\Entity
*/
class Document
{
    /**
    * @ORM\Id
    * @ORM\Column(type="integer")
    * @ORM\GeneratedValue(strategy="AUTO")
    */
    public $id;
    /**
    * @ORM\Column(type="string", length=255)
    * @Assert\NotBlank
    */
    public $name;
    /**
    * @ORM\Column(type="string", length=255, nullable=true)
    */
    public $path;
    public function getAbsolutePath()
    {
        return null === $this->path
        ? null
        : $this->getUploadRootDir().'/'.$this->path;
    }
    public function getWebPath()
    {
        return null === $this->path
        ? null
        : $this->getUploadDir().'/'.$this->path;
    }
    protected function getUploadRootDir()
    {
        // the absolute directory path where uploaded
        // documents should be saved
        return __DIR__.'/../../../../web/'.$this->getUploadDir();
    }
    protected function getUploadDir()
    {
        // get rid of the __DIR__ so it doesn't screw up
        // when displaying uploaded doc/image in the view.
        return 'uploads/documents';
    }
}

# In the form, use a "virtual" file field
public function uploadAction()
{
    $form = $this->createFormBuilder($document)
    ->add('name')
    ->add('file')
    ->getForm();
}

// Create this property on your Document
use Symfony\Component\HttpFoundation\File\UploadedFile;
// ...
class Document
{
    /**
    * @Assert\File(maxSize="6000000")
    */
    private $file;
    /**
    * Sets file.
    *
    * @param UploadedFile $file
    */
    public function setFile(UploadedFile $file = null)
    {
        $this->file = $file;
    }
    /**
    * Get file.
    *
    * @return UploadedFile
    */
    public function getFile()
    {
        return $this->file;
    }
}

// src/AppBundle/Entity/Document.php
namespace AppBundle\Entity;
// ...
use Symfony\Component\Validator\Constraints as Assert;
class Document
{
    /**
    * @Assert\File(maxSize="6000000")
    */
    private $file;
}

// Handle the entire process
use AppBundle\Entity\Document;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Template;
use Symfony\Component\HttpFoundation\Request;
// ...
/**
* @Template()
*/
public function uploadAction(Request $request)
{
    $document = new Document();
    $form = $this->createFormBuilder($document)
        ->add('name')
        ->add('file')
        ->getForm();
    $form->handleRequest($request);
    if ($form->isValid()) {
        $em = $this->getDoctrine()->getManager();
        $em->persist($document);
        $em->flush();
        return $this->redirectToRoute(...);
    }
    return array('form' => $form->createView());
}

if ($form->isValid()) {
    $em = $this->getDoctrine()->getManager();
    $document->upload();
    $em->persist($document);
    $em->flush();
    return $this->redirectToRoute(...);
}

public function upload()
{
    // the file property can be empty if the field is not required
    if (null === $this->getFile()) {
        return;
    }
    // use the original file name here but you should
    // sanitize it at least to avoid any security issues
    // move takes the target directory and then the
    // target filename to move to
    $this->getFile()->move(
        $this->getUploadRootDir(),
        $this->getFile()->getClientOriginalName()
    );
    // set the path property to the filename where you've saved the file
    $this->path = $this->getFile()->getClientOriginalName();
    // clean up the file property as you won't need it anymore
    $this->file = null;
}

// Using the id as the Filename
use Symfony\Component\HttpFoundation\File\UploadedFile;
/**
* @ORM\Entity
* @ORM\HasLifecycleCallbacks
*/
class Document
{
    private $temp;
    /**
    * Sets file.
    *
    * @param UploadedFile $file
    */
    public function setFile(UploadedFile $file = null)
    {
        $this->file = $file;
        // check if we have an old image path
        if (is_file($this->getAbsolutePath())) {
            // store the old name to delete after the update
            $this->temp = $this->getAbsolutePath();
        } else {
            $this->path = 'initial';
        }
    }
    /**
    * @ORM\PrePersist()
    * @ORM\PreUpdate()
    */
    public function preUpload()
    {
        if (null !== $this->getFile()) {
            $this->path = $this->getFile()->guessExtension();
        }
    }
    /**
    * @ORM\PostPersist()
    * @ORM\PostUpdate()
    */
    public function upload()
    {
        if (null === $this->getFile()) {
           return;
        }
        // check if we have an old image
        if (isset($this->temp)) {
            // delete the old image
            unlink($this->temp);
            // clear the temp image path
            $this->temp = null;
        }
        // you must throw an exception here if the file cannot be moved
        // so that the entity is not persisted to the database
        // which the UploadedFile move() method does
        $this->getFile()->move(
            $this->getUploadRootDir(),
            $this->id.'.'.$this->getFile()->guessExtension()
        );
        $this->setFile(null);
    }
    /**
    * @ORM\PreRemove()
    */
    public function storeFilenameForRemove()
    {
        $this->temp = $this->getAbsolutePath();
    }
    /**
    * @ORM\PostRemove()
    */
    public function removeUpload()
    {
        if (isset($this->temp)) {
        unlink($this->temp);
    }
    }
    public function getAbsolutePath()
    {
        return null === $this->path
        ? null
        : $this->getUploadRootDir().'/'.$this->id.'.'.$this->path;
    }
}
````

####Define Relationships with Abstract Classes and Interfaces
````php
// src/Acme/AppBundle/Entity/Customer.php
namespace Acme\AppBundle\Entity;
use Doctrine\ORM\Mapping as ORM;
use Acme\CustomerBundle\Entity\Customer as BaseCustomer;
use Acme\InvoiceBundle\Model\InvoiceSubjectInterface;
/**
 * @ORM\Entity
 * @ORM\Table(name="customer")
 */
class Customer extends BaseCustomer implements InvoiceSubjectInterface
{
    // In this example, any methods defined in the InvoiceSubjectInterface
    // are already implemented in the BaseCustomer
}

// src/Acme/InvoiceBundle/Entity/Invoice.php
namespace Acme\InvoiceBundle\Entity;
use Doctrine\ORM\Mapping AS ORM;
use Acme\InvoiceBundle\Model\InvoiceSubjectInterface;
/**
* Represents an Invoice.
*
* @ORM\Entity
* @ORM\Table(name="invoice")
*/
class Invoice
{
    /**
     * @ORM\ManyToOne(targetEntity="Acme\InvoiceBundle\Model\InvoiceSubjectInterface")
     * @var InvoiceSubjectInterface
     */
    protected $subject;
}

// src/Acme/InvoiceBundle/Model/InvoiceSubjectInterface.php
namespace Acme\InvoiceBundle\Model;
/**
* An interface that the invoice Subject object should implement.
* In most circumstances, only a single object should implement
* this interface as the ResolveTargetEntityListener can only
* change the target to a single object.
*/
interface InvoiceSubjectInterface
{
    // List any additional methods that your InvoiceBundle
    // will need to access on the subject so that you can
    // be sure that you have access to those methods.
    /**
     * @return string
     */
    public function getName();
}

# app/config/config.yml
doctrine:
    orm:
        resolve_target_entities:
            Acme\InvoiceBundle\Model\InvoiceSubjectInterface: Acme\AppBundle\Entity\Customer
````

####Store sessions in the database
````yml
# app/config/config.yml
framework:
    session:
        # ...
        handler_id: session.handler.pdo
services:
    session.handler.pdo:
        class: Symfony\Component\HttpFoundation\Session\Storage\Handler\PdoSessionHandler
        public: false
        arguments:
            - "mysql:dbname=mydatabase"
            - { db_username: myuser, db_password: mypassword }

// Configuring the Table and Column Names
# app/config/config.yml
services:
    # ...
    session.handler.pdo:
        class: Symfony\Component\HttpFoundation\Session\Storage\Handler\PdoSessionHandler
        public: false
        arguments:
            - "mysql:dbname=mydatabase"
            - { db_table: sessions, db_username: myuser, db_password: mypassword }

// Sharing your Database Connection Information

services:
session.handler.pdo:
    class: Symfony\Component\HttpFoundation\Session\Storage\Handler\PdoSessionHandler
    public: false
    arguments:
        - "mysql:host=%database_host%;port=%database_port%;dbname=%database_name%"
        - { db_username: %database_user%, db_password: %database_password% }

// MySQL
CREATE TABLE `sessions` (
    `sess_id` VARBINARY(128) NOT NULL PRIMARY KEY,
    `sess_data` BLOB NOT NULL,
    `sess_time` INTEGER UNSIGNED NOT NULL,
    `sess_lifetime` MEDIUMINT NOT NULL
) COLLATE utf8_bin, ENGINE = InnoDB;
````

####Register Event Listeners and Subscribers
````php
# Configuring
doctrine:
    dbal:
        default_connection: default
        connections:
            default:
                driver: pdo_sqlite
                memory: true
services:
my.listener:
    class: Acme\SearchBundle\EventListener\SearchIndexer
    tags:
        - { name: doctrine.event_listener, event: postPersist }
my.listener2:
    class: Acme\SearchBundle\EventListener\SearchIndexer2
    tags:
        - { name: doctrine.event_listener, event: postPersist, connection: default }
my.subscriber:
    class: Acme\SearchBundle\EventListener\SearchIndexerSubscriber
    tags:
        - { name: doctrine.event_subscriber, connection: default }

# Listener Class
// src/Acme/SearchBundle/EventListener/SearchIndexer.php
namespace Acme\SearchBundle\EventListener;
use Doctrine\ORM\Event\LifecycleEventArgs;
use Acme\StoreBundle\Entity\Product;
class SearchIndexer
{
    public function postPersist(LifecycleEventArgs $args)
    {
        $entity = $args->getEntity();
        $entityManager = $args->getEntityManager();
        // perhaps you only want to act on some "Product" entity
        if ($entity instanceof Product) {
            // ... do something with the Product
        }
    }
}

# Subscriber Class
// src/Acme/SearchBundle/EventListener/SearchIndexerSubscriber.php
namespace Acme\SearchBundle\EventListener;
use Doctrine\Common\EventSubscriber;
use Doctrine\ORM\Event\LifecycleEventArgs;
// for Doctrine 2.4: Doctrine\Common\Persistence\Event\LifecycleEventArgs;
use Acme\StoreBundle\Entity\Product;
class SearchIndexerSubscriber implements EventSubscriber
{
    public function getSubscribedEvents()
    {
        return array(
            'postPersist',
            'postUpdate',
        );
    }
    public function postUpdate(LifecycleEventArgs $args)
    {
        $this->index($args);
    }
    public function postPersist(LifecycleEventArgs $args)
    {
        $this->index($args);
    }
    public function index(LifecycleEventArgs $args)
    {
        $entity = $args->getEntity();
        $entityManager = $args->getEntityManager();
        // perhaps you only want to act on some "Product" entity
        if ($entity instanceof Product) {
            // ... do something with the Product
        }
    }
}
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

page:181
