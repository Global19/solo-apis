// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	gateway_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the Gateway Resource
// DEPRECATED: Prefer reconciler pattern.
type GatewayEventHandler interface {
	CreateGateway(obj *gateway_solo_io_v1.Gateway) error
	UpdateGateway(old, new *gateway_solo_io_v1.Gateway) error
	DeleteGateway(obj *gateway_solo_io_v1.Gateway) error
	GenericGateway(obj *gateway_solo_io_v1.Gateway) error
}

type GatewayEventHandlerFuncs struct {
	OnCreate  func(obj *gateway_solo_io_v1.Gateway) error
	OnUpdate  func(old, new *gateway_solo_io_v1.Gateway) error
	OnDelete  func(obj *gateway_solo_io_v1.Gateway) error
	OnGeneric func(obj *gateway_solo_io_v1.Gateway) error
}

func (f *GatewayEventHandlerFuncs) CreateGateway(obj *gateway_solo_io_v1.Gateway) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *GatewayEventHandlerFuncs) DeleteGateway(obj *gateway_solo_io_v1.Gateway) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *GatewayEventHandlerFuncs) UpdateGateway(objOld, objNew *gateway_solo_io_v1.Gateway) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *GatewayEventHandlerFuncs) GenericGateway(obj *gateway_solo_io_v1.Gateway) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type GatewayEventWatcher interface {
	AddEventHandler(ctx context.Context, h GatewayEventHandler, predicates ...predicate.Predicate) error
}

type gatewayEventWatcher struct {
	watcher events.EventWatcher
}

func NewGatewayEventWatcher(name string, mgr manager.Manager) GatewayEventWatcher {
	return &gatewayEventWatcher{
		watcher: events.NewWatcher(name, mgr, &gateway_solo_io_v1.Gateway{}),
	}
}

func (c *gatewayEventWatcher) AddEventHandler(ctx context.Context, h GatewayEventHandler, predicates ...predicate.Predicate) error {
	handler := genericGatewayHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericGatewayHandler implements a generic events.EventHandler
type genericGatewayHandler struct {
	handler GatewayEventHandler
}

func (h genericGatewayHandler) Create(object client.Object) error {
	obj, ok := object.(*gateway_solo_io_v1.Gateway)
	if !ok {
		return errors.Errorf("internal error: Gateway handler received event for %T", object)
	}
	return h.handler.CreateGateway(obj)
}

func (h genericGatewayHandler) Delete(object client.Object) error {
	obj, ok := object.(*gateway_solo_io_v1.Gateway)
	if !ok {
		return errors.Errorf("internal error: Gateway handler received event for %T", object)
	}
	return h.handler.DeleteGateway(obj)
}

func (h genericGatewayHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*gateway_solo_io_v1.Gateway)
	if !ok {
		return errors.Errorf("internal error: Gateway handler received event for %T", old)
	}
	objNew, ok := new.(*gateway_solo_io_v1.Gateway)
	if !ok {
		return errors.Errorf("internal error: Gateway handler received event for %T", new)
	}
	return h.handler.UpdateGateway(objOld, objNew)
}

func (h genericGatewayHandler) Generic(object client.Object) error {
	obj, ok := object.(*gateway_solo_io_v1.Gateway)
	if !ok {
		return errors.Errorf("internal error: Gateway handler received event for %T", object)
	}
	return h.handler.GenericGateway(obj)
}

// Handle events for the RouteTable Resource
// DEPRECATED: Prefer reconciler pattern.
type RouteTableEventHandler interface {
	CreateRouteTable(obj *gateway_solo_io_v1.RouteTable) error
	UpdateRouteTable(old, new *gateway_solo_io_v1.RouteTable) error
	DeleteRouteTable(obj *gateway_solo_io_v1.RouteTable) error
	GenericRouteTable(obj *gateway_solo_io_v1.RouteTable) error
}

type RouteTableEventHandlerFuncs struct {
	OnCreate  func(obj *gateway_solo_io_v1.RouteTable) error
	OnUpdate  func(old, new *gateway_solo_io_v1.RouteTable) error
	OnDelete  func(obj *gateway_solo_io_v1.RouteTable) error
	OnGeneric func(obj *gateway_solo_io_v1.RouteTable) error
}

func (f *RouteTableEventHandlerFuncs) CreateRouteTable(obj *gateway_solo_io_v1.RouteTable) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *RouteTableEventHandlerFuncs) DeleteRouteTable(obj *gateway_solo_io_v1.RouteTable) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *RouteTableEventHandlerFuncs) UpdateRouteTable(objOld, objNew *gateway_solo_io_v1.RouteTable) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *RouteTableEventHandlerFuncs) GenericRouteTable(obj *gateway_solo_io_v1.RouteTable) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type RouteTableEventWatcher interface {
	AddEventHandler(ctx context.Context, h RouteTableEventHandler, predicates ...predicate.Predicate) error
}

type routeTableEventWatcher struct {
	watcher events.EventWatcher
}

func NewRouteTableEventWatcher(name string, mgr manager.Manager) RouteTableEventWatcher {
	return &routeTableEventWatcher{
		watcher: events.NewWatcher(name, mgr, &gateway_solo_io_v1.RouteTable{}),
	}
}

func (c *routeTableEventWatcher) AddEventHandler(ctx context.Context, h RouteTableEventHandler, predicates ...predicate.Predicate) error {
	handler := genericRouteTableHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericRouteTableHandler implements a generic events.EventHandler
type genericRouteTableHandler struct {
	handler RouteTableEventHandler
}

func (h genericRouteTableHandler) Create(object client.Object) error {
	obj, ok := object.(*gateway_solo_io_v1.RouteTable)
	if !ok {
		return errors.Errorf("internal error: RouteTable handler received event for %T", object)
	}
	return h.handler.CreateRouteTable(obj)
}

func (h genericRouteTableHandler) Delete(object client.Object) error {
	obj, ok := object.(*gateway_solo_io_v1.RouteTable)
	if !ok {
		return errors.Errorf("internal error: RouteTable handler received event for %T", object)
	}
	return h.handler.DeleteRouteTable(obj)
}

func (h genericRouteTableHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*gateway_solo_io_v1.RouteTable)
	if !ok {
		return errors.Errorf("internal error: RouteTable handler received event for %T", old)
	}
	objNew, ok := new.(*gateway_solo_io_v1.RouteTable)
	if !ok {
		return errors.Errorf("internal error: RouteTable handler received event for %T", new)
	}
	return h.handler.UpdateRouteTable(objOld, objNew)
}

func (h genericRouteTableHandler) Generic(object client.Object) error {
	obj, ok := object.(*gateway_solo_io_v1.RouteTable)
	if !ok {
		return errors.Errorf("internal error: RouteTable handler received event for %T", object)
	}
	return h.handler.GenericRouteTable(obj)
}

// Handle events for the VirtualService Resource
// DEPRECATED: Prefer reconciler pattern.
type VirtualServiceEventHandler interface {
	CreateVirtualService(obj *gateway_solo_io_v1.VirtualService) error
	UpdateVirtualService(old, new *gateway_solo_io_v1.VirtualService) error
	DeleteVirtualService(obj *gateway_solo_io_v1.VirtualService) error
	GenericVirtualService(obj *gateway_solo_io_v1.VirtualService) error
}

type VirtualServiceEventHandlerFuncs struct {
	OnCreate  func(obj *gateway_solo_io_v1.VirtualService) error
	OnUpdate  func(old, new *gateway_solo_io_v1.VirtualService) error
	OnDelete  func(obj *gateway_solo_io_v1.VirtualService) error
	OnGeneric func(obj *gateway_solo_io_v1.VirtualService) error
}

func (f *VirtualServiceEventHandlerFuncs) CreateVirtualService(obj *gateway_solo_io_v1.VirtualService) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *VirtualServiceEventHandlerFuncs) DeleteVirtualService(obj *gateway_solo_io_v1.VirtualService) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *VirtualServiceEventHandlerFuncs) UpdateVirtualService(objOld, objNew *gateway_solo_io_v1.VirtualService) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *VirtualServiceEventHandlerFuncs) GenericVirtualService(obj *gateway_solo_io_v1.VirtualService) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type VirtualServiceEventWatcher interface {
	AddEventHandler(ctx context.Context, h VirtualServiceEventHandler, predicates ...predicate.Predicate) error
}

type virtualServiceEventWatcher struct {
	watcher events.EventWatcher
}

func NewVirtualServiceEventWatcher(name string, mgr manager.Manager) VirtualServiceEventWatcher {
	return &virtualServiceEventWatcher{
		watcher: events.NewWatcher(name, mgr, &gateway_solo_io_v1.VirtualService{}),
	}
}

func (c *virtualServiceEventWatcher) AddEventHandler(ctx context.Context, h VirtualServiceEventHandler, predicates ...predicate.Predicate) error {
	handler := genericVirtualServiceHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericVirtualServiceHandler implements a generic events.EventHandler
type genericVirtualServiceHandler struct {
	handler VirtualServiceEventHandler
}

func (h genericVirtualServiceHandler) Create(object client.Object) error {
	obj, ok := object.(*gateway_solo_io_v1.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return h.handler.CreateVirtualService(obj)
}

func (h genericVirtualServiceHandler) Delete(object client.Object) error {
	obj, ok := object.(*gateway_solo_io_v1.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return h.handler.DeleteVirtualService(obj)
}

func (h genericVirtualServiceHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*gateway_solo_io_v1.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", old)
	}
	objNew, ok := new.(*gateway_solo_io_v1.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", new)
	}
	return h.handler.UpdateVirtualService(objOld, objNew)
}

func (h genericVirtualServiceHandler) Generic(object client.Object) error {
	obj, ok := object.(*gateway_solo_io_v1.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return h.handler.GenericVirtualService(obj)
}
