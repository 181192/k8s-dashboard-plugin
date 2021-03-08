package main

import (
	"github.com/181192/k8s-dashboard-plugin/cmd"
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
)

func main() {
	cmd.Execute()
}
