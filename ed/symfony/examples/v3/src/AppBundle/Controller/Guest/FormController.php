<?php

namespace AppBundle\Controller\Guest;

use AppBundle\Form\MyCsrfType;
use AppBundle\Form\CategoryType;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Security;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\Form\Extension\Core\Type\SubmitType;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Security\Csrf\CsrfToken;

class FormController extends Controller
{
    /**
     * @Route("/embeded_form", name="embeded_form")
     */
    public function embededFormAction(Request $request)
    {
        $form = $this->createFormBuilder()
            ->add('msg')
            ->add('category', CategoryType::class)
            ->add('submit', SubmitType::class)
            ->getForm()
        ;
        $form->handleRequest($request);
        $data = $form->getData();
        dump($data);
        $response = $this->render('AppBundle:form:my_csrf.html.twig', [
            'form' => $form->createView(),
        ]);
        return $response;
    }

    /**
     * @Route("/my_csrf", name="my_csrf")
     */
    public function indexAction(Request $request)
    {
        $form = $this->createForm(MyCsrfType::class);
        $form->handleRequest($request);
        $message = '';
        if ($form->isSubmitted()) {
            $args = $request->request->get('my_csrf');
            $csrfToken = new CsrfToken('my_csrf', $args['_token']);
            $isTokenValid1 = $this->get('security.csrf.token_manager')->isTokenValid($csrfToken);
            $isTokenValid2 = $this->isCsrfTokenValid('my_csrf', $args['_token']);
            $message = $this->get('translator')->trans(
                'Form submitted with message: %msg%! Token valid: method 1 = %v1%, method 2 = %v2%',
                [
                    '%msg%' => $args['message'],
                    '%v1%' => var_export($isTokenValid1, true),
                    '%v2%' => var_export($isTokenValid2, true),
                ]
            );
        }
        $response = $this->render('AppBundle:form:my_csrf.html.twig', [
            'form' => $form->createView(),
            'message' => $message,
        ]);
        return $response;
    }
}
