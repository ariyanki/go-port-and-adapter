package user

func (s *userHandler) DeleteUser(id int) error {
	return s.userRepository.DeleteUser(id)
}
