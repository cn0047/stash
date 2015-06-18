<?php

class TestForm extends CFormModel
{
    public $ids;
    public $adIds;

    public function rules()
    {
        return [
            ['ids', 'arrayOfInt', 'allowEmpty' => false],
            ['adIds', 'ArrayOfChar32Validator', 'allowEmpty' => false],
        ];
    }

    public function arrayOfInt($attributeName, $params)
    {
        $allowEmpty = false;
        if (isset($params['allowEmpty']) and is_bool($params['allowEmpty'])) {
            $allowEmpty = $params['allowEmpty'];
        }
        if (!is_array($this->$attributeName)) {
            $this->addError($attributeName, "$attributeName must be array.");
        }
        if (empty($this->$attributeName) and !$allowEmpty) {
            $this->addError($attributeName, "$attributeName cannot be empty array.");
        }
        foreach ($this->$attributeName as $key => $value) {
            if (!is_int($value)) {
                $this->addError($attributeName, "$attributeName contains invalid value: $value.");
            }
        }
    }
}
