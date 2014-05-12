// родительский конструктор
function Article() {
    this.tags = [‘js’, ‘css’];
}
var article = new Article();
// объект сообщения в блоге наследует свойства объекта article
// через классический шаб­он No1
function BlogPost() {}
BlogPost.prototype = article;
var blog = new BlogPost();
// обратите внимание, что выше нет необходимости
// использовать выражение `new Article()`,
// потому что уже имеется доступный экземпляр
// статическая страница наследует свойства объекта article
// через шаб­он заимствования конструктора
function StaticPage() {
    Article.call(this);
}

var page = new StaticPage();
alert(article.hasOwnProperty(‘tags’)); // true
alert(blog.hasOwnProperty(‘tags’));    // false
alert(page.hasOwnProperty(‘tags’));    // true

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
blog.tags.push(‘html’);
page.tags.push(‘php’);
alert(article.tags.join(‘, ‘)); // “js, css, html”
