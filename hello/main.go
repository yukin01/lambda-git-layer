package main

import (
	"context"
	"log"
	"os/exec"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context) error {
	log.Println("[info] git version check")

	log.Println("[info] ls -la /opt/bin")
	output, err := exec.Command("ls", "-la", "/opt/bin").CombinedOutput()
	log.Println(string(output))
	if err != nil {
		return err
	}

	log.Println("[info] exec LookPath git")
	path, err := exec.LookPath("git")
	log.Println(path)
	if err != nil {
		return err
	}

	log.Println("[info] git --version")
	output, err = exec.Command("git", "--version").CombinedOutput()
	log.Println(string(output))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
