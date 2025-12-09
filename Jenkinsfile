pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Tests') {
            steps {
                sh 'go version'                     // покажет, что Go уже есть
                sh 'go test -v ./...'               // запустит ВСЕ тесты, включая интеграционный
            }
        }

        stage('Build & Deploy (опционально)') {
            when { branch 'main' }
            steps {
                sh '''
                    go build -o app main.go
                    docker build -t jenkins-lab:${GIT_COMMIT.take(7)} .
                    docker rm -f jenkins-lab || true
                    docker run -d -p 9000:8080 --name jenkins-lab jenkins-lab:${GIT_COMMIT.take(7)}
                '''
            }
        }
    }

    post {
        always {
            echo "Билд завершён. Интеграционный тест должен быть в логе выше ↑"
        }
    }
}