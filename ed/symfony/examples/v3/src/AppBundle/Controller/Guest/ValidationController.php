<?php

namespace AppBundle\Controller\Guest;

use AppBundle\Entity\UserForValidation;
use AppBundle\Form\MyCsrfType;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Security;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Security\Csrf\CsrfToken;

class ValidationController extends Controller
{
    /**
     * @Route("/validate_user_by_group", name="validate_user_by_group")
     */
    public function validateUserByGroupAction(Request $request)
    {
        /** @var \Symfony\Component\Validator\Validator\RecursiveValidator */
        $validator = $this->get('validator');
        $user = new UserForValidation();
        $errors = $validator->validate($user, null, array('g1'));
        var_dump((string)$errors);
        return $this->render('AppBundle:default:index.html.twig');
    }
}
