package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	matchapi "mastercard-match/app/services/mastercard/match"
	"net/http"
	"os"
	"sync"

	"github.com/mastercard/oauth1-signer-go/interceptor"
	"github.com/revel/revel"
)

var once sync.Once
var api_client *matchapi.APIClient

type RenderByte []byte

func (r RenderByte) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "application/json")
	resp.GetWriter().Write([]byte(r))
}

type ApiMatchV1 struct {
	*revel.Controller
}

func (c ApiMatchV1) PostRetroInquiryDetails() revel.Result {
	format := "JSON"                                                      // string | Describes format of the response you wants to serverd, response can be delevired either as XML or JSON.
	acquirerId := "1996"                                                  // string | The Member ICA number which is used to validate that the application has permission to hit the MATCH database. This number must be obtained from a participating MATCH acquiring bank or entity before a developer can access the MATCH resource.
	retroInquiryRequestSchema := *matchapi.NewRetroInquiryRequestSchema() // RetroInquiryRequestSchema | The retro inquiry request object
	retroInquiryRequest := matchapi.NewRetroInquiryRequest()

	c.Params.BindJSON(&retroInquiryRequest)

	retroInquiryRequestSchema.SetRetroInquiryRequest(*retroInquiryRequest)
	api_client := getMatchApiClient()
	resp, r, err := api_client.RetroactiveInquiryDetailsRequestApi.RetroRetroInquiryDetailsPost(context.Background()).Format(format).AcquirerId(acquirerId).RetroInquiryRequestSchema(retroInquiryRequestSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetroactiveInquiryDetailsRequestApi.RetroRetroInquiryDetailsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		errors := new(matchapi.ErrorsResponse)
		err = json.NewDecoder(r.Body).Decode(errors)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error en marshaljson: %v\n", r)
		}
		return c.RenderJSON(errors)
	}
	// response from `RetroRetroInquiryDetailsPost`: RetroInquiryResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `RetroactiveInquiryDetailsRequestApi.RetroRetroInquiryDetailsPost`: %v\n", resp)

	response, err := resp.MarshalJSON()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error en marshaljson: %v", err)
	}

	return RenderByte(response)
}

func (c ApiMatchV1) PostRetroList() revel.Result {
	format := "JSON"                                        // string | Describes format of the response you wants to serverd, response can be delevired either as XML or JSON.
	retroRequestSchema := *matchapi.NewRetroRequestSchema() // RetroRequestSchema | The retro request object
	retroRequest := matchapi.NewRetroRequest()

	c.Params.BindJSON(&retroRequest)

	retroRequestSchema.SetRetroRequest(*retroRequest)
	api_client := getMatchApiClient()
	resp, r, err := api_client.RetroactiveInquiryRequestApi.RetroRetroListPost(context.Background()).Format(format).RetroRequestSchema(retroRequestSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetroactiveInquiryRequestApi.RetroRetroListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		errors := new(matchapi.ErrorsResponse)
		err = json.NewDecoder(r.Body).Decode(errors)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error en marshaljson: %v\n", r)
		}
		return c.RenderJSON(errors)
	}
	// response from `RetroRetroListPost`: RetroResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `RetroactiveInquiryRequestApi.RetroRetroListPost`: %v\n", resp)

	response, err := resp.MarshalJSON()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error en marshaljson: %v", err)
	}

	return RenderByte(response)
}

func (c ApiMatchV1) PostCommonContactDetails() revel.Result {
	format := "JSON"                                            // string | Describes format of the response you wants to serverd, response can be delevired either as XML or JSON.
	contactRequestSchema := *matchapi.NewContactRequestSchema() // ContactRequestSchema | The contact request object
	contactRequest := matchapi.NewContactRequest("12086")

	c.Params.BindJSON(&contactRequest)

	contactRequestSchema.SetContactRequest(*contactRequest)
	api_client := getMatchApiClient()
	resp, r, err := api_client.AcquirerContactRequestApi.CommonContactDetailsPost(context.Background()).Format(format).ContactRequestSchema(contactRequestSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcquirerContactRequestApi.CommonContactDetailsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		errors := new(matchapi.ErrorsResponse)
		err = json.NewDecoder(r.Body).Decode(errors)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error en marshaljson: %v\n", r)
		}
		return c.RenderJSON(errors)
	}
	// response from `CommonContactDetailsPost`: ContactResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `AcquirerContactRequestApi.CommonContactDetailsPost`: %v\n", resp)

	response, err := resp.MarshalJSON()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error en marshaljson: %v", err)
	}

	return RenderByte(response)
}

func (c ApiMatchV1) PostAddMerchant() revel.Result {
	format := "JSON"                                                                        // string | Describes format of the response you wants to serverd, response can be delevired either as XML or JSON.
	addTerminatedMerchantRequestSchema := *matchapi.NewAddTerminatedMerchantRequestSchema() // AddTerminatedMerchantRequestSchema | Body of the Add Terminated Merchant Request
	addTerminatedMerchantRequest := matchapi.NewAddMerchantRequestWithDefaults()

	c.Params.BindJSON(&addTerminatedMerchantRequest)
	fmt.Printf("body: %v\n\n\n", addTerminatedMerchantRequest)

	addTerminatedMerchantRequestSchema.SetAddMerchantRequest(*addTerminatedMerchantRequest)

	api_client := getMatchApiClient()
	resp, r, err := api_client.AddTerminatedMerchantApi.AddMerchantPost(context.Background()).Format(format).AddTerminatedMerchantRequestSchema(addTerminatedMerchantRequestSchema).Execute()
	defer r.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddTerminatedMerchantApi.AddMerchantPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		errors := new(matchapi.ErrorsResponse)
		err = json.NewDecoder(r.Body).Decode(errors)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error en marshaljson: %v\n", r)
		}
		return c.RenderJSON(errors)
	}
	// response from `AddMerchantPost`: AddTerminatedMerchantResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `AddTerminatedMerchantApi.AddMerchantPost`: %v\n", resp)

	response, err := resp.MarshalJSON()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error en marshaljson: %v", err)
	}

	return RenderByte(response)
}

func (c ApiMatchV1) PostTerminationInquiry() revel.Result {
	pageOffset := int32(0)                                                            // int32 | The zero-based offset to start at. The actual start position is this value +1. An offset of 10 starts at item 11. Combined with the PageLength option this allows pagination to be supported through the service requests.
	pageLength := int32(10)                                                           // int32 | The maximum number of items to retrieve within the current \"page\" of results.
	format := "JSON"                                                                  // string | Describes format of the response you wants to serverd, response can be delevired either as XML or JSON.
	terminationInquiryRequestSchema := *matchapi.NewTerminationInquiryRequestSchema() // TerminationInquiryRequestSchema | Body of the Termination Inquiry Request
	terminationInquiryRequest := matchapi.NewTerminationInquiryRequest("1996")

	c.Params.BindJSON(&terminationInquiryRequest)

	terminationInquiryRequestSchema.SetTerminationInquiryRequest(*terminationInquiryRequest)

	api_client := getMatchApiClient()
	resp, r, err := api_client.TerminationInquiryRequestApi.TerminationInquiryPost(context.Background()).PageOffset(pageOffset).PageLength(pageLength).Format(format).TerminationInquiryRequestSchema(terminationInquiryRequestSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerminationInquiryRequestApi.TerminationInquiryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		errors := new(matchapi.ErrorsResponse)
		err = json.NewDecoder(r.Body).Decode(errors)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error en marshaljson: %v\n", r)
		}
		return c.RenderJSON(errors)
	}
	// response from `TerminationInquiryPost`: TerminationInquirySchema
	fmt.Fprintf(os.Stdout, "Response from `TerminationInquiryRequestApi.TerminationInquiryPost`: %v\n\n\n", resp.TerminationInquiry.GetRef())
	// fmt.Fprintf(os.Stdout, "Response from `TerminationInquiryRequestApi.TerminationInquiryPost`: %v\n\n\n", r.Body)

	defer r.Body.Close()

	response, err := resp.MarshalJSON()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error en marshaljson: %v", err)
	}

	return RenderByte(response)
	// return c.RenderJSON(response)
}

func getMatchApiClient() *matchapi.APIClient {
	once.Do(func() {
		configuration := matchapi.NewConfiguration()
		configuration.Host = revel.Config.StringDefault("mastercard.API_HOST", "sandbox.api.mastercard.com")
		consumerKey := revel.Config.StringDefault("mastercard.CONSUMER_KEY", "")
		KeyPath := revel.Config.StringDefault("mastercard.SIGNING_KEY_PATH", "")
		keyPass := revel.Config.StringDefault("mastercard.SIGNING_KEY_SECRET", "")
		httpClient, _ := interceptor.GetHttpClient(consumerKey, KeyPath, keyPass)
		configuration.HTTPClient = httpClient
		api_client = matchapi.NewAPIClient(configuration)
	})

	return api_client
}
