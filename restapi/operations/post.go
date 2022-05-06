// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostHandlerFunc turns a function with the right signature into a post handler
type PostHandlerFunc func(PostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostHandlerFunc) Handle(params PostParams) middleware.Responder {
	return fn(params)
}

// PostHandler interface for that can handle valid post params
type PostHandler interface {
	Handle(PostParams) middleware.Responder
}

// NewPost creates a new http.Handler for the post operation
func NewPost(ctx *middleware.Context, handler PostHandler) *Post {
	return &Post{Context: ctx, Handler: handler}
}

/* Post swagger:route POST / post

Post post API

*/
type Post struct {
	Context *middleware.Context
	Handler PostHandler
}

func (o *Post) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostBody post body
//
// swagger:model PostBody
type PostBody struct {

	// List of commands to run
	// Example: convert mypic.png -resize 200x100 mypic.jpg
	// Required: true
	Commands []string `json:"commands"`

	// Returns the images as base64
	// Example: myimage.jpg
	Return []string `json:"return"`
}

// Validate validates this post body
func (o *PostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommands(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostBody) validateCommands(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"commands", "body", o.Commands); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post body based on context it is used
func (o *PostBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBody) UnmarshalBinary(b []byte) error {
	var res PostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostOKBody post o k body
// Example: {"greeting":"Hello YourName"}
//
// swagger:model PostOKBody
type PostOKBody struct {

	// commands
	Commands []*PostOKBodyCommandsItems0 `json:"commands"`

	// images
	Images []*PostOKBodyImagesItems0 `json:"images"`
}

// Validate validates this post o k body
func (o *PostOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommands(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOKBody) validateCommands(formats strfmt.Registry) error {
	if swag.IsZero(o.Commands) { // not required
		return nil
	}

	for i := 0; i < len(o.Commands); i++ {
		if swag.IsZero(o.Commands[i]) { // not required
			continue
		}

		if o.Commands[i] != nil {
			if err := o.Commands[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postOK" + "." + "commands" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postOK" + "." + "commands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostOKBody) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(o.Images) { // not required
		return nil
	}

	for i := 0; i < len(o.Images); i++ {
		if swag.IsZero(o.Images[i]) { // not required
			continue
		}

		if o.Images[i] != nil {
			if err := o.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postOK" + "." + "images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postOK" + "." + "images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post o k body based on the context it is used
func (o *PostOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCommands(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOKBody) contextValidateCommands(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Commands); i++ {

		if o.Commands[i] != nil {
			if err := o.Commands[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postOK" + "." + "commands" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postOK" + "." + "commands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostOKBody) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Images); i++ {

		if o.Images[i] != nil {
			if err := o.Images[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postOK" + "." + "images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postOK" + "." + "images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOKBody) UnmarshalBinary(b []byte) error {
	var res PostOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostOKBodyCommandsItems0 post o k body commands items0
//
// swagger:model PostOKBodyCommandsItems0
type PostOKBodyCommandsItems0 struct {

	// result
	Result interface{} `json:"result,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this post o k body commands items0
func (o *PostOKBodyCommandsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post o k body commands items0 based on context it is used
func (o *PostOKBodyCommandsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOKBodyCommandsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOKBodyCommandsItems0) UnmarshalBinary(b []byte) error {
	var res PostOKBodyCommandsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostOKBodyImagesItems0 post o k body images items0
//
// swagger:model PostOKBodyImagesItems0
type PostOKBodyImagesItems0 struct {

	// result
	Result string `json:"result,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this post o k body images items0
func (o *PostOKBodyImagesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post o k body images items0 based on context it is used
func (o *PostOKBodyImagesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOKBodyImagesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOKBodyImagesItems0) UnmarshalBinary(b []byte) error {
	var res PostOKBodyImagesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}