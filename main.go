package main

import (
	"fmt"
	"jws-signature-util/cmd"
	"jws-signature-util/utils"
	bootstrap "k8s.io/cluster-bootstrap/token/jws"
	"os"
)

func main() {
	command := cmd.NewJwsCommand()
	err := command.Execute()
	if err != nil {
		os.Exit(1)
	}

	//check()
}

func check() {
	clusterInfo := utils.GetClusterInfo()
	tokens := utils.GetTokenSecret()
	fmt.Printf("verify content: \n%v\n", clusterInfo.Kubeconfig)
	for tokenId, tokenSecret := range tokens {
		fmt.Printf("verify tokenId %s and token Secret %s \n", tokenId, tokenSecret)
		jwsToken, err := bootstrap.ComputeDetachedSignature(clusterInfo.Kubeconfig, tokenId, tokenSecret)
		if err != nil {
			panic(err.Error())
		}
		oldJwsToken := clusterInfo.Tokens[tokenId]
		fmt.Printf("jwsToken:%s\noldJwsToken:%s\n", jwsToken, oldJwsToken)
		if jwsToken != oldJwsToken {
			fmt.Print("fail\n")
		} else {
			fmt.Println("success\n")
		}
		// another way
		if !bootstrap.DetachedTokenIsValid(oldJwsToken, clusterInfo.Kubeconfig, tokenId, tokenSecret) {
			fmt.Print("fail2\n")
		} else {
			fmt.Println("success2\n")
		}
	}
}
