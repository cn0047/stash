HTTP
-

## Routing

````
Special Routing Parameters:
_controller - which controller is executed
_format     - format
_locale     - locale

// Visualizing & Debugging Routes.
php bin/console debug:router
php bin/console debug:router article_show
php bin/console router:match /blog/my-latest-post
````

## Request

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

## Response

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

## Controller

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
