package main

import (
	connector "github.com/chernienkoii/Connector_1C_Enterprise/connector"
	handlers "github.com/chernienkoii/Connector_1C_Enterprise/handlers"
	rootsctuct "github.com/chernienkoii/Connector_1C_Enterprise/rootdescription"
)

func main() {

	rootsctuct.LoggerConnV.InitLog()
	connector.ConnectorV.LoggerConn = rootsctuct.LoggerConnV

	rootsctuct.Global_settingsV.LoadSettingsFromDisk()
	err := connector.ConnectorV.SetSettings(rootsctuct.Global_settingsV)

	if err != nil {
		connector.ConnectorV.LoggerConn.ErrorLogger.Println(err.Error())
	}

	handlers.StratHandlers()
}
