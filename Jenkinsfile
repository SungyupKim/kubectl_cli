pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh '''
                echo "Start Build"
                def root = tool type: 'go', name: 'Go 1.15'
                // Export environment variables pointing to the directory where Go was installed
                withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                        sh 'go version'
                }
                go build -o main
                '''
            }
        }
        stage('Test') {
            steps {
                sh '''
                echo "Start Test"
                ./main
                '''
            }
        }
        stage('Archive') {
            steps {
                sh '''
                echo "Start Archiving"
                mv ./main ./artifact/
                '''
                archiveArtifacts artifacts: 'artifact/*', fingerprint: true 
            }
        }
    }
}
