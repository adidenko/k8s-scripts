package k8s

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	"github.com/impossiblecloud/pd-cert-assistant/internal/cfg"
	"github.com/impossiblecloud/pd-cert-assistant/internal/utils"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	cmapi "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	cmmeta "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	cmclient "github.com/cert-manager/cert-manager/pkg/client/clientset/versioned"
)

type Client struct {
	Config *rest.Config
}

func loadKubeConfig(path string) (*rest.Config, error) {
	file, err := clientcmd.LoadFromFile(path)
	if err != nil {
		return nil, err
	}
	// Just in case we decide to do some overrides in the future
	override := &clientcmd.ConfigOverrides{}
	return clientcmd.NewDefaultClientConfig(*file, override).ClientConfig()
}

func (c *Client) Init(kubeConfigPath string) error {
	var err error

	// Load either the kubeconfig file or in-cluster config
	if kubeConfigPath != "" {
		c.Config, err = loadKubeConfig(kubeConfigPath)
	} else {
		c.Config, err = rest.InClusterConfig()
	}
	return err
}

// GetCiliumNodes retrieves a list of CiliumNode resources from the cilium.io/v2 API
// and returns a list of CiliumInternalIP addresses.
func (c *Client) GetCiliumNodes() ([]string, error) {

	ciliumNodeGVR := schema.GroupVersionResource{
		Group:    "cilium.io",
		Version:  "v2",
		Resource: "ciliumnodes",
	}

	dynamicClient, err := dynamic.NewForConfig(c.Config)
	if err != nil {
		return nil, fmt.Errorf("failed to create dynamic client: %v", err)
	}

	ciliumNodes, err := dynamicClient.Resource(ciliumNodeGVR).Namespace("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list CiliumNode resources: %v", err)
	}

	// Extract CiliumInternalIP addresses
	var internalIPs []string
	for _, item := range ciliumNodes.Items {
		glog.V(8).Infof("Processing CiliumNode: %s", item.GetName())
		spec, ok := item.Object["spec"].(map[string]interface{})
		if !ok {
			continue
		}

		addresses, ok := spec["addresses"].([]interface{})
		if !ok {
			continue
		}

		for _, addr := range addresses {
			addressMap, ok := addr.(map[string]interface{})
			if !ok {
				continue
			}

			if addressMap["type"] == "CiliumInternalIP" {
				if ip, ok := addressMap["ip"].(string); ok {
					internalIPs = append(internalIPs, ip)
				}
			}
		}
	}

	return internalIPs, nil
}
