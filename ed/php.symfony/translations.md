Translations
-

````php
echo $translator->trans('Hello World');

# app/config/config.yml
framework:
    translator: { fallbacks: [en] }
    default_locale: en

# app/config/routing.yml
contact:
    path: /{_locale}/contact
        defaults: { _controller: AppBundle:Contact:index }
        requirements:
            _locale: en|fr|de

public function indexAction()
{
    $locale = $request->getLocale();
    $request->getSession()->set('_locale', $locale);
    $translated = $this->get('translator')->trans('Symfony is great');
    return new Response($translated);
}

<!-- messages.fr.xlf -->
<?xml version="1.0"?>
<xliff version="1.2" xmlns="urn:oasis:names:tc:xliff:document:1.2">
    <file source-language="en" datatype="plaintext" original="file.ext">
        <body>
            <trans-unit id="1">
                <source>Symfony is great</source>
                <target>J'aime Symfony</target>
            </trans-unit>
        </body>
    </file>
</xliff>
````

````twig
{% trans %}Hello %name%{% endtrans %}
{% transchoice count %}
    {0} There are no apples|{1} There is one apple|]1,Inf] There are %count% apples
{% endtranschoice %}

{% trans with {'%name%': 'Fabien'} from "app" %}Hello %name%{% endtrans %}
{% trans with {'%name%': 'Fabien'} from "app" into "fr" %}Hello %name%{% endtrans %}
{% transchoice count with {'%name%': 'Fabien'} from "app" %}
    {0} %name%, there are no apples|{1} %name%, there is one apple|]1,Inf] %name%,
    there are %count% apples
{% endtranschoice %}

{{ message|trans }}
{{ message|transchoice(5) }}
{{ message|trans({'%name%': 'Fabien'}, "app") }}
{{ message|transchoice(5, {'%name%': 'Fabien'}, 'app') }}

{# text translated between tags is never escaped #}
{% trans %}
    <h3>foo</h3>
{% endtrans %}
{% set message = '<h3>foo</h3>' %}
{# strings and variables translated via a filter are escaped by default #}
{{ message|trans|raw }}
{{ '<h3>bar</h3>'|trans|raw }}

{% trans_default_domain "app" %}

<?php echo $view['translator']->trans('Symfony is great') ?>
<?php echo $view['translator']->transChoice(
    '{0} There are no apples|{1} There is one apple|]1,Inf[ There are %count% apples',
    10,
    array('%count%' => 10)
) ?>

// Each time you create a new translation resource
php bin/console cache:clear

class Author
{
    /**
    * @Assert\NotBlank(message = "author.name.not_blank")
    */
    public $name;
}

{% trans %}Symfony2 is great{% endtrans %}
{{ 'Symfony2 is great'|trans }}
{{ 'Symfony2 is great'|transchoice(1) }}
{% transchoice 1 %}Symfony2 is great{% endtranschoice %}

$view['translator']->trans("Symfony2 is great");
$view['translator']->transChoice('Symfony2 is great', 1);

{% set message = 'Symfony2 is great' %}
{{ message|trans }}
````

````bash
// To inspect all messages in the fr locale for the AcmeDemoBundle, run:
php bin/console debug:translation fr AcmeDemoBundle
php bin/console debug:translation en AcmeDemoBundle --domain=messages

php bin/console debug:translation en AcmeDemoBundle --only-unused
php bin/console debug:translation en AcmeDemoBundle --only-missing
````
