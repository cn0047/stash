<?php

namespace AppBundle\Controller\Guest;

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\Request;

class CacheController extends Controller
{
    /**
     * @Route("/http_cache", name="http_cache")
     */
    public function httpCacheAction(Request $request)
    {
        $html = <<<"HEREDOC"
<pre>
HTTP Cache
3.0 version

The nature of rich web applications means that they're dynamic. No matter how efficient your application,
each request will always contain more overhead than serving a static file.

And for most Web applications, that's fine. Symfony is lightning fast,
and unless you're doing some serious heavy-lifting,
each request will come back quickly without putting too much stress on your server.

But as your site grows, that overhead can become a problem.
The processing that's normally performed on every request should be done only once.
This is exactly what caching aims to accomplish.
</pre>
HEREDOC;
        $response = $this->render('AppBundle:default:index.html.twig', [
            'message' => $html,
        ]);
        $response->setPublic();
        $response->setMaxAge(600);
        $response->setSharedMaxAge(600);
        $response->headers->addCacheControlDirective('must-revalidate', true);
        return $response;
    }
}
