<?php

namespace AppBundle\Controller;

use AppBundle\Form\MyCsrfType;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Security;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Security\Csrf\CsrfToken;

class DoctrineController extends Controller
{
    /**
     * @Route("/plain_sql", name="plain_sql")
     */
    public function plainSqlAction(Request $request)
    {
        return $this->render('AppBundle:default:index.html.twig');
    }
}
