/*
 * Copyright 2023 John Douglas Pritchard
 *
 * Redistribution and use in source and binary forms, with
 * or without modification, are permitted provided that the
 * following conditions are met:
 *
 * 1. Redistributions of source code must retain the above
 * copyright notice, this list of conditions and the
 * following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the
 * above copyright notice, this list of conditions and the
 * following disclaimer in the documentation and/or other
 * materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND
 * CONTRIBUTORS “AS IS” AND ANY EXPRESS OR IMPLIED
 * WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A
 * PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 * COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
 * OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH
 * DAMAGE.
 */
/*
 * https://go.dev/play/p/5mr5M0luZ9k
 *
 * Exploration of complex type abstraction.
 */
package main

import (
	"fmt"
	"os"
)

/*
 * Type classification operator.
 */
type TypeIdentifier uint8

const (
	TypeIdentifierAbstract   TypeIdentifier = iota
	TypeIdentifierConcrete   TypeIdentifier = TypeIdentifierAbstract + 1
	TypeIdentifierDerivative TypeIdentifier = TypeIdentifierConcrete + 1
)

/*
 * Type class abstraction.
 */
type Abstract interface {
	Type() TypeIdentifier

	List() string
}

/*
 * Type class application.
 */
type Concrete struct {
	Abstract
}

/*
 * Type class derivation.
 */
type Derivative struct {
	Concrete
}

/*
 * Type class principal operator.
 */
func Init() Abstract {

	var concrete Concrete

	return concrete
}

/*
 * Type class derivative operator.
 */
func (concrete Concrete) Type() TypeIdentifier {

	return TypeIdentifierConcrete
}

/*
 * Type class derivative operator.
 */
func (concrete Concrete) List() string {

	return "Concrete"
}

/*
 * Type class derivative operator.
 */
func (derivative Derivative) Type() TypeIdentifier {

	return TypeIdentifierDerivative
}

/*
 * Type class derivative operator.
 */
func (derivative Derivative) List() string {

	return "Derivative"
}

func main() {

	var abstract Abstract = Init()

	fmt.Fprintf(os.Stderr, "type: %s\n", abstract.List())

	os.Exit(0)
}

