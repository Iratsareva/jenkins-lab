pipeline {
    agent any

    environment {
        GO_INSTALL_DIR = "/var/jenkins_home/go"
        PATH = "$GO_INSTALL_DIR/bin:${env.PATH}"
    }

    stages {
        stage('Установка Go') {
            steps {
                sh '''
                    if ! command -v go >/dev/null 2>&1; then
                        mkdir -p "$GO_INSTALL_DIR"
                        mkdir -p /tmp/go-install
                        cd /tmp/go-install
                        curl -L https://go.dev/dl/go1.22.8.linux-amd64.tar.gz -o go.tar.gz
                        tar -xzf go.tar.gz
                        cp -r go/* "$GO_INSTALL_DIR/"
                    fi
                    go version
                '''
            }
        }

        stage('Получение кода') {
            steps {
                checkout scm
            }
        }

        stage('Модульные тесты') {
            steps {
                sh '''
                    echo "Запуск модульных тестов..."
                    go test -v -run="TestHandler" .
                    go test -v -run="TestHandler2" .
                '''
            }
        }

        stage('Интеграционное тестирование') {
            steps {
                sh '''
                    echo "Запуск интеграционного тестирования..."
                    go test -v -run="TestIntegration" .
                    echo "Интеграционное тестирование завершено"
                    echo "Результат: оба сервиса работают корректно"
                '''
            }
        }
    }

    post {
        success {
            echo "Сборка завершена успешно"
            echo "Интеграционное тестирование пройдено"
        }
        failure {
            echo "Сборка завершилась с ошибкой"
        }
    }
}