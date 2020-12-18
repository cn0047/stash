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
environment {
    PROJECT_NAME="test"
    BRANCH="${env.BRANCH_NAME}"
}
stages {
    stage('Build Docker Image') {
        when {branch 'master'}
        when {
            expression {
                return (env.BRANCH_NAME == "master" || env.BRANCH_NAME == "dev")
            }
        }
        steps {
            sh "docker build . -t $PROJECT_NAME:${GIT_COMMIT}-build"
        }
    }
}

pipeline {
    agent any
    triggers {
        cron('H */4 * * 1-5')
    }
    stages {
        stage('Example') {
            steps {
                echo 'Hello World'
            }
        }
    }
}
````
