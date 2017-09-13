Security
-

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
