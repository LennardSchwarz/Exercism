package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

// Returns a random big.Int (1, p)
func PrivateKey(p *big.Int) *big.Int {
	// We want to generate a number in the range [0, p-3] and then add 2
	max := new(big.Int).Sub(p, big.NewInt(3))

	n, _ := rand.Int(rand.Reader, max)

	// Add 1 to the result to ensure it is in the range (1, p)
	result := n.Add(n, big.NewInt(2))
	return result
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	// Convert g to big.Int
	bigG := big.NewInt(g)

	// Allocate a big.Int for the result
	public := new(big.Int)

	// Calculate g^private mod p
	public.Exp(bigG, private, p)

	return public
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	privateKey := PrivateKey(p)
	publicKey := PublicKey(privateKey, p, g)
	return privateKey, publicKey

}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	// Allocate a big.Int for the result
	secretKey := new(big.Int)

	// calculate public^private mod p
	secretKey.Exp(public2, private1, p)

	return secretKey
}
