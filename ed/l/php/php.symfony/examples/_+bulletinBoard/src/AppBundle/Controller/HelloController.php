<?php

namespace AppBundle\Controller;

use Symfony\Component\HttpFoundation\Response;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;

class HelloController extends Controller
{
    /**
     *  @Route("/hello/{name}", name="hello")
     */
    public function indexAction($name)
    {
        return new Response('<html><body>Hello '.$name.'!</body></html>', Response::HTTP_OK);
    }

    /**
    * @ Route("/hello/{firstName}/{lastName}", name="hello")
    */
    public function hiAction($firstName, $lastName, $ending = '!!!')
    {
        return new Response("<html><body>Hello $firstName $lastName $ending</body></html>");
    }
}
