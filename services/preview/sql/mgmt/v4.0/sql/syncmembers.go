package sql

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SyncMembersClient is the the Azure SQL Database management API provides a RESTful set of web services that interact
// with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update, and
// delete databases.
type SyncMembersClient struct {
	BaseClient
}

// NewSyncMembersClient creates an instance of the SyncMembersClient client.
func NewSyncMembersClient(subscriptionID string) SyncMembersClient {
	return NewSyncMembersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSyncMembersClientWithBaseURI creates an instance of the SyncMembersClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSyncMembersClientWithBaseURI(baseURI string, subscriptionID string) SyncMembersClient {
	return SyncMembersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a sync member.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database on which the sync group is hosted.
// syncGroupName - the name of the sync group on which the sync member is hosted.
// syncMemberName - the name of the sync member.
// parameters - the requested sync member resource state.
func (client SyncMembersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters SyncMember) (result SyncMembersCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SyncMembersClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SyncMembersClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters SyncMember) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":     autorest.Encode("path", syncGroupName),
		"syncMemberName":    autorest.Encode("path", syncMemberName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SyncMembersClient) CreateOrUpdateSender(req *http.Request) (future SyncMembersCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SyncMembersClient) CreateOrUpdateResponder(resp *http.Response) (result SyncMember, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a sync member.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database on which the sync group is hosted.
// syncGroupName - the name of the sync group on which the sync member is hosted.
// syncMemberName - the name of the sync member.
func (client SyncMembersClient) Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (result SyncMembersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SyncMembersClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SyncMembersClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":     autorest.Encode("path", syncGroupName),
		"syncMemberName":    autorest.Encode("path", syncMemberName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SyncMembersClient) DeleteSender(req *http.Request) (future SyncMembersDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SyncMembersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a sync member.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database on which the sync group is hosted.
// syncGroupName - the name of the sync group on which the sync member is hosted.
// syncMemberName - the name of the sync member.
func (client SyncMembersClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (result SyncMember, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SyncMembersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SyncMembersClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":     autorest.Encode("path", syncGroupName),
		"syncMemberName":    autorest.Encode("path", syncMemberName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SyncMembersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SyncMembersClient) GetResponder(resp *http.Response) (result SyncMember, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySyncGroup lists sync members in the given sync group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database on which the sync group is hosted.
// syncGroupName - the name of the sync group.
func (client SyncMembersClient) ListBySyncGroup(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string) (result SyncMemberListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SyncMembersClient.ListBySyncGroup")
		defer func() {
			sc := -1
			if result.smlr.Response.Response != nil {
				sc = result.smlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySyncGroupNextResults
	req, err := client.ListBySyncGroupPreparer(ctx, resourceGroupName, serverName, databaseName, syncGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "ListBySyncGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySyncGroupSender(req)
	if err != nil {
		result.smlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "ListBySyncGroup", resp, "Failure sending request")
		return
	}

	result.smlr, err = client.ListBySyncGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "ListBySyncGroup", resp, "Failure responding to request")
		return
	}
	if result.smlr.hasNextLink() && result.smlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySyncGroupPreparer prepares the ListBySyncGroup request.
func (client SyncMembersClient) ListBySyncGroupPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":     autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySyncGroupSender sends the ListBySyncGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SyncMembersClient) ListBySyncGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySyncGroupResponder handles the response to the ListBySyncGroup request. The method always
// closes the http.Response Body.
func (client SyncMembersClient) ListBySyncGroupResponder(resp *http.Response) (result SyncMemberListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySyncGroupNextResults retrieves the next set of results, if any.
func (client SyncMembersClient) listBySyncGroupNextResults(ctx context.Context, lastResults SyncMemberListResult) (result SyncMemberListResult, err error) {
	req, err := lastResults.syncMemberListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.SyncMembersClient", "listBySyncGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySyncGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.SyncMembersClient", "listBySyncGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySyncGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "listBySyncGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySyncGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SyncMembersClient) ListBySyncGroupComplete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string) (result SyncMemberListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SyncMembersClient.ListBySyncGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySyncGroup(ctx, resourceGroupName, serverName, databaseName, syncGroupName)
	return
}

// ListMemberSchemas gets a sync member database schema.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database on which the sync group is hosted.
// syncGroupName - the name of the sync group on which the sync member is hosted.
// syncMemberName - the name of the sync member.
func (client SyncMembersClient) ListMemberSchemas(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (result SyncFullSchemaPropertiesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SyncMembersClient.ListMemberSchemas")
		defer func() {
			sc := -1
			if result.sfsplr.Response.Response != nil {
				sc = result.sfsplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listMemberSchemasNextResults
	req, err := client.ListMemberSchemasPreparer(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "ListMemberSchemas", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListMemberSchemasSender(req)
	if err != nil {
		result.sfsplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "ListMemberSchemas", resp, "Failure sending request")
		return
	}

	result.sfsplr, err = client.ListMemberSchemasResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "ListMemberSchemas", resp, "Failure responding to request")
		return
	}
	if result.sfsplr.hasNextLink() && result.sfsplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListMemberSchemasPreparer prepares the ListMemberSchemas request.
func (client SyncMembersClient) ListMemberSchemasPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":     autorest.Encode("path", syncGroupName),
		"syncMemberName":    autorest.Encode("path", syncMemberName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}/schemas", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListMemberSchemasSender sends the ListMemberSchemas request. The method will close the
// http.Response Body if it receives an error.
func (client SyncMembersClient) ListMemberSchemasSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListMemberSchemasResponder handles the response to the ListMemberSchemas request. The method always
// closes the http.Response Body.
func (client SyncMembersClient) ListMemberSchemasResponder(resp *http.Response) (result SyncFullSchemaPropertiesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listMemberSchemasNextResults retrieves the next set of results, if any.
func (client SyncMembersClient) listMemberSchemasNextResults(ctx context.Context, lastResults SyncFullSchemaPropertiesListResult) (result SyncFullSchemaPropertiesListResult, err error) {
	req, err := lastResults.syncFullSchemaPropertiesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.SyncMembersClient", "listMemberSchemasNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListMemberSchemasSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.SyncMembersClient", "listMemberSchemasNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListMemberSchemasResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "listMemberSchemasNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListMemberSchemasComplete enumerates all values, automatically crossing page boundaries as required.
func (client SyncMembersClient) ListMemberSchemasComplete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (result SyncFullSchemaPropertiesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SyncMembersClient.ListMemberSchemas")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListMemberSchemas(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName)
	return
}

// RefreshMemberSchema refreshes a sync member database schema.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database on which the sync group is hosted.
// syncGroupName - the name of the sync group on which the sync member is hosted.
// syncMemberName - the name of the sync member.
func (client SyncMembersClient) RefreshMemberSchema(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (result SyncMembersRefreshMemberSchemaFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SyncMembersClient.RefreshMemberSchema")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RefreshMemberSchemaPreparer(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "RefreshMemberSchema", nil, "Failure preparing request")
		return
	}

	result, err = client.RefreshMemberSchemaSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "RefreshMemberSchema", result.Response(), "Failure sending request")
		return
	}

	return
}

// RefreshMemberSchemaPreparer prepares the RefreshMemberSchema request.
func (client SyncMembersClient) RefreshMemberSchemaPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":     autorest.Encode("path", syncGroupName),
		"syncMemberName":    autorest.Encode("path", syncMemberName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}/refreshSchema", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RefreshMemberSchemaSender sends the RefreshMemberSchema request. The method will close the
// http.Response Body if it receives an error.
func (client SyncMembersClient) RefreshMemberSchemaSender(req *http.Request) (future SyncMembersRefreshMemberSchemaFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// RefreshMemberSchemaResponder handles the response to the RefreshMemberSchema request. The method always
// closes the http.Response Body.
func (client SyncMembersClient) RefreshMemberSchemaResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates an existing sync member.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database on which the sync group is hosted.
// syncGroupName - the name of the sync group on which the sync member is hosted.
// syncMemberName - the name of the sync member.
// parameters - the requested sync member resource state.
func (client SyncMembersClient) Update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters SyncMember) (result SyncMembersUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SyncMembersClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SyncMembersClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters SyncMember) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":     autorest.Encode("path", syncGroupName),
		"syncMemberName":    autorest.Encode("path", syncMemberName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SyncMembersClient) UpdateSender(req *http.Request) (future SyncMembersUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SyncMembersClient) UpdateResponder(resp *http.Response) (result SyncMember, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
