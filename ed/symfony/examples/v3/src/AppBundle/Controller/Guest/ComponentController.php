<?php

namespace AppBundle\Controller\Guest;

use AppBundle\Component\Common;
use AppBundle\Component\Main;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\Request;

class ComponentController extends Controller
{
    /**
     * @Route("/component", name="component")
     */
    public function componentAction(Request $request)
    {
        $d = $this->get('event_dispatcher');
        $c = new Common();
        $d->addListener('component.main.method_is_not_found', array($c, 'onMainMethodIsNotFound'));
        $m = new Main($d);
        $r = $m->check('this');
        return $this->render('AppBundle:default:index.html.twig', [
            'message' => $r,
        ]);
    }
}
