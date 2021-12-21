pipeline {
    agent any

    stages {
        stage('Build') {
            node("testing"){
                def root = tool type: 'go', name: 'go 1.17.5'
                // Export environment variables pointing to the directory where Go was installed
                withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                        sh 'go version'
                        sh 'go build -o main'
                }
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
