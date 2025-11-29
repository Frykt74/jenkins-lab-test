pipeline {
    agent any
    
    environment {
        PATH = "/var/jenkins_home/go/bin:${env.PATH}"
        IMAGE_TAG = sh(script: 'git rev-parse --short HEAD', returnStdout: true).trim()
    }
    
    stages {
        // Стадия 1: Unit-тесты
        stage('Unit Tests') {
            steps {
                echo 'Running unit tests...'
                sh 'go test -v -run="^(TestHandler|TestHandler2)$"'
            }
        }
        
        // Стадия 2: Интеграционный тест
        stage('Integration Test') {
            steps {
                echo 'Running integration test...'
                sh 'go test -v -run="^TestIntegration$"'
            }
        }
        
        // Стадия 3: Сборка образа
        stage('Build') {
            steps {
                echo 'Building Docker image...'
                sh 'docker build -t test:${IMAGE_TAG} .'
            }
        }
        
        // Стадия 4: Деплой
        stage('Deploy') {
            steps {
                echo 'Deploying container...'
                sh 'docker rm -f test || true'
                sh 'docker run -p 9000:8080 -d --name test test:${IMAGE_TAG}'
            }
        }
    }
}
