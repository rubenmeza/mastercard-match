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

// TerminationInquiryRequestApiService TerminationInquiryRequestApi service
type TerminationInquiryRequestApiService service

type ApiTerminationInquiryPostRequest struct {
	ctx _context.Context
	ApiService *TerminationInquiryRequestApiService
	pageOffset *int32
	pageLength *int32
	format *string
	terminationInquiryRequestSchema *TerminationInquiryRequestSchema
}

func (r ApiTerminationInquiryPostRequest) PageOffset(pageOffset int32) ApiTerminationInquiryPostRequest {
	r.pageOffset = &pageOffset
	return r
}
func (r ApiTerminationInquiryPostRequest) PageLength(pageLength int32) ApiTerminationInquiryPostRequest {
	r.pageLength = &pageLength
	return r
}
func (r ApiTerminationInquiryPostRequest) Format(format string) ApiTerminationInquiryPostRequest {
	r.format = &format
	return r
}
func (r ApiTerminationInquiryPostRequest) TerminationInquiryRequestSchema(terminationInquiryRequestSchema TerminationInquiryRequestSchema) ApiTerminationInquiryPostRequest {
	r.terminationInquiryRequestSchema = &terminationInquiryRequestSchema
	return r
}

func (r ApiTerminationInquiryPostRequest) Execute() (TerminationInquirySchema, *_nethttp.Response, error) {
	return r.ApiService.TerminationInquiryPostExecute(r)
}

/*
 * TerminationInquiryPost ##### Retrieves terminated merchant information based on the criteria defined in the API request.
 * Returns information on merchants that have been terminated and merchants which have been inquired upon. MATCH can provide the acquiring bank with information if a Merchant has been terminated by another acquirer already, the reason for termination, the history of fraudulent or risky business practices that led to that termination and the inquiry that was made already on the Merchant or individual by own or another acquiring bank. The response also includes an inquiry reference number to be used for future reference in Termination Inquiry History Request API. For further details refer the documentation on [Use Cases.](/match/documentation/use-cases)

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiTerminationInquiryPostRequest
 */
func (a *TerminationInquiryRequestApiService) TerminationInquiryPost(ctx _context.Context) ApiTerminationInquiryPostRequest {
	return ApiTerminationInquiryPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return TerminationInquirySchema
 */
func (a *TerminationInquiryRequestApiService) TerminationInquiryPostExecute(r ApiTerminationInquiryPostRequest) (TerminationInquirySchema, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TerminationInquirySchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TerminationInquiryRequestApiService.TerminationInquiryPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/termination-inquiry"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pageOffset == nil {
		return localVarReturnValue, nil, reportError("pageOffset is required and must be specified")
	}
	if r.pageLength == nil {
		return localVarReturnValue, nil, reportError("pageLength is required and must be specified")
	}
	if r.format == nil {
		return localVarReturnValue, nil, reportError("format is required and must be specified")
	}
	if r.terminationInquiryRequestSchema == nil {
		return localVarReturnValue, nil, reportError("terminationInquiryRequestSchema is required and must be specified")
	}

	localVarQueryParams.Add("PageOffset", parameterToString(*r.pageOffset, ""))
	localVarQueryParams.Add("PageLength", parameterToString(*r.pageLength, ""))
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
	localVarPostBody = r.terminationInquiryRequestSchema
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