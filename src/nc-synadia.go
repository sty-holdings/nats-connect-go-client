// Package src
// /*
// Copyright 1/2024 STY Holdings Inc
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the “Software”), to deal in
// the Software without restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the
// Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.
// */
package src

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"

	ctv "github.com/sty-holdings/constant-type-vars-go/v2024"
	ncs "github.com/sty-holdings/nats-connect-shared/v2024"
	jwts "github.com/sty-holdings/sty-shared/v2024/jwtServices"
	ns "github.com/sty-holdings/sty-shared/v2024/natsSerices"
	pi "github.com/sty-holdings/sty-shared/v2024/programInfo"
)

// getPersonalAccessToken - will provide information about your token
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func getPersonalAccessToken(
	clientId, secretKey, username, instanceName string,
	request ncs.GetPersonalAccessTokenRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_GET_PERSONAL_ACCESS_TOKEN))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_GET_PERSONAL_ACCESS_TOKEN,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}

// getSystem - will provide information about the system
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func getSystem(
	clientId, secretKey, username, instanceName string,
	request ncs.GetSystemRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_GET_SYSTEM))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_GET_SYSTEM,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}

// getSystemLimits - will provide information about the system limits
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func getSystemLimits(
	clientId, secretKey, username, instanceName string,
	request ncs.GetSystemLimitsRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_GET_SYSTEM_LIMITS))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_GET_SYSTEM_LIMITS,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}

// getTeam - will provide information about the team
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func getTeam(
	clientId, secretKey, username, instanceName string,
	request ncs.GetTeamRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_GET_TEAM))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_GET_TEAM,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}

// getTeamLimits - will provide information about the team's limits
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func getTeamLimits(
	clientId, secretKey, username, instanceName string,
	request ncs.GetTeamLimitsRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_GET_TEAM_LIMITS))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_GET_TEAM_LIMITS,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}

// getVersion - will provide the version information
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func getVersion(
	clientId, secretKey, username, instanceName string,
	request ncs.GetVersionRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_GET_VERSION))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_GET_VERSION,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}

// listAccounts - will list the account for a system id
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func listAccounts(
	clientId, secretKey, username, instanceName string,
	request ncs.ListAccountsRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_LIST_ACCOUNT))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_LIST_ACCOUNT,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}

// listInfoAppUsersTeam - will list the user account for a team id
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func listInfoAppUsersTeam(
	clientId, secretKey, username, instanceName string,
	request ncs.ListInfoAppUserTeamRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_LIST_INFO_APP_USERS_TEAM))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_LIST_INFO_APP_USERS_TEAM,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}

// listNATSUsers - will list the NATS user for a team id
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func listNATSUsers(
	clientId, secretKey, username, instanceName string,
	request ncs.ListNATSUsersRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_LIST_NATS_USERS))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_LIST_NATS_USERS,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}

// listPersonalAccessTokens - will list your personal access tokens
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func listPersonalAccessTokens(
	clientId, secretKey, username, instanceName string,
	request ncs.ListPersonalAccessTokensRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_LIST_PERSONAL_ACCESS_TOKENS))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_LIST_PERSONAL_ACCESS_TOKENS,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}

// listSystems - will list systems for a team
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func listSystems(
	clientId, secretKey, username, instanceName string,
	request ncs.ListSystemsRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_LIST_SYSTEMS))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_LIST_SYSTEMS,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}

// listSystemAccountInfo - will list system account info
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func listSystemAccountInfo(
	clientId, secretKey, username, instanceName string,
	request ncs.ListSystemAccountInfoRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_LIST_SYSTEM_ACCOUN_TINFO))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_LIST_SYSTEM_ACCOUN_TINFO,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}

// listSystemServerInfo - will list server information for a server
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func listSystemServerInfo(
	clientId, secretKey, username, instanceName string,
	request ncs.ListSystemServerInfoRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_LIST_SYSTEM_SERVER_INFO))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_LIST_SYSTEM_SERVER_INFO,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}

// listTeamServerAccounts - will list all service accounts for the team
// This appears to be a restricted API. Only tested using a personal account.
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func listTeamServerAccounts(
	clientId, secretKey, username, instanceName string,
	request ncs.ListTeamServerAccountsRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_LIST_TEAM_SERVER_ACCOUNTS))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_LIST_TEAM_SERVER_ACCOUNTS,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}

// listTeams - returns information about all your teams
//
//	Customer Messages: None
//	Errors: returned from json.Marshal, jwts.Encrypt
//	Verifications: None
func listTeams(
	clientId, secretKey, username, instanceName string,
	request ncs.ListTeamsRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_LIST_TEAMS))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_LIST_TEAMS,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}
