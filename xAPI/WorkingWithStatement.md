Working with statement
====

## For steps for xAPI

### 1. Define the LRS and the username and password to authenticate.

```js
var conf = {
  "endpoint": "https://…",
  "user": "xapi-tools",
  "password": "xxxxxxx"
}
```

### 2. Tell the browser to use that LRS.

```js
ADL,XAPIWrapper.changeConfig(conf);
```

### 3. Define the xAPI statement.

```js
const statement = {
  "actor": …
  "verb": …
  "object": …
}
```

### 4. Send the statement.

```js
const result = ADL.XAPIWrapper.sendStatement(statement);
```
