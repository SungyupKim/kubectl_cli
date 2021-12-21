pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh '''
                echo "Start Build"
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
