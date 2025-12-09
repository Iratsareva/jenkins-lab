pipeline {
    agent any

    // Эта часть — почти как в лабораторной, но с fallback-установкой Go
    environment {
        // Сначала пробуем путь из лабораторной
        PATH = "/var/jenkins_home/go/bin:${env.PATH}"
    }

    stages {
        stage('Prepare Go') {
            steps {
                sh '''
                    # Если Go нет по пути из лабы — ставим сами
                    if ! command -v go >/dev/null 2>&1; then
                        echo "Go не найден в /var/jenkins_home/go — устанавливаем в /usr/local/go"
                        rm -rf /usr/local/go
                        wget -q https://go.dev/dl/go1.22.8.linux-amd64.tar.gz -O /tmp/go.tar.gz || \
                        curl -k -L https://go.dev/dl/go1.22.8.linux-amd64.tar.gz -o /tmp/go.tar.gz
                        tar -C /usr/local -xzf /tmp/go.tar.gz
                        export PATH="/usr/local/go/bin:$PATH"
                    fi
                    go version
                '''
            }
        }

        stage('Unit Tests') {
            steps {
                sh 'go test -v -run="TestHandler|TestHandler2"'
            }
        }

        stage('Integration Test') {
            steps {
                sh 'go test -v -run="TestIntegration"'
            }
        }

        // Эти стадии — точно как в лабораторной, просто для красоты
        stage('Build (опционально)') {
            when { branch 'main' }
            steps {
                sh 'go build -o app main.go'
            }
        }

        stage('Docker Build & Deploy (опционально)') {
            when { branch 'main' }
            steps {
                sh '''
                    docker build -t test:${GIT_COMMIT.take(7)} .
                    docker rm -f test || true
                    docker run -d -p 9000:8080 --name test test:${GIT_COMMIT.take(7)} || true
                '''
            }
        }
    }

    post {
        success {
            echo "Интеграционное тестирование успешно пройдено!"
            echo "Оба сервиса отвечают корректно"
        }
    }
}