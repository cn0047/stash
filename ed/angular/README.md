Angular
-

Decorator - it's function that adds metadata to a class (prefixed with `@`).
Example: @Component({ selector: '', template: `<div></div>`, templateUrl: '', styleUrls: ['']});

Binding - coordinates communication between the component's class and its template.

Directive - custom html element or attribute used to power up and extend our html.
`*ngIf`, `*ngFor` - stuctural built-in directives.

Lifecycle Hooks:
* OnInit
* OnChanges
* OnDestroy

@Input - pass data to nested component.
@Output - raising an event.

````
observable.subscribe(valueFn, errFn, completeFn) # http

# in template
{{product?.productName}} # safe navigation
````

Routes Guards:
* CanActivate
* CanDeactivate
* Resolve
* CanLoad

````
# cli
# npm i -D @angular/cli

angular-cli new hello-world

angular-cli g m products/product --flat -m app.module
````

    https://app.pluralsight.com/library/courses/angular-2-getting-started-update/table-of-contents
https://app.pluralsight.com/library/courses/angular-cli
https://app.pluralsight.com/library/courses/angular-2-forms
https://app.pluralsight.com/library/courses/angular-fundamentals
https://app.pluralsight.com/library/courses/angular-routing
https://app.pluralsight.com/library/courses/angular-2-reactive-forms
https://app.pluralsight.com/library/courses/best-practices-angular
