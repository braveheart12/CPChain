package temp

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	Mc02 = MainnetConfig{
		// contract
		ContractRpt:       "0x000",
		ContractRnode:     "0x000",
		ContractAdmission: "0x000",
		ContractCampaign:  "0x000",
		ContractNetwork:   "0x000",

		Ca: "0xad8c6c5c0ea5aa5c34c91f731667cca295fca81f",

		// proposer
		P1:  "0xbcbe2655e84e1f12717a70b34ae5c613a651c1ac",
		P2:  "0xae6e61c3dea67294dc1a030f73b0500b7fe2d168",
		P3:  "0x8f7013a77b50bd100b0ad018e7380a3038210f7c",
		P4:  "0x9d82a4c51bc8461bdaf142d039e45f22aa7629f3",
		P5:  "0x4b744613fb3cd4821588e511c93ef4e1739d0c0a",
		P6:  "0xe535e99bc88bc738b36f8fdfe901b8309a88adb7",
		P14: "0x61185658e12ed22923b027ec8586e182690ef897",
		P15: "0x610a857dd2a676694545771eb23938f3474e2159",
		P16: "0x2a26c4bce965b56f6ba5c7fc5d4f187e64c2dcdb",
		P17: "0xb470a8d7fcb1105ae3d37400549d9caf7a942d96",
		P18: "0xa9f8d074c1de45e1e8378c971a5a4860dc0a17f0",
		P19: "0x4ab362d632b3758eca7da3b8b0c4db899e85010d",
		// validator
		V7:  "0x40697d6b8d5d539ce4fb77fd1f01abcce35cf6a5",
		V8:  "0x0a882b4ae83c8ca8af46711481b16ef96a504c7f",
		V9:  "0xf7b5f0af95e3ac5620ac7c8328051b8f29c7736d",
		V10: "0x7554eae3ccea3d6fa22ed3bcee527f716c03a8f8",
		V11: "0x1fcbee2ac5c28f4a7ac934aa8535ae62cc123376",
		V12: "0x6e23821bb2493687f7b5a44cca3d256fae832bf9",
		V13: "0x14e5ffa43798ea835fcdaec057f4e864d5c9cc94",
		// bootnode
		B1: "62c12325bc5e4f2b8e8be381df4d1b1f661d086c655e410f269e656ba7975e3e92da0864060ee120b0cae9bd739f223711a4a78887a3e6463ff1340cac74ab73",
		B2: "1489be549547e50c51730221cd04d834a7b9cb5467d88e2f81d4b33bde343007dbcdc4d9bc75bc7d7a109d0838a7b8d3ac32a487c7e57a5fab066dd3909af31e",
		B3: "8d075acb61149a506c3ff207c4831d71301af3d9667d43b0b9239af2b8bc42016b0ad919b348ef6a74df1007bc3b742283b0b5810f21d1af6f8c199f56480be8",
		// validator enode
		Ve1: "b0fb401804d48d992733bcb01812698cf2da94790c9f15f69bffd7c5188ac1c7b214f0fa1d7896816199c6080acb93b9f50b4c50017b5d32a190478c8eca2600",
		Ve2: "6033d60bf1daa21bf2a99f933bd1307fafb050d9a13268bb5a694d66632d5675926f953bf426f63065e376cf43f0ad9692e216bc2163f4a2dd17712e202d41c0",
		Ve3: "aadf96c7b6172596275859ffd773a61e3bfa0668b5eb9cc4c5c647af35512f6ba528011c1760fa2e6af444cb591f9d300d11ea9864a39ce1c1c509199a9ff32f",
		Ve4: "ddb02779aca4baa7348bc47069b5eb8e66a5249a9c63c91509472080f7aa35dd87afba1ce38902785f88e284f2216c547dde09090d372d0b84c0ad9220ef358b",
		Ve5: "7e7057aa0faa6b17d70c1d84db6d592acbe342a0f9e8c3f79993cbce3ef8fd6051faf50f7066d22c71e96b32700ea9a2291a2107b14decf9a7a70a388fb447a4",
		Ve6: "a07e2dbdef23b9f2565e288493c156b0c1cb3615b7b1c558b96a13658c059453e3c5efcf08458a04c4a95a18d0b17bfcbc23bba8b736b0c5c96d829ccc3a5d7e",
		Ve7: "097cec8e68bbd1e3fa4aa928c8052cc4a1d54d62942cb73cd195f5ff91d86b59897465da4ae7114b78123674a7e3b417fa41dc86cea2df152db3474a9e9b7f2d",
	}
)

func init() {
	caAddr := common.HexToAddress(Mc02.Ca)
	Mc02.ContractRnode = crypto.CreateAddress(caAddr, 0).Hex()
	Mc02.ContractAdmission = crypto.CreateAddress(caAddr, 1).Hex()
	Mc02.ContractCampaign = crypto.CreateAddress(caAddr, 2).Hex()
	Mc02.ContractRpt = crypto.CreateAddress(caAddr, 3).Hex()
	Mc02.ContractNetwork = crypto.CreateAddress(caAddr, 4).Hex()
}
