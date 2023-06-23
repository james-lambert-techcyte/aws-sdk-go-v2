// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Contains information about an alias.
type Alias struct {

	// A friendly name that you can use to refer to a key. The value must begin with
	// alias/ . Do not include confidential or sensitive information in this field.
	// This field may be displayed in plaintext in CloudTrail logs and other output.
	//
	// This member is required.
	AliasName *string

	// The KeyARN of the key associated with the alias.
	KeyArn *string

	noSmithyDocumentSerde
}

// Parameter information for key material export from Amazon Web Services Payment
// Cryptography.
//
// The following types satisfy this interface:
//
//	ExportKeyMaterialMemberTr31KeyBlock
//	ExportKeyMaterialMemberTr34KeyBlock
type ExportKeyMaterial interface {
	isExportKeyMaterial()
}

// Parameter information for key material export using TR-31 standard.
type ExportKeyMaterialMemberTr31KeyBlock struct {
	Value ExportTr31KeyBlock

	noSmithyDocumentSerde
}

func (*ExportKeyMaterialMemberTr31KeyBlock) isExportKeyMaterial() {}

// Parameter information for key material export using TR-34 standard.
type ExportKeyMaterialMemberTr34KeyBlock struct {
	Value ExportTr34KeyBlock

	noSmithyDocumentSerde
}

func (*ExportKeyMaterialMemberTr34KeyBlock) isExportKeyMaterial() {}

// Parameter information for key material export using TR-31 standard.
type ExportTr31KeyBlock struct {

	// The KeyARN of the the wrapping key. This key encrypts or wraps the key under
	// export for TR-31 key block generation.
	//
	// This member is required.
	WrappingKeyIdentifier *string

	noSmithyDocumentSerde
}

// Parameter information for key material export using TR-34 standard.
type ExportTr34KeyBlock struct {

	// The KeyARN of the certificate chain that signs the wrapping key certificate
	// during TR-34 key export.
	//
	// This member is required.
	CertificateAuthorityPublicKeyIdentifier *string

	// The export token to initiate key export from Amazon Web Services Payment
	// Cryptography. It also contains the signing key certificate that will sign the
	// wrapped key during TR-34 key block generation. Call GetParametersForExport to
	// receive an export token. It expires after 7 days. You can use the same export
	// token to export multiple keys from the same service account.
	//
	// This member is required.
	ExportToken *string

	// The format of key block that Amazon Web Services Payment Cryptography will use
	// during key export.
	//
	// This member is required.
	KeyBlockFormat Tr34KeyBlockFormat

	// The KeyARN of the wrapping key certificate. Amazon Web Services Payment
	// Cryptography uses this certificate to wrap the key under export.
	//
	// This member is required.
	WrappingKeyCertificate *string

	// A random number value that is unique to the TR-34 key block generated using 2
	// pass. The operation will fail, if a random nonce value is not provided for a
	// TR-34 key block generated using 2 pass.
	RandomNonce *string

	noSmithyDocumentSerde
}

// Parameter information for key material import.
//
// The following types satisfy this interface:
//
//	ImportKeyMaterialMemberRootCertificatePublicKey
//	ImportKeyMaterialMemberTr31KeyBlock
//	ImportKeyMaterialMemberTr34KeyBlock
//	ImportKeyMaterialMemberTrustedCertificatePublicKey
type ImportKeyMaterial interface {
	isImportKeyMaterial()
}

// Parameter information for root public key certificate import.
type ImportKeyMaterialMemberRootCertificatePublicKey struct {
	Value RootCertificatePublicKey

	noSmithyDocumentSerde
}

func (*ImportKeyMaterialMemberRootCertificatePublicKey) isImportKeyMaterial() {}

// Parameter information for key material import using TR-31 standard.
type ImportKeyMaterialMemberTr31KeyBlock struct {
	Value ImportTr31KeyBlock

	noSmithyDocumentSerde
}

func (*ImportKeyMaterialMemberTr31KeyBlock) isImportKeyMaterial() {}

// Parameter information for key material import using TR-34 standard.
type ImportKeyMaterialMemberTr34KeyBlock struct {
	Value ImportTr34KeyBlock

	noSmithyDocumentSerde
}

func (*ImportKeyMaterialMemberTr34KeyBlock) isImportKeyMaterial() {}

// Parameter information for trusted public key certificate import.
type ImportKeyMaterialMemberTrustedCertificatePublicKey struct {
	Value TrustedCertificatePublicKey

	noSmithyDocumentSerde
}

func (*ImportKeyMaterialMemberTrustedCertificatePublicKey) isImportKeyMaterial() {}

// Parameter information for key material import using TR-31 standard.
type ImportTr31KeyBlock struct {

	// The TR-34 wrapped key block to import.
	//
	// This member is required.
	WrappedKeyBlock *string

	// The KeyARN of the key that will decrypt or unwrap a TR-31 key block during
	// import.
	//
	// This member is required.
	WrappingKeyIdentifier *string

	noSmithyDocumentSerde
}

// Parameter information for key material import using TR-34 standard.
type ImportTr34KeyBlock struct {

	// The KeyARN of the certificate chain that signs the signing key certificate
	// during TR-34 key import.
	//
	// This member is required.
	CertificateAuthorityPublicKeyIdentifier *string

	// The import token that initiates key import into Amazon Web Services Payment
	// Cryptography. It expires after 7 days. You can use the same import token to
	// import multiple keys to the same service account.
	//
	// This member is required.
	ImportToken *string

	// The key block format to use during key import. The only value allowed is
	// X9_TR34_2012 .
	//
	// This member is required.
	KeyBlockFormat Tr34KeyBlockFormat

	// The public key component in PEM certificate format of the private key that
	// signs the KDH TR-34 wrapped key block.
	//
	// This member is required.
	SigningKeyCertificate *string

	// The TR-34 wrapped key block to import.
	//
	// This member is required.
	WrappedKeyBlock *string

	// A random number value that is unique to the TR-34 key block generated using 2
	// pass. The operation will fail, if a random nonce value is not provided for a
	// TR-34 key block generated using 2 pass.
	RandomNonce *string

	noSmithyDocumentSerde
}

// Metadata about an Amazon Web Services Payment Cryptography key.
type Key struct {

	// The date and time when the key was created.
	//
	// This member is required.
	CreateTimestamp *time.Time

	// Specifies whether the key is enabled.
	//
	// This member is required.
	Enabled *bool

	// Specifies whether the key is exportable. This data is immutable after the key
	// is created.
	//
	// This member is required.
	Exportable *bool

	// The Amazon Resource Name (ARN) of the key.
	//
	// This member is required.
	KeyArn *string

	// The role of the key, the algorithm it supports, and the cryptographic
	// operations allowed with the key. This data is immutable after the key is
	// created.
	//
	// This member is required.
	KeyAttributes *KeyAttributes

	// The key check value (KCV) is used to check if all parties holding a given key
	// have the same key or to detect that a key has changed. Amazon Web Services
	// Payment Cryptography calculates the KCV by using standard algorithms, typically
	// by encrypting 8 or 16 bytes or "00" or "01" and then truncating the result to
	// the first 3 bytes, or 6 hex digits, of the resulting cryptogram.
	//
	// This member is required.
	KeyCheckValue *string

	// The algorithm used for calculating key check value (KCV) for DES and AES keys.
	// For a DES key, Amazon Web Services Payment Cryptography computes the KCV by
	// encrypting 8 bytes, each with value '00', with the key to be checked and
	// retaining the 3 highest order bytes of the encrypted result. For an AES key,
	// Amazon Web Services Payment Cryptography computes the KCV by encrypting 8 bytes,
	// each with value '01', with the key to be checked and retaining the 3 highest
	// order bytes of the encrypted result.
	//
	// This member is required.
	KeyCheckValueAlgorithm KeyCheckValueAlgorithm

	// The source of the key material. For keys created within Amazon Web Services
	// Payment Cryptography, the value is AWS_PAYMENT_CRYPTOGRAPHY . For keys imported
	// into Amazon Web Services Payment Cryptography, the value is EXTERNAL .
	//
	// This member is required.
	KeyOrigin KeyOrigin

	// The state of key that is being created or deleted.
	//
	// This member is required.
	KeyState KeyState

	// The date and time after which Amazon Web Services Payment Cryptography will
	// delete the key. This value is present only when KeyState is DELETE_PENDING and
	// the key is scheduled for deletion.
	DeletePendingTimestamp *time.Time

	// The date and time after which Amazon Web Services Payment Cryptography will
	// delete the key. This value is present only when when the KeyState is
	// DELETE_COMPLETE and the Amazon Web Services Payment Cryptography key is deleted.
	DeleteTimestamp *time.Time

	// The date and time after which Amazon Web Services Payment Cryptography will
	// start using the key material for cryptographic operations.
	UsageStartTimestamp *time.Time

	// The date and time after which Amazon Web Services Payment Cryptography will
	// stop using the key material for cryptographic operations.
	UsageStopTimestamp *time.Time

	noSmithyDocumentSerde
}

// The role of the key, the algorithm it supports, and the cryptographic
// operations allowed with the key. This data is immutable after the key is
// created.
type KeyAttributes struct {

	// The key algorithm to be use during creation of an Amazon Web Services Payment
	// Cryptography key. For symmetric keys, Amazon Web Services Payment Cryptography
	// supports AES and TDES algorithms. For asymmetric keys, Amazon Web Services
	// Payment Cryptography supports RSA and ECC_NIST algorithms.
	//
	// This member is required.
	KeyAlgorithm KeyAlgorithm

	// The type of Amazon Web Services Payment Cryptography key to create, which
	// determines the classiﬁcation of the cryptographic method and whether Amazon Web
	// Services Payment Cryptography key contains a symmetric key or an asymmetric key
	// pair.
	//
	// This member is required.
	KeyClass KeyClass

	// The list of cryptographic operations that you can perform using the key.
	//
	// This member is required.
	KeyModesOfUse *KeyModesOfUse

	// The cryptographic usage of an Amazon Web Services Payment Cryptography key as
	// deﬁned in section A.5.2 of the TR-31 spec.
	//
	// This member is required.
	KeyUsage KeyUsage

	noSmithyDocumentSerde
}

// The list of cryptographic operations that you can perform using the key. The
// modes of use are deﬁned in section A.5.3 of the TR-31 spec.
type KeyModesOfUse struct {

	// Speciﬁes whether an Amazon Web Services Payment Cryptography key can be used to
	// decrypt data.
	Decrypt bool

	// Speciﬁes whether an Amazon Web Services Payment Cryptography key can be used to
	// derive new keys.
	DeriveKey bool

	// Speciﬁes whether an Amazon Web Services Payment Cryptography key can be used to
	// encrypt data.
	Encrypt bool

	// Speciﬁes whether an Amazon Web Services Payment Cryptography key can be used to
	// generate and verify other card and PIN verification keys.
	Generate bool

	// Speciﬁes whether an Amazon Web Services Payment Cryptography key has no special
	// restrictions other than the restrictions implied by KeyUsage .
	NoRestrictions bool

	// Speciﬁes whether an Amazon Web Services Payment Cryptography key can be used
	// for signing.
	Sign bool

	// Speciﬁes whether an Amazon Web Services Payment Cryptography key can be used to
	// unwrap other keys.
	Unwrap bool

	// Speciﬁes whether an Amazon Web Services Payment Cryptography key can be used to
	// verify signatures.
	Verify bool

	// Speciﬁes whether an Amazon Web Services Payment Cryptography key can be used to
	// wrap other keys.
	Wrap bool

	noSmithyDocumentSerde
}

// Metadata about an Amazon Web Services Payment Cryptography key.
type KeySummary struct {

	// Specifies whether the key is enabled.
	//
	// This member is required.
	Enabled *bool

	// Specifies whether the key is exportable. This data is immutable after the key
	// is created.
	//
	// This member is required.
	Exportable *bool

	// The Amazon Resource Name (ARN) of the key.
	//
	// This member is required.
	KeyArn *string

	// The role of the key, the algorithm it supports, and the cryptographic
	// operations allowed with the key. This data is immutable after the key is
	// created.
	//
	// This member is required.
	KeyAttributes *KeyAttributes

	// The key check value (KCV) is used to check if all parties holding a given key
	// have the same key or to detect that a key has changed. Amazon Web Services
	// Payment Cryptography calculates the KCV by using standard algorithms, typically
	// by encrypting 8 or 16 bytes or "00" or "01" and then truncating the result to
	// the first 3 bytes, or 6 hex digits, of the resulting cryptogram.
	//
	// This member is required.
	KeyCheckValue *string

	// The state of an Amazon Web Services Payment Cryptography that is being created
	// or deleted.
	//
	// This member is required.
	KeyState KeyState

	noSmithyDocumentSerde
}

// Parameter information for root public key certificate import.
type RootCertificatePublicKey struct {

	// The role of the key, the algorithm it supports, and the cryptographic
	// operations allowed with the key. This data is immutable after the root public
	// key is imported.
	//
	// This member is required.
	KeyAttributes *KeyAttributes

	// Parameter information for root public key certificate import.
	//
	// This member is required.
	PublicKeyCertificate *string

	noSmithyDocumentSerde
}

// A structure that contains information about a tag.
type Tag struct {

	// The key of the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag.
	Value *string

	noSmithyDocumentSerde
}

// Parameter information for trusted public key certificate import.
type TrustedCertificatePublicKey struct {

	// The KeyARN of the root public key certificate or certificate chain that signs
	// the trusted public key certificate import.
	//
	// This member is required.
	CertificateAuthorityPublicKeyIdentifier *string

	// The role of the key, the algorithm it supports, and the cryptographic
	// operations allowed with the key. This data is immutable after a trusted public
	// key is imported.
	//
	// This member is required.
	KeyAttributes *KeyAttributes

	// Parameter information for trusted public key certificate import.
	//
	// This member is required.
	PublicKeyCertificate *string

	noSmithyDocumentSerde
}

// Parameter information for generating a wrapped key using TR-31 or TR-34
// standard.
type WrappedKey struct {

	// Parameter information for generating a wrapped key using TR-31 or TR-34
	// standard.
	//
	// This member is required.
	KeyMaterial *string

	// The key block format of a wrapped key.
	//
	// This member is required.
	WrappedKeyMaterialFormat WrappedKeyMaterialFormat

	// The KeyARN of the wrapped key.
	//
	// This member is required.
	WrappingKeyArn *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isExportKeyMaterial() {}
func (*UnknownUnionMember) isImportKeyMaterial() {}