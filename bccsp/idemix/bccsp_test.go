
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:49</date>
//</624455928338452480>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package idemix_test

import (
	"crypto/rand"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/idemix"
	"github.com/hyperledger/fabric/bccsp/sw"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Idemix Bridge", func() {

	Describe("setting up the environment with one issuer and one user", func() {
		var (
			CSP             bccsp.BCCSP
			IssuerKey       bccsp.Key
			IssuerPublicKey bccsp.Key
			AttributeNames  []string

			UserKey      bccsp.Key
			NymKey       bccsp.Key
			NymPublicKey bccsp.Key

			IssuerNonce []byte
			credRequest []byte

			credential []byte

			RevocationKey       bccsp.Key
			RevocationPublicKey bccsp.Key
			cri                 []byte
		)

		BeforeEach(func() {
			var err error
			CSP, err = idemix.New(sw.NewDummyKeyStore())
			Expect(err).NotTo(HaveOccurred())

//发行人
			AttributeNames = []string{"Attr1", "Attr2", "Attr3", "Attr4", "Attr5"}
			IssuerKey, err = CSP.KeyGen(&bccsp.IdemixIssuerKeyGenOpts{Temporary: true, AttributeNames: AttributeNames})
			Expect(err).NotTo(HaveOccurred())
			IssuerPublicKey, err = IssuerKey.PublicKey()
			Expect(err).NotTo(HaveOccurred())

//用户
			UserKey, err = CSP.KeyGen(&bccsp.IdemixUserSecretKeyGenOpts{Temporary: true})
			Expect(err).NotTo(HaveOccurred())

//用户NYM密钥
			NymKey, err = CSP.KeyDeriv(UserKey, &bccsp.IdemixNymKeyDerivationOpts{Temporary: true, IssuerPK: IssuerPublicKey})
			Expect(err).NotTo(HaveOccurred())
			NymPublicKey, err = NymKey.PublicKey()
			Expect(err).NotTo(HaveOccurred())

			IssuerNonce = make([]byte, 32)
			n, err := rand.Read(IssuerNonce)
			Expect(n).To(BeEquivalentTo(32))
			Expect(err).NotTo(HaveOccurred())

//用户的凭据请求
			credRequest, err = CSP.Sign(
				UserKey,
				nil,
				&bccsp.IdemixCredentialRequestSignerOpts{IssuerPK: IssuerPublicKey, IssuerNonce: IssuerNonce},
			)
			Expect(err).NotTo(HaveOccurred())

//凭据
			credential, err = CSP.Sign(
				IssuerKey,
				credRequest,
				&bccsp.IdemixCredentialSignerOpts{
					Attributes: []bccsp.IdemixAttribute{
						{Type: bccsp.IdemixBytesAttribute, Value: []byte{0}},
						{Type: bccsp.IdemixBytesAttribute, Value: []byte{0, 1}},
						{Type: bccsp.IdemixIntAttribute, Value: 1},
						{Type: bccsp.IdemixBytesAttribute, Value: []byte{0, 1, 2}},
						{Type: bccsp.IdemixBytesAttribute, Value: []byte{0, 1, 2, 3}},
					},
				},
			)
			Expect(err).NotTo(HaveOccurred())

//撤销
			RevocationKey, err = CSP.KeyGen(&bccsp.IdemixRevocationKeyGenOpts{Temporary: true})
			Expect(err).NotTo(HaveOccurred())
			RevocationPublicKey, err = RevocationKey.PublicKey()
			Expect(err).NotTo(HaveOccurred())

//CRI
			cri, err = CSP.Sign(
				RevocationKey,
				nil,
				&bccsp.IdemixCRISignerOpts{},
			)
			Expect(err).NotTo(HaveOccurred())

		})

		It("the environment is properly set", func() {
//验证CredRequest
			valid, err := CSP.Verify(
				IssuerPublicKey,
				credRequest,
				nil,
				&bccsp.IdemixCredentialRequestSignerOpts{IssuerNonce: IssuerNonce},
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(valid).To(BeTrue())

//验证凭据
			valid, err = CSP.Verify(
				UserKey,
				credential,
				nil,
				&bccsp.IdemixCredentialSignerOpts{
					IssuerPK: IssuerPublicKey,
					Attributes: []bccsp.IdemixAttribute{
						{Type: bccsp.IdemixBytesAttribute, Value: []byte{0}},
						{Type: bccsp.IdemixBytesAttribute, Value: []byte{0, 1}},
						{Type: bccsp.IdemixIntAttribute, Value: 1},
						{Type: bccsp.IdemixBytesAttribute, Value: []byte{0, 1, 2}},
						{Type: bccsp.IdemixBytesAttribute, Value: []byte{0, 1, 2, 3}},
					},
				},
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(valid).To(BeTrue())

//验证CRI
			valid, err = CSP.Verify(
				RevocationPublicKey,
				cri,
				nil,
				&bccsp.IdemixCRISignerOpts{},
			)
			Expect(err).NotTo(HaveOccurred())
		})

		Describe("producing an idemix signature with no disclosed attribute", func() {
			var (
				digest    []byte
				signature []byte
			)

			BeforeEach(func() {
				var err error

				digest = []byte("a digest")

				signature, err = CSP.Sign(
					UserKey,
					digest,
					&bccsp.IdemixSignerOpts{
						Credential: credential,
						Nym:        NymKey,
						IssuerPK:   IssuerPublicKey,
						Attributes: []bccsp.IdemixAttribute{
							{Type: bccsp.IdemixHiddenAttribute},
							{Type: bccsp.IdemixHiddenAttribute},
							{Type: bccsp.IdemixHiddenAttribute},
							{Type: bccsp.IdemixHiddenAttribute},
							{Type: bccsp.IdemixHiddenAttribute},
						},
						RhIndex: 4,
						Epoch:   0,
						CRI:     cri,
					},
				)
				Expect(err).NotTo(HaveOccurred())
			})

			It("the signature is valid", func() {
				valid, err := CSP.Verify(
					IssuerPublicKey,
					signature,
					digest,
					&bccsp.IdemixSignerOpts{
						RevocationPublicKey: RevocationPublicKey,
						Attributes: []bccsp.IdemixAttribute{
							{Type: bccsp.IdemixHiddenAttribute},
							{Type: bccsp.IdemixHiddenAttribute},
							{Type: bccsp.IdemixHiddenAttribute},
							{Type: bccsp.IdemixHiddenAttribute},
							{Type: bccsp.IdemixHiddenAttribute},
						},
						RhIndex: 4,
						Epoch:   0,
					},
				)
				Expect(err).NotTo(HaveOccurred())
				Expect(valid).To(BeTrue())
			})

		})

		Describe("producing an idemix signature with disclosed attributes", func() {
			var (
				digest    []byte
				signature []byte
			)

			BeforeEach(func() {
				var err error

				digest = []byte("a digest")

				signature, err = CSP.Sign(
					UserKey,
					digest,
					&bccsp.IdemixSignerOpts{
						Credential: credential,
						Nym:        NymKey,
						IssuerPK:   IssuerPublicKey,
						Attributes: []bccsp.IdemixAttribute{
							{Type: bccsp.IdemixBytesAttribute},
							{Type: bccsp.IdemixHiddenAttribute},
							{Type: bccsp.IdemixIntAttribute},
							{Type: bccsp.IdemixHiddenAttribute},
							{Type: bccsp.IdemixHiddenAttribute},
						},
						RhIndex: 4,
						Epoch:   0,
						CRI:     cri,
					},
				)
				Expect(err).NotTo(HaveOccurred())
			})

			It("the signature is valid", func() {
				valid, err := CSP.Verify(
					IssuerPublicKey,
					signature,
					digest,
					&bccsp.IdemixSignerOpts{
						RevocationPublicKey: RevocationPublicKey,
						Attributes: []bccsp.IdemixAttribute{
							{Type: bccsp.IdemixBytesAttribute, Value: []byte{0}},
							{Type: bccsp.IdemixHiddenAttribute},
							{Type: bccsp.IdemixIntAttribute, Value: 1},
							{Type: bccsp.IdemixHiddenAttribute},
							{Type: bccsp.IdemixHiddenAttribute},
						},
						RhIndex: 4,
						Epoch:   0,
					},
				)
				Expect(err).NotTo(HaveOccurred())
				Expect(valid).To(BeTrue())
			})

		})

		Describe("producing an idemix nym signature", func() {
			var (
				digest    []byte
				signature []byte
			)

			BeforeEach(func() {
				var err error

				digest = []byte("a digest")

				signature, err = CSP.Sign(
					UserKey,
					digest,
					&bccsp.IdemixNymSignerOpts{
						Nym:      NymKey,
						IssuerPK: IssuerPublicKey,
					},
				)
				Expect(err).NotTo(HaveOccurred())
			})

			It("the signature is valid", func() {
				valid, err := CSP.Verify(
					NymPublicKey,
					signature,
					digest,
					&bccsp.IdemixNymSignerOpts{
						IssuerPK: IssuerPublicKey,
					},
				)
				Expect(err).NotTo(HaveOccurred())
				Expect(valid).To(BeTrue())
			})

		})
	})
})

