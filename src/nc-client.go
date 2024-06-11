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
	"github.com/nats-io/nats.go"

	ncs "github.com/sty-holdings/nats-connect-shared"
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

func (clientPtr *NCClient) SynaidaListTeams(request ncs.ListTeamsRequest) (reply []byte, errorInfo pi.ErrorInfo) {

	var (
		tReply *nats.Msg
	)

	tReply, errorInfo = listTeams(
		clientPtr.styhCustomerConfig.clientId, clientPtr.styhCustomerConfig.secretKey, clientPtr.styhCustomerConfig.username, clientPtr.natsService.InstanceName, request,
		clientPtr.natsService.ConnPtr,
	)

	reply = tReply.Data

	return
}
