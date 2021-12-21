
node{
    def root = tool type: 'go', name: 'go 1.17.5'
        // Export environment variables pointing to the directory where Go was installed

    stage('Build'){
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
            sh 'go version'
            sh 'go build -o main'
        }
    }
 
    stage('Test') {
        sh '''
        echo "Start Test"
        ./main
        '''
    }
    stage('Archive') {
        sh '''
        echo "Start Archiving"
        mv ./main ./artifact/
        '''
        archiveArtifacts artifacts: 'artifact/*', fingerprint: true 
    }
}

