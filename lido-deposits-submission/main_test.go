package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var (
	badDepositJSONTest = `[
  {
    "pubkey": "908487156aab37d60a0a4da525d66b243f52942840877d252f0f22325c07a2785596bacef8c4875e54f82a8dd2db88a2",
    "withdrawal_credentials": "00c508123dff12579b0f2a4e0cad9131768e97459d63292a55573b92ed893757",
    "amount": 32000000000,
    "signature": "853e7b017e6b6f398abeb3715dd69ed3f72f7564f8e0b052dd36cd60362b5af6e3c01a93488bdca13962aca558f2e53a035a38c07ef106284dce0b7aa46c7b46f389ac40de6a993410ed0fe66930d465af11ddff30e46a9202e1c17ec59186c6",
    "deposit_message_root": "594ee2d074a2d3739f6fed2d59b930ae54e67151ff383134358475da973bde88",
    "deposit_data_root": "e26e4f11390c62afffac7e85ca447907361b03ce4f552d4f8d79e189ecef6d6f",
    "fork_version": "00001020",
    "network_name": "prater",
    "deposit_cli_version": "2.2.0"
  }]`
	goodDepositJSONTest = `[
  {
    "pubkey": "908487156aab37d60a0a4da525d66b243f52942840877d252f0f22325c07a2785596bacef8c4875e54f82a8dd2db88a2",
    "withdrawal_credentials": "0100000000000000000000009d4af1ee19dad8857db3a45b0374c81c8a1c6320",
    "amount": 32000000000,
    "signature": "853e7b017e6b6f398abeb3715dd69ed3f72f7564f8e0b052dd36cd60362b5af6e3c01a93488bdca13962aca558f2e53a035a38c07ef106284dce0b7aa46c7b46f389ac40de6a993410ed0fe66930d465af11ddff30e46a9202e1c17ec59186c6",
    "deposit_message_root": "594ee2d074a2d3739f6fed2d59b930ae54e67151ff383134358475da973bde88",
    "deposit_data_root": "e26e4f11390c62afffac7e85ca447907361b03ce4f552d4f8d79e189ecef6d6f",
    "fork_version": "00001020",
    "network_name": "prater",
    "deposit_cli_version": "2.2.0"
  }]`
)

func TestCheckWithdrawalCredentialsContainsRegistry(t *testing.T) {
	t.Run("bad data", func(t *testing.T) {
		decoded := []*Deposit{}
		buf := bytes.NewBuffer([]byte(badDepositJSONTest))
		if err := json.NewDecoder(buf).Decode(&decoded); err != nil {
			t.Fatal(err)
		}
		praterAddr := registryContracts["prater"]
		_, _, err := extractPubkeysAndSigsFromData(decoded, praterAddr, "prater")
		if err == nil {
			t.Fatal("Expected error in test json")
		}
		lowerCase := strings.ToLower(strings.TrimPrefix(praterAddr, "0x"))
		want := fmt.Sprintf("withdrawal credentials %s for deposit at index %d does not contain expected "+
			"Lido operator registry address for %v network: %v", decoded[0].WithdrawalCredentials, 0, lowerCase, "prater")
		if err.Error() != want {
			t.Fatalf("Wanted %v, got %v", want, err)
		}
	})
	t.Run("good data", func(t *testing.T) {
		decoded := []*Deposit{}
		buf := bytes.NewBuffer([]byte(goodDepositJSONTest))
		if err := json.NewDecoder(buf).Decode(&decoded); err != nil {
			t.Fatal(err)
		}
		praterAddr := registryContracts["prater"]
		_, _, err := extractPubkeysAndSigsFromData(decoded, praterAddr, "prater")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
	})
}
