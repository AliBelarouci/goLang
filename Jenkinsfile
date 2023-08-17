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
        stage('Checkout') {
            steps {
                // Checkout your source code from the repository
                // Checkout your source code from the repository
                checkout([$class: 'GitSCM', 
                    branches: [[name: '*/main']], // or '*/master' depending on your default branch
                    userRemoteConfigs: [[url: 'https://github.com/AliBelarouci/goLang.git']]])
            }
        }
        
        stage('Build') {
            steps {
                // Build your Go project
                sh 'go build'
            }
        }
        
        stage('Test') {
            steps {
                // Run tests for your Go project
                sh 'go test ./...'
            }
        }

    }
}
