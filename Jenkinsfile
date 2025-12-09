pipeline {
    agent any                                    // запускаем прямо в самом Jenkins

    stages {
        stage('Install Go') {
            steps {
                // Устанавливаем Go один раз (если ещё нет)
                sh '''
                    if ! command -v go > /dev/null; then
                        rm -rf /tmp/go && rm -rf /usr/local/go
                        wget -q https://go.dev/dl/go1.22.8.linux-amd64.tar.gz -O /tmp/go.tar.gz
                        tar -C /usr/local -xzf /tmp/go.tar.gz
                    fi
                    /usr/local/go/bin/go version
                '''
            }
        }

        stage('Tests') {
            environment {
                PATH = "/usr/local/go/bin:${env.PATH}"
            }
            steps {
                sh 'go version'
                sh 'go test -v ./...'
            }
        }
    }

    post {
        success {
            echo "ЛАБОРАТОРНАЯ ВЫПОЛНЕНА УСПЕШНО!"
            echo "Интеграционный тест прошёл — оба сервиса отвечают"
        }
        failure {
            echo "Где-то ошибка, но ты уже очень близко"
        }
    }
}