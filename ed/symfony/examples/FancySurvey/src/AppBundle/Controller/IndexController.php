<?php

namespace AppBundle\Controller;

use AppBundle\Entity\User;
use AppBundle\Entity\UserInfo;
use AppBundle\Form\Type\Step1;
use AppBundle\Form\Type\Step2;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Session\Session;

class IndexController extends Controller
{
    /**
     * @Route("/", name="step1")
     */
    public function step1Action(Request $request)
    {
        $user = new User();
        $form = $this->createForm(new Step1(), $user);
        return $this->render(
            'AppBundle:Index:Step1.html.twig',
            ['form' => $form->createView()]
        );
    }

    /**
     * @Route("/makeStep1", name="makeStep1")
     */
    public function makeStep1Action(Request $request)
    {
        if (!$request->isXmlHttpRequest()) {
            return new JsonResponse(['message' => 'You should use ajax.'], 400);
        }
        $user = new User();
        $form = $this->createForm(new Step1(), $user);
        $form->handleRequest($request);
        if ($form->isValid()) {
            $em = $this->getDoctrine()->getManager();
            $em->persist($user);
            $em->flush();
            $session = $this->getRequest()->getSession();
            $session->set('user', $user->getId());
            return new JsonResponse(
                ['message' => 'Success!', 'route' => 'step2'],
                200
            );
        }
        $formHtml = $this->renderView(
            'AppBundle:Index:Step1.html.twig',
            ['form' => $form->createView()]
        );
        $response = new JsonResponse(
            ['message' => 'Error', 'form' => $formHtml],
            400
        );
        return $response;
    }

    /**
     * @Route("/step2", name="step2")
     */
    public function step2Action(Request $request)
    {
        $userInfo = new UserInfo();
        $form = $this->createForm(new Step2(), $userInfo);
        return $this->render(
            'AppBundle:Index:Step2.html.twig',
            ['form' => $form->createView()]
        );
    }

    /**
     * @Route("/makeStep2", name="makeStep2")
     */
    public function makeStep2Action(Request $request)
    {
        if (!$request->isXmlHttpRequest()) {
            return new JsonResponse(['message' => 'You should use ajax.'], 400);
        }
        $userInfo = new UserInfo();
        $form = $this->createForm(new Step2(), $userInfo);
        $form->handleRequest($request);
        if ($form->isValid()) {
            $session = $this->getRequest()->getSession();
            $user = $this->getDoctrine()
                ->getRepository('AppBundle:User')
                ->find($session->get('user'))
            ;
            $userInfo->setuser($user);
            $em = $this->getDoctrine()->getManager();
            $em->persist($userInfo);
            $em->flush();
            return new JsonResponse(
                ['message' => 'Success!', 'route' => 'step3'],
                200
            );
        }
        $formHtml = $this->renderView(
            'AppBundle:Index:Step3.html.twig',
            ['form' => $form->createView()]
        );
        $response = new JsonResponse(
            ['message' => 'Error', 'form' => $formHtml],
            400
        );
        return $response;
    }

    /**
     * @Route("/step3", name="step3")
     */
    public function step3Action(Request $request)
    {
        return $this->render('AppBundle:Index:Step3.html.twig');
    }
}
