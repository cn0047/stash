<?php

namespace AppBundle\Form\Type;

use Symfony\Component\Form\AbstractType;
use Symfony\Component\Form\FormBuilderInterface;
use Symfony\Component\OptionsResolver\OptionsResolver;

class Step2 extends AbstractType
{
    public function buildForm(FormBuilderInterface $builder, array $options)
    {
        $builder
            ->add('iceCream', 'text')
            ->add('superHero', 'text')
            ->add('movieStar', 'text')
            ->add('worldEnd', 'text')
            ->add('whoWinSuperBowl', 'text')
            ->add('submit', 'submit')
            ;
    }

    public function configureOptions(OptionsResolver $resolver)
    {
        $resolver->setDefaults(array(
            'data_class' => 'AppBundle\Entity\UserInfo',
        ));
    }

    public function getName()
    {
        return 'Step2';
    }
}