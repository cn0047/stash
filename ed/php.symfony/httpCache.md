HTTP Cache
-

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
