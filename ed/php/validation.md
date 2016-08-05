Validation
-

#### Custom validation

````php
<?php

namespace W3\Ziipr\Phalcon\Validation;

use Phalcon\Validation;
use W3\Ziipr\Exceptions\ValidationException;

class BaseValidation extends Validation
{
    /**
     * Wrapper for core phalcon method validate that helps to throw ValidationException.
     *
     * @param array $data Data being validated.
     *
     * @throws ValidationException In case when validation failed.
     */
    public function tryValidate($data)
    {
        $messages = parent::validate($data);
        if (count($messages) > 0) {
            throw new ValidationException($messages);
        }
    }
}
````

````php
<?php

namespace W3\Ziipr\Services\v1\SummaryScreen\VO;

use Phalcon\Validation\Message;
use Phalcon\Validation\Message\Group;
use W3\Ziipr\Phalcon\Validation\BaseValidation;

class UsersUnziipedValidation extends BaseValidation
{
    /**
     * Validate data.
     *
     * @param array $data Data being validated.
     * @param object $entity Entity.
     * @param Group $messages Collection of messages.
     *
     * @return bool Success of validation.
     */
    public function afterValidation(array $data, $entity, Group $messages)
    {
        $success = true;
        //
        if (!array_key_exists('userId', $data)) {
            $messages->appendMessage(new Message('Parameter userId - cannot be blank.'));
            $success = false;
        }
        if (!array_key_exists('toUserId', $data)) {
            $messages->appendMessage(new Message('Parameter toUserId - cannot be blank.'));
            $success = false;
        }
        if (!array_key_exists('from', $data)) {
            $messages->appendMessage(new Message('Parameter from - cannot be blank.'));
            $success = false;
        }
        if (!array_key_exists('limit', $data)) {
            $messages->appendMessage(new Message('Parameter limit - cannot be blank.'));
            $success = false;
        }
        if (!$success) {
            return false;
        }
        //
        if ($data['userId'] === null && $data['toUserId'] === null) {
            $messages->appendMessage(new Message('Both parameters userId and toUserId cannot be null simultaneously.'));
            $success = false;
        }
        if ($data['userId'] !== null && !is_int($data['userId'])) {
            $messages->appendMessage(new Message('Parameter userId - must be null or integer.'));
            $success = false;
        }
        if ($data['toUserId'] !== null && !is_int($data['toUserId'])) {
            $messages->appendMessage(new Message('Parameter toUserId - must be null or integer.'));
            $success = false;
        }
        if (!ctype_digit($data['from'])) {
            $messages->appendMessage(new Message('Parameter from - must consist of digits.'));
            $success = false;
        }
        if (!ctype_digit($data['limit'])) {
            $messages->appendMessage(new Message('Parameter limit - must consist of digits.'));
            $success = false;
        }
        return $success;
    }
}
````

````php
$validation = new UsersUnziipedValidation();
$validation->tryValidate($args);
````
