# UrlGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExactMatchUrls** | Pointer to [**Url**](Url.md) |  | [optional] 
**CloseMatchUrls** | Pointer to [**Url**](Url.md) |  | [optional] 
**NoMatchUrls** | Pointer to [**Url**](Url.md) |  | [optional] 

## Methods

### NewUrlGroup

`func NewUrlGroup() *UrlGroup`

NewUrlGroup instantiates a new UrlGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlGroupWithDefaults

`func NewUrlGroupWithDefaults() *UrlGroup`

NewUrlGroupWithDefaults instantiates a new UrlGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExactMatchUrls

`func (o *UrlGroup) GetExactMatchUrls() Url`

GetExactMatchUrls returns the ExactMatchUrls field if non-nil, zero value otherwise.

### GetExactMatchUrlsOk

`func (o *UrlGroup) GetExactMatchUrlsOk() (*Url, bool)`

GetExactMatchUrlsOk returns a tuple with the ExactMatchUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactMatchUrls

`func (o *UrlGroup) SetExactMatchUrls(v Url)`

SetExactMatchUrls sets ExactMatchUrls field to given value.

### HasExactMatchUrls

`func (o *UrlGroup) HasExactMatchUrls() bool`

HasExactMatchUrls returns a boolean if a field has been set.

### GetCloseMatchUrls

`func (o *UrlGroup) GetCloseMatchUrls() Url`

GetCloseMatchUrls returns the CloseMatchUrls field if non-nil, zero value otherwise.

### GetCloseMatchUrlsOk

`func (o *UrlGroup) GetCloseMatchUrlsOk() (*Url, bool)`

GetCloseMatchUrlsOk returns a tuple with the CloseMatchUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseMatchUrls

`func (o *UrlGroup) SetCloseMatchUrls(v Url)`

SetCloseMatchUrls sets CloseMatchUrls field to given value.

### HasCloseMatchUrls

`func (o *UrlGroup) HasCloseMatchUrls() bool`

HasCloseMatchUrls returns a boolean if a field has been set.

### GetNoMatchUrls

`func (o *UrlGroup) GetNoMatchUrls() Url`

GetNoMatchUrls returns the NoMatchUrls field if non-nil, zero value otherwise.

### GetNoMatchUrlsOk

`func (o *UrlGroup) GetNoMatchUrlsOk() (*Url, bool)`

GetNoMatchUrlsOk returns a tuple with the NoMatchUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoMatchUrls

`func (o *UrlGroup) SetNoMatchUrls(v Url)`

SetNoMatchUrls sets NoMatchUrls field to given value.

### HasNoMatchUrls

`func (o *UrlGroup) HasNoMatchUrls() bool`

HasNoMatchUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


