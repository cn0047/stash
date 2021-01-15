Continuous Integration/Continuous Delivery
-

How long would it take your organization to deploy a change that involves just one single line of code?
Do you do this on a repeatable, reliable basis?

A deployment pipeline is, in essence, an automated implementation of your application’s
build, deploy, test, and release process.

The practice of building and testing your application on every check-in
is known as continuous integration.

Continuous Integration - it is a software development practice
where members of a team use a version control system
and integrate their work frequently to the same location, such as a master branch.
Each change is built and verified by tests and other verifications
in order to detect any integration errors as quickly as possible.

Continuous Delivery - it is a software development methodology where the release process is automated.
Every software change is automatically built, tested, and deployed to production.

Continuous Deployment - is a synonym to Continuous Delivery.

<br>Pre-alpha ⇒ Alpha ⇒ Beta ⇒ Release candidate ⇒ Gold
<br> Beta ⇒ Dev; Gamma ⇒ QA;

Problems:
* db migrations (slow migrations, few steps migrations)
* `cron`

#### Deployments Strategies

Dark launching - deploying the very first version of a service
so no users available yet.

Highlander (most traditional deployment patten) - all instances are deploying simultaneously.

Canary Deployment - deploys to only a small portion of the available servers.
Some kind of A/B testing.

Rolling Deploy (continuation of the canary deploy) - update one server after another.

Blue-Green - Once you have deployed and fully tested the software in Green,
you switch the router so all incoming requests now go to Green instead of Blue.
Green is now live, and Blue is idle.

Canary with two groups - Blue-Green and add 1 node from new cluster into old one.

Rolling Deploy with two groups - continuation of the canary with two groups.

Don't forget about:
Users may see V1 of a page on one click, then see V2 on refresh and back to V1 on yet another refresh.
As solution you can suggest:
* to use separated site for some users (http://beta.yourcompany.com)
* or use "Feature Toggles"
* or A/B
