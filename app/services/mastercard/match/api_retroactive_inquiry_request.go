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

// RetroactiveInquiryRequestApiService RetroactiveInquiryRequestApi service
type RetroactiveInquiryRequestApiService service

type ApiRetroRetroListPostRequest struct {
	ctx _context.Context
	ApiService *RetroactiveInquiryRequestApiService
	format *string
	retroRequestSchema *RetroRequestSchema
}

func (r ApiRetroRetroListPostRequest) Format(format string) ApiRetroRetroListPostRequest {
	r.format = &format
	return r
}
func (r ApiRetroRetroListPostRequest) RetroRequestSchema(retroRequestSchema RetroRequestSchema) ApiRetroRetroListPostRequest {
	r.retroRequestSchema = &retroRequestSchema
	return r
}

func (r ApiRetroRetroListPostRequest) Execute() (RetroResponseSchema, *_nethttp.Response, error) {
	return r.ApiService.RetroRetroListPostExecute(r)
}

/*
 * RetroRetroListPost ##### The retroactive inquiry helps acquirer to retrieve list of termination inquiry matches made within 360 days of inquiry initiation.
 * Return the summary of Retroactive Inquiry matches for the given Acquirer. Even if initial inquiry (TerminationInquiryRequest) does not result in a possible match, there after it's still possible that the merchant listing could appear sometime in the next 360 days. This may occur when another/same acquirer enters the merchant into the MATCH database. The system will automatically continue to search the MATCH system for 360 days. To view these notifications, acquirers must use the Retroactive Inquiry service. For further details refer the documentation on [Use Cases.](/match/documentation/use-cases)
<br><br>
#### Note:
(a)  Each day, MATCH automatically conducts a retroactive Inquiry search on every merchant on which the acquirer has submitted an inquiry. MATCH automatically deletes each retroactive inquiry match after seven days. Therefore, acquirers should view retroactive inquiries every day to be sure that they do not miss a possible match.
<br>
(b)  When there are no results for RetroActive inquiry, users will get in response HTTP status code of 400 and a response code of “RESPONSE_DATA_NORESULT”. For information around all error codes refer [Code and Formats](/match/documentation/code-and-formats) documentation page
<br><br>
To understand the working of RetroActive API use case, follow the steps mentioned below:
<br>
1.  Acquirer A (AcquirerId = 1996) uses a TerminationInquiryRequest for a merchant “X Inc”  but doesn’t find a match.
<br>
2.  Within the next 360 days, same/another Acquirer (use same AcquirerId = 1996, as we have only one acquirer in sandbox) adds a terminated merchant “X Inc”, who matches the search parameters used by Acquirer A. 
<br>
3.  After adding the merchant, it takes one day for Match Service to process retroactive inquiry match.
<br>
4.  Post processing, within the next 7 days, Acquirer A should use the Retroactive Inquiry API to retrieve a list of retroactive matches. 
<br>
5.  This list of retroactive inquiry matches includes an Inquiry reference number, which the Acquirer A can use to get more details of match using RetroActive Inquiry Details API. 

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiRetroRetroListPostRequest
 */
func (a *RetroactiveInquiryRequestApiService) RetroRetroListPost(ctx _context.Context) ApiRetroRetroListPostRequest {
	return ApiRetroRetroListPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return RetroResponseSchema
 */
func (a *RetroactiveInquiryRequestApiService) RetroRetroListPostExecute(r ApiRetroRetroListPostRequest) (RetroResponseSchema, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RetroResponseSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RetroactiveInquiryRequestApiService.RetroRetroListPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/retro/retro-list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.format == nil {
		return localVarReturnValue, nil, reportError("format is required and must be specified")
	}
	if r.retroRequestSchema == nil {
		return localVarReturnValue, nil, reportError("retroRequestSchema is required and must be specified")
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
	localVarPostBody = r.retroRequestSchema
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
