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
        /** @var \Symfony\Component\Form\Form */
        $form = $this->createForm(MyCsrfType::class);
        $form->handleRequest($request);
        if ($form->isSubmitted()) {
            print($this->get('translator')->trans('Form submited! Token valid: '));
            $args = $request->request->get('my_csrf');
            $csrfToken = new CsrfToken('my_csrf', $args['_token']);
            var_dump($this->get('security.csrf.token_manager')->isTokenValid($csrfToken));
            var_dump($this->isCsrfTokenValid('my_csrf', $args['_token']));
        }
        $response = $this->render('AppBundle:form:my_csrf.html.twig', [
            'form' => $form->createView(),
        ]);
        return $response;
    }
}
