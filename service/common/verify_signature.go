package common

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/storyicon/sigverify"
)

func VerifySignature(msgHex, sigHex, walletAddressHex string) (bool, error) {
	msg, err := hexutil.Decode(msgHex)
	if err != nil {
		return false, errors.WithMessage(err, "msg hex decode failed")
	}
	valid, err := sigverify.VerifyEllipticCurveHexSignatureEx(
		ethcommon.HexToAddress(walletAddressHex),
		msg,
		sigHex,
	)
	if err != nil {
		return false, err
	}
	return valid, nil
}
