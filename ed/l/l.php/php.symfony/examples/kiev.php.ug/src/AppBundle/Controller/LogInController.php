<?php

namespace AppBundle\Controller;

use AppBundle\Entity\User;
use AppBundle\Form\Type\ForgotPassword;
use AppBundle\Form\Type\LogIn;
use AppBundle\Form\Type\Registration;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\Form\FormError;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Security\Core\Authentication\Token\UsernamePasswordToken;

class LogInController extends Controller
{
    public function githubAction()
    {
        return $this->redirectToRoute('home');
    }

    /**
     * @Route("/registration", name="registration")
     */
    public function registrationAction(Request $request)
    {
        $user = new User();
        $form = $this->createForm(new Registration(), $user);
        $form->handleRequest($request);
        if ($form->isValid() and $form->isSubmitted()) {
            $user->setPassword(password_hash($user->getPlainPassword(), PASSWORD_DEFAULT));
            $em = $this->getDoctrine()->getManager();
            $em->persist($user);
            $em->flush();
            $request->getSession()
                ->getFlashBag()
                ->add(
                    'success',
                    $this->get('translator')->trans('kpug.success.user_registered')
                )
            ;
            return $this->redirectToRoute('logIn');
        }
        return $this->render(
            'AppBundle:LogIn:Registration.html.twig',
            ['form' => $form->createView()]
        );
    }

    /**
     * @Route("/logInOwn", name="logInOwn")
     */
    public function logInOwnAction(Request $request)
    {
        $user = new User();
        $form = $this->createForm(new LogIn(), $user);
        return $this->render(
            'AppBundle:LogIn:LogInOwn.html.twig',
            ['form' => $form->createView()]
        );
    }

    /**
     * @Route("/logInOwnCheck", name="logInOwnCheck")
     */
    public function logInOwnCheckAction(Request $request)
    {
        $t = $this->get('translator');
        $user = new User();
        $form = $this->createForm(new LogIn(), $user);
        $form->handleRequest($request);
        if ($form->isValid() and $form->isSubmitted()) {
            $foundUser = $this->get('doctrine')
                ->getManager()
                ->getRepository('AppBundle:User')
                ->findOneByEmail($user->getEmail());
            if (is_null($foundUser)) {
                $form->get('email')->addError(new FormError(
                    $t->trans('kpug.error.not_found_email')
                ));
            } else {
                $isPasswordVerified = password_verify($user->getPassword(), $foundUser->getPassword());
                if ($isPasswordVerified) {
                    // OK.
                    $token = new UsernamePasswordToken(
                        $foundUser,
                        $foundUser->getPassword(),
                        'default',
                        $user->getRoles()
                    );
                    $this->get('security.context')->setToken($token);
                    return $this->redirectToRoute('home');
                } else {
                    $form->get('password')->addError(new FormError(
                        $t->trans('kpug.error.wrong_password')
                    ));
                }
            }
        }
        return $this->render(
            'AppBundle:LogIn:LogInOwn.html.twig',
            ['form' => $form->createView()]
        );
    }

    /**
     * @Route("/logOut", name="logOut")
     */
    public function logOutAction(Request $request)
    {
        return $this->redirectToRoute('home');
    }

    /**
     * @Route("/forgotPassword", name="forgotPassword")
     */
    public function forgotPasswordAction(Request $request)
    {
        $passwordUpdated = false;
        $user = new User();
        $form = $this->createForm(new ForgotPassword(), $user);
        $form->handleRequest($request);
        if ($form->isValid() and $form->isSubmitted()) {
            $foundUser = $this->get('doctrine')
                ->getManager()
                ->getRepository('AppBundle:User')
                ->findOneByEmail($user->getEmail());
            if (is_null($foundUser)) {
                $form->get('email')->addError(new FormError(
                    $this->get('translator')->trans('kpug.error.not_found_email')
                ));
            } else {
                $foundUser->setPlainPassword(uniqid());
                $foundUser->setPassword(password_hash($foundUser->getPlainPassword(), PASSWORD_DEFAULT));
                $em = $this->getDoctrine()->getManager();
                $em->persist($foundUser);
                $em->flush();
                $m = $this->get('mail');
                $m->sendMail(
                    'forgotPassword',
                    ['email' => $foundUser->getEmail(), 'plainPassword' => $foundUser->getPlainPassword()]
                );
                $passwordUpdated = true;
            }
        }
        return $this->render(
            'AppBundle:LogIn:ForgotPassword.html.twig',
            ['form' => $form->createView(), 'passwordUpdated' => $passwordUpdated]
        );
    }
}
