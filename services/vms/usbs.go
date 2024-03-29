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

// USBsClient is the VM Client
type USBsClient struct {
	BaseClient
}

// NewUSBsClient creates an instance of the USBsClient client.
func NewUSBsClient() USBsClient {
	return NewUSBsClientWithBaseURI(DefaultBaseURI)
}

// NewUSBsClientWithBaseURI creates an instance of the USBsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewUSBsClientWithBaseURI(baseURI string) USBsClient {
	return USBsClient{NewWithBaseURI(baseURI)}
}

// Create attach USB to VM
func (client USBsClient) Create(ctx context.Context, vmid int32, data *WriteUSBSerializer) (result USBCreateResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/USBsClient.Create")
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
				Chain: []validation.Constraint{{Target: "data.BusNum", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "data.DevNum", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "data.ProductID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "data.VendorID", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("vms.USBsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, vmid, data)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client USBsClient) CreatePreparer(ctx context.Context, vmid int32, data *WriteUSBSerializer) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/usbs", pathParameters))
	if data != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(data))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client USBsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client USBsClient) CreateResponder(resp *http.Response) (result USBCreateResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Destroy detach USB to VM
func (client USBsClient) Destroy(ctx context.Context, vmid int32, ID int32) (result USBDestroyResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/USBsClient.Destroy")
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
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "Destroy", nil, "Failure preparing request")
		return
	}

	resp, err := client.DestroySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "Destroy", resp, "Failure sending request")
		return
	}

	result, err = client.DestroyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "Destroy", resp, "Failure responding to request")
		return
	}

	return
}

// DestroyPreparer prepares the Destroy request.
func (client USBsClient) DestroyPreparer(ctx context.Context, vmid int32, ID int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":    autorest.Encode("path", ID),
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/usbs/{id}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DestroySender sends the Destroy request. The method will close the
// http.Response Body if it receives an error.
func (client USBsClient) DestroySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DestroyResponder handles the response to the Destroy request. The method always
// closes the http.Response Body.
func (client USBsClient) DestroyResponder(resp *http.Response) (result USBDestroyResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list USBs of vm
func (client USBsClient) List(ctx context.Context, vmid int32) (result USBListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/USBsClient.List")
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
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client USBsClient) ListPreparer(ctx context.Context, vmid int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/usbs", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client USBsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client USBsClient) ListResponder(resp *http.Response) (result USBListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Retrieve find a USB
func (client USBsClient) Retrieve(ctx context.Context, vmid int32, ID int32) (result USBRetrieveResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/USBsClient.Retrieve")
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
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "Retrieve", nil, "Failure preparing request")
		return
	}

	resp, err := client.RetrieveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "Retrieve", resp, "Failure sending request")
		return
	}

	result, err = client.RetrieveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "Retrieve", resp, "Failure responding to request")
		return
	}

	return
}

// RetrievePreparer prepares the Retrieve request.
func (client USBsClient) RetrievePreparer(ctx context.Context, vmid int32, ID int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":    autorest.Encode("path", ID),
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/usbs/{id}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RetrieveSender sends the Retrieve request. The method will close the
// http.Response Body if it receives an error.
func (client USBsClient) RetrieveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RetrieveResponder handles the response to the Retrieve request. The method always
// closes the http.Response Body.
func (client USBsClient) RetrieveResponder(resp *http.Response) (result USBRetrieveResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update modify USB setting
func (client USBsClient) Update(ctx context.Context, vmid int32, ID int32, body *WriteUSBSerializer) (result USBPartialUpdateResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/USBsClient.Update")
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
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.USBsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client USBsClient) UpdatePreparer(ctx context.Context, vmid int32, ID int32, body *WriteUSBSerializer) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":    autorest.Encode("path", ID),
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/usbs/{id}", pathParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client USBsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client USBsClient) UpdateResponder(resp *http.Response) (result USBPartialUpdateResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
