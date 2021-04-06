# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | This is the unique identifier that attempts to define the field in error when available.  If a specific field cant be identified System will be returned. | [optional] 
**ReasonCode** | Pointer to **string** | This will identify the reason for the error. | [optional] 
**Description** | Pointer to **string** | This is the text description of the error. This is optional and will only be displayed if more information is available than is stored in the data identifier and reason code. | [optional] 
**Recoverable** | Pointer to **bool** | This is a true/false presentation to explain if the transaction was submitted again would it be successful or not. | [optional] 
**Details** | Pointer to **string** | Details description of the issue | [optional] 

## Methods

### NewError

`func NewError() *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *Error) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Error) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Error) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Error) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetReasonCode

`func (o *Error) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *Error) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *Error) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *Error) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetDescription

`func (o *Error) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Error) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Error) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Error) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRecoverable

`func (o *Error) GetRecoverable() bool`

GetRecoverable returns the Recoverable field if non-nil, zero value otherwise.

### GetRecoverableOk

`func (o *Error) GetRecoverableOk() (*bool, bool)`

GetRecoverableOk returns a tuple with the Recoverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverable

`func (o *Error) SetRecoverable(v bool)`

SetRecoverable sets Recoverable field to given value.

### HasRecoverable

`func (o *Error) HasRecoverable() bool`

HasRecoverable returns a boolean if a field has been set.

### GetDetails

`func (o *Error) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Error) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Error) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Error) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


