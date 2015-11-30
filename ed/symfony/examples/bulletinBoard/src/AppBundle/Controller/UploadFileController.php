<?php

namespace AppBundle\Controller;

use AppBundle\Entity\Document;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Template;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;

class UploadFileController extends Controller
{
    /**
    * @Route("/uploadFile", name="uploadFile")
    * @Template()
    */
    public function uploadFileAction(Request $request)
    {
        $document = new Document();
        $form = $this->createFormBuilder($document)
            ->add('name')
            ->add('file')
            ->add('save', 'submit')
            ->getForm()
        ;
        $form->handleRequest($request);
        if ($form->isValid()) {
            $em = $this->getDoctrine()->getManager();
            $document->upload();
            $em->persist($document);
            $em->flush();
            var_dump('OK 200');
            var_dump($document);
        }
        return $this->render('default/newTask.html.twig', array(
            'form' => $form->createView(),
        ));
    }
}
