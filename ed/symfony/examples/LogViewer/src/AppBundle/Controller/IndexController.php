<?php

namespace AppBundle\Controller;

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\Request;
use APY\DataGridBundle\Grid\Source\Vector;

class IndexController extends Controller
{
    /**
     * @Route("/", name="index")
     */
    public function indexAction(Request $request)
    {
        return $this->render('AppBundle:Index:Index.html.twig');
    }
}
