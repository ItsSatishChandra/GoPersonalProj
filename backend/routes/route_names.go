package routes

type FunctionName string

func RouteNameResolver(functionName string) string {

	routeNames := map[string]string{}

	//Add package.function as key and  resolve path as value
	// Example routeNames["package.function"]="/resolves/route"
	routeNames["hello.Hello"] = "/"

	return routeNames[functionName]
}
