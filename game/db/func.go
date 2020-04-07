package db

import (
	"database/sql"
	"fmt"

	"github.com/YWJSonic/ServerUtility/dbinfo"
	"github.com/YWJSonic/ServerUtility/dbservice"
	"github.com/YWJSonic/ServerUtility/foundation"
	"github.com/YWJSonic/ServerUtility/messagehandle"
)

// GetSetting get db setting
func GetSetting(db *sql.DB) ([]map[string]interface{}, messagehandle.ErrorMsg) {
	result, err := dbservice.CallReadOutMap(db, "SettingGet_Read")
	return result, err
}

// GetSettingKey get db setting
func GetSettingKey(db *sql.DB, key string) ([]map[string]interface{}, messagehandle.ErrorMsg) {
	result, err := dbservice.CallReadOutMap(db, "SettingKeyGet_Read", key)
	return result, err
}

// NewSetting ...
func NewSetting(db *sql.DB, args ...interface{}) {
	dbservice.CallWrite(
		db,
		dbservice.MakeProcedureQueryStr("SettingNew_Write", len(args)),
		args...,
	)
}

// UpdateSetting ...
func UpdateSetting(db *sql.DB, args ...interface{}) messagehandle.ErrorMsg {
	_, err := dbservice.CallWrite(db, dbservice.MakeProcedureQueryStr("SettingSet_Update", len(args)), args...)
	return err
}

// ReflushSetting ...
func ReflushSetting(db *sql.DB, args ...interface{}) messagehandle.ErrorMsg {
	args = append(args, foundation.ServerNowTime())
	_, err := dbservice.CallWrite(db, dbservice.MakeProcedureQueryStr("SettingSet_Update_v2", len(args)), args...)
	return err
}

// GetAttachTypeRange ...
func GetAttachTypeRange(db *sql.DB, playerid, kind, miniAttType, maxAttType int64) ([]map[string]interface{}, messagehandle.ErrorMsg) {
	result, err := dbservice.CallReadOutMap(db, "AttachTypeRangeGet_Read", playerid, kind, miniAttType, maxAttType)
	return result, err
}

// GetAttachType ...
func GetAttachType(db *sql.DB, playerid int64, kind int64, attType int64) ([]map[string]interface{}, messagehandle.ErrorMsg) {
	result, err := dbservice.CallReadOutMap(db, "AttachTypeGet_Read", playerid, kind, attType)
	return result, err
}

// GetAttachKind get db attach kind
func GetAttachKind(db *sql.DB, playerid int64, kind int64) ([]map[string]interface{}, messagehandle.ErrorMsg) {
	result, err := dbservice.CallReadOutMap(db, "AttachKindGet_Read", playerid, kind)
	return result, err
}

// NewAttach ...
func NewAttach(db *sql.DB, args ...interface{}) {
	dbservice.CallWrite(
		db,
		dbservice.MakeProcedureQueryStr("AttachNew_Write", len(args)),
		args...,
	)
}

// UpdateAttach ...
func UpdateAttach(db *sql.DB, args ...interface{}) messagehandle.ErrorMsg {
	_, err := dbservice.CallWrite(db, dbservice.MakeProcedureQueryStr("AttachSet_Write", len(args)), args...)
	return err
}

// SetLog new goruting set log
func SetLog(db *sql.DB, account string, playerID, time int64, activityEvent uint8, iValue1, iValue2, iValue3 int64, sValue1, sValue2, sValue3, msg string) messagehandle.ErrorMsg {
	tableName := foundation.ServerNow().Format("20060102")
	query := fmt.Sprintf("INSERT INTO `%s` VALUE(NULL,\"%s\",%d,%d, %d, %d,%d,%d,\"%s\",\"%s\",\"%s\",\"%s\");", tableName, account, playerID, time, activityEvent, iValue1, iValue2, iValue3, sValue1, sValue2, sValue3, msg)
	_, err := dbinfo.CallWrite(db, query)
	return err
}
