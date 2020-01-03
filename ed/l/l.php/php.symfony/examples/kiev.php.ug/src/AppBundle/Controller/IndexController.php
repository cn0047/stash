<?php

namespace AppBundle\Controller;

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\BinaryFileResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;

class IndexController extends Controller
{
    /**
     * @Route("/", name="home")
     */
    public function homeAction(Request $request)
    {
        $response = $this->render('AppBundle:Index:Index.html.twig');
        $response->setETag(md5($response->getContent()));
        $response->setPublic();
        $response->isNotModified($request);
        return $response;
    }

    /**
     * @Route("/pdf/{pdf}", name="pdf")
     */
    public function pdfAction($pdf)
    {
        return new BinaryFileResponse("bundles/app/pdf/$pdf");
    }

    /**
     * @Route("/testMail", name="testMail")
     */
    public function testMailAction()
    {
        $m = $this->get('mail');
        $r = $m->sendMail('testMail', ['mail' => 'cn007b@gmail.com']);
        return new Response(var_export($r, true));
    }

    /**
     * @Route("/testMail2", name="testMail2")
     */
    public function testMail2Action()
    {
        $m = $this->get('mail');
        $r = $m->sendMail('testMail2', ['mail' => 'cn007b@gmail.com']);
        return new Response(var_export($r, true));
    }

    /**
     * @Route("/test", name="test")
     */
    public function testAction()
    {
    }
}
