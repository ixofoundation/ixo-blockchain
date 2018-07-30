node {
    def app

    stage('Clone repository') {
        /* Let's make sure we have the repository cloned to our workspace */

        checkout scm
    }

    stage('Build source') {
        /* Let's make sure we have the repository cloned to our workspace */

        sh 'make build'
    }

    stage('Build blockchain image') {
        /* This builds the actual image; synonymous to
         * docker build on the command line */
        blockchain = docker.build("trustlab/ixo-blockchain", "./docker/blockchain/")
    }

    stage('Test image') {
        /* Ideally, we would run a test framework against our image.
         * For this example, we're using a Volkswagen-type approach ;-) */

        blockchain.inside {
            sh 'echo "Tests passed"'
        }
    }

    stage('Push image') {
        /* Finally, we'll push the image with two tags:
         * First, the incremental build number from Jenkins
         * Second, the 'latest' tag.
         * Pushing multiple tags is cheap, as all the layers are reused. */
        docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
            blockchain.push("${env.BUILD_NUMBER}")
            blockchain.push("latest")
        }
    }

}