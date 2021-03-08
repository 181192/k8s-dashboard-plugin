module github.com/181192/k8s-dashboard-plugin

go 1.15

require (
	github.com/181192/ops-cli v0.1.17
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v1.0.0
	k8s.io/client-go v0.18.0
)

replace github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.11.15

replace github.com/Azure/go-autorest/autorest/adal => github.com/Azure/go-autorest/autorest/adal v0.9.10
