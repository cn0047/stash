Validation
-

````php
// Basic Constraints
• NotBlank
• Blank
• NotNull
• Null
• True
• False
• Type

// String Constraints
• Email
• Length
• Url
• Regex
• Ip
• Uuid

// Number Constraints
• Range

// Comparison Constraints
• EqualTo
• NotEqualTo
• IdenticalTo
• NotIdenticalTo
• LessThan
• LessThanOrEqual
• GreaterThan
• GreaterThanOrEqual

// Date Constraints
• Date
• DateTime
• Time

// Collection Constraints
• Choice
• Collection
• Count
• UniqueEntity
• Language
• Locale
• Country

// File Constraints
• File(mimeTypes={ "application/pdf" })
• Image

// Financial and other Number Constraints
• CardScheme
• Currency
• Luhn
• Iban
• Isbn
• Issn

// Other Constraints
• Callback
• Expression
• All
• UserPassword
• Valid

/**
 * @Assert\Choice(
 * choices = { "male", "female" },
 * message = "Choose a valid gender."
 * )
 */
public $gender;

/**
 * @Assert\True(message = "The password cannot match your first name")
 */
public function isPasswordLegal() { return $this->firstName !== $this->password; }

// group
$errors = $validator->validate($author, null, array('registration'));

$violations = $this->get('validator')->validate(
    $request->get('id'),
    [
        new \Symfony\Component\Validator\Constraints\NotNull(),
        new \Symfony\Component\Validator\Constraints\Type(['type' => 'digit']),
    ]
);
if ($violations->count() > 0) {
    $error = '';
    foreach ($violations as $v) {
        $error .= ' '.$v->getMessage();
    }
    throw new \UnexpectedValueException($m);
}
$errors = [];
foreach ($form->all() as $field) {
    $errors[$field->getName()] = trim((string)$field->getErrors());
}
var_dump($errors);
````
