<?php

class ArrayOfChar32Validator extends CValidator
{
    public $allowEmpty;

    public function validateAttribute($object, $attributeName)
    {
        if (!is_array($object->$attributeName)) {
            $this->addError($object, $attributeName, "$attributeName must be array.");
            return;
        }
        if (empty($object->$attributeName) and !$this->allowEmpty) {
            $this->addError($object, $attributeName, "$attributeName cannot be empty array.");
            return;
        }
        foreach ($object->$attributeName as $key => $value) {
            if (!preg_match('/^[\w\d]{32}$/', $value)) {
                $this->addError($object, $attributeName, "$attributeName contains invalid value: $value.");
            }
        }
    }
}
