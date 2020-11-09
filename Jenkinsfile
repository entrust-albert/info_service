pipeline {
    agent any
    tools {
        go 'golang'
        dockerTool 'docker'
    }
    stages {
        stage ('Installing dependencies'){
            steps {
                echo 'Installing dependencies'
                sh 'go get -u  github.com/go-sql-driver/mysql'                
            }
            
        }

        stage ('Git'){
            steps {
                echo 'Getting Git'
                git url: 'https://github.com/entrust-albert/info_service'
            }
        }
        
        stage ('Building'){
            steps {
                echo 'Compiling and building'
                sh 'go build -o getter main.go'
            }
            
        }

        stage('Dockerize') {
            steps{
                sh 'docker build -t get:v1.0 .'
                sh 'docker tag get:v1.0 192.168.176.144:5000/get1:latest'
                sh 'docker push 192.168.176.144:5000/get1:latest'
            }
            
        }
        
    }
    
}
