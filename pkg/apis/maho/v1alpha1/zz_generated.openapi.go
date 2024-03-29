// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/takara-operator/pkg/apis/maho/v1alpha1.Takara":       schema_pkg_apis_maho_v1alpha1_Takara(ref),
		"github.com/takara-operator/pkg/apis/maho/v1alpha1.TakaraSpec":   schema_pkg_apis_maho_v1alpha1_TakaraSpec(ref),
		"github.com/takara-operator/pkg/apis/maho/v1alpha1.TakaraStatus": schema_pkg_apis_maho_v1alpha1_TakaraStatus(ref),
	}
}

func schema_pkg_apis_maho_v1alpha1_Takara(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Takara is the Schema for the takaras API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/takara-operator/pkg/apis/maho/v1alpha1.TakaraSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/takara-operator/pkg/apis/maho/v1alpha1.TakaraStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/takara-operator/pkg/apis/maho/v1alpha1.TakaraSpec", "github.com/takara-operator/pkg/apis/maho/v1alpha1.TakaraStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_maho_v1alpha1_TakaraSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TakaraSpec defines the desired state of Takara",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_maho_v1alpha1_TakaraStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TakaraStatus defines the observed state of Takara",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
