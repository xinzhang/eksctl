package defaultaddons

import (
	"bytes"
	"io"

	"github.com/pkg/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

func init() {
	apiextensionsv1beta1.AddToScheme(scheme.Scheme)
}

// LoadAsset return embedded manifest as a runtime.Object
func LoadAsset(name, ext string) (*metav1.List, error) {
	data, err := Asset(name + "." + ext)
	if err != nil {
		return nil, errors.Wrapf(err, "decoding embedded manifest for %q", name)
	}

	list := metav1.List{}
	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewBuffer(data), 4096)

	for {
		obj := new(runtime.RawExtension)
		err := decoder.Decode(obj)
		if err != nil {
			if err == io.EOF {
				return &list, nil
			}
			return nil, errors.Wrapf(err, "loading individual resources from manifest for %q", name)
		}
		obj.Object, err = runtime.Decode(scheme.Codecs.UniversalDeserializer(), obj.Raw)
		if err != nil {
			return nil, errors.Wrapf(err, "converting object")
		}
		list.Items = append(list.Items, *obj)
	}
}
