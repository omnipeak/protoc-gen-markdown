# Models

## Pet

Pet is an example enum describing a type of pet

| Value             | Description                                                            |
| ----------------- | ---------------------------------------------------------------------- |
| `PET_UNSPECIFIED` | The default value for a Pet field, usually treated as an invalid value |
| `PET_CAT`         | ≽^•⩊•^≼                                                                |
| `PET_DOG`         | ૮・ﻌ・ა                                                                  |
| `PET_FISH`        | 𓆟                                                                      |

## Person

Person is an example message describing a person

| Name   | Type                        | Required? | Description                       |
| ------ | --------------------------- | --------- | --------------------------------- |
| `id`   | `string`                    | ✅         | The unique identifier of a Person |
| `data` | [`PersonData`](#persondata) | ✅         | The actual data for the Person    |

### `Person.id` validation

The following validation rules apply to the `id` field:

- Must be a valid UUID

## PersonData

Person Data is an example message containing the actual data for a person

| Name           | Type                    | Required? | Description                |
| -------------- | ----------------------- | --------- | -------------------------- |
| `name`         | `string`                | ✅         | The Person's name          |
| `email`        | `string`                | ✅         | The Person's email address |
| `addresses`    | [`Address[]`](#address) | ✅         | The Person's Address(es)   |
| `favorite_pet` | [`Pet`](#pet)           | ❌         | The Person's favorite pet  |

### `PersonData.name` validation

The following validation rules apply to the `name` field:

- Must be at least 1 character long
- Must be 100 or fewer characters long

### `PersonData.email` validation

The following validation rules apply to the `email` field:

- Must be a valid email address

## Address

Address is an example message describing an address

| Name   | Type                          | Required? | Description                         |
| ------ | ----------------------------- | --------- | ----------------------------------- |
| `id`   | `string`                      | ✅         | The unique identifier of an Address |
| `data` | [`AddressData`](#addressdata) | ✅         | The actual data for the Address     |

### `Address.id` validation

The following validation rules apply to the `id` field:

- Must be a valid UUID

## AddressData

Address Data is an example message containing the actual data for an Address

| Name           | Type     | Required? | Description                                       |
| -------------- | -------- | --------- | ------------------------------------------------- |
| `person_id`    | `string` | ✅         | The ID of the Person to which the Address belongs |
| `line1`        | `string` | ✅         | The first line of the Address                     |
| `line2`        | `string` | ❌         | The second line of the Address                    |
| `line3`        | `string` | ❌         | The third line of the Address                     |
| `city`         | `string` | ✅         | The Address's city                                |
| `region`       | `string` | ❌         | The Address's region                              |
| `postal_code`  | `string` | ❌         | The postcode or zip of the Address                |
| `country_code` | `string` | ✅         | The 3-character ISO country code of the Address   |

### `AddressData.region` validation

The following validation rules apply to the `region` field:

- Must be at least 0 characters long
- Must be 40 or fewer characters long

### `AddressData.postal_code` validation

The following validation rules apply to the `postal_code` field:

- Must be at least 0 characters long
- Must be 40 or fewer characters long

### `AddressData.country_code` validation

The following validation rules apply to the `country_code` field:

- Must be at least 3 characters long
- Must be 3 or fewer characters long

### `AddressData.person_id` validation

The following validation rules apply to the `person_id` field:

- Must be a valid UUID

### `AddressData.line1` validation

The following validation rules apply to the `line1` field:

- Must be at least 1 character long
- Must be 40 or fewer characters long

### `AddressData.line2` validation

The following validation rules apply to the `line2` field:

- Must be at least 0 characters long
- Must be 40 or fewer characters long

### `AddressData.line3` validation

The following validation rules apply to the `line3` field:

- Must be at least 0 characters long
- Must be 40 or fewer characters long

### `AddressData.city` validation

The following validation rules apply to the `city` field:

- Must be at least 1 character long
- Must be 40 or fewer characters long

