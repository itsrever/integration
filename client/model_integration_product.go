/*
Integration stubs

Stubs for implementing a REVER integration

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// IntegrationProduct A product from the e-commerce
type IntegrationProduct struct {
	Categories []IntegrationProductCategory `json:"categories,omitempty"`
	// Product description
	Description *string `json:"description,omitempty"`
	// Unique identifier for the resource in the source platform
	Id *string `json:"id,omitempty"`
	// Images associated to the product
	Images []IntegrationProductImage `json:"images,omitempty"`
	// Product name
	Name *string `json:"name,omitempty"`
	// platform from where this product comes from
	Platform *int32 `json:"platform,omitempty"`
	// price per unit (following MoneyFormat)
	Price *int32 `json:"price,omitempty"`
	// Product short description
	ShortDescription *string `json:"short_description,omitempty"`
	// unique identifier of the product
	Sku *string `json:"sku,omitempty"`
	Tags []IntegrationTag `json:"tags,omitempty"`
	// Optional: product variations if a line-item references a variation, the attributes of the variation have to be used instead of the ones of the product
	Variants []IntegrationProductVariant `json:"variants,omitempty"`
}

// NewIntegrationProduct instantiates a new IntegrationProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationProduct() *IntegrationProduct {
	this := IntegrationProduct{}
	return &this
}

// NewIntegrationProductWithDefaults instantiates a new IntegrationProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationProductWithDefaults() *IntegrationProduct {
	this := IntegrationProduct{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *IntegrationProduct) GetCategories() []IntegrationProductCategory {
	if o == nil || isNil(o.Categories) {
		var ret []IntegrationProductCategory
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationProduct) GetCategoriesOk() ([]IntegrationProductCategory, bool) {
	if o == nil || isNil(o.Categories) {
    return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *IntegrationProduct) HasCategories() bool {
	if o != nil && !isNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []IntegrationProductCategory and assigns it to the Categories field.
func (o *IntegrationProduct) SetCategories(v []IntegrationProductCategory) {
	o.Categories = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IntegrationProduct) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationProduct) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IntegrationProduct) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IntegrationProduct) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegrationProduct) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationProduct) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntegrationProduct) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IntegrationProduct) SetId(v string) {
	o.Id = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *IntegrationProduct) GetImages() []IntegrationProductImage {
	if o == nil || isNil(o.Images) {
		var ret []IntegrationProductImage
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationProduct) GetImagesOk() ([]IntegrationProductImage, bool) {
	if o == nil || isNil(o.Images) {
    return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *IntegrationProduct) HasImages() bool {
	if o != nil && !isNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []IntegrationProductImage and assigns it to the Images field.
func (o *IntegrationProduct) SetImages(v []IntegrationProductImage) {
	o.Images = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IntegrationProduct) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationProduct) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IntegrationProduct) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IntegrationProduct) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *IntegrationProduct) GetPlatform() int32 {
	if o == nil || isNil(o.Platform) {
		var ret int32
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationProduct) GetPlatformOk() (*int32, bool) {
	if o == nil || isNil(o.Platform) {
    return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *IntegrationProduct) HasPlatform() bool {
	if o != nil && !isNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given int32 and assigns it to the Platform field.
func (o *IntegrationProduct) SetPlatform(v int32) {
	o.Platform = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *IntegrationProduct) GetPrice() int32 {
	if o == nil || isNil(o.Price) {
		var ret int32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationProduct) GetPriceOk() (*int32, bool) {
	if o == nil || isNil(o.Price) {
    return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *IntegrationProduct) HasPrice() bool {
	if o != nil && !isNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given int32 and assigns it to the Price field.
func (o *IntegrationProduct) SetPrice(v int32) {
	o.Price = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *IntegrationProduct) GetShortDescription() string {
	if o == nil || isNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationProduct) GetShortDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.ShortDescription) {
    return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *IntegrationProduct) HasShortDescription() bool {
	if o != nil && !isNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *IntegrationProduct) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *IntegrationProduct) GetSku() string {
	if o == nil || isNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationProduct) GetSkuOk() (*string, bool) {
	if o == nil || isNil(o.Sku) {
    return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *IntegrationProduct) HasSku() bool {
	if o != nil && !isNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *IntegrationProduct) SetSku(v string) {
	o.Sku = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IntegrationProduct) GetTags() []IntegrationTag {
	if o == nil || isNil(o.Tags) {
		var ret []IntegrationTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationProduct) GetTagsOk() ([]IntegrationTag, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IntegrationProduct) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []IntegrationTag and assigns it to the Tags field.
func (o *IntegrationProduct) SetTags(v []IntegrationTag) {
	o.Tags = v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *IntegrationProduct) GetVariants() []IntegrationProductVariant {
	if o == nil || isNil(o.Variants) {
		var ret []IntegrationProductVariant
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationProduct) GetVariantsOk() ([]IntegrationProductVariant, bool) {
	if o == nil || isNil(o.Variants) {
    return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *IntegrationProduct) HasVariants() bool {
	if o != nil && !isNil(o.Variants) {
		return true
	}

	return false
}

// SetVariants gets a reference to the given []IntegrationProductVariant and assigns it to the Variants field.
func (o *IntegrationProduct) SetVariants(v []IntegrationProductVariant) {
	o.Variants = v
}

func (o IntegrationProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !isNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !isNil(o.ShortDescription) {
		toSerialize["short_description"] = o.ShortDescription
	}
	if !isNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Variants) {
		toSerialize["variants"] = o.Variants
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationProduct struct {
	value *IntegrationProduct
	isSet bool
}

func (v NullableIntegrationProduct) Get() *IntegrationProduct {
	return v.value
}

func (v *NullableIntegrationProduct) Set(val *IntegrationProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationProduct(val *IntegrationProduct) *NullableIntegrationProduct {
	return &NullableIntegrationProduct{value: val, isSet: true}
}

func (v NullableIntegrationProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


