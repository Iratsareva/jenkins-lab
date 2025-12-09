pipeline {
    agent any

    stages {
        stage('Install Go') {
            steps {
                sh '''
                    if ! command -v go >/dev/null 2>&1; then
                        wget -q https://go.dev/dl/go1.22.8.linux-amd64.tar.gz
                        tar -C /tmp -xzf go1.22.8.linux-amd64.tar.gz
                        mv /tmp/go /usr/local/go
                    fi
                    /usr/local/go/bin/go version
                '''
            }
        }

        stage('Test') {
            environment {
                PATH = "/usr/local/go/bin:${env.PATH}"
            }
            steps {
                sh 'go test -v ./...'

                // Или отдельно:
                // sh 'go test -v -run TestHandler'
                // sh 'go test -v -run TestIntegration'
            }
        }

        stage('Build & Deploy (опционально)') {
            when { branch 'main' }
            environment {
                PATH = "/usr/local/go/bin:${env.PATH}"
            }
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
}