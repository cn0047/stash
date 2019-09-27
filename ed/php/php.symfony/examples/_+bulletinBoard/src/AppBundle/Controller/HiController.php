<?php

namespace AppBundle\Controller;

use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;

class HiController extends Controller
{
    /**
    * @Route("/hi/{firstName}/{lastName}", name="hi")
    */
    public function indexAction($firstName, $lastName, Request $request, $ending = '!!!')
    {
        var_dump([
            $request->query->get('name', 'none'),
            $request->query->get('page', 1),
        ]);
        return new Response("<html><body>Hello $firstName $lastName $ending</body></html>");
    }
}
