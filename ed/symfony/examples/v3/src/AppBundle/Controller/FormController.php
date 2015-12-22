<?php

namespace AppBundle\Controller;

use AppBundle\Form\MyCsrfType;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Security;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;

class FormController extends Controller
{
    /**
     * @Route("/my_csrf", name="my_csrf")
     */
    public function indexAction(Request $request)
    {
        /** @var Form */
        $form = $this->createForm(MyCsrfType::class);
        $form->handleRequest($request);
        if ($form->isSubmitted()) {
            print($this->get('translator')->trans('Form submited!'));
        }
        $response = $this->render('AppBundle:form:my_csrf.html.twig', [
            'form' => $form->createView(),
        ]);
        return $response;
    }
}
