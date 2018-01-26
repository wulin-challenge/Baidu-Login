package bdcrypto

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"math/big"
)

// RsaEncrypt 针对百度的 RSA 加密
func RsaEncrypt(rsaString string, origData []byte) ([]byte, error) {
	// var rsaString = "AE47B04D3A55A5FDABC612A426D84484BCB1C29C63BBAC33544A1BB94D930772E6E201CF2B39B5B6EDED1CCCBB5E4DCE713B87C6DD88C3DBBEE3A1FBE220723F01E2AA81ED9497C8FFB05FF54A3E982A76D682B0AABC60DBF9D1A8243FE2922E43DD5DF9C259442147BBF4717E5ED8D4C1BD5344DD1A8F35B631D80AB45A9BC7"
	var i = new(big.Int)
	_, ok := i.SetString(rsaString, 16)
	if !ok {
		return nil, errors.New("rsaString is invalid")
	}

	pub := &rsa.PublicKey{
		N: i,
		E: 0x10001,
	}

	ol := len(origData)
	reverseOrigData := make([]byte, ol)
	for k := range origData {
		reverseOrigData[ol-k-1] = origData[k]
	}

	c := new(big.Int).SetBytes(reverseOrigData)
	return c.Exp(c, big.NewInt(int64(pub.E)), pub.N).Bytes(), nil
}

// RsaDecrypt 解密(非针对百度)
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCuR7BNOlWl/avGEqQm2ESEvLHCnGO7rDNUShu5TZMHcubiAc8r
ObW27e0czLteTc5xO4fG3YjD277jofviIHI/AeKqge2Ul8j/sF/1Sj6YKnbWgrCq
vGDb+dGoJD/iki5D3V35wllEIUe79HF+XtjUwb1TRN0ajzW2MdgKtFqbxwIDAQAB
AoGAW0CoHFe9/tLq/SRHlRtKDSJsBRUz11Fb8vd2urjWkmDkaVQ/MEfgUK8Vpy2/
saoVvQ5JkqPud3b45WGsbINGrb8saugZ1h5huDbuxVXKDj1ZWyJPkmxHLUK2+7iL
5c7F7+v2C+n6polIgMV9SbLXD6YIXUJ+GengWQffhTRE7WECQQDj/g5x7Rj5vc7X
o3i0SQmyN4RcxxOWfiLe5OUASKM2UPVBQKI3CugkmiTaXTi7auuG3I4GVPRHVHw9
y/Ekz7J3AkEAw7B5+uI60MwcDMeGoXAMAEYe/s7LhyBICarY6cNwySb46B7OHEUz
ooFV2qx31I6ivpMRwCqrRKXEvjPEAfPlMQJAGrXi/1nluSyRlRXjyEteRXDXov73
voPclfx/D79yz6RAd3qZBpXSiKc+dg7B3MM0AMLKKNe/HrQ5Mgw4njVvFQJAPKvX
ddh0Qc42mCO4cw8JOYCEFZ5J7fAtRYoJzJhCvKrvmxAJ+SvfcW/GDZFRab57aLiy
VTElfpgiopHsIGrc0QJBAMliJywM9BNYn9Q4aqKN/dR22W/gctfa6bxU1m9SfJ5t
5G8MR8HsB/9Cafv8f+KnFzp2SncEu0zuFc/S8n5X5v0=
-----END RSA PRIVATE KEY-----
`))
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	c := new(big.Int).SetBytes(ciphertext)
	// return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	return c.Exp(c, priv.D, priv.N).Bytes(), nil
}

func Base64Decode(raw []byte) []byte {
	var buf bytes.Buffer
	decoded := make([]byte, 215)
	buf.Write(raw)
	decoder := base64.NewDecoder(base64.StdEncoding, &buf)
	decoder.Read(decoded)
	return decoded
}

func Base64Encode(raw []byte) []byte {
	var encoded bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &encoded)
	encoder.Write(raw)
	encoder.Close()
	return encoded.Bytes()
}
