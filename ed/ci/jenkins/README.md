Jenkins
-

````sh
sudo su -s /bin/bash jenkins
ssh-keygen -t rsa
````

````sh
git ls-remote -h https://github.com/org/repo HEAD
````

Jenkins file syntax - groovy.

````
stage('Build Docker Image') {
   when {branch 'master'}
   steps {
       sh "docker build . -t $PROJECT_NAME:${GIT_COMMIT}-build"
   }
}
````
