
node{
    def docker = "docker"
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
}

