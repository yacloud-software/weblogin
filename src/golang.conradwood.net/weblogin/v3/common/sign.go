package common

/*

This signs a userobjects' fields.
Note that the "public" part is in the common go-easyops "auth" package.
Amongst other things it contains the conversion from proto to bytes

*/
import (
	ed "crypto/ed25519"
	"fmt"
	os "golang.conradwood.net/apis/objectstore"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/tokens"
	"golang.conradwood.net/go-easyops/utils"
)

const (
	KEY_ID = "weblogin_private_key_seed"
)

var (
	private_seed []byte
	ostore       os.ObjectStoreClient
)

func InitKey() {
	if private_seed != nil && len(private_seed) == 32 {
		return
	}
	if ostore == nil {
		ostore = os.NewObjectStoreClient(client.Connect("objectstore.ObjectStore"))
	}
	ctx := authremote.Context()
	ctx = tokens.ContextWithToken()
	p, err := ostore.TryGet(ctx, &os.GetRequest{ID: KEY_ID})
	if err != nil {
		fmt.Printf("Cannot get key: %s\n", utils.ErrorString(err))
		panic("no key")
	}
	if p.DoesExist {
		fmt.Printf("Retrieved private key\n")
		private_seed = p.Object.Content
		return
	}

	nk := []byte(utils.RandomString(32))
	// create
	pe, err := ostore.PutIfNotExists(ctx, &os.PutWithIDRequest{ID: KEY_ID, Content: nk})
	if err != nil {
		panic(fmt.Sprintf("unable to create new key: %s", utils.ErrorString(err)))
	}
	if pe.WasAdded {
		private_seed = nk
		fmt.Printf("Created private key.\n")
		return
	}

}
func signPublicKey() ed.PublicKey {
	pk := signPrivateKey().Public().(ed.PublicKey)
	return pk
}
func signPrivateKey() ed.PrivateKey {
	InitKey()
	if len(private_seed) != 32 {
		panic(fmt.Sprintf("Private key seed has invalid length (%d != 32)", len(private_seed)))
	}
	pk := ed.NewKeyFromSeed(private_seed)
	return pk
}

// sign a string
func SignString(text string) []byte {
	b := []byte(text)
	s := ed.Sign(signPrivateKey(), b)
	return s
}
func Verify(message, signature []byte) bool {
	v := ed.Verify(signPublicKey(), message, signature)
	return v
}
