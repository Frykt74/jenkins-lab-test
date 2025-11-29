pipeline {
    agent any
    
    environment {
        PATH = "/var/jenkins_home/go/bin:${env.PATH}"
        IMAGE_TAG = sh(script: 'git rev-parse --short HEAD', returnStdout: true).trim()
    }
    
    stages {
        stage('Test') {
            steps {
                sh 'go test -v'
            }
        }
        
        stage('Build') {
            steps {
                sh 'docker build -t test:${IMAGE_TAG} .'
            }
        }
        
        stage('Deploy') {
            steps {
                sh 'docker rm -f test || true'
                sh 'docker run -p 9000:8080 -d --name test test:${IMAGE_TAG}'
            }
        }
    }
}
