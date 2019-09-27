<?php

namespace AppBundle\Controller;

use Symfony\Component\HttpFoundation\Response;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Method;

class BlogController extends Controller
{
    /**
    * @Route("/blog/{page}", defaults={"page": 1}, requirements={
    * "page": "\d+"
    * })
    */
    public function indexAction($page)
    {
        return new Response("Page: $page");
    }

    /**
    * @Route("/blog/{slug}", name="blog_show")
    */
    public function showAction($slug)
    {
        return new Response("Slug: $slug");
    }

    /**
    * @Route("/{_locale}", defaults={"_locale": "en"}, requirements={
    * "_locale": "en|fr"
    * })
    */
    public function homepageAction($_locale)
    {
        return $this->render('default/index.html.twig');
    }

    /**
    * @Route("/contact")
    * @Method("GET")
    */
    public function contactAction()
    {
    }

    /**
    * @Route("/contact")
    * @Method("POST")
    */
    public function processContactAction()
    {
    }

    /**
    * @Route(
    * "/articles/{_locale}/{year}/{title}.{_format}",
    * defaults={"_format": "html"},
    * requirements={
    * "_locale": "en|fr",
    * "_format": "html|rss",
    * "year": "\d+"
    * }
    * )
    */
    public function renderAction($_locale, $year, $title)
    {
    }
}
