package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

///////////////////////////////////////////////////
// Create a Custom Field
// POST /contactdb/custom_fields

func CreateaCustomField() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/custom_fields", host, "v3")
  request.Method = "POST"
  request.RequestBody = []byte(` {
  "name": "pet", 
  "type": "text"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve all custom fields
// GET /contactdb/custom_fields

func Retrieveallcustomfields() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/custom_fields", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve a Custom Field
// GET /contactdb/custom_fields/{custom_field_id}

func RetrieveaCustomField() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/custom_fields/{custom_field_id}", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Delete a Custom Field
// DELETE /contactdb/custom_fields/{custom_field_id}

func DeleteaCustomField() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/custom_fields/{custom_field_id}", host, "v3")
  request.Method = "DELETE"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Create a List
// POST /contactdb/lists

func CreateaList() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/lists", host, "v3")
  request.Method = "POST"
  request.RequestBody = []byte(` {
  "name": "your list name"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve all lists
// GET /contactdb/lists

func Retrievealllists() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/lists", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Delete Multiple lists
// DELETE /contactdb/lists

func DeleteMultiplelists() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/lists", host, "v3")
  request.Method = "DELETE"
  request.RequestBody = []byte(` [
  1, 
  2, 
  3, 
  4
]`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update a List
// PATCH /contactdb/lists/{list_id}

func UpdateaList() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/lists/{list_id}", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "name": "newlistname"
}`)
  queryParams := make(map[string]string)
  queryParams["list_id"] = "0"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve a single list
// GET /contactdb/lists/{list_id}

func Retrieveasinglelist() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/lists/{list_id}", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["list_id"] = "0"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Delete a List
// DELETE /contactdb/lists/{list_id}

func DeleteaList() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/lists/{list_id}", host, "v3")
  request.Method = "DELETE"
  queryParams := make(map[string]string)
  queryParams["delete_contacts"] = "true"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Add Multiple Recipients to a List
// POST /contactdb/lists/{list_id}/recipients

func AddMultipleRecipientstoaList() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/lists/{list_id}/recipients", host, "v3")
  request.Method = "POST"
  request.RequestBody = []byte(` [
  "recipient_id1", 
  "recipient_id2"
]`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve all recipients on a List
// GET /contactdb/lists/{list_id}/recipients

func RetrieveallrecipientsonaList() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/lists/{list_id}/recipients", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["page"] = "1"
queryParams["page_size"] = "1"
queryParams["list_id"] = "0"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Add a Single Recipient to a List
// POST /contactdb/lists/{list_id}/recipients/{recipient_id}

func AddaSingleRecipienttoaList() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/lists/{list_id}/recipients/{recipient_id}", host, "v3")
  request.Method = "POST"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Delete a Single Recipient from a Single List
// DELETE /contactdb/lists/{list_id}/recipients/{recipient_id}

func DeleteaSingleRecipientfromaSingleList() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/lists/{list_id}/recipients/{recipient_id}", host, "v3")
  request.Method = "DELETE"
  queryParams := make(map[string]string)
  queryParams["recipient_id"] = "0"
queryParams["list_id"] = "0"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update Recipient
// PATCH /contactdb/recipients

func UpdateRecipient() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/recipients", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` [
  {
    "email": "jones@example.com", 
    "first_name": "Guy", 
    "last_name": "Jones"
  }
]`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Add recipients
// POST /contactdb/recipients

func Addrecipients() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/recipients", host, "v3")
  request.Method = "POST"
  request.RequestBody = []byte(` [
  {
    "age": 25, 
    "email": "example@example.com", 
    "first_name": "", 
    "last_name": "User"
  }, 
  {
    "age": 25, 
    "email": "example2@example.com", 
    "first_name": "Example", 
    "last_name": "User"
  }
]`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve recipients
// GET /contactdb/recipients

func Retrieverecipients() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/recipients", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["page"] = "1"
queryParams["page_size"] = "1"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Delete Recipient
// DELETE /contactdb/recipients

func DeleteRecipient() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/recipients", host, "v3")
  request.Method = "DELETE"
  request.RequestBody = []byte(` [
  "recipient_id1", 
  "recipient_id2"
]`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve the count of billable recipients
// GET /contactdb/recipients/billable_count

func Retrievethecountofbillablerecipients() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/recipients/billable_count", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve a Count of Recipients
// GET /contactdb/recipients/count

func RetrieveaCountofRecipients() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/recipients/count", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve recipients matching search criteria
// GET /contactdb/recipients/search

func Retrieverecipientsmatchingsearchcriteria() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/recipients/search", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["{field_name}"] = "test_string"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve a single recipient
// GET /contactdb/recipients/{recipient_id}

func Retrieveasinglerecipient() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/recipients/{recipient_id}", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Delete a Recipient
// DELETE /contactdb/recipients/{recipient_id}

func DeleteaRecipient() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/recipients/{recipient_id}", host, "v3")
  request.Method = "DELETE"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve the lists that a recipient is on
// GET /contactdb/recipients/{recipient_id}/lists

func Retrievetheliststhatarecipientison() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/recipients/{recipient_id}/lists", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve reserved fields
// GET /contactdb/reserved_fields

func Retrievereservedfields() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/reserved_fields", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Create a Segment
// POST /contactdb/segments

func CreateaSegment() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/segments", host, "v3")
  request.Method = "POST"
  request.RequestBody = []byte(` {
  "conditions": [
    {
      "and_or": "", 
      "field": "last_name", 
      "operator": "eq", 
      "value": "Miller"
    }, 
    {
      "and_or": "and", 
      "field": "last_clicked", 
      "operator": "gt", 
      "value": "01/02/2015"
    }, 
    {
      "and_or": "or", 
      "field": "clicks.campaign_identifier", 
      "operator": "eq", 
      "value": "513"
    }
  ], 
  "list_id": 4, 
  "name": "Last Name Miller"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve all segments
// GET /contactdb/segments

func Retrieveallsegments() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/segments", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update a segment
// PATCH /contactdb/segments/{segment_id}

func Updateasegment() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/segments/{segment_id}", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "conditions": [
    {
      "and_or": "", 
      "field": "last_name", 
      "operator": "eq", 
      "value": "Miller"
    }
  ], 
  "list_id": 5, 
  "name": "The Millers"
}`)
  queryParams := make(map[string]string)
  queryParams["segment_id"] = "test_string"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve a segment
// GET /contactdb/segments/{segment_id}

func Retrieveasegment() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/segments/{segment_id}", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["segment_id"] = "0"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Delete a segment
// DELETE /contactdb/segments/{segment_id}

func Deleteasegment() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/segments/{segment_id}", host, "v3")
  request.Method = "DELETE"
  queryParams := make(map[string]string)
  queryParams["delete_contacts"] = "true"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve recipients on a segment
// GET /contactdb/segments/{segment_id}/recipients

func Retrieverecipientsonasegment() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/contactdb/segments/{segment_id}/recipients", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["page"] = "1"
queryParams["page_size"] = "1"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

