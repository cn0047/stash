<?php

namespace AppBundle\Controller;

use Symfony\Component\HttpFoundation\Request;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Method;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Template;
use AppBundle\Entity\BookAutor;
use AppBundle\Form\BookAutorType;

/**
 * BookAutor controller.
 *
 * @Route("/bookautor")
 */
class BookAutorController extends Controller
{

    /**
     * Lists all BookAutor entities.
     *
     * @Route("/", name="bookautor")
     * @Method("GET")
     * @Template()
     */
    public function indexAction()
    {
        $em = $this->getDoctrine()->getManager();

        $entities = $em->getRepository('AppBundle:BookAutor')->findAll();

        return array(
            'entities' => $entities,
        );
    }
    /**
     * Creates a new BookAutor entity.
     *
     * @Route("/", name="bookautor_create")
     * @Method("POST")
     * @Template("AppBundle:BookAutor:new.html.twig")
     */
    public function createAction(Request $request)
    {
        $entity = new BookAutor();
        $form = $this->createCreateForm($entity);
        $form->handleRequest($request);

        if ($form->isValid()) {
            $em = $this->getDoctrine()->getManager();
            $em->persist($entity);
            $em->flush();

            return $this->redirect($this->generateUrl('bookautor_show', array('id' => $entity->getId())));
        }

        return array(
            'entity' => $entity,
            'form'   => $form->createView(),
        );
    }

    /**
     * Creates a form to create a BookAutor entity.
     *
     * @param BookAutor $entity The entity
     *
     * @return \Symfony\Component\Form\Form The form
     */
    private function createCreateForm(BookAutor $entity)
    {
        $form = $this->createForm(new BookAutorType(), $entity, array(
            'action' => $this->generateUrl('bookautor_create'),
            'method' => 'POST',
        ));

        $form->add('submit', 'submit', array('label' => 'Create'));

        return $form;
    }

    /**
     * Displays a form to create a new BookAutor entity.
     *
     * @Route("/new", name="bookautor_new")
     * @Method("GET")
     * @Template()
     */
    public function newAction()
    {
        $entity = new BookAutor();
        $form   = $this->createCreateForm($entity);

        return array(
            'entity' => $entity,
            'form'   => $form->createView(),
        );
    }

    /**
     * Finds and displays a BookAutor entity.
     *
     * @Route("/{id}", name="bookautor_show")
     * @Method("GET")
     * @Template()
     */
    public function showAction($id)
    {
        $em = $this->getDoctrine()->getManager();

        $entity = $em->getRepository('AppBundle:BookAutor')->find($id);

        if (!$entity) {
            throw $this->createNotFoundException('Unable to find BookAutor entity.');
        }

        $deleteForm = $this->createDeleteForm($id);

        return array(
            'entity'      => $entity,
            'delete_form' => $deleteForm->createView(),
        );
    }

    /**
     * Displays a form to edit an existing BookAutor entity.
     *
     * @Route("/{id}/edit", name="bookautor_edit")
     * @Method("GET")
     * @Template()
     */
    public function editAction($id)
    {
        $em = $this->getDoctrine()->getManager();

        $entity = $em->getRepository('AppBundle:BookAutor')->find($id);

        if (!$entity) {
            throw $this->createNotFoundException('Unable to find BookAutor entity.');
        }

        $editForm = $this->createEditForm($entity);
        $deleteForm = $this->createDeleteForm($id);

        return array(
            'entity'      => $entity,
            'edit_form'   => $editForm->createView(),
            'delete_form' => $deleteForm->createView(),
        );
    }

    /**
    * Creates a form to edit a BookAutor entity.
    *
    * @param BookAutor $entity The entity
    *
    * @return \Symfony\Component\Form\Form The form
    */
    private function createEditForm(BookAutor $entity)
    {
        $form = $this->createForm(new BookAutorType(), $entity, array(
            'action' => $this->generateUrl('bookautor_update', array('id' => $entity->getId())),
            'method' => 'PUT',
        ));

        $form->add('submit', 'submit', array('label' => 'Update'));

        return $form;
    }
    /**
     * Edits an existing BookAutor entity.
     *
     * @Route("/{id}", name="bookautor_update")
     * @Method("PUT")
     * @Template("AppBundle:BookAutor:edit.html.twig")
     */
    public function updateAction(Request $request, $id)
    {
        $em = $this->getDoctrine()->getManager();

        $entity = $em->getRepository('AppBundle:BookAutor')->find($id);

        if (!$entity) {
            throw $this->createNotFoundException('Unable to find BookAutor entity.');
        }

        $deleteForm = $this->createDeleteForm($id);
        $editForm = $this->createEditForm($entity);
        $editForm->handleRequest($request);

        if ($editForm->isValid()) {
            $em->flush();

            return $this->redirect($this->generateUrl('bookautor_edit', array('id' => $id)));
        }

        return array(
            'entity'      => $entity,
            'edit_form'   => $editForm->createView(),
            'delete_form' => $deleteForm->createView(),
        );
    }
    /**
     * Deletes a BookAutor entity.
     *
     * @Route("/{id}", name="bookautor_delete")
     * @Method("DELETE")
     */
    public function deleteAction(Request $request, $id)
    {
        $form = $this->createDeleteForm($id);
        $form->handleRequest($request);

        if ($form->isValid()) {
            $em = $this->getDoctrine()->getManager();
            $entity = $em->getRepository('AppBundle:BookAutor')->find($id);

            if (!$entity) {
                throw $this->createNotFoundException('Unable to find BookAutor entity.');
            }

            $em->remove($entity);
            $em->flush();
        }

        return $this->redirect($this->generateUrl('bookautor'));
    }

    /**
     * Creates a form to delete a BookAutor entity by id.
     *
     * @param mixed $id The entity id
     *
     * @return \Symfony\Component\Form\Form The form
     */
    private function createDeleteForm($id)
    {
        return $this->createFormBuilder()
            ->setAction($this->generateUrl('bookautor_delete', array('id' => $id)))
            ->setMethod('DELETE')
            ->add('submit', 'submit', array('label' => 'Delete'))
            ->getForm()
        ;
    }
}
