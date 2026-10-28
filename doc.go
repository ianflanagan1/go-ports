// Package ports holds no code. Its subpackages are the plain values that
// ports exchange: requests and responses, queries and rows. They do no I/O,
// so a project's core can build and read them, and its shell can declare
// ports in them, while go-adapters does the I/O behind those ports.
package ports
