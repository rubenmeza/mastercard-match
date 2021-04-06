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

// RetroactiveInquiryDetailsRequestApiService RetroactiveInquiryDetailsRequestApi service
type RetroactiveInquiryDetailsRequestApiService service

type ApiRetroRetroInquiryDetailsPostRequest struct {
	ctx _context.Context
	ApiService *RetroactiveInquiryDetailsRequestApiService
	format *string
	acquirerId *string
	retroInquiryRequestSchema *RetroInquiryRequestSchema
}

func (r ApiRetroRetroInquiryDetailsPostRequest) Format(format string) ApiRetroRetroInquiryDetailsPostRequest {
	r.format = &format
	return r
}
func (r ApiRetroRetroInquiryDetailsPostRequest) AcquirerId(acquirerId string) ApiRetroRetroInquiryDetailsPostRequest {
	r.acquirerId = &acquirerId
	return r
}
func (r ApiRetroRetroInquiryDetailsPostRequest) RetroInquiryRequestSchema(retroInquiryRequestSchema RetroInquiryRequestSchema) ApiRetroRetroInquiryDetailsPostRequest {
	r.retroInquiryRequestSchema = &retroInquiryRequestSchema
	return r
}

func (r ApiRetroRetroInquiryDetailsPostRequest) Execute() (RetroInquiryResponseSchema, *_nethttp.Response, error) {
	return r.ApiService.RetroRetroInquiryDetailsPostExecute(r)
}

/*
 * RetroRetroInquiryDetailsPost ##### Retrieve retro inquiry search results
 * The acquirer uses the Retroactive Inquiry API to retrieve a list of Retroactive Inquiry Reference Numbers which matched the acquirer’s previous termination inquiries. The acquirer then uses Retroactive Inquiry Details API to fetch details of the matches for each from previous inquiries. Details about merchant match includes information, such as, if a Merchant has been terminated by another acquirer after an inquiry was made, the reason for termination, and the history of fraudulent or risky business practices that led to that termination. For further details refer the documentation on [Use Cases.](/match/documentation/use-cases)

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiRetroRetroInquiryDetailsPostRequest
 */
func (a *RetroactiveInquiryDetailsRequestApiService) RetroRetroInquiryDetailsPost(ctx _context.Context) ApiRetroRetroInquiryDetailsPostRequest {
	return ApiRetroRetroInquiryDetailsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return RetroInquiryResponseSchema
 */
func (a *RetroactiveInquiryDetailsRequestApiService) RetroRetroInquiryDetailsPostExecute(r ApiRetroRetroInquiryDetailsPostRequest) (RetroInquiryResponseSchema, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RetroInquiryResponseSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RetroactiveInquiryDetailsRequestApiService.RetroRetroInquiryDetailsPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/retro/retro-inquiry-details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.format == nil {
		return localVarReturnValue, nil, reportError("format is required and must be specified")
	}
	if r.acquirerId == nil {
		return localVarReturnValue, nil, reportError("acquirerId is required and must be specified")
	}
	if r.retroInquiryRequestSchema == nil {
		return localVarReturnValue, nil, reportError("retroInquiryRequestSchema is required and must be specified")
	}

	localVarQueryParams.Add("Format", parameterToString(*r.format, ""))
	localVarQueryParams.Add("AcquirerId", parameterToString(*r.acquirerId, ""))
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
	localVarPostBody = r.retroInquiryRequestSchema
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
