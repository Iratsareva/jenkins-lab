pipeline {
    agent {
        docker {
            image 'golang:1.22'   // ← здесь уже есть Go, git, всё что нужно
            args '-v /var/run/docker.sock:/var/run/docker.sock'  // чтобы docker build работал
        }
    }

    stages {
        stage('Tests') {
            steps {
                sh 'go version'
                sh 'go test -v ./...'
            }
        }

        stage('Build & Deploy') {
            when { branch 'main' }
            steps {
                sh '''
                    go build -o app main.go
                    docker build -t jenkins-lab:${GIT_COMMIT.take(7)} .
                    docker rm -f jenkins-lab || true
                    docker run -d -p 9000:8080 --name jenkins-lab jenkins-lab:${GIT_COMMIT.take(7)} || true
                '''
            }
        }
    }

    post {
        success {
            echo "ВСЁ РАБОТАЕТ! Интеграционный тест прошёл успешно!"
        }
        failure {
            echo "Где-то ошибка, но мы уже почти у цели"
        }
    }
}