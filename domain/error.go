package domain

import "errors"

var ErrorAccountNotFound = errors.New("Account Not found !")
var ErrorInvalidCredential = errors.New("Invalid Credential !")
var ErrorEmailIsAlreadyExists = errors.New("Email is already taken !")
var ErrorInternalServerError = errors.New("Something went wrong !")
