/*
 * MATCH API
 *
 * Helps acquirers identify potentially high-risk merchants before entering to a merchant agreement.
 *
 * API version: 1.0.0
 * Contact: apisupport@mastercard.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package match

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// AcquirerContactRequestApiService AcquirerContactRequestApi service
type AcquirerContactRequestApiService service

type ApiCommonContactDetailsPostRequest struct {
	ctx _context.Context
	ApiService *AcquirerContactRequestApiService
	format *string
	contactRequestSchema *ContactRequestSchema
}

func (r ApiCommonContactDetailsPostRequest) Format(format string) ApiCommonContactDetailsPostRequest {
	r.format = &format
	return r
}
func (r ApiCommonContactDetailsPostRequest) ContactRequestSchema(contactRequestSchema ContactRequestSchema) ApiCommonContactDetailsPostRequest {
	r.contactRequestSchema = &contactRequestSchema
	return r
}

func (r ApiCommonContactDetailsPostRequest) Execute() (ContactResponseSchema, *_nethttp.Response, error) {
	return r.ApiService.CommonContactDetailsPostExecute(r)
}

/*
 * CommonContactDetailsPost ##### Retrieve contact information of another acquirer.
 * Returns the contact information for the acquirer id requested. When MATCH returns a possible merchant match, this resource assists users by retrieving the contact information associated with the Acquirer ID/ICA that added the merchant to MATCH. For further details refer the documentation on [Use Cases.](/match/documentation/use-cases) 

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCommonContactDetailsPostRequest
 */
func (a *AcquirerContactRequestApiService) CommonContactDetailsPost(ctx _context.Context) ApiCommonContactDetailsPostRequest {
	return ApiCommonContactDetailsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ContactResponseSchema
 */
func (a *AcquirerContactRequestApiService) CommonContactDetailsPostExecute(r ApiCommonContactDetailsPostRequest) (ContactResponseSchema, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ContactResponseSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AcquirerContactRequestApiService.CommonContactDetailsPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/common/contact-details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.format == nil {
		return localVarReturnValue, nil, reportError("format is required and must be specified")
	}
	if r.contactRequestSchema == nil {
		return localVarReturnValue, nil, reportError("contactRequestSchema is required and must be specified")
	}

	localVarQueryParams.Add("Format", parameterToString(*r.format, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.contactRequestSchema
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorsResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
