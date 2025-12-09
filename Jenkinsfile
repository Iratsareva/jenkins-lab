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
                '''
            }
        }

        stage('Создание Dockerfile') {
            steps {
                sh '''
                    echo "Создание Dockerfile..."
                    
                    cat > Dockerfile << 'EOF'
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go test -v -run="TestIntegration" .
RUN go build -o service1 main.go
RUN go build -o service2 main_2.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/service1 .
COPY --from=builder /app/service2 .
EXPOSE 8080 8081
CMD ["sh", "-c", "./service1 & ./service2 & wait"]
EOF
                    
                    echo "Dockerfile создан"
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