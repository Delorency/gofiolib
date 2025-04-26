package personservice

func (ps *personService) Delete(id uint) error {
	return ps.repo.Delete(id)
}
