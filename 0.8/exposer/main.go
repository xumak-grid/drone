package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/xumak-grid/go-grid/pkg/client"
	discovery "github.com/xumak-grid/go-grid/pkg/service-discovery"
)

const (
	// NAMESPACE ENV VARIABLE
	NAMESPACE = "K8_NAMESPACE"
)

func exposeBalancer(s discovery.Selector, envVar string) {
	ns := os.Getenv(NAMESPACE)
	if ns == "default" {
		ns = ""
	}

	var services []discovery.Service
	var err error
	for {
		client := client.MustNewKubeClient()
		services, err = discovery.Locate(client, s, ns, false)
		if err != nil {
			fmt.Printf("service/balancer not found for %v and  %s \n", s, envVar)
			fmt.Println(err.Error(), "next try in 5 seconds")
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}

	if len(services) == 0 || services[0].Balancer == "" {
		fmt.Printf("service/balancer not found for %v \n", s)
		os.Exit(1)
	}
	service := services[0]
	fmt.Printf("export %s=%s\n", envVar, service.Balancer)
}
func main() {
	viper.SetConfigFile("/meta/labels.properties")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	appName := strings.Trim(viper.GetString("app"), "\"")
	droneSelector := discovery.Selector{
		"app": appName,
	}
	exposeBalancer(droneSelector, "DRONE_HOST")
	gogsSelector := discovery.Selector{
		"app": "gogs-server",
	}
	exposeBalancer(gogsSelector, "DRONE_GOGS_URL")
}
