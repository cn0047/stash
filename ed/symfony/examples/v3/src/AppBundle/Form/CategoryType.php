<?php

namespace AppBundle\Form;

use Symfony\Component\Form\AbstractType;
use Symfony\Component\Form\Extension\Core\Type\TextType;
use Symfony\Component\Form\FormBuilderInterface;
use Symfony\Component\OptionsResolver\OptionsResolver;
use Symfony\Component\OptionsResolver\OptionsResolverInterface;
use Symfony\Component\Form\FormEvent;
use Symfony\Component\Form\FormEvents;

class CategoryType extends AbstractType
{
    public function getName()
    {
        return 'CategoryType';
    }

    public function setDefaultOptions(OptionsResolverInterface $resolver)
    {
        $resolver->setDefaults(array(
            'data_class' => 'AppBundle\VO\Category',
        ));
    }

    public function buildForm(FormBuilderInterface $builder, array $options)
    {
        $builder->add('name', TextType::class);
        $builder->addEventListener(FormEvents::PRE_SET_DATA, function (FormEvent $event) {
            dump('form PRE_SET_DATA');
        });
        $builder->addEventListener(FormEvents::POST_SET_DATA, function (FormEvent $event) {
            dump('form POST_SET_DATA');
        });
        $builder->addEventListener(FormEvents::PRE_SUBMIT, function (FormEvent $event) {
            dump('form PRE_SUBMIT');
        });
        $builder->addEventListener(FormEvents::SUBMIT, function (FormEvent $event) {
            dump('form SUBMIT');
        });
        $builder->addEventListener(FormEvents::POST_SUBMIT, function (FormEvent $event) {
            dump('form POST_SUBMIT');
        });
    }

    public function configureOptions(OptionsResolver $resolver)
    {
        $resolver->setDefaults(array(
            'data_class' => 'AppBundle\VO\Category',
        ));
    }
}
