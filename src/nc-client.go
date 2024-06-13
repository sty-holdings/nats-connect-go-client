// Package src
// /*
// ====> This is a sample usage of NATS Connect. The CLI part is to allow an easy place to start.
// ====> The run function is the code to drop into you program.
//
// Copyright 6/5/24 STY Holdings Inc
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

	"github.com/nats-io/nats.go"

	ctv "github.com/sty-holdings/constant-type-vars-go/v2024"
	ncs "github.com/sty-holdings/nats-connect-shared/v2024"
	awss "github.com/sty-holdings/sty-shared/v2024/awsServices"
	ns "github.com/sty-holdings/sty-shared/v2024/natsSerices"
	pi "github.com/sty-holdings/sty-shared/v2024/programInfo"
)

//goland:noinspection ALL
const ()

type NCClient struct {
	awsSettings        awss.AWSSettings
	environment        string
	natsService        ns.NATSService
	natsConfig         ns.NATSConfiguration
	styhCustomerConfig styhCustomerConfig
	tempDirectory      string
}

type SaaSKeysTokens struct {
	SendGridKey  string `json:"sendgrid_key"`
	StripeKey    string `json:"stripe_key"`
	SynadiaToken string `json:"synadia_token"`
}

// SynaidaGetPersonalAccessToken - will provide information about your token
func (clientPtr *NCClient) SynaidaGetPersonalAccessToken(request interface{}) (reply ncs.GetPersonalAccessTokenReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = getPersonalAccessToken(
		clientPtr.styhCustomerConfig.clientId,
		clientPtr.styhCustomerConfig.secretKey,
		clientPtr.styhCustomerConfig.username,
		clientPtr.natsService.InstanceName,
		request.(ncs.GetPersonalAccessTokenRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}

// SynaidaGetSystem - will provide information about the system
func (clientPtr *NCClient) SynaidaGetSystem(request interface{}) (reply ncs.GetSystemReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = getSystem(
		clientPtr.styhCustomerConfig.clientId,
		clientPtr.styhCustomerConfig.secretKey,
		clientPtr.styhCustomerConfig.username,
		clientPtr.natsService.InstanceName,
		request.(ncs.GetSystemRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}

// SynaidaGetSystemLimits - will provide information about the system limits
func (clientPtr *NCClient) SynaidaGetSystemLimits(request interface{}) (reply ncs.GetSystemLimitsReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = getSystemLimits(
		clientPtr.styhCustomerConfig.clientId,
		clientPtr.styhCustomerConfig.secretKey,
		clientPtr.styhCustomerConfig.username,
		clientPtr.natsService.InstanceName,
		request.(ncs.GetSystemLimitsRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}

// SynaidaGetTeam - will provide information about the team
func (clientPtr *NCClient) SynaidaGetTeam(request interface{}) (reply ncs.GetTeamReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = getTeam(
		clientPtr.styhCustomerConfig.clientId,
		clientPtr.styhCustomerConfig.secretKey,
		clientPtr.styhCustomerConfig.username,
		clientPtr.natsService.InstanceName,
		request.(ncs.GetTeamRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}

// SynaidaGetTeamLimits - will provide information about the team's limits
func (clientPtr *NCClient) SynaidaGetTeamLimits(request interface{}) (reply ncs.GetTeamLimitsReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = getTeamLimits(
		clientPtr.styhCustomerConfig.clientId,
		clientPtr.styhCustomerConfig.secretKey,
		clientPtr.styhCustomerConfig.username,
		clientPtr.natsService.InstanceName,
		request.(ncs.GetTeamLimitsRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}

// SynaidaGetVersion - will provide the version information
func (clientPtr *NCClient) SynaidaGetVersion(request interface{}) (reply ncs.GetVersionReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = getVersion(
		clientPtr.styhCustomerConfig.clientId,
		clientPtr.styhCustomerConfig.secretKey,
		clientPtr.styhCustomerConfig.username,
		clientPtr.natsService.InstanceName,
		request.(ncs.GetVersionRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}

// SynaidaListAccounts - will list the account for a system id
func (clientPtr *NCClient) SynaidaListAccounts(request interface{}) (reply ncs.ListAccountsReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = listAccounts(
		clientPtr.styhCustomerConfig.clientId,
		clientPtr.styhCustomerConfig.secretKey,
		clientPtr.styhCustomerConfig.username,
		clientPtr.natsService.InstanceName,
		request.(ncs.ListAccountsRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}

// SynaidaListInfoAppUsersTeam - will list the user account for a team id
func (clientPtr *NCClient) SynaidaListInfoAppUsersTeam(request interface{}) (reply ncs.ListInfoAppUsersTeamReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = listInfoAppUsersTeam(
		clientPtr.styhCustomerConfig.clientId,
		clientPtr.styhCustomerConfig.secretKey,
		clientPtr.styhCustomerConfig.username,
		clientPtr.natsService.InstanceName,
		request.(ncs.ListInfoAppUserTeamRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}

func (clientPtr *NCClient) SynaidaListNATSUsers(request interface{}) (reply ncs.ListNATSUsersReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = listNATSUsers(
		clientPtr.styhCustomerConfig.clientId, clientPtr.styhCustomerConfig.secretKey, clientPtr.styhCustomerConfig.username, clientPtr.natsService.InstanceName,
		request.(ncs.ListNATSUsersRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}

// SynaidaListPersonalAccessTokens - will list your personal access tokens
func (clientPtr *NCClient) SynaidaListPersonalAccessTokens(request interface{}) (reply ncs.ListPersonalAccessTokensReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = listPersonalAccessTokens(
		clientPtr.styhCustomerConfig.clientId,
		clientPtr.styhCustomerConfig.secretKey,
		clientPtr.styhCustomerConfig.username,
		clientPtr.natsService.InstanceName,
		request.(ncs.ListPersonalAccessTokensRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}

// SynaidaListSystems - will list systems for a team
func (clientPtr *NCClient) SynaidaListSystems(request interface{}) (reply ncs.ListSystemsReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = listSystems(
		clientPtr.styhCustomerConfig.clientId,
		clientPtr.styhCustomerConfig.secretKey,
		clientPtr.styhCustomerConfig.username,
		clientPtr.natsService.InstanceName,
		request.(ncs.ListSystemsRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}

// SynaidaListSystemAccountInfo - will list system account info
func (clientPtr *NCClient) SynaidaListSystemAccountInfo(request interface{}) (reply ncs.ListSystemAccountInfoReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = listSystemAccountInfo(
		clientPtr.styhCustomerConfig.clientId,
		clientPtr.styhCustomerConfig.secretKey,
		clientPtr.styhCustomerConfig.username,
		clientPtr.natsService.InstanceName,
		request.(ncs.ListSystemAccountInfoRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}

func (clientPtr *NCClient) SynaidaListSystemServerInfo(request interface{}) (reply ncs.ListSystemServerInfoReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = listSystemServerInfo(
		clientPtr.styhCustomerConfig.clientId,
		clientPtr.styhCustomerConfig.secretKey,
		clientPtr.styhCustomerConfig.username,
		clientPtr.natsService.InstanceName,
		request.(ncs.ListSystemServerInfoRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}

// SynaidaListTeamServerAccounts - will list all service accounts for the team
// This appears to be a restricted API. Only tested using a personal account.
func (clientPtr *NCClient) SynaidaListTeamServerAccounts(request interface{}) (reply ncs.ListTeamServerAccountsReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = listTeamServerAccounts(
		clientPtr.styhCustomerConfig.clientId,
		clientPtr.styhCustomerConfig.secretKey,
		clientPtr.styhCustomerConfig.username,
		clientPtr.natsService.InstanceName,
		request.(ncs.ListTeamServerAccountsRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}

// SynaidaListTeams - returns information about all your teams
func (clientPtr *NCClient) SynaidaListTeams(request interface{}) (reply ncs.ListTeamsReply, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	if tReply, errorInfo = listTeams(
		clientPtr.styhCustomerConfig.clientId,
		clientPtr.styhCustomerConfig.secretKey,
		clientPtr.styhCustomerConfig.username,
		clientPtr.natsService.InstanceName,
		request.(ncs.ListTeamsRequest),
		clientPtr.natsService.ConnPtr,
	); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	if errorInfo.Error = json.Unmarshal(tReply.Data, &reply); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
	}

	return
}
