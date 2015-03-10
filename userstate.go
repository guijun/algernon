package main

import (
	"net/http"

	"github.com/xyproto/permissions2"
	"github.com/yuin/gopher-lua"
)

func exportUserstate(w http.ResponseWriter, req *http.Request, L *lua.LState, userstate *permissions.UserState) {
	// Check if the current user has "user rights", returns bool
	// Takes no arguments
	L.SetGlobal("UserRights", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(userstate.UserRights(req)))
		return 1 // number of results
	}))
	// Check if the given username exists, returns bool
	// Takes a username
	L.SetGlobal("HasUser", L.NewFunction(func(L *lua.LState) int {
		username := L.ToString(1)
		L.Push(lua.LBool(userstate.HasUser(username)))
		return 1 // number of results
	}))
	// Get the value from the given boolean field, returns bool
	// Takes a username and fieldname
	L.SetGlobal("BooleanField", L.NewFunction(func(L *lua.LState) int {
		username := L.ToString(1)
		fieldname := L.ToString(2)
		L.Push(lua.LBool(userstate.BooleanField(username, fieldname)))
		return 1 // number of results
	}))
	// Save a value as a boolean field, returns nothing
	// Takes a username, fieldname and boolean value
	L.SetGlobal("SetBooleanField", L.NewFunction(func(L *lua.LState) int {
		username := L.ToString(1)
		fieldname := L.ToString(2)
		value := L.ToBool(3)
		userstate.SetBooleanField(username, fieldname, value)
		return 0 // number of results
	}))
	// Check if a given username is confirmed, returns a bool
	// Takes a username
	L.SetGlobal("IsConfirmed", L.NewFunction(func(L *lua.LState) int {
		username := L.ToString(1)
		L.Push(lua.LBool(userstate.IsConfirmed(username)))
		return 1 // number of results
	}))
	// Check if a given username is logged in, returns a bool
	// Takes a username
	L.SetGlobal("IsLoggedIn", L.NewFunction(func(L *lua.LState) int {
		username := L.ToString(1)
		L.Push(lua.LBool(userstate.IsLoggedIn(username)))
		return 1 // number of results
	}))
	// Check if the current user has "admin rights", returns a bool
	// Takes no arguments.
	L.SetGlobal("AdminRights", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(userstate.AdminRights(req)))
		return 1 // number of results
	}))
	// Check if a given username is an admin, returns a bool
	// Takes a username
	L.SetGlobal("IsAdmin", L.NewFunction(func(L *lua.LState) int {
		username := L.ToString(1)
		L.Push(lua.LBool(userstate.IsAdmin(username)))
		return 1 // number of results
	}))
	// Get the username stored in a cookie, or an empty string
	// Takes no arguments
	L.SetGlobal("UsernameCookie", L.NewFunction(func(L *lua.LState) int {
		username, err := userstate.UsernameCookie(req)
		if err != nil {
			L.Push(lua.LString(""))
		} else {
			L.Push(lua.LString(username))
		}
		return 1 // number of results
	}))
	// Store the username in a cookie, returns true if successful
	// Takes a username
	L.SetGlobal("SetUsernameCookie", L.NewFunction(func(L *lua.LState) int {
		username := L.ToString(1)
		L.Push(lua.LBool(nil == userstate.SetUsernameCookie(w, username)))
		return 1 // number of results
	}))

	// TODO: The rest of the functions from permissions2/userstate.go
}