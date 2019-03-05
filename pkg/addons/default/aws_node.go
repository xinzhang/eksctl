package defaultaddons

import (
	"github.com/kris-nova/logger"
	"github.com/pkg/errors"

	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	AWSNode = "aws-node"
)

func UpdateAWSNode(clientSet *kubernetes.Clientset, dryRun bool) error {
	_, err := clientSet.Apps().DaemonSets(metav1.NamespaceSystem).Get(AWSNode, metav1.GetOptions{})
	if err != nil {
		if apierrs.IsNotFound(err) {
			logger.Warning("%q was not found", AWSNode)
			return nil
		}
		return errors.Wrapf(err, "getting %s", AWSNode)
	}

	// if DaemonSets is present, go through our list of assets
	list, err := LoadAsset(AWSNode, "yaml")
	if err != nil {
		return err
	}

	for _, obj := range list.Items {
		logger.Debug("%#v", obj.Object.GetObjectKind())
	}
	return nil
}
