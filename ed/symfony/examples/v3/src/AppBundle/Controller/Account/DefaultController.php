<?php

namespace AppBundle\Controller\Account;

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Security;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\Request;

/**
 * @Route("/account")
 * @Security("has_role('ROLE_USER')")
 */
class DefaultController extends Controller
{
    /**
     * @Route("/", name="account_homepage")
     */
    public function indexAction(Request $request)
    {
        var_dump(__CLASS__);
        return $this->render('AppBundle:default:index.html.twig');
    }
}
