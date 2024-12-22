module github.com/ojlogic/GO/multi_module_shop/app

go 1.23.4

require (
	github.com/ojlogic/GO/multi_module_shop/orders v0.0.0
	github.com/ojlogic/GO/multi_module_shop/users v0.0.0
)

replace github.com/ojlogic/GO/multi_module_shop/orders => ../orders
replace github.com/ojlogic/GO/multi_module_shop/users => ../users