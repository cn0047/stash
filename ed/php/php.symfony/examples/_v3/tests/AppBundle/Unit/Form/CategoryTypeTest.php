<?php

namespace Tests\AppBundle\Form\Type;

use AppBundle\Form\CategoryType;
use AppBundle\Model\TestObject;

use Symfony\Component\Form\Test\TypeTestCase;

class CategoryTypeTest extends TypeTestCase
{
    public function testSubmitValidData()
    {
        $formData = array(
            'name' => 'test',
        );
        $type = new CategoryType();
        $type = 'AppBundle\Form\CategoryType';
        $form = $this->factory->create($type);
        $form->submit($formData);
        $this->assertTrue($form->isSynchronized());
    }
}
