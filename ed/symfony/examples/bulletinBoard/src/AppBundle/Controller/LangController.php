<?php

namespace AppBundle\Controller;

use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;

class LangController extends Controller
{
    public function a1Action()
    {
        var_export(__METHOD__);
        return $this->render('default/index.html.twig');
    }

    public function a2Action()
    {
        var_export(__METHOD__);
        return $this->render('default/index.html.twig');
    }
}
