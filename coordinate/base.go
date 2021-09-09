package coordinate

type (
	BaseCoordinateManager interface {
		RegisterHttp(...interface{}) string
		RegisterWs(...interface{}) string
		RegisterRpc(...interface{}) string
		DeregisterHttp(string)
		DeregisterWs(string)
		DeregisterRpc(string)
		GetHttpService()interface{}
	}
)
