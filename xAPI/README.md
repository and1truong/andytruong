xAPI
====

## xAPI statement

- I -> actor: actor taking the action
- DID -> verb: Describe the action taken
- THIS -> object: The target of the action

### Example

Tom attended "xAPI Foundations"

- Tom: actor `mailto:tom@example.com`
- Attended: verb `https://adlnet.gov/expapi/verbs/Attended`
- "xAPI foundations": object `https://www.linkedin.com/learning/xapi-foundations`

Tom played "Caminandes Video"

- Tom: actor `mailto:tom@example.com`
- Played: verb `https://adlnet.gov/expapi/verbs/Played`
- "Caminandes Video": object `https://example.com/caminandes1.mp4`


```json
{
	"actor": {
		"mbox":"mailto:tom@example.com",
		"name":"Tom Doe",
		"objectType": "Agent"
	},
	"verb": {
		"id": "https://adlnet.gov/expapi/verbs/attended",
		"display": {
			"en-US": "attended"
		}
	},
	"object": {
		"id": "https://www.linkedin.com/learning/xapi-foundations/",
		"objectType": "Activity",
		"definition": {
			"name": { "en-US": "xAPI Foundations"},
			"description": {
				"en-US": "An introduction to the Experience API"
			}
		}
	}
}
```

## xAPI Unique Identifiers

- IRI: Internationalized Resource Identifier
- URI: Uniform Resource Identifier
- URL: Uniform Resource Locator

## Actor

example 1

```json
{
    "actor": {
        "mbox":"mailto:tom@example.com",
        "name":"Tom Doe",
        "objectType": "Agent"
    }
}
```

example 2

```json
{
    "actor": {
        "account": {
            "name":"TomDoe",
            "homePage": "https://www.myLMS.com",
            "mbox_sha1sum": "41f0631eedc0664d61adcff8d0a248d7a141aa15",
            "openid": "https://d_thompson.myopenID.com"
        },
        "objectType": "Agent"
    }
}
```

## Result

```json
{
	"result": {
		"completion": true,
		"success": true,
		"score": {
			"raw": 95,
			"scaled": 0.95
		}
	}
}
```

## Context

1. *registration* - UUID that links the statements to the same registration/enrollment.
2. *instructor* - an agent that describes who provided the instruction, or taught the class.
3. *revision* - free-form string value that is used to denote updates to the activity, such as bug fixes and typo corrections; however, extensive updates should be references with a new activity/object ID.
4. *platform* - free-form string value, which can be used to list the method of delivery of the activity.
5. *language* - describes the language used during the activity.
6. *Team* (advanced)
7. *contextActivities* (advanced)
8. *Statement* (advanced)

Example

```json
{
	"context": {
		"registration": "01F1MH0JFKW830816P6BJH6MC3",
		"instructor": {
			"account": {
				"name": "aa_doe",
				"homePage": "https://twitter.com/aa_doe"
			}
		},
		"revision": "1",
		"platform": "LinkedIn.com/learning",
		"language": "en-US"
	}
}
```

## Extension

- Extensions add custom information to create more detailed definitions/statements.
- They can extend the definition of the Object, Result, or Context elements.
- Keys must be fully qualified IRIs (usually will be URLs).
  - Extract opposite of Object IDs
  - URLs should resolve to a working webpage to provide further informations, but it is not required.Sin

```json
{
	"object": {
		"id": "https://www.linkedin.com/learning/xapi-foundations/",
		"objectType": "Activity",
		"definition": {
			"name": { "en-US": "xAPI Foundations"},
			"description": {
				"en-US": "An introduction to the Experience API"
			}
		},
		"extensions": {
			"https://example.com/xAPI_ext/catalog": "12345"
		}
	}
}
```
