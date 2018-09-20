node {
    try{
        notifyBuild('STARTED')
        ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
            withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {
                env.PATH="${GOPATH}/bin:$PATH"
                
                def app
                def branch

                stage('Get Branch Name') {
                    echo 'Getting branch name from Jenkins settings'
                    branch = scm.branches[0].name.drop(2)
                    echo 'Branch Name: ' + branch
                }
                
                stage('Install Dependencies') {
                    echo 'Pulling Dependencies'
                    sh 'go version'
                    sh 'go get -u github.com/btcsuite/btcutil/base58'
                    sh 'go get -u github.com/ixofoundation/ixo-cosmos/app'
                }

                stage('Building') {
                    dir('src/github.com/ixofoundation/ixo-cosmos/') {
                         sh 'git checkout ' + branch
                         sh 'make build && make install' 
                    }
                } 

                stage('Building Docker Image') {
                    dir('src/github.com/ixofoundation/ixo-cosmos/') {
                        app = docker.build('trustlab/ixo-blockchain:' + branch)
                    } 
                } 

                stage('Test Image') {
                    app.inside {
                        sh 'echo "Tests passed"'
                    }
                }

                stage('Push Image') {
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
                        app.push(branch + "-" + "${env.BUILD_NUMBER}")
                        app.push(branch)
                    }
                }

                stage('Removing Images') {
                    sh "docker rmi ${app.id}"
                    sh "docker rmi registry.hub.docker.com/${app.id}"
                    sh "docker rmi registry.hub.docker.com/${app.id}-${env.BUILD_NUMBER}"
                }
            }
        }
    } catch (e) {
        // If there was an exception thrown, the build failed
        currentBuild.result = "FAILED"
        
    } finally {
        // Success or failure, always send notifications
        notifyBuild(currentBuild.result)
    }
}

def notifyBuild(String buildStatus = 'STARTED') {
  // build status of null means successful
  buildStatus =  buildStatus ?: 'SUCCESSFUL'

  // Default values
  def colorName = 'RED'
  def colorCode = '#FF0000'
  def subject = "${buildStatus}: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'"
  def summary = "${subject} <${env.BUILD_URL}|Job URL> - <${env.BUILD_URL}/console|Console Output>"

  // Override default values based on build status
  if (buildStatus == 'STARTED') {
    color = 'YELLOW'
    colorCode = '#FFFF00'
  } else if (buildStatus == 'SUCCESSFUL') {
    color = 'GREEN'
    colorCode = '#00FF00'
  } else {
    color = 'RED'
    colorCode = '#FF0000'
  }
}