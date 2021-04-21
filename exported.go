// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

package errorlogger

import "github.com/sirupsen/logrus"

// Log implements the default logrus error logger
//
// Reference: https://github.com/sirupsen/logrus/
var Log *logrus.Logger = DefaultLogger

var EL = New()
var Err = EL.Err
