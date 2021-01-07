package vms

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AdaptersClient is the VM Client
type AdaptersClient struct {
	BaseClient
}

// NewAdaptersClient creates an instance of the AdaptersClient client.
func NewAdaptersClient() AdaptersClient {
	return NewAdaptersClientWithBaseURI(DefaultBaseURI)
}

// NewAdaptersClientWithBaseURI creates an instance of the AdaptersClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAdaptersClientWithBaseURI(baseURI string) AdaptersClient {
	return AdaptersClient{NewWithBaseURI(baseURI)}
}

// Create create a new Adapter
func (client AdaptersClient) Create(ctx context.Context, vmid int32, data *WriteAdapterSerializer) (result AdapterCreateResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdaptersClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: data,
			Constraints: []validation.Constraint{{Target: "data", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "data.Bridge", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("vms.AdaptersClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, vmid, data)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client AdaptersClient) CreatePreparer(ctx context.Context, vmid int32, data *WriteAdapterSerializer) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/adapters", pathParameters))
	if data != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(data))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client AdaptersClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client AdaptersClient) CreateResponder(resp *http.Response) (result AdapterCreateResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Destroy delete a Adapter
func (client AdaptersClient) Destroy(ctx context.Context, vmid int32, ID int32) (result AdapterDestroyResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdaptersClient.Destroy")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DestroyPreparer(ctx, vmid, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "Destroy", nil, "Failure preparing request")
		return
	}

	resp, err := client.DestroySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "Destroy", resp, "Failure sending request")
		return
	}

	result, err = client.DestroyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "Destroy", resp, "Failure responding to request")
		return
	}

	return
}

// DestroyPreparer prepares the Destroy request.
func (client AdaptersClient) DestroyPreparer(ctx context.Context, vmid int32, ID int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":    autorest.Encode("path", ID),
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/adapters/{id}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DestroySender sends the Destroy request. The method will close the
// http.Response Body if it receives an error.
func (client AdaptersClient) DestroySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DestroyResponder handles the response to the Destroy request. The method always
// closes the http.Response Body.
func (client AdaptersClient) DestroyResponder(resp *http.Response) (result AdapterDestroyResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list Adapters of vm
func (client AdaptersClient) List(ctx context.Context, vmid int32) (result AdapterListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdaptersClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, vmid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AdaptersClient) ListPreparer(ctx context.Context, vmid int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/adapters", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AdaptersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AdaptersClient) ListResponder(resp *http.Response) (result AdapterListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListIPs get Adapter IPs by qemu-img agent
func (client AdaptersClient) ListIPs(ctx context.Context, vmid int32, ID int32) (result AdapterIpsResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdaptersClient.ListIPs")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListIPsPreparer(ctx, vmid, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "ListIPs", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListIPsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "ListIPs", resp, "Failure sending request")
		return
	}

	result, err = client.ListIPsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "ListIPs", resp, "Failure responding to request")
		return
	}

	return
}

// ListIPsPreparer prepares the ListIPs request.
func (client AdaptersClient) ListIPsPreparer(ctx context.Context, vmid int32, ID int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":    autorest.Encode("path", ID),
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/adapters/{id}/ips", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListIPsSender sends the ListIPs request. The method will close the
// http.Response Body if it receives an error.
func (client AdaptersClient) ListIPsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListIPsResponder handles the response to the ListIPs request. The method always
// closes the http.Response Body.
func (client AdaptersClient) ListIPsResponder(resp *http.Response) (result AdapterIpsResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Retrieve find a Adapter
func (client AdaptersClient) Retrieve(ctx context.Context, vmid int32, ID int32) (result AdapterRetrieveResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdaptersClient.Retrieve")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RetrievePreparer(ctx, vmid, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "Retrieve", nil, "Failure preparing request")
		return
	}

	resp, err := client.RetrieveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "Retrieve", resp, "Failure sending request")
		return
	}

	result, err = client.RetrieveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "Retrieve", resp, "Failure responding to request")
		return
	}

	return
}

// RetrievePreparer prepares the Retrieve request.
func (client AdaptersClient) RetrievePreparer(ctx context.Context, vmid int32, ID int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":    autorest.Encode("path", ID),
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/adapters/{id}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RetrieveSender sends the Retrieve request. The method will close the
// http.Response Body if it receives an error.
func (client AdaptersClient) RetrieveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RetrieveResponder handles the response to the Retrieve request. The method always
// closes the http.Response Body.
func (client AdaptersClient) RetrieveResponder(resp *http.Response) (result AdapterRetrieveResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update modify Adapter setting
func (client AdaptersClient) Update(ctx context.Context, vmid int32, ID int32, body *WriteAdapterSerializer) (result AdapterPartialUpdateResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdaptersClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, vmid, ID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.AdaptersClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client AdaptersClient) UpdatePreparer(ctx context.Context, vmid int32, ID int32, body *WriteAdapterSerializer) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":    autorest.Encode("path", ID),
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/adapters/{id}", pathParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AdaptersClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AdaptersClient) UpdateResponder(resp *http.Response) (result AdapterPartialUpdateResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}