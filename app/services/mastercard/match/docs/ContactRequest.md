# ContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquirerId** | **string** | The acquiring bank ICA number whose contact information is to be fetched. | 

## Methods

### NewContactRequest

`func NewContactRequest(acquirerId string, ) *ContactRequest`

NewContactRequest instantiates a new ContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactRequestWithDefaults

`func NewContactRequestWithDefaults() *ContactRequest`

NewContactRequestWithDefaults instantiates a new ContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcquirerId

`func (o *ContactRequest) GetAcquirerId() string`

GetAcquirerId returns the AcquirerId field if non-nil, zero value otherwise.

### GetAcquirerIdOk

`func (o *ContactRequest) GetAcquirerIdOk() (*string, bool)`

GetAcquirerIdOk returns a tuple with the AcquirerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerId

`func (o *ContactRequest) SetAcquirerId(v string)`

SetAcquirerId sets AcquirerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


