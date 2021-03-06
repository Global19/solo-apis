// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the gloo.solo.io/v1 APIs
type MulticlusterClientset interface {
	// Cluster returns a Clientset for the given cluster
	Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
	client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
	return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

// clienset for the gloo.solo.io/v1 APIs
type Clientset interface {
	// clienset for the gloo.solo.io/v1/v1 APIs
	Settings() SettingsClient
	// clienset for the gloo.solo.io/v1/v1 APIs
	Upstreams() UpstreamClient
	// clienset for the gloo.solo.io/v1/v1 APIs
	UpstreamGroups() UpstreamGroupClient
	// clienset for the gloo.solo.io/v1/v1 APIs
	Proxies() ProxyClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the gloo.solo.io/v1/v1 APIs
func (c *clientSet) Settings() SettingsClient {
	return NewSettingsClient(c.client)
}

// clienset for the gloo.solo.io/v1/v1 APIs
func (c *clientSet) Upstreams() UpstreamClient {
	return NewUpstreamClient(c.client)
}

// clienset for the gloo.solo.io/v1/v1 APIs
func (c *clientSet) UpstreamGroups() UpstreamGroupClient {
	return NewUpstreamGroupClient(c.client)
}

// clienset for the gloo.solo.io/v1/v1 APIs
func (c *clientSet) Proxies() ProxyClient {
	return NewProxyClient(c.client)
}

// Reader knows how to read and list Settingss.
type SettingsReader interface {
	// Get retrieves a Settings for the given object key
	GetSettings(ctx context.Context, key client.ObjectKey) (*Settings, error)

	// List retrieves list of Settingss for a given namespace and list options.
	ListSettings(ctx context.Context, opts ...client.ListOption) (*SettingsList, error)
}

// SettingsTransitionFunction instructs the SettingsWriter how to transition between an existing
// Settings object and a desired on an Upsert
type SettingsTransitionFunction func(existing, desired *Settings) error

// Writer knows how to create, delete, and update Settingss.
type SettingsWriter interface {
	// Create saves the Settings object.
	CreateSettings(ctx context.Context, obj *Settings, opts ...client.CreateOption) error

	// Delete deletes the Settings object.
	DeleteSettings(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Settings object.
	UpdateSettings(ctx context.Context, obj *Settings, opts ...client.UpdateOption) error

	// Patch patches the given Settings object.
	PatchSettings(ctx context.Context, obj *Settings, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Settings objects matching the given options.
	DeleteAllOfSettings(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Settings object.
	UpsertSettings(ctx context.Context, obj *Settings, transitionFuncs ...SettingsTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Settings object.
type SettingsStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Settings object.
	UpdateSettingsStatus(ctx context.Context, obj *Settings, opts ...client.UpdateOption) error

	// Patch patches the given Settings object's subresource.
	PatchSettingsStatus(ctx context.Context, obj *Settings, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Settingss.
type SettingsClient interface {
	SettingsReader
	SettingsWriter
	SettingsStatusWriter
}

type settingsClient struct {
	client client.Client
}

func NewSettingsClient(client client.Client) *settingsClient {
	return &settingsClient{client: client}
}

func (c *settingsClient) GetSettings(ctx context.Context, key client.ObjectKey) (*Settings, error) {
	obj := &Settings{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *settingsClient) ListSettings(ctx context.Context, opts ...client.ListOption) (*SettingsList, error) {
	list := &SettingsList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *settingsClient) CreateSettings(ctx context.Context, obj *Settings, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *settingsClient) DeleteSettings(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Settings{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *settingsClient) UpdateSettings(ctx context.Context, obj *Settings, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *settingsClient) PatchSettings(ctx context.Context, obj *Settings, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *settingsClient) DeleteAllOfSettings(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Settings{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *settingsClient) UpsertSettings(ctx context.Context, obj *Settings, transitionFuncs ...SettingsTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*Settings), desired.(*Settings)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *settingsClient) UpdateSettingsStatus(ctx context.Context, obj *Settings, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *settingsClient) PatchSettingsStatus(ctx context.Context, obj *Settings, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides SettingsClients for multiple clusters.
type MulticlusterSettingsClient interface {
	// Cluster returns a SettingsClient for the given cluster
	Cluster(cluster string) (SettingsClient, error)
}

type multiclusterSettingsClient struct {
	client multicluster.Client
}

func NewMulticlusterSettingsClient(client multicluster.Client) MulticlusterSettingsClient {
	return &multiclusterSettingsClient{client: client}
}

func (m *multiclusterSettingsClient) Cluster(cluster string) (SettingsClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewSettingsClient(client), nil
}

// Reader knows how to read and list Upstreams.
type UpstreamReader interface {
	// Get retrieves a Upstream for the given object key
	GetUpstream(ctx context.Context, key client.ObjectKey) (*Upstream, error)

	// List retrieves list of Upstreams for a given namespace and list options.
	ListUpstream(ctx context.Context, opts ...client.ListOption) (*UpstreamList, error)
}

// UpstreamTransitionFunction instructs the UpstreamWriter how to transition between an existing
// Upstream object and a desired on an Upsert
type UpstreamTransitionFunction func(existing, desired *Upstream) error

// Writer knows how to create, delete, and update Upstreams.
type UpstreamWriter interface {
	// Create saves the Upstream object.
	CreateUpstream(ctx context.Context, obj *Upstream, opts ...client.CreateOption) error

	// Delete deletes the Upstream object.
	DeleteUpstream(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Upstream object.
	UpdateUpstream(ctx context.Context, obj *Upstream, opts ...client.UpdateOption) error

	// Patch patches the given Upstream object.
	PatchUpstream(ctx context.Context, obj *Upstream, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Upstream objects matching the given options.
	DeleteAllOfUpstream(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Upstream object.
	UpsertUpstream(ctx context.Context, obj *Upstream, transitionFuncs ...UpstreamTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Upstream object.
type UpstreamStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Upstream object.
	UpdateUpstreamStatus(ctx context.Context, obj *Upstream, opts ...client.UpdateOption) error

	// Patch patches the given Upstream object's subresource.
	PatchUpstreamStatus(ctx context.Context, obj *Upstream, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Upstreams.
type UpstreamClient interface {
	UpstreamReader
	UpstreamWriter
	UpstreamStatusWriter
}

type upstreamClient struct {
	client client.Client
}

func NewUpstreamClient(client client.Client) *upstreamClient {
	return &upstreamClient{client: client}
}

func (c *upstreamClient) GetUpstream(ctx context.Context, key client.ObjectKey) (*Upstream, error) {
	obj := &Upstream{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *upstreamClient) ListUpstream(ctx context.Context, opts ...client.ListOption) (*UpstreamList, error) {
	list := &UpstreamList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *upstreamClient) CreateUpstream(ctx context.Context, obj *Upstream, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *upstreamClient) DeleteUpstream(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Upstream{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *upstreamClient) UpdateUpstream(ctx context.Context, obj *Upstream, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *upstreamClient) PatchUpstream(ctx context.Context, obj *Upstream, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *upstreamClient) DeleteAllOfUpstream(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Upstream{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *upstreamClient) UpsertUpstream(ctx context.Context, obj *Upstream, transitionFuncs ...UpstreamTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*Upstream), desired.(*Upstream)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *upstreamClient) UpdateUpstreamStatus(ctx context.Context, obj *Upstream, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *upstreamClient) PatchUpstreamStatus(ctx context.Context, obj *Upstream, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides UpstreamClients for multiple clusters.
type MulticlusterUpstreamClient interface {
	// Cluster returns a UpstreamClient for the given cluster
	Cluster(cluster string) (UpstreamClient, error)
}

type multiclusterUpstreamClient struct {
	client multicluster.Client
}

func NewMulticlusterUpstreamClient(client multicluster.Client) MulticlusterUpstreamClient {
	return &multiclusterUpstreamClient{client: client}
}

func (m *multiclusterUpstreamClient) Cluster(cluster string) (UpstreamClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewUpstreamClient(client), nil
}

// Reader knows how to read and list UpstreamGroups.
type UpstreamGroupReader interface {
	// Get retrieves a UpstreamGroup for the given object key
	GetUpstreamGroup(ctx context.Context, key client.ObjectKey) (*UpstreamGroup, error)

	// List retrieves list of UpstreamGroups for a given namespace and list options.
	ListUpstreamGroup(ctx context.Context, opts ...client.ListOption) (*UpstreamGroupList, error)
}

// UpstreamGroupTransitionFunction instructs the UpstreamGroupWriter how to transition between an existing
// UpstreamGroup object and a desired on an Upsert
type UpstreamGroupTransitionFunction func(existing, desired *UpstreamGroup) error

// Writer knows how to create, delete, and update UpstreamGroups.
type UpstreamGroupWriter interface {
	// Create saves the UpstreamGroup object.
	CreateUpstreamGroup(ctx context.Context, obj *UpstreamGroup, opts ...client.CreateOption) error

	// Delete deletes the UpstreamGroup object.
	DeleteUpstreamGroup(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given UpstreamGroup object.
	UpdateUpstreamGroup(ctx context.Context, obj *UpstreamGroup, opts ...client.UpdateOption) error

	// Patch patches the given UpstreamGroup object.
	PatchUpstreamGroup(ctx context.Context, obj *UpstreamGroup, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all UpstreamGroup objects matching the given options.
	DeleteAllOfUpstreamGroup(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the UpstreamGroup object.
	UpsertUpstreamGroup(ctx context.Context, obj *UpstreamGroup, transitionFuncs ...UpstreamGroupTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a UpstreamGroup object.
type UpstreamGroupStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given UpstreamGroup object.
	UpdateUpstreamGroupStatus(ctx context.Context, obj *UpstreamGroup, opts ...client.UpdateOption) error

	// Patch patches the given UpstreamGroup object's subresource.
	PatchUpstreamGroupStatus(ctx context.Context, obj *UpstreamGroup, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on UpstreamGroups.
type UpstreamGroupClient interface {
	UpstreamGroupReader
	UpstreamGroupWriter
	UpstreamGroupStatusWriter
}

type upstreamGroupClient struct {
	client client.Client
}

func NewUpstreamGroupClient(client client.Client) *upstreamGroupClient {
	return &upstreamGroupClient{client: client}
}

func (c *upstreamGroupClient) GetUpstreamGroup(ctx context.Context, key client.ObjectKey) (*UpstreamGroup, error) {
	obj := &UpstreamGroup{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *upstreamGroupClient) ListUpstreamGroup(ctx context.Context, opts ...client.ListOption) (*UpstreamGroupList, error) {
	list := &UpstreamGroupList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *upstreamGroupClient) CreateUpstreamGroup(ctx context.Context, obj *UpstreamGroup, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *upstreamGroupClient) DeleteUpstreamGroup(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &UpstreamGroup{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *upstreamGroupClient) UpdateUpstreamGroup(ctx context.Context, obj *UpstreamGroup, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *upstreamGroupClient) PatchUpstreamGroup(ctx context.Context, obj *UpstreamGroup, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *upstreamGroupClient) DeleteAllOfUpstreamGroup(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &UpstreamGroup{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *upstreamGroupClient) UpsertUpstreamGroup(ctx context.Context, obj *UpstreamGroup, transitionFuncs ...UpstreamGroupTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*UpstreamGroup), desired.(*UpstreamGroup)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *upstreamGroupClient) UpdateUpstreamGroupStatus(ctx context.Context, obj *UpstreamGroup, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *upstreamGroupClient) PatchUpstreamGroupStatus(ctx context.Context, obj *UpstreamGroup, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides UpstreamGroupClients for multiple clusters.
type MulticlusterUpstreamGroupClient interface {
	// Cluster returns a UpstreamGroupClient for the given cluster
	Cluster(cluster string) (UpstreamGroupClient, error)
}

type multiclusterUpstreamGroupClient struct {
	client multicluster.Client
}

func NewMulticlusterUpstreamGroupClient(client multicluster.Client) MulticlusterUpstreamGroupClient {
	return &multiclusterUpstreamGroupClient{client: client}
}

func (m *multiclusterUpstreamGroupClient) Cluster(cluster string) (UpstreamGroupClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewUpstreamGroupClient(client), nil
}

// Reader knows how to read and list Proxys.
type ProxyReader interface {
	// Get retrieves a Proxy for the given object key
	GetProxy(ctx context.Context, key client.ObjectKey) (*Proxy, error)

	// List retrieves list of Proxys for a given namespace and list options.
	ListProxy(ctx context.Context, opts ...client.ListOption) (*ProxyList, error)
}

// ProxyTransitionFunction instructs the ProxyWriter how to transition between an existing
// Proxy object and a desired on an Upsert
type ProxyTransitionFunction func(existing, desired *Proxy) error

// Writer knows how to create, delete, and update Proxys.
type ProxyWriter interface {
	// Create saves the Proxy object.
	CreateProxy(ctx context.Context, obj *Proxy, opts ...client.CreateOption) error

	// Delete deletes the Proxy object.
	DeleteProxy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Proxy object.
	UpdateProxy(ctx context.Context, obj *Proxy, opts ...client.UpdateOption) error

	// Patch patches the given Proxy object.
	PatchProxy(ctx context.Context, obj *Proxy, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Proxy objects matching the given options.
	DeleteAllOfProxy(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Proxy object.
	UpsertProxy(ctx context.Context, obj *Proxy, transitionFuncs ...ProxyTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Proxy object.
type ProxyStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Proxy object.
	UpdateProxyStatus(ctx context.Context, obj *Proxy, opts ...client.UpdateOption) error

	// Patch patches the given Proxy object's subresource.
	PatchProxyStatus(ctx context.Context, obj *Proxy, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Proxys.
type ProxyClient interface {
	ProxyReader
	ProxyWriter
	ProxyStatusWriter
}

type proxyClient struct {
	client client.Client
}

func NewProxyClient(client client.Client) *proxyClient {
	return &proxyClient{client: client}
}

func (c *proxyClient) GetProxy(ctx context.Context, key client.ObjectKey) (*Proxy, error) {
	obj := &Proxy{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *proxyClient) ListProxy(ctx context.Context, opts ...client.ListOption) (*ProxyList, error) {
	list := &ProxyList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *proxyClient) CreateProxy(ctx context.Context, obj *Proxy, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *proxyClient) DeleteProxy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Proxy{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *proxyClient) UpdateProxy(ctx context.Context, obj *Proxy, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *proxyClient) PatchProxy(ctx context.Context, obj *Proxy, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *proxyClient) DeleteAllOfProxy(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Proxy{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *proxyClient) UpsertProxy(ctx context.Context, obj *Proxy, transitionFuncs ...ProxyTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*Proxy), desired.(*Proxy)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *proxyClient) UpdateProxyStatus(ctx context.Context, obj *Proxy, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *proxyClient) PatchProxyStatus(ctx context.Context, obj *Proxy, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides ProxyClients for multiple clusters.
type MulticlusterProxyClient interface {
	// Cluster returns a ProxyClient for the given cluster
	Cluster(cluster string) (ProxyClient, error)
}

type multiclusterProxyClient struct {
	client multicluster.Client
}

func NewMulticlusterProxyClient(client multicluster.Client) MulticlusterProxyClient {
	return &multiclusterProxyClient{client: client}
}

func (m *multiclusterProxyClient) Cluster(cluster string) (ProxyClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewProxyClient(client), nil
}
