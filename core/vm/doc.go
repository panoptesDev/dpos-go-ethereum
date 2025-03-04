// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

/*
Package vm implements the Ethereum Virtual Machine.

The vm package implements one EVM, a byte code VM. The BC (Byte Code) VM loops
over a set of bytes and executes them according to the set of rules defined
in the Ethereum yellow paper.

EVM是支持栈的虚拟机，为了方便进行密码学计算，采用了32字节（256比特）的字长。EVM是大端机器
EVM栈以字（Word）为单位进行操作，最多可以容纳1024个字。
由于操作码被限制在一个字节以内，所以EVM指令集最多只能容纳256条指令

EVM栈操作指令
	POP指令、PUSHx系列指令、DUPx系列指令、SWAPx系列指令

*/
package vm
