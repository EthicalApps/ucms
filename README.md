# uCMS - A headless CMS

The status of this project is: Alpha - in active development.

uCMS aims to be a lightweight, standalone, stable, fast headless CMS.  It is distributed as a single self-contained binary available for macOS, Windows and Linux.

A future vision of this project is that it could be used as a personal content store where you could grant third party applications access to store your data in a well defined structure.  You could then grant access to another third party if you wanted to change providers as long as they are type-compatible.  To achieve this, we will produce a standard library of Content Types for common data structures so they can easily be shared.

## Dynamic Content Types

The core of the CMS revolves around Content Types.  Content Types can be created dynamically using a custom JSON format, the editor UI is then auto-generated from the definition and a JSON Schema is created to validate documents against.

A simplified `Article` Content Type might look like this

```json
{
  "id": "article",
  "type": "Article",
  "name": "Article",
  "description": "A simplified article",
  "displayField": "title",
  "fields": [
    {
      "id": "title",
      "type": "Text",
      "required": true
    },
    {
      "id": "body",
      "type": "TextArea",
      "required": true
    }
  ]
}
```

The corresponding generated JSON Schema would be

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "additionalProperties": false,
  "description": "A simplified article",
  "id": "/schema/article.json",
  "properties": {
      "body": {
          "type": "string"
      },
      "title": {
          "type": "string"
      }
  },
  "required": [
      "title",
      "body"
  ],
  "type": "object"
}
```

The CMS will now allow you to create and store `Article` documents that validate against this JSON Schema, either through the editor interface or securely via the API for third party access.

## Dynamic Editor

uCMS will come with a default editor implementation so you can create/edit content directly in the application, but the types and schemas are served via the API so third party editors could be built and/or embedded into another web application communicating with uCMS via the API.

## TODO

- [ ] Reference Types
- [ ] Full-text search
- [ ] Backups
- [ ] LetsEncrypt SSL certs
- [ ] Versioning
- [ ] Multi-language support
- [ ] Owner-authentication (3rd Party)
- [ ] Demo-frontend apps (Gatsby, Hugo, ...)


