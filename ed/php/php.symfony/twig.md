Twig
-

````php
// Symfony actually looks in two different locations for the template:
1. app/Resources/AcmeBlogBundle/views/Blog/index.html.twig
2. src/Acme/BlogBundle/Resources/views/Blog/index.html.twig

// Output Escaping
// If you're using twig templates, then output escaping is ON by default.
{{ article.body|raw }} - render normally
Hello <?php echo $view->escape($name) ?>
var myMsg = 'Hello <?php echo $view->escape($name, 'js') ?>';

// Syntax Checking.
php bin/console twig:lint app/Resources/views/article/recent_list.html.twig
php bin/console twig:lint app/Resources/views

// Debug.
{{ dump(articles) }}

// Global Template Variables.
app.security
app.user
app.request
app.session
app.environment
app.debug

{{ ... }}         - variable or the result of an expression
{% ... %}         - a tag that controls the logic of the template
{# ... #}         - comment
{{ title|upper }} - filters
{{ parent() }}    - block from the parent template

{{ include('article/article_details.html.twig', { 'article': article }) }}

{% if items is not empty %}
{% for item in items  %}
    {{ item.tariffNum }}
{% endfor %}
{% endif %}

// Links.
<a href="{{ path('_welcome') }}">Home</a>
<a href="{{ path('article_show', {'slug': article.slug}) }}"> {{ article.title }} </a>
// Absolute URL.
<a href="{{ url('_welcome') }}">Home</a>
<a href="<?php echo $view['router']->generate('_welcome', array(), true) ?>">Home</a>

// Assets.
<img src="{{ asset('images/logo.png') }}" alt="Symfony!" />
<link href="{{ asset('css/blog.css') }}" rel="stylesheet" type="text/css" />
<img src="{{ asset('images/logo.png', version='3.0') }}" alt="Symfony!" />
<img src="{{ asset('images/logo.png', absolute=true) }}" alt="Symfony!" />

// Stylesheets and JavaScripts
{% block stylesheets %}
    <link href="{{ asset('css/main.css') }}" rel="stylesheet" />
{% endblock %}
{% block javascripts %}
    <script src="{{ asset('js/main.js') }}"></script>
{% endblock %}
<link href="{{ asset('bundles/acmedemo/css/contact.css') }}" rel="stylesheet" />

<script src="{{ asset('js/script.js') }}"></script>
{% set seriesSendTimeUtc = (series.send_time_utc is defined) ? series.send_time_utc :  null %}
````
````sh
# Assets
app/console assets:install web/

# Dumping Asset Files
php bin/console assetic:dump --env=prod --no-debug
php bin/console assetic:dump
php bin/console assetic:watch
````

Template Suffix:

| Filename | Format | Engine |
|--------|------|------|
|blog/index.html.twig|HTML|Twig|
|blog/index.html.php|HTML|PHP|
|blog/index.css.twig|CSS|Twig|

