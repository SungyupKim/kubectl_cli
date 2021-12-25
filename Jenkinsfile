podTemplate(label: 'build',
  containers: [
    containerTemplate(
      name: 'argo',
      image: 'argoproj/argo-cd-ci-builder:latest',
      command: 'cat',
      ttyEnabled: true
    ),
  ],
  volumes: [ 
    hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock'), 
  ]
) {
    node('build'){
        stage("Fix the permission issue") {
            sh "sudo chown root:jenkins /run/docker.sock"
        }

        stage('Build'){
            sh "sudo su"
            checkout([$class: 'GitSCM',
                    branches: [[name: '*/main' ]],
                    extensions: [[$class: 'RelativeTargetDirectory', relativeTargetDir: 'source']],
                    userRemoteConfigs: [[
                        url: 'git@192.168.219.116:paas/kubectl-cli.git',
                        credentialsId: 'd5e8a4a0-9ad4-4ccb-a17d-2691121e762c',
                    ]]
            ])
            sh "git checkout main"
            app = docker.build("sungyupv/kubectl_cli")
        }
    
        stage('Test') {
            sh '''
            echo "Start Test"
            '''
        }
        stage('Archive') {
            stage('Push image') {
                docker.withRegistry('https://registry.hub.docker.com/', '8160a729-efa4-4177-97b3-666feac4bb75') {
                    app.push("${env.BUILD_NUMBER}")
                    app.push("latest")
                }
            }
        }


        stage('Deploy'){
            container('argo'){
                checkout([$class: 'GitSCM',
                        branches: [[name: '*/main' ]],
                        extensions: [[$class: 'RelativeTargetDirectory', relativeTargetDir: 'deploy']],,
                        userRemoteConfigs: [[
                            url: 'git@192.168.219.116:paas/kubectl-cli-deployment.git',
                            credentialsId: 'd5e8a4a0-9ad4-4ccb-a17d-2691121e762c',
                        ]]
                ])
                sshagent(credentials: ['d5e8a4a0-9ad4-4ccb-a17d-2691121e762c']){
                    sh("""
                        #!/usr/bin/env bash
                        set +x
                        export GIT_SSH_COMMAND="ssh -oStrictHostKeyChecking=no"
                        git config --global user.email "sungyupv@gmail.com"
                        git checkout main
                        cd env/dev && kustomize edit set image sungyupv/kubectl_cli:${BUILD_NUMBER}
                        git commit -a -m "updated the image tag"
                        git push
                    """)
                }
            }
        }
    }
}

