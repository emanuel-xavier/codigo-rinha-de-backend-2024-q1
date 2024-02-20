package main

import "errors"

var (
	ErrNotFound            = errors.New("Not found")
	ErrDatabaseFailure     = errors.New("Database failure")
	ErrInsufficientBalance = errors.New("Insufficient balance")
	ErrInsufficientLimit   = errors.New("Insufficient limit")
	ErrUnknow              = errors.New("Unknow error")
	ErrParsingError        = errors.New("Failure to parse entity")
	ErrUnknowPaymentMethod = errors.New("Unknow payment method")
)
