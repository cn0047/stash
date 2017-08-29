<?php

namespace AppBundle\Controller;

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;

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
     * @Route("/voter/view", name="voterView")
     * @throws \Symfony\Component\Security\Core\Exception\AccessDeniedException
     */
    public function voterViewGuestAction(Request $request)
    {
        $this->denyAccessUnlessGranted('view', ['page' => 'VoterViewPage']);
        return $this->render('default/voter.html.twig', [
            'code' => 204,
        ]);
    }

    /**
     * @Route("/voter/edit", name="voterEdit")
     * @throws \Symfony\Component\Security\Core\Exception\AccessDeniedException
     */
    public function voterEditGuestAction(Request $request)
    {
        $this->denyAccessUnlessGranted('edit', ['page' => 'VoterEditPage']);
        return $this->render('default/voter.html.twig', [
            'code' => 200,
        ]);
    }

    /**
     * @Route("/anotation", name="anotation")
     * @throws \InvalidArgumentException
     * @throws \UnexpectedValueException
     */
    public function anotationAction(Request $request)
    {
        $response = new Response();
        $response->setContent('Anotation.');
        return $response;
    }

    public function yamlAction(Request $request)
    {
        $response = new Response();
        $response->setContent('YAML.');
        return $response;
    }
}
