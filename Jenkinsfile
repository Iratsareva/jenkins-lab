pipeline {
    agent any

    environment {
        // Устанавливаем переменные для Go
        GO_INSTALL_DIR = "/usr/local/go"
        PATH = "$GO_INSTALL_DIR/bin:${env.PATH}"
    }

    stages {
        stage('Setup Go') {
            steps {
                sh '''
                    # Устанавливаем Go если нет
                    if ! command -v go >/dev/null 2>&1; then
                        echo "Установка Go 1.22.8..."
                        
                        # Создаем временную директорию с правами
                        mkdir -p /tmp/go-install
                        cd /tmp/go-install
                        
                        # Скачиваем Go
                        wget -q https://go.dev/dl/go1.22.8.linux-amd64.tar.gz || \
                        curl -L https://go.dev/dl/go1.22.8.linux-amd64.tar.gz -o go.tar.gz
                        
                        # Распаковываем ВО временную директорию
                        tar -xzf go*.tar.gz
                        
                        # Копируем только нужное
                        mkdir -p /usr/local
                        cp -r go /usr/local/
                        
                        # Очистка
                        rm -rf /tmp/go-install
                    fi
                    
                    # Проверяем
                    go version
                '''
            }
        }

        stage('Checkout Code') {
            steps {
                checkout scm
            }
        }

        stage('Unit Tests') {
            steps {
                sh 'go test -v -run="TestHandler|TestHandler2" ./...'
            }
        }

        stage('Integration Test') {
            steps {
                sh 'go test -v -run="TestIntegration" ./...'
            }
        }

        stage('Build') {
            when { branch 'main' }
            steps {
                sh '''
                    go build -o app main.go
                    ls -la
                '''
            }
        }
    }

    post {
        success {
            echo "Все тесты пройдены успешно!"
        }
        failure {
            echo "Сборка завершилась с ошибкой"
        }
    }
}