
node{
    stage("Fix the permission issue") {
        sh "sudo chown root:jenkins /run/docker.sock"
    }

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
            docker.withRegistry('https://registry.hub.docker.com/', '44cd9687-e273-488b-986e-2d608da5fe27') {
                app.push("latest")
            }
        }
    }
}

