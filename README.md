
# image-magick 1.0.0

Image modifications in Direktiv

---
- #### Category: Tools
- #### Image: direktiv/image-magick 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/image-magick/issues
- #### URL: https://github.com/direktiv-apps/image-magick
- #### Maintainer: [direktiv.io](https://www.direktiv.io)
---

## About image-magick

This app can run multiple [Image Magick](https://imagemagick.org/index.php) commands. 
The results can either be stored in the output folder of Direktiv to store them as variables or returned as base64.

### Example(s)
  #### Function Configuration
  ```yaml
  functions:
  - id: image-magick
    image: direktiv/image-magick
    type: knative-workflow
  ```
   #### Basic
   ```yaml
   - id: img
     type: action
     action:
       function: image-magick
       input:
         commands:
           - convert mypic.png json:
           - convert mypic.png -fuzz 25% -fill red -opaque white -flatten mypic.png
           - convert mypic.png -resize 200x100 mypic.jpg
         return:
           - mypic.png
   ```
   #### File Resize
   ```yaml
   - id: set
     type: setter
     variables:
     - key: mypic.png
       scope: workflow
       mimeType: application/octet-stream
       value: iVBORw0KGgoAAAANSUhEUgAAABQAAAAUCAYAAACNiR0NAAAAH0lEQVR42mNk+P+/noGKgHHUwFEDRw0cNXDUwJFqIAAczzHZPJWe1QAAAABJRU5ErkJggg==
     transition: modify
   - id: modify 
     type: action
     action:
       function: image
       files:
         - key: mypic.png
           scope: workflow
       input: 
         commands:
           # stores the image in workflow scope variable `resized.png`
           - convert mypic.png -resize 200% out/workflow/resized.png 
   ```

### Responses
  Response contains the results of the command and the requested images as base64.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| commands | [][PostOKBodyCommandsItems](#post-o-k-body-commands-items)| `[]*PostOKBodyCommandsItems` |  | | List of results for the commands. |  |
| images | [][PostOKBodyImagesItems](#post-o-k-body-images-items)| `[]*PostOKBodyImagesItems` |  | | List of base64 encoded images. |  |


#### <span id="post-o-k-body-commands-items"></span> postOKBodyCommandsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` |  | | Empty for conversion commands and<br> JSON for json returning commands. | `{"image":{"alpha":"sss"}}` |
| success | boolean| `bool` |  | | Indicates success | `true` |


#### <span id="post-o-k-body-images-items"></span> postOKBodyImagesItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | string| `string` |  | | Base64 encoded image string | `iVBORw0KGgoA...AABJRU5ErkJggg==` |
| success | boolean| `bool` |  | | Indicates success | `true` |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| commands | []string| `[]string` | ✓ | | List of commands to run | `convert mypic.png -resize 200x100 mypic.jpg` |
| return | []string| `[]string` |  | | Returns the images as base64 | `myimage.jpg` |

 
