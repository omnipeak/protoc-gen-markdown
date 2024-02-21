# Service

## PersonService

PersonService is the service for managing Person objects

### Methods

| Method           | Inputs                                            | Response                                            | Description                         |
| ---------------- | ------------------------------------------------- | --------------------------------------------------- | ----------------------------------- |
| `CreatePerson`   | [`CreatePersonRequest`](#createpersonrequest)     | [`CreatePersonResponse`](#createpersonresponse)     | Creates a Person                    |
| `RetrievePerson` | [`RetrievePersonRequest`](#retrievepersonrequest) | [`RetrievePersonResponse`](#retrievepersonresponse) | Retrieves a Person                  |
| `ListPeople`     | [`ListPeopleRequest`](#listpeoplerequest)         | [`ListPeopleResponse`](#listpeopleresponse)         | Lists people                        |
| `UpdatePerson`   | [`UpdatePersonRequest`](#updatepersonrequest)     | [`UpdatePersonResponse`](#updatepersonresponse)     | Updates a Person                    |
| `DeletePeople`   | [`DeletePeopleRequest`](#deletepeoplerequest)     | [`DeletePeopleResponse`](#deletepeopleresponse)     | Deletes one or more Person objects  |

## CreatePersonRequest

CreatePersonRequest is the request for creating a Person

| Name           | Type                            | Required? | Description                |
| -------------- | ------------------------------- | --------- | -------------------------- |
| `name`         | `string`                        | ✅         | The Person's Name          |
| `email`        | `string`                        | ✅         | The Person's email address |
| `addresses`    | [`AddressData[]`](#addressdata) | ✅         | The Person's Address(es)   |
| `favorite_pet` | [`Pet`](#pet)                   | ❌         | The Person's favorite pet  |

### `CreatePersonRequest.name` validation

The following validation rules apply to the `name` field:

- Must be at least 1 character long
- Must be 100 or fewer characters long

### `CreatePersonRequest.email` validation

The following validation rules apply to the `email` field:

- Must be a valid email address

## CreatePersonResponse

CreatePersonResponse is the response for creating a Person

| Name     | Type                | Required? | Description                               |
| -------- | ------------------- | --------- | ----------------------------------------- |
| `error`  | `string`            | ❌         | The error which occurred, if any          |
| `person` | [`Person`](#person) | ❌         | The Person which was created, if no error |

## RetrievePersonRequest

RetrievePersonRequest is the request for retrieving a Person

| Name | Type     | Required? | Description                      |
| ---- | -------- | --------- | -------------------------------- |
| `id` | `string` | ❌         | The ID of the Person to retrieve |

## RetrievePersonResponse

RetrievePersonResponse is the response for retrieving a Person

| Name     | Type                | Required? | Description                               |
| -------- | ------------------- | --------- | ----------------------------------------- |
| `error`  | `string`            | ❌         | The error which occurred, if any          |
| `person` | [`Person`](#person) | ❌         | The Person which was created, if no error |

## ListPeopleRequest

ListPeopleRequest is the request for listing people

| Name        | Type    | Required? | Description                           |
| ----------- | ------- | --------- | ------------------------------------- |
| `page`      | `int32` | ❌         | The page of data to retrieve          |
| `page_size` | `int32` | ❌         | The number of Person objects per page |

## ListPeopleResponse

ListPeopleResponse is the response for listing Person objects

| Name     | Type                  | Required? | Description                                              |
| -------- | --------------------- | --------- | -------------------------------------------------------- |
| `error`  | `string`              | ❌         | The error which occurred, if any                         |
| `people` | [`Person[]`](#person) | ❌         | A list of Person objects matching the request parameters |

## UpdatePersonRequest

UpdatePersonResponse is the response for updating a Person

| Name   | Type                | Required? | Description                        |
| ------ | ------------------- | --------- | ---------------------------------- |
| `data` | [`Person`](#person) | ❌         | The data to update the Person with |

## UpdatePersonResponse

UpdatePersonResponse is the response for updating a Person

| Name     | Type                | Required? | Description                      |
| -------- | ------------------- | --------- | -------------------------------- |
| `error`  | `string`            | ❌         | The error which occurred, if any |
| `person` | [`Person`](#person) | ❌         | Person is the updated person     |

## DeletePeopleRequest

DeletePeopleRequest is the request for deleting a Person

| Name | Type     | Required? | Description                          |
| ---- | -------- | --------- | ------------------------------------ |
| `id` | `string` | ❌         | Id is the id of the Person to delete |

## DeletePeopleResponse

DeletePeopleResponse is the response for deleting People

| Name    | Type       | Required? | Description                                     |
| ------- | ---------- | --------- | ----------------------------------------------- |
| `error` | `string`   | ❌         | The error which occurred, if any                |
| `ids`   | `string[]` | ❌         | The IDs of the Person objects that were deleted |

