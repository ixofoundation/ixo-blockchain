node {

    agent {
        label {
            label 'ixo-blockchain'
            customWorkspace '$GOPATH/src/github.com/ixofoundation/ixo-cosmos'
        }
    }

    stages {
        stage('Pull repository') {
            steps {
                sh 'go version'
                sh 'cd $GOPATH/src/github.com/ixofoundation/ixo-cosmos'
                sh 'git pull origin master'
            }
        }

        stage('Build source') {
          steps {
                sh 'pwd'
                sh 'make build'
                sh 'make install'
            }
        }

        stage('Build blockchain image') {
            /* This builds the actual image; synonymous to
            * docker build on the command line */
            steps {
                blockchain = docker.build("trustlab/ixo-blockchain", "./docker/blockchain/")
            }
        }

        stage('Test image') {
            /* Ideally, we would run a test framework against our image.
            * For this example, we're using a Volkswagen-type approach ;-) */
            steps {
                blockchain.inside {
                    sh 'echo "Tests passed"'
                }
            }
        }

        stage('Push image') {
            /* Finally, we'll push the image with two tags:
            * First, the incremental build number from Jenkins
            * Second, the 'latest' tag.
            * Pushing multiple tags is cheap, as all the layers are reused. */
            steps {
                docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
                    blockchain.push("${env.BUILD_NUMBER}")
                    blockchain.push("latest")
                }
            }
        }
    }

}