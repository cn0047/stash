<?php

namespace AppBundle\Form\Type;

use Symfony\Component\Form\AbstractType;
use Symfony\Component\Form\FormBuilderInterface;
use Symfony\Component\OptionsResolver\OptionsResolverInterface;
use AppBundle\Entity\Task;

class TaskAddType extends AbstractType
{
    public function setDefaultOptions(OptionsResolverInterface $resolver)
    {
        $resolver->setDefaults(array(
            'data_class' => 'AppBundle\Entity\Task',
        ));
    }

    public function buildForm(FormBuilderInterface $builder, array $options)
    {
        $builder
            ->add('task', 'text', ['constraints' => [
                new \Symfony\Component\Validator\Constraints\NotBlank(),
                new \Symfony\Component\Validator\Constraints\Length(['min' => 3]),
                new \AppBundle\Validator\Constraints\ContainsAlphanumeric(),
            ]])
            ->add('dueDate', null, array('widget' => 'single_text'))
            ->add('dueDate2', null, array('mapped' => false))
            ->add('save', 'submit');
    }

    public function getName()
    {
        return 'taskAdd';
    }
}
