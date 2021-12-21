
node{
    def root = tool type: 'go', name: 'go 1.17.5'
        // Export environment variables pointing to the directory where Go was installed

    stage('Build'){
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
            sh 'go version'
            sh 'sudo docker build -t sungyupv/kubectl_cli:latest .'
        }
    }
 
    stage('Test') {
        sh '''
        echo "Start Test"
        '''
    }
    stage('Archive') {
        sh 'sudo docker push sungyupv/kubectl_cli:latest'
    }
}

