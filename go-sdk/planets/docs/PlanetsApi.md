# \PlanetsApi

All URIs are relative to *http://localhost:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllPlanets**](PlanetsApi.md#AllPlanets) | **Get** /planets | List all planets
[**OnePlanet**](PlanetsApi.md#OnePlanet) | **Get** /planets/{planetId} | Fetch one planet by position



## AllPlanets

> []Planet AllPlanets(ctx, )

List all planets

Returns a list of all the planets that are stored in the system.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Planet**](planet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnePlanet

> Planet OnePlanet(ctx, planetId)

Fetch one planet by position

Get the data relating to one planet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planetId** | **float32**|  | 

### Return type

[**Planet**](planet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

