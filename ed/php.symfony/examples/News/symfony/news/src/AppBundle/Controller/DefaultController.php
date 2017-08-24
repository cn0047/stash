<?php

namespace AppBundle\Controller;

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\Request;

class DefaultController extends Controller
{
    /**
     * @Route("/", name="homepage")
     */
    public function indexAction(Request $request)
    {
        // replace this example code with whatever you need
        return $this->render('default/index.html.twig', [
            'base_dir' => realpath($this->getParameter('kernel.root_dir').'/..'),
        ]);
    }

    /**
     * @Route("/voter", name="voter")
     */
    public function voterGuestAction(Request $request)
    {
        $this->denyAccessUnlessGranted('view', ['page' => 'voterGuest']);
        return $this->render('default/voter.html.twig', [
            'code' => 200,
        ]);
    }
}
