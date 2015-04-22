<?php

namespace AppBundle\Controller;

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\HttpFoundation\Request;
use AppBundle\Entity\Category;
use AppBundle\Entity\Product;
use AppBundle\Entity\Task;
use AppBundle\Form\Type\TaskType;

class DefaultController extends Controller
{
    /**
     * @Route("/app/example", name="homepage")
     */
    public function indexAction()
    {
        return $this->render('default/index.html.twig');
    }

    /**
     * @Route("/doctrine/product/initialize", name="doctrineProductInitialize")
     */
    public function productInitialize()
    {
        /*
        drop table product;
        create table product (
            id int auto_increment,
            name varchar(100),
            price decimal,
            description text,
            category_id int,
            primary key (id)
        );
        create table category (
            id int auto_increment,
            name varchar(100),
            primary key (id)
        );
        */
        return $this->render('default/message.html.twig', ['m' => 100]);
    }

    /**
     * @Route("/app/example/create", name="doctrineProductCreate")
     */
    public function createAction()
    {
        $product = new Product();
        $product->setName('A Foo Bar');
        $product->setPrice('19.99');
        $product->setDescription('Lorem ipsum dolor');
        $em = $this->getDoctrine()->getManager();
        $em->persist($product);
        $em->flush();
        return $this->render(
            'default/message.html.twig',
            ['m' => 'Created product id '.$product->getId()]
        );
    }

    /**
     * @Route("/app/example/show/{id}", defaults={"id": 2}, name="doctrineProductShow")
     */
    public function showAction($id)
    {
        $product = $this->getDoctrine()
            ->getRepository('AppBundle:Product')
            ->find($id);
        if (!$product) {
            throw $this->createNotFoundException('No product found for id '.$id);
        }
        dump($product);
        return $this->render(
            'default/message.html.twig',
            ['m' => var_export($product, 1)]
        );
    }

    /**
     * @Route("/app/example/show2", name="doctrineProductShow2")
     */
    public function show2Action()
    {
        $repository = $this->getDoctrine()
            ->getRepository('AppBundle:Product');
        $id = 3;
        $product = $repository->find($id);
        dump($product);
        $product = $repository->findOneById($id);
        dump($product);
        $product = $repository->findOneByName('foo');
        dump($product);
        $product = $repository->findOneByName('A Foo Bar');
        dump($product);
        // find *all* products
        $products = $repository->findAll();
        dump($products);
        // // find a group of products based on an arbitrary column value
        $products = $repository->findByPrice('19.99');
        dump($products);
        // // query for one product matching by name and price
        $product = $repository->findOneBy(
            array('name' => 'A Foo Bar', 'price' => '19.99')
        );
        dump($product);
        // // query for all products matching the name, ordered by price
        $products = $repository->findBy(
            array('name' => 'foo'),
            array('price' => 'ASC')
        );
        return $this->render('default/message.html.twig', ['m' => '']);
    }

    /**
     * @Route("/app/example/show3", name="doctrineProductShow3")
     */
    public function show3Action()
    {
        $repository = $this->getDoctrine()
           ->getRepository('AppBundle:Product');
        $query = $repository->createQueryBuilder('p')
            ->where('p.price > :price')
            ->setParameter('price', '19.99')
            ->orderBy('p.price', 'ASC')
            ->getQuery();
        $products = $query->getResult();
        dump($products);
        return $this->render(
            'default/message.html.twig',
            ['m' => var_export($products, 1)]
        );
    }

    /**
     * @Route("/app/example/show4", name="doctrineProductShow4")
     */
    public function show4Action()
    {
        $em = $this->getDoctrine()->getManager();
        $query = $em->createQuery(
            'SELECT p
            FROM AppBundle:Product p
            WHERE p.price > :price
            ORDER BY p.price ASC'
        )->setParameter('price', '19.99');
        $products = $query->getResult();
        dump($products);
        return $this->render(
            'default/message.html.twig',
            ['m' => var_export($products, 1)]
        );
    }

    /**
     * @Route("/app/example/update/{id}", defaults={"id": 1}, name="doctrineProductUpdate")
     */
    public function updateAction($id)
    {
        $em = $this->getDoctrine()->getManager();
        $product = $em->getRepository('AppBundle:Product')->find($id);
        if (!$product) {
            throw $this->createNotFoundException('No product found for id '.$id);
        }
        $product->setName('New product name!');
        $r = $em->flush();
        return $this->render(
            'default/message.html.twig',
            ['m' => var_export($r, 1)]
        );
    }

    public function deleteAction()
    {
        $em->remove($product);
        $em->flush();
    }

    /**
     * @Route("/doctrine/product/initialize2", name="doctrineProductInitialize2")
     */
    public function createProductAction()
    {
        $category = new Category();
        $category->setName('Main Products');
        $product = new Product();
        $product->setName('Foo');
        $product->setPrice(19.99);
        $product->setDescription('Lorem ipsum dolor');
        // relate this product to the category
        $product->setCategory($category);
        $em = $this->getDoctrine()->getManager();
        $em->persist($category);
        $em->persist($product);
        $em->flush();
        return $this->render(
            'default/message.html.twig',
            ['m' => 'Created product id: '.$product->getId().' and category id: '.$category->getId()]
        );
    }

    /**
     * @Route("/app/example/show5/{id}", defaults={"id": 1}, name="doctrineProductShow5")
     *
     * @todo FIX BUG HERE.
     */
    public function show5Action($id)
    {
        $categoryName = 'none';
        $product = $this->getDoctrine()
           ->getRepository('AppBundle:Product')
            ->find($id);
        $categoryName = $product->getCategory()->getName();
        dump($product);
        $category = $this->getDoctrine()
            ->getRepository('AppBundle:Category')
            ->find($id);
        $products = $category->getProducts();
        dump($products);
        return $this->render(
            'default/message.html.twig',
            ['m' => '']
        );
    }

    public function show6Action($id)
    {
        $product = $this->getDoctrine()
           ->getRepository('AppBundle:Product')
            ->findOneByIdJoinedToCategory($id);
        $category = $product->getCategory();
    }

    /**
     * @Route("/newTask", name="newTask")
     */
    public function newTaskAction(Request $request)
    {
        // create a task and give it some dummy data for this example
        $task = new Task();
        $task->setTask('Write a blog post');
        $task->setDueDate(new \DateTime('tomorrow'));
        $form = $this->createFormBuilder($task)
           ->add('task', 'text')
            ->add('dueDate', 'date')
            ->add('save', 'submit', array('label' => 'Create Task'))
            ->add('saveAndAdd', 'submit', array('label' => 'Save and Add'))
            ->getForm();
        $form->handleRequest($request);
        if ($form->isValid()) {
            $nextAction = $form->get('saveAndAdd')->isClicked()
                ? 'task_new'
                : 'task_success';
            var_dump($nextAction);
            // perform some action, such as saving the task to the database
            // return $this->redirectToRoute('task_success');
        }
        return $this->render('default/newTask.html.twig', array(
            'form' => $form->createView(),
        ));
    }

    /**
     * @Route("/newTask2", name="newTask2")
     */
    public function newTaskAction2()
    {
        $task = new Task();
        $form = $this->createForm(new TaskType(), $task);
        return $this->render('default/newTask.html.twig', array(
            'form' => $form->createView(),
        ));
    }

    /**
     * @Route("/newTask3", name="newTask3")
     */
    public function newTaskAction3()
    {
        $task = new Task();
        $form = $this->createForm('task', $task);
        return $this->render('default/newTask.html.twig', array(
            'form' => $form->createView(),
        ));
    }
}
