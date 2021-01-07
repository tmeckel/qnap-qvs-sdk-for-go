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

// PCIsClient is the VM Client
type PCIsClient struct {
	BaseClient
}

// NewPCIsClient creates an instance of the PCIsClient client.
func NewPCIsClient() PCIsClient {
	return NewPCIsClientWithBaseURI(DefaultBaseURI)
}

// NewPCIsClientWithBaseURI creates an instance of the PCIsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPCIsClientWithBaseURI(baseURI string) PCIsClient {
	return PCIsClient{NewWithBaseURI(baseURI)}
}

// CreateGroup attach a group of PCI devices to VM
func (client PCIsClient) CreateGroup(ctx context.Context, vmid int32, data *PCIGroupSerializer) (result PCIGroupCreateResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PCIsClient.CreateGroup")
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
				Chain: []validation.Constraint{{Target: "data.ID", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("vms.PCIsClient", "CreateGroup", err.Error())
	}

	req, err := client.CreateGroupPreparer(ctx, vmid, data)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.PCIsClient", "CreateGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.PCIsClient", "CreateGroup", resp, "Failure sending request")
		return
	}

	result, err = client.CreateGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.PCIsClient", "CreateGroup", resp, "Failure responding to request")
		return
	}

	return
}

// CreateGroupPreparer prepares the CreateGroup request.
func (client PCIsClient) CreateGroupPreparer(ctx context.Context, vmid int32, data *PCIGroupSerializer) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/pcigroup", pathParameters))
	if data != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(data))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateGroupSender sends the CreateGroup request. The method will close the
// http.Response Body if it receives an error.
func (client PCIsClient) CreateGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateGroupResponder handles the response to the CreateGroup request. The method always
// closes the http.Response Body.
func (client PCIsClient) CreateGroupResponder(resp *http.Response) (result PCIGroupCreateResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DestroyGroup detach a group of PCI devices from VM
func (client PCIsClient) DestroyGroup(ctx context.Context, vmid int32, ID int32) (result PCIGroupDestroyResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PCIsClient.DestroyGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DestroyGroupPreparer(ctx, vmid, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.PCIsClient", "DestroyGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.DestroyGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.PCIsClient", "DestroyGroup", resp, "Failure sending request")
		return
	}

	result, err = client.DestroyGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.PCIsClient", "DestroyGroup", resp, "Failure responding to request")
		return
	}

	return
}

// DestroyGroupPreparer prepares the DestroyGroup request.
func (client PCIsClient) DestroyGroupPreparer(ctx context.Context, vmid int32, ID int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":    autorest.Encode("path", ID),
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/pcigroup/{id}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DestroyGroupSender sends the DestroyGroup request. The method will close the
// http.Response Body if it receives an error.
func (client PCIsClient) DestroyGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DestroyGroupResponder handles the response to the DestroyGroup request. The method always
// closes the http.Response Body.
func (client PCIsClient) DestroyGroupResponder(resp *http.Response) (result PCIGroupDestroyResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list PCI devices of VM
func (client PCIsClient) List(ctx context.Context, vmid int32) (result PCIListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PCIsClient.List")
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
		err = autorest.NewErrorWithError(err, "vms.PCIsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.PCIsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.PCIsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client PCIsClient) ListPreparer(ctx context.Context, vmid int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/pci", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PCIsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PCIsClient) ListResponder(resp *http.Response) (result PCIListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Retrieve find a PCI device of VM
func (client PCIsClient) Retrieve(ctx context.Context, vmid int32, ID int32) (result PCIRetrieveResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PCIsClient.Retrieve")
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
		err = autorest.NewErrorWithError(err, "vms.PCIsClient", "Retrieve", nil, "Failure preparing request")
		return
	}

	resp, err := client.RetrieveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vms.PCIsClient", "Retrieve", resp, "Failure sending request")
		return
	}

	result, err = client.RetrieveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vms.PCIsClient", "Retrieve", resp, "Failure responding to request")
		return
	}

	return
}

// RetrievePreparer prepares the Retrieve request.
func (client PCIsClient) RetrievePreparer(ctx context.Context, vmid int32, ID int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":    autorest.Encode("path", ID),
		"vm_id": autorest.Encode("path", vmid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vms/{vm_id}/pci/{id}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RetrieveSender sends the Retrieve request. The method will close the
// http.Response Body if it receives an error.
func (client PCIsClient) RetrieveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RetrieveResponder handles the response to the Retrieve request. The method always
// closes the http.Response Body.
func (client PCIsClient) RetrieveResponder(resp *http.Response) (result PCIRetrieveResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}