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
