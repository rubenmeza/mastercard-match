# SearchCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchAll** | **string** | Determines if the inquiry is worldwide or not.  The Inquiry request can be either worldwide or Regions and/or Countries based. It cannot be both.  If it is not worldwide search (SearchAll &#x3D; N) and if both  Region and Country are not specified, then search will happen for USA | 
**Region** | Pointer to **[]string** |  | [optional] 
**Country** | Pointer to **[]string** |  | [optional] 
**MinPossibleMatchCount** | Pointer to **string** | Determines how many minimum matches present for a merchant or inquiry to appear in the results. | [optional] 

## Methods

### NewSearchCriteria

`func NewSearchCriteria(searchAll string, ) *SearchCriteria`

NewSearchCriteria instantiates a new SearchCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchCriteriaWithDefaults

`func NewSearchCriteriaWithDefaults() *SearchCriteria`

NewSearchCriteriaWithDefaults instantiates a new SearchCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchAll

`func (o *SearchCriteria) GetSearchAll() string`

GetSearchAll returns the SearchAll field if non-nil, zero value otherwise.

### GetSearchAllOk

`func (o *SearchCriteria) GetSearchAllOk() (*string, bool)`

GetSearchAllOk returns a tuple with the SearchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAll

`func (o *SearchCriteria) SetSearchAll(v string)`

SetSearchAll sets SearchAll field to given value.


### GetRegion

`func (o *SearchCriteria) GetRegion() []string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SearchCriteria) GetRegionOk() (*[]string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SearchCriteria) SetRegion(v []string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SearchCriteria) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCountry

`func (o *SearchCriteria) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SearchCriteria) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SearchCriteria) SetCountry(v []string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SearchCriteria) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetMinPossibleMatchCount

`func (o *SearchCriteria) GetMinPossibleMatchCount() string`

GetMinPossibleMatchCount returns the MinPossibleMatchCount field if non-nil, zero value otherwise.

### GetMinPossibleMatchCountOk

`func (o *SearchCriteria) GetMinPossibleMatchCountOk() (*string, bool)`

GetMinPossibleMatchCountOk returns a tuple with the MinPossibleMatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPossibleMatchCount

`func (o *SearchCriteria) SetMinPossibleMatchCount(v string)`

SetMinPossibleMatchCount sets MinPossibleMatchCount field to given value.

### HasMinPossibleMatchCount

`func (o *SearchCriteria) HasMinPossibleMatchCount() bool`

HasMinPossibleMatchCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


