<?php

namespace AppBundle\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;

class FlashController extends Controller
{
    /**
    * @Route("/flash/in", name="in")
    */
    public function inAction(Request $request)
    {
        $this->addFlash('notice', 'Everything ok!');
        return $this->redirect('out');
    }

    /**
    * @Route("/flash/out", name="out")
    */
    public function outAction(Request $request)
    {
        return $this->render('default/flash.html.twig');
    }
}
