// Parent constructor
function Article() {
    this.tags = ['js', 'css'];
}
// Child
function BlogPost() {}
var article = new Article();
BlogPost.prototype = article;
function StaticPage () {
    Article.call(this);
}

var blog = new BlogPost();
var page = new StaticPage();
blog.tags.push('html');
page.tags.push('php');
console.log(article.hasOwnProperty('tags')); // true
console.log(blog.hasOwnProperty('tags'));    // false
console.log(page.hasOwnProperty('tags'));    // true
console.log(article.tags.join(', ')); // js, css, html
console.log(page.tags.join(', ')); // js, css, php
