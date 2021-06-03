// Cloud Info Manager's Rest Runtime of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// by CB-Spider Team, 2019.09.

package restruntime

import (
	"strconv"

	im "github.com/cloud-barista/poc-cicd-spider/cloud-info-manager"
	ccim "github.com/cloud-barista/poc-cicd-spider/cloud-info-manager/connection-config-info-manager"
	cim "github.com/cloud-barista/poc-cicd-spider/cloud-info-manager/credential-info-manager"
	dim "github.com/cloud-barista/poc-cicd-spider/cloud-info-manager/driver-info-manager"
	rim "github.com/cloud-barista/poc-cicd-spider/cloud-info-manager/region-info-manager"

	// REST API (echo)
	"net/http"

	"github.com/labstack/echo/v4"
)

//================ List of support CloudOS
func ListCloudOS(c echo.Context) error {
	cblog.Info("call listCloudOS()")

	infoList := im.ListCloudOS()

	var jsonResult struct {
		Result []string `json:"cloudos"`
	}
	jsonResult.Result = infoList
	return c.JSON(http.StatusOK, &jsonResult)
}

//================ CloudDriver Handler
func RegisterCloudDriver(c echo.Context) error {
	cblog.Info("call registerCloudDriver()")
	req := &dim.CloudDriverInfo{}
	if err := c.Bind(req); err != nil {
		cblog.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	cldinfoList, err := dim.RegisterCloudDriverInfo(*req)
	if err != nil {
		cblog.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &cldinfoList)
}

func ListCloudDriver(c echo.Context) error {
	cblog.Info("call listCloudDriver()")

	infoList, err := dim.ListCloudDriver()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var jsonResult struct {
		Result []*dim.CloudDriverInfo `json:"driver"`
	}
	if infoList == nil {
		infoList = []*dim.CloudDriverInfo{}
	}
	jsonResult.Result = infoList
	return c.JSON(http.StatusOK, &jsonResult)

}

func GetCloudDriver(c echo.Context) error {
	cblog.Info("call getCloudDriver()")

	cldinfo, err := dim.GetCloudDriver(c.Param("DriverName"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &cldinfo)
}

func UnRegisterCloudDriver(c echo.Context) error {
	cblog.Info("call unRegisterCloudDriver()")

	result, err := dim.UnRegisterCloudDriver(c.Param("DriverName"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resultInfo := BooleanInfo{
		Result: strconv.FormatBool(result),
	}

	return c.JSON(http.StatusOK, &resultInfo)
}

//================ Credential Handler
func RegisterCredential(c echo.Context) error {
	cblog.Info("call registerCredential()")

	req := &cim.CredentialInfo{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	crdinfoList, err := cim.RegisterCredentialInfo(*req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &crdinfoList)
}

func ListCredential(c echo.Context) error {
	cblog.Info("call listCredential()")

	infoList, err := cim.ListCredential()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var jsonResult struct {
		Result []*cim.CredentialInfo `json:"credential"`
	}
	if infoList == nil {
		infoList = []*cim.CredentialInfo{}
	}
	jsonResult.Result = infoList
	return c.JSON(http.StatusOK, &jsonResult)
}

func GetCredential(c echo.Context) error {
	cblog.Info("call getCredential()")

	crdinfo, err := cim.GetCredential(c.Param("CredentialName"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &crdinfo)
}

func UnRegisterCredential(c echo.Context) error {
	cblog.Info("call unRegisterCredential()")

	result, err := cim.UnRegisterCredential(c.Param("CredentialName"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resultInfo := BooleanInfo{
		Result: strconv.FormatBool(result),
	}

	return c.JSON(http.StatusOK, &resultInfo)
}

//================ Region Handler
func RegisterRegion(c echo.Context) error {
	cblog.Info("call registerRegion()")

	req := &rim.RegionInfo{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	crdinfoList, err := rim.RegisterRegionInfo(*req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &crdinfoList)
}

func ListRegion(c echo.Context) error {
	cblog.Info("call listRegion()")

	infoList, err := rim.ListRegion()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var jsonResult struct {
		Result []*rim.RegionInfo `json:"region"`
	}
	if infoList == nil {
		infoList = []*rim.RegionInfo{}
	}
	jsonResult.Result = infoList
	return c.JSON(http.StatusOK, &jsonResult)
}

func GetRegion(c echo.Context) error {
	cblog.Info("call getRegion()")

	crdinfo, err := rim.GetRegion(c.Param("RegionName"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &crdinfo)
}

func UnRegisterRegion(c echo.Context) error {
	cblog.Info("call unRegisterRegion()")

	result, err := rim.UnRegisterRegion(c.Param("RegionName"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resultInfo := BooleanInfo{
		Result: strconv.FormatBool(result),
	}

	return c.JSON(http.StatusOK, &resultInfo)
}

//================ ConnectionConfig Handler
func CreateConnectionConfig(c echo.Context) error {
	cblog.Info("call registerConnectionConfig()")

	req := &ccim.ConnectionConfigInfo{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	crdinfoList, err := ccim.CreateConnectionConfigInfo(*req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &crdinfoList)
}

func ListConnectionConfig(c echo.Context) error {
	cblog.Info("call listConnectionConfig()")

	infoList, err := ccim.ListConnectionConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var jsonResult struct {
		Result []*ccim.ConnectionConfigInfo `json:"connectionconfig"`
	}
	if infoList == nil {
		infoList = []*ccim.ConnectionConfigInfo{}
	}
	jsonResult.Result = infoList
	return c.JSON(http.StatusOK, &jsonResult)
}

func GetConnectionConfig(c echo.Context) error {
	cblog.Info("call getConnectionConfig()")

	crdinfo, err := ccim.GetConnectionConfig(c.Param("ConfigName"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &crdinfo)
}

func DeleteConnectionConfig(c echo.Context) error {
	cblog.Info("call deleteConnectionConfig()")

	result, err := ccim.DeleteConnectionConfig(c.Param("ConfigName"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resultInfo := BooleanInfo{
		Result: strconv.FormatBool(result),
	}

	return c.JSON(http.StatusOK, &resultInfo)
}
