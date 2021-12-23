
node{
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


    stage('Deploy'){
        container('argo'){
            checkout([$class: 'GitSCM',
                    branches: [[name: '*/main' ]],
                    extensions: scm.extensions,
                    userRemoteConfigs: [[
                        url: 'git@github.com:cure4itches/docker-hello-world-deployment.git',
                        credentialsId: 'jenkins-ssh-private',
                    ]]
            ])
            sshagent(credentials: ['jenkins-ssh-private']){
                sh("""
                    #!/usr/bin/env bash
                    set +x
                    export GIT_SSH_COMMAND="ssh -oStrictHostKeyChecking=no"
                    git config --global user.email "cure4itches@gmail.com"
                    git checkout main
                    cd env/dev && kustomize edit set image arm7tdmi/node-hello-world:${BUILD_NUMBER}
                    git commit -a -m "updated the image tag"
                    git push
                """)
            }
        }
    }
}

