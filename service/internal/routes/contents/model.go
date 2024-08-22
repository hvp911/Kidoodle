package contents

// Device represents a device object
type Device struct {
	ID                 uint   `json:"id" example:"1"`                          // Returns ID of the device
	Name               string `json:"name,omitempty" example:"abc"`            // Returns name of the device
	ProtectionSystemID uint   `json:"protection_system,omitempty" example:"1"` // Returns ID of the protection system of the device
}

// ProtectionSystem represents a protection system object
type ProtectionSystem struct {
	ID             uint   `json:"id" example:"1"`                        // Returns ID of the protection system
	Name           string `json:"name,omitempty" example:"abc"`          // Returns name of the protection system
	EncryptionMode string `json:"encryption_mode,omitempty" example:"1"` // Returns encryption mode associated with the protection system
}

// Content represents a content object
type Content struct {
	ID                 uint   `json:"id" example:"1"`                            // Returns ID of the content
	ProtectionSystemID uint   `json:"protection_system,omitempty" example:"1"`   // Returns ID of the protection system of the content
	EncryptionKey      string `json:"encryption_key,omitempty" example:"abc"`    // Returns encryption key to decipher content
	EncryptedPayload   string `json:"encrypted_payload,omitempty" example:"abc"` // Returns encrypted payload
}

// DecryptedContent represents a decrypted content object
type DecryptedContent struct {
	ID                 uint   `json:"id" example:"1"`                          // Returns ID of the content
	ProtectionSystemID uint   `json:"protection_system,omitempty" example:"1"` // Returns ID of the protection system of the content
	Payload            string `json:"payload,omitempty" example:"abc"`         // Returns decrypted payload
}
