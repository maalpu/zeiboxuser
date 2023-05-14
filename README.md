# zeiboxuser
1. Incicializar proyecto
   ```
   $ git mod init
   ```
2. Creamos main.go
  - Instalamos github.com/aws/aws-lambda-go/lambda
    ```
    $ go get github.com/aws/aws-lambda-go/lambda
    ```
  - Escribimos main.go
    ```go
     package main

    import "github.com/aws/aws-lambda-go/lambda"

    func main() {
        lambda.Start(EjecutoLambda)
    }

    func EjecutoLambda() {

    }
    ```
  - Importar github.com/aws/aws-lambda-go/events
    ```
    $ go get github.com/aws/aws-lambda-go/evemts
    ```
3. Creamos en la carpeta /awsgo/awsgo.go
    ```go
    package awsgo

    import (
        "context"

        "github.com/aws/aws-sdk-go-v2/aws"
        "github.com/aws/aws-sdk-go-v2/config"
    )

    var Ctx context.Context
    var Cfg aws.Config
    var err error

    func InicializoAWS() {
        Ctx = context.TODO()
        Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
        if err != nil {
            panic("Error al cargar la fonfiguración .aws/config " + err.Error())
        }
    }
    ```
4. Instalamos las siguientes dependencias
    ```
    $ go get github.com/aws/aws-sdk-go-v2/aws
    $ go get github.com/aws/aws-sdk-go-v2/config
    ```
