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

#### Security
````php
# app/config/security.yml
security:
    providers:
        in_memory:
            memory: ~
    firewalls:
        dev:
            pattern: ^/(_(profiler|wdt)|css|images|js)/
            security: false
        default:
            anonymous: ~
            http_basic: ~

public function helloAction($name)
{
    // The second parameter is used to specify on what object the role is tested.
    $this->denyAccessUnlessGranted('ROLE_ADMIN', null, 'Unable to access this page!');
}

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Security;
/**
 * @Security("has_role('ROLE_ADMIN')")
 */
public function helloAction($name)
{
}

public function helloAction($name)
{
    if (!$this->get('security.authorization_checker')->isGranted('IS_AUTHENTICATED_FULLY')) {
        throw $this->createAccessDeniedException();
    }
    $user = $this->getUser();
    // the above is a shortcut for this
    $user = $this->get('security.token_storage')->getToken()->getUser();
    return new Response('Well hi there '.$user->getFirstName());
}

{% if is_granted('ROLE_ADMIN') %}
    <a href="...">Delete</a>
{% endif %}

{% if app.user and is_granted('ROLE_ADMIN') %}

{% if is_granted(expression('"ROLE_ADMIN" in roles or (user and user.isSuperAdmin())')) %}
    <a href="...">Delete</a>
{% endif %}

{% if is_granted('IS_AUTHENTICATED_FULLY') %}
    <p>Username: {{ app.user.username }}</p>
{% endif %}

# Hierarchical Roles
security:
    role_hierarchy:
        ROLE_ADMIN: ROLE_USER
        ROLE_SUPER_ADMIN: [ROLE_ADMIN, ROLE_ALLOWED_TO_SWITCH]
````
