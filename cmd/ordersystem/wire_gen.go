// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/luizhenrique-dev/go-order-manager/internal/entity"
	"github.com/luizhenrique-dev/go-order-manager/internal/event"
	"github.com/luizhenrique-dev/go-order-manager/internal/infra/database"
	"github.com/luizhenrique-dev/go-order-manager/internal/infra/web"
	"github.com/luizhenrique-dev/go-order-manager/internal/usecase"
	"github.com/luizhenrique-dev/go-order-manager/pkg/events"
	"github.com/google/wire"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, orderCreated, eventDispatcher)
	return createOrderUseCase
}

func NewListOrderUseCase(db *sql.DB) *usecase.ListOrderUseCase {
	orderRepository := database.NewOrderRepository(db)
	listOrderUseCase := usecase.NewListOrderUseCase(orderRepository)
	return listOrderUseCase
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	webOrderHandler := web.NewWebOrderHandler(eventDispatcher, orderRepository, orderCreated)
	return webOrderHandler
}

// wire.go:

var setOrderRepositoryDependency = wire.NewSet(database.NewOrderRepository, wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)))

var setEventDispatcherDependency = wire.NewSet(events.NewEventDispatcher, event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)), wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))

var setOrderCreatedEvent = wire.NewSet(event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)))
