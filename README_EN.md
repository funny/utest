Introduction
============

This is a unit test tool library for Go program.

Utility API
===========

Let's make a comparation. The test code without unitest：

```go
func VerifyBuffer(t *testing.T, buffer InBuffer) {
	if buffer.ReadUint8() != 1 {
		t.Fatal("buffer.ReadUint8() != 1")
	}

	if buffer.ReadByte() != 99 {
		t.Fatal("buffer.ReadByte() != 99")
	}

	if buffer.ReadInt8() != -2 {
		t.Fatal("buffer.ReadInt8() != -2")
	}

	if buffer.ReadUint16() != 0xFFEE {
		t.Fatal("buffer.ReadUint16() != 0xFFEE")
	}

	if buffer.ReadInt16() != 0x7FEE {
		t.Fatal("buffer.ReadInt16() != 0x7FEE")
	}

	if buffer.ReadUint32() != 0xFFEEDDCC {
		t.Fatal("buffer.ReadUint32() != 0xFFEEDDCC")
	}

	if buffer.ReadInt32() != 0x7FEEDDCC {
		t.Fatal("buffer.ReadInt32() != 0x7FEEDDCC")
	}

	if buffer.ReadUint64() != 0xFFEEDDCCBBAA9988 {
		t.Fatal("buffer.ReadUint64() != 0xFFEEDDCCBBAA9988")
	}

	if buffer.ReadInt64() != 0x7FEEDDCCBBAA9988 {
		t.Fatal("buffer.ReadInt64() != 0x7FEEDDCCBBAA9988")
	}

	if buffer.ReadRune() != '好' {
		t.Fatal(`buffer.ReadRune() != '好'`)
	}

	if buffer.ReadString(6) != "Hello1" {
		t.Fatal(`buffer.ReadString() != "Hello"`)
	}

	if bytes.Equal(buffer.ReadBytes(6), []byte("Hello2")) != true {
		t.Fatal(`bytes.Equal(buffer.ReadBytes(5), []byte("Hello")) != true`)
	}

	if bytes.Equal(buffer.ReadSlice(6), []byte("Hello3")) != true {
		t.Fatal(`bytes.Equal(buffer.ReadSlice(5), []byte("Hello")) != true`)
	}
}
```

The test code with unitest：

```go
func VerifyBuffer(t *testing.T, buffer InBuffer) {
	unitest.Pass(t, buffer.ReadByte() == 99)
	unitest.Pass(t, buffer.ReadInt8() == -2)
	unitest.Pass(t, buffer.ReadUint8() == 1)
	unitest.Pass(t, buffer.ReadInt16() == 0x7FEE)
	unitest.Pass(t, buffer.ReadUint16() == 0xFFEE)
	unitest.Pass(t, buffer.ReadInt32() == 0x7FEEDDCC)
	unitest.Pass(t, buffer.ReadUint32() == 0xFFEEDDCC)
	unitest.Pass(t, buffer.ReadInt64() == 0x7FEEDDCCBBAA9988)
	unitest.Pass(t, buffer.ReadUint64() == 0xFFEEDDCCBBAA9988)
	unitest.Pass(t, buffer.ReadRune() == '好')
	unitest.Pass(t, buffer.ReadString(6) == "Hello1")
	unitest.Pass(t, bytes.Equal(buffer.ReadBytes(6), []byte("Hello2")))
	unitest.Pass(t, bytes.Equal(buffer.ReadSlice(6), []byte("Hello3")))
}
```

When the unit test failed, unitest will extract the failed line from source code.

Process Monitor
===============

Sometimes we need to diagnosis deadlock or monitor memory in unit test or benchmark.

Unitest provide a method to monitor test process：

```shell
echo 'lookup goroutine' > unitest.cmd
```

The shell script let unitest save stack trace for all goroutines into `unitest.goroutine` file.

Unitest support these monitor command：

```
lookup goroutine  -  Save stack trace for all goroutines into unitest.goroutine file
lookup heap       -  Save heap and GC information into unitest.heap file
lookup threadcreate - Save thread information into unitest.thread file
```

And, you can register `unitest.CommandHandler` callback to add custom monitor command support.
