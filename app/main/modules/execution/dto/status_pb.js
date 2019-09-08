// source: status.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.JobStatus', null, global);
/**
 * @enum {number}
 */
proto.JobStatus = {
    UNKNOWN: 0,
    QUEUED: 1,
    RUNNING: 2,
    COMPLETED: 3,
    CANCELLED: 4,
    ERRORED: 5,
};

goog.object.extend(exports, proto);
