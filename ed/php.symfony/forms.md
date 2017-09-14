Forms
-

````php
{{ form(form, {'attr': {'novalidate': 'novalidate'}}) }}

$form = $this->createFormBuilder($users, array(
    'validation_groups' => array('registration'),
))->add(...);

$form = $this->createFormBuilder($task)
    ->add('nextStep', 'submit')
    ->add('previousStep', 'submit')
    ->add('dueDate', 'date', array('widget' => 'single_text'))
    ->add('dueDate', 'date', array(
        'widget' => 'single_text',
        'label' => 'Due Date',
    ))
    ->getForm();

$form = $this->createFormBuilder($task)
    ->setAction($this->generateUrl('target_route'))
    ->setMethod('GET')
    ->add('task', 'text')
    ->add('dueDate', 'date')
    ->add('save', 'submit')
    ->add('message_format', 'choice', [
        'required'    => false,
        'empty_data'  => null,
        'choices' => [
            'raw_html' => 'raw_html',
            'block' => 'block',
            'text' => 'text',
        ],
    ])
    ->getForm();

$form = $this->createForm(new TaskType(), $task, array(
    'action' => $this->generateUrl('target_route'),
    'method' => 'GET',
));

if ($form->isValid()) {
    $em = $this->getDoctrine()->getManager();
    $em->persist($task);
    $em->flush();
    return $this->redirectToRoute('task_success');
}

# CSRF Protection
use Symfony\Component\OptionsResolver\OptionsResolverInterface;
class TaskType extends AbstractType
{
    public function setDefaultOptions(OptionsResolverInterface $resolver)
    {
        $resolver->setDefaults(array(
            'data_class' => 'AppBundle\Entity\Task',
            'csrf_protection' => true,
            'csrf_field_name' => '_token',
            // a unique key to help generate the secret token
            'intention' => 'task_item',
        ));
    }
}

# Using a Form without a Class
use Symfony\Component\HttpFoundation\Request;
// make sure you've imported the Request namespace above the class
public function contactAction(Request $request)
{
    $defaultData = array('message' => 'Type your message here');
    $form = $this->createFormBuilder($defaultData)
        ->add('name', 'text')
        ->add('email', 'email')
        ->add('message', 'textarea')
        ->add('send', 'submit')
        ->getForm();
    $form->handleRequest($request);
    if ($form->isValid()) {
        // data is an array with "name", "email", and "message" keys
        $data = $form->getData();
    }
}
````

Form Events:

* PRE_SUBMIT - at the beginning of the Form::submit()
* SUBMIT - just before the Form::submit()
* POST_SUBMIT - after the Form::submit()
* PRE_SET_DATA - at the beginning of the Form::setData()
* POST_SET_DATA - at the end of the Form::setData()

setData - `$form->get('firstname')->setData('John');`

````
$form = $formFactory->createBuilder()
    // ... add form fields
    ->addEventListener(FormEvents::PRE_SUBMIT, $listener);
````

````twig
# Built-in Field Types:
// Text Fields
• text
• textarea
• email
• integer
• money
• number
• password
• percent
• search
• url

// Choice Fields
• choice
• entity
• country
• language
• locale
• timezone
• currency

// Date and Time Fields
• date
• datetime
• time
• birthday

// Other Fields
• checkbox
• file
• radio

// Field Groups
• collection
• repeated

// Hidden Fields
• hidden

// Buttons
• button
• reset
• submit

// Base Fields
• form

# Rendering a Form in a Template
{{ form_start(form) }}
    {{ form_errors(form) }}
    {{ form_row(form.task) }}
    {{ form_row(form.dueDate) }}
    {{ form.vars.value.task }}
{{ form_end(form) }}

{{ form_start(form) }}
    {{ form_errors(form) }}
    <div>
        {{ form_label(form.task) }}
        {{ form_errors(form.task) }}
        {{ form_widget(form.task) }}
    </div>
    <div>
        {{ form_label(form.dueDate) }}
        {{ form_errors(form.dueDate) }}
        {{ form_widget(form.dueDate) }}
    </div>
    <div>
        {{ form_widget(form.save) }}
    </div>
{{ form_end(form) }}

{{ form_label(form.task, 'Task Description') }}

{{ form_widget(form.task, {'attr': {'class': 'task_field'}}) }}

{{ form.task.vars.id }}
{{ form.task.vars.full_name }}

{{ form_start(form, {'action': path('target_route'), 'method': 'GET'}) }}

$form->get('dueDate')->getData();
$form->get('dueDate')->setData(new \DateTime());

# Global Form Theming
// app/config/config.yml
twig:
    form_themes:
        - 'form/fields.html.twig'

// Customizing Form Output all in a Single File with Twig
{% extends 'base.html.twig' %}

{# import "_self" as the form theme #}
{% form_theme form _self %}

{# make the form fragment customization #}
{% block form_row %}
    {# custom field row output #}
{% endblock form_row %}

{% block content %}
    {{ form_row(form.task) }}
{% endblock %}
````
