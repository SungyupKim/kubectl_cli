
node{
    def root = tool type: 'go', name: 'go 1.17.5'
        // Export environment variables pointing to the directory where Go was installed

    stage('Build'){
        app = docker.build("sungyupv/kubectl_cli")
    }
 
    stage('Test') {
        sh '''
        echo "Start Test"
        '''
    }
    stage('Archive') {
        stage('Push image') {
            docker.withRegistry('https://hub.docker.com/repository/docker/sungyupv/kubectl_cli', 'docker-hub') {
                app.push("latest")
        }
    }
}

