// +build !ignore_autogenerated

/*
Copyright 2018 The Jetstack cert-manager contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMECertificateConfig) DeepCopyInto(out *ACMECertificateConfig) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]DomainSolverConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMECertificateConfig.
func (in *ACMECertificateConfig) DeepCopy() *ACMECertificateConfig {
	if in == nil {
		return nil
	}
	out := new(ACMECertificateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuer) DeepCopyInto(out *ACMEIssuer) {
	*out = *in
	out.PrivateKey = in.PrivateKey
	if in.HTTP01 != nil {
		in, out := &in.HTTP01, &out.HTTP01
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuerHTTP01Config)
			**out = **in
		}
	}
	if in.DNS01 != nil {
		in, out := &in.DNS01, &out.DNS01
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuerDNS01Config)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuer.
func (in *ACMEIssuer) DeepCopy() *ACMEIssuer {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01Config) DeepCopyInto(out *ACMEIssuerDNS01Config) {
	*out = *in
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]ACMEIssuerDNS01Provider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01Config.
func (in *ACMEIssuerDNS01Config) DeepCopy() *ACMEIssuerDNS01Config {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01Provider) DeepCopyInto(out *ACMEIssuerDNS01Provider) {
	*out = *in
	if in.Akamai != nil {
		in, out := &in.Akamai, &out.Akamai
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuerDNS01ProviderAkamai)
			**out = **in
		}
	}
	if in.CloudDNS != nil {
		in, out := &in.CloudDNS, &out.CloudDNS
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuerDNS01ProviderCloudDNS)
			**out = **in
		}
	}
	if in.Cloudflare != nil {
		in, out := &in.Cloudflare, &out.Cloudflare
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuerDNS01ProviderCloudflare)
			**out = **in
		}
	}
	if in.Route53 != nil {
		in, out := &in.Route53, &out.Route53
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuerDNS01ProviderRoute53)
			**out = **in
		}
	}
	if in.AzureDNS != nil {
		in, out := &in.AzureDNS, &out.AzureDNS
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuerDNS01ProviderAzureDNS)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01Provider.
func (in *ACMEIssuerDNS01Provider) DeepCopy() *ACMEIssuerDNS01Provider {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01Provider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderAkamai) DeepCopyInto(out *ACMEIssuerDNS01ProviderAkamai) {
	*out = *in
	out.ClientToken = in.ClientToken
	out.ClientSecret = in.ClientSecret
	out.AccessToken = in.AccessToken
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderAkamai.
func (in *ACMEIssuerDNS01ProviderAkamai) DeepCopy() *ACMEIssuerDNS01ProviderAkamai {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderAkamai)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderAzureDNS) DeepCopyInto(out *ACMEIssuerDNS01ProviderAzureDNS) {
	*out = *in
	out.ClientSecret = in.ClientSecret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderAzureDNS.
func (in *ACMEIssuerDNS01ProviderAzureDNS) DeepCopy() *ACMEIssuerDNS01ProviderAzureDNS {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderAzureDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderCloudDNS) DeepCopyInto(out *ACMEIssuerDNS01ProviderCloudDNS) {
	*out = *in
	out.ServiceAccount = in.ServiceAccount
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderCloudDNS.
func (in *ACMEIssuerDNS01ProviderCloudDNS) DeepCopy() *ACMEIssuerDNS01ProviderCloudDNS {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderCloudDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderCloudflare) DeepCopyInto(out *ACMEIssuerDNS01ProviderCloudflare) {
	*out = *in
	out.APIKey = in.APIKey
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderCloudflare.
func (in *ACMEIssuerDNS01ProviderCloudflare) DeepCopy() *ACMEIssuerDNS01ProviderCloudflare {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderCloudflare)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderRoute53) DeepCopyInto(out *ACMEIssuerDNS01ProviderRoute53) {
	*out = *in
	out.SecretAccessKey = in.SecretAccessKey
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderRoute53.
func (in *ACMEIssuerDNS01ProviderRoute53) DeepCopy() *ACMEIssuerDNS01ProviderRoute53 {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderRoute53)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerHTTP01Config) DeepCopyInto(out *ACMEIssuerHTTP01Config) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerHTTP01Config.
func (in *ACMEIssuerHTTP01Config) DeepCopy() *ACMEIssuerHTTP01Config {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerHTTP01Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerStatus) DeepCopyInto(out *ACMEIssuerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerStatus.
func (in *ACMEIssuerStatus) DeepCopy() *ACMEIssuerStatus {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEOrderChallenge) DeepCopyInto(out *ACMEOrderChallenge) {
	*out = *in
	in.SolverConfig.DeepCopyInto(&out.SolverConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEOrderChallenge.
func (in *ACMEOrderChallenge) DeepCopy() *ACMEOrderChallenge {
	if in == nil {
		return nil
	}
	out := new(ACMEOrderChallenge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEOrderStatus) DeepCopyInto(out *ACMEOrderStatus) {
	*out = *in
	if in.Challenges != nil {
		in, out := &in.Challenges, &out.Challenges
		*out = make([]ACMEOrderChallenge, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEOrderStatus.
func (in *ACMEOrderStatus) DeepCopy() *ACMEOrderStatus {
	if in == nil {
		return nil
	}
	out := new(ACMEOrderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAIssuer) DeepCopyInto(out *CAIssuer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAIssuer.
func (in *CAIssuer) DeepCopy() *CAIssuer {
	if in == nil {
		return nil
	}
	out := new(CAIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Certificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateACMEStatus) DeepCopyInto(out *CertificateACMEStatus) {
	*out = *in
	in.Order.DeepCopyInto(&out.Order)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateACMEStatus.
func (in *CertificateACMEStatus) DeepCopy() *CertificateACMEStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateACMEStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateCondition) DeepCopyInto(out *CertificateCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateCondition.
func (in *CertificateCondition) DeepCopy() *CertificateCondition {
	if in == nil {
		return nil
	}
	out := new(CertificateCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateList) DeepCopyInto(out *CertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Certificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateList.
func (in *CertificateList) DeepCopy() *CertificateList {
	if in == nil {
		return nil
	}
	out := new(CertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpec) DeepCopyInto(out *CertificateSpec) {
	*out = *in
	if in.DNSNames != nil {
		in, out := &in.DNSNames, &out.DNSNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.IssuerRef = in.IssuerRef
	if in.ACME != nil {
		in, out := &in.ACME, &out.ACME
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMECertificateConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpec.
func (in *CertificateSpec) DeepCopy() *CertificateSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateStatus) DeepCopyInto(out *CertificateStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CertificateCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ACME != nil {
		in, out := &in.ACME, &out.ACME
		if *in == nil {
			*out = nil
		} else {
			*out = new(CertificateACMEStatus)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateStatus.
func (in *CertificateStatus) DeepCopy() *CertificateStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIssuer) DeepCopyInto(out *ClusterIssuer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIssuer.
func (in *ClusterIssuer) DeepCopy() *ClusterIssuer {
	if in == nil {
		return nil
	}
	out := new(ClusterIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterIssuer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIssuerList) DeepCopyInto(out *ClusterIssuerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterIssuer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIssuerList.
func (in *ClusterIssuerList) DeepCopy() *ClusterIssuerList {
	if in == nil {
		return nil
	}
	out := new(ClusterIssuerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterIssuerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNS01SolverConfig) DeepCopyInto(out *DNS01SolverConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNS01SolverConfig.
func (in *DNS01SolverConfig) DeepCopy() *DNS01SolverConfig {
	if in == nil {
		return nil
	}
	out := new(DNS01SolverConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSolverConfig) DeepCopyInto(out *DomainSolverConfig) {
	*out = *in
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.SolverConfig.DeepCopyInto(&out.SolverConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSolverConfig.
func (in *DomainSolverConfig) DeepCopy() *DomainSolverConfig {
	if in == nil {
		return nil
	}
	out := new(DomainSolverConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTP01SolverConfig) DeepCopyInto(out *HTTP01SolverConfig) {
	*out = *in
	if in.IngressClass != nil {
		in, out := &in.IngressClass, &out.IngressClass
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTP01SolverConfig.
func (in *HTTP01SolverConfig) DeepCopy() *HTTP01SolverConfig {
	if in == nil {
		return nil
	}
	out := new(HTTP01SolverConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Issuer) DeepCopyInto(out *Issuer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Issuer.
func (in *Issuer) DeepCopy() *Issuer {
	if in == nil {
		return nil
	}
	out := new(Issuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Issuer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerCondition) DeepCopyInto(out *IssuerCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerCondition.
func (in *IssuerCondition) DeepCopy() *IssuerCondition {
	if in == nil {
		return nil
	}
	out := new(IssuerCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerConfig) DeepCopyInto(out *IssuerConfig) {
	*out = *in
	if in.ACME != nil {
		in, out := &in.ACME, &out.ACME
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuer)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.CA != nil {
		in, out := &in.CA, &out.CA
		if *in == nil {
			*out = nil
		} else {
			*out = new(CAIssuer)
			**out = **in
		}
	}
	if in.Vault != nil {
		in, out := &in.Vault, &out.Vault
		if *in == nil {
			*out = nil
		} else {
			*out = new(VaultIssuer)
			**out = **in
		}
	}
	if in.SelfSigned != nil {
		in, out := &in.SelfSigned, &out.SelfSigned
		if *in == nil {
			*out = nil
		} else {
			*out = new(SelfSignedIssuer)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerConfig.
func (in *IssuerConfig) DeepCopy() *IssuerConfig {
	if in == nil {
		return nil
	}
	out := new(IssuerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerList) DeepCopyInto(out *IssuerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Issuer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerList.
func (in *IssuerList) DeepCopy() *IssuerList {
	if in == nil {
		return nil
	}
	out := new(IssuerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IssuerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerSpec) DeepCopyInto(out *IssuerSpec) {
	*out = *in
	in.IssuerConfig.DeepCopyInto(&out.IssuerConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerSpec.
func (in *IssuerSpec) DeepCopy() *IssuerSpec {
	if in == nil {
		return nil
	}
	out := new(IssuerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerStatus) DeepCopyInto(out *IssuerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]IssuerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ACME != nil {
		in, out := &in.ACME, &out.ACME
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuerStatus)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerStatus.
func (in *IssuerStatus) DeepCopy() *IssuerStatus {
	if in == nil {
		return nil
	}
	out := new(IssuerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeySelector) DeepCopyInto(out *SecretKeySelector) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeySelector.
func (in *SecretKeySelector) DeepCopy() *SecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(SecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSignedIssuer) DeepCopyInto(out *SelfSignedIssuer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSignedIssuer.
func (in *SelfSignedIssuer) DeepCopy() *SelfSignedIssuer {
	if in == nil {
		return nil
	}
	out := new(SelfSignedIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolverConfig) DeepCopyInto(out *SolverConfig) {
	*out = *in
	if in.HTTP01 != nil {
		in, out := &in.HTTP01, &out.HTTP01
		if *in == nil {
			*out = nil
		} else {
			*out = new(HTTP01SolverConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.DNS01 != nil {
		in, out := &in.DNS01, &out.DNS01
		if *in == nil {
			*out = nil
		} else {
			*out = new(DNS01SolverConfig)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolverConfig.
func (in *SolverConfig) DeepCopy() *SolverConfig {
	if in == nil {
		return nil
	}
	out := new(SolverConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAppRole) DeepCopyInto(out *VaultAppRole) {
	*out = *in
	out.SecretRef = in.SecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAppRole.
func (in *VaultAppRole) DeepCopy() *VaultAppRole {
	if in == nil {
		return nil
	}
	out := new(VaultAppRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuth) DeepCopyInto(out *VaultAuth) {
	*out = *in
	out.TokenSecretRef = in.TokenSecretRef
	out.AppRole = in.AppRole
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuth.
func (in *VaultAuth) DeepCopy() *VaultAuth {
	if in == nil {
		return nil
	}
	out := new(VaultAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultIssuer) DeepCopyInto(out *VaultIssuer) {
	*out = *in
	out.Auth = in.Auth
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultIssuer.
func (in *VaultIssuer) DeepCopy() *VaultIssuer {
	if in == nil {
		return nil
	}
	out := new(VaultIssuer)
	in.DeepCopyInto(out)
	return out
}
