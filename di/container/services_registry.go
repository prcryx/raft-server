package container

import "github.com/prcryx/raft-server/internal/infrastructure/jwt"

// import "github.com/prcryx/raft-server/internal/infrastructure/jwt"

type ServicesRegistry struct {
	JwtStrategy *jwt.JwtStrategy
}

func NewServicesRegistry(jwtStrategy *jwt.JwtStrategy) *ServicesRegistry {
	return &ServicesRegistry{
		JwtStrategy: jwtStrategy,
	}
}

func GetJwtService(serviceRegistry *ServicesRegistry) jwt.IJwtStrategy {
	return serviceRegistry.JwtStrategy
}
