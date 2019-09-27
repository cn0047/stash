<?php

namespace AppBundle\Form\Type;

use Symfony\Component\Form\AbstractType;
use Symfony\Component\Form\FormBuilderInterface;
use Symfony\Component\OptionsResolver\OptionsResolver;

class Step1 extends AbstractType
{
    public function buildForm(FormBuilderInterface $builder, array $options)
    {
        $builder
            ->add('firstName', 'text')
            ->add('lastName', 'text')
            ->add('email', 'email')
            ->add('birthDay', 'date')
            ->add('shoeSize', 'integer')
            ->add('submit', 'submit')
            ;
    }

    public function buildFormUs(FormBuilderInterface $builder, array $options)
    {
        $builder
            ->add('date_from', 'text_not_null', [
                'required'    => true,
                'constraints' => [
                    new Constraints\NotBlank(['message' => t('param_exception_missing', 'date_from')]),
                    new DateTime(['message' => t('param_exception_invalid_datetime_value', 'date_from')]),
                ]
            ])
            ->add('date_to', 'text_not_null', [
                'required'    => true,
                'constraints' => [
                    new Constraints\NotBlank(['message' => t('param_exception_missing', 'date_to')]),
                    new DateTime(['message' => t('param_exception_invalid_datetime_value', 'date_to')]),
                ]
            ])
            ->add('limit', 'text_not_null', [
                'empty_data'  => 50,
                'required'    => false,
                'constraints' => [
                    new Constraints\Range(
                        [
                            'minMessage'     => t('param_exception_invalid_more_value', 'limit'),
                            'maxMessage'     => t('param_exception_invalid_less_value', 'limit'),
                            'invalidMessage' => t('param_exception_invalid_number_value', 'limit'),
                            'min'            => 1,
                            'max'            => 100,
                        ]),
                ]
            ])
            ->add('offset', 'text_not_null', [
                'empty_data'  => 0,
                'required'    => false,
                'constraints' => [
                    new Constraints\GreaterThanOrEqual(
                        [
                            'message' => t('param_exception_invalid_greater_or_equal_value', 'offset'),
                            'value'   => 0,
                        ]),
                ]
            ]);

        parent::buildForm($builder, $options);
    }

    public function configureOptions(OptionsResolver $resolver)
    {
        $resolver->setDefaults(array(
            'data_class' => 'AppBundle\Entity\User',
        ));
    }

    public function getName()
    {
        return 'Step1';
    }
}