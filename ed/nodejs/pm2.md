pm2
-

PM2 is a production process manager for + built-in load balancer.
Allows keep applications alive forever, reload without downtime etc.

````
cd ed/nodejs/examples/coursera.one && \
node_modules/.bin/pm2 start server.4.js

# Start 5 instances
node_modules/.bin/pm2 start server.4.js -i 5

node_modules/.bin/pm2 reload all

node_modules/.bin/pm2 list
node_modules/.bin/pm2 monit
node_modules/.bin/pm2 show server.4

node_modules/.bin/pm2 stop all
````
