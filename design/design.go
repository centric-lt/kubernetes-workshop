package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("pods", func() {
	Title("pods service")
	Description("Service for pod info")
	Server("PodInfo", func() {
		Host("0.0.0.0", func() {
			URI("http://0.0.0.0:8080")
		})
	})
})

var _ = Service("podinfo", func() {
	Description("The podinfo service pulls info about current pod")

	Method("get", func() {
		Result(PodInfoResult)

		HTTP(func() {
			GET("/pod")
			Response(StatusOK)
			Response(StatusInternalServerError)
		})
	})

	Files("/", "static/index.html")
	Files("/ui/{*path}", "static/")
})

var PodInfoResult = ResultType("PodInfoResult", func() {
	Description("Pod Info Response")
	Attributes(func() {
		Attribute("ip", String, "POD ip address")
		Attribute("hostname", String, "POD hostname")
		Required("ip", "hostname")
	})
})
