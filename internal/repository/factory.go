package repository

// NewBaseRepository wires all individual repositories using shared access
func NewBaseRepository(access *RepositoryAccess) *BaseRepository {
	return &BaseRepository{
		// User: NewUserRepository(access),
		// Add others similarly:
		// Order: NewOrderRepository(access),
		// Auth: NewAuthRepository(access),
	}
}
