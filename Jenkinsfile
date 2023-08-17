pipeline {
    agent any
    tools {
        go '1.18.1'
    }
    stages{
        stage("Go example"){
            steps{
                sh "go version"
            }
        }

    }
}
