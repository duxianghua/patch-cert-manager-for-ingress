/*
Copyright 2021 jack.du.
*/

package v1beta1

import (
	"context"
	"encoding/json"
	"net/http"

	v1beta1 "k8s.io/api/extensions/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var ingresslog = logf.Log.WithName("ingress-resource")

// +kubebuilder:webhook:verbs=create;update,path=/networking.k8s.io/v1/ingress,mutating=true,failurePolicy=fail,groups=extensions,resources=ingresses,sideEffects=none,versions=v1beta1,name=test.syncbug.io,admissionReviewVersions=v1

type IngressMutate struct {
	Client  client.Client
	decoder *admission.Decoder
}

func (i *IngressMutate) Handle(ctx context.Context, req admission.Request) admission.Response {
	ingress := &v1beta1.Ingress{}

	err := i.decoder.Decode(req, ingress)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}
	ingresslog.Info("default", "name", ingress.Name)
	//var annotation map[string]string
	//annotation["asfasdf.io"] = "asfdasfasdfa"
	ingress.Annotations["asdfasd.io"] = "asfdasdfas"

	// mutate the fields in pod

	marshaledIngress, err := json.Marshal(ingress)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}

	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledIngress)
}

func (i *IngressMutate) InjectDecoder(d *admission.Decoder) error {
	i.decoder = d
	return nil
}
