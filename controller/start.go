package controller

func Run() error {
	c := NewLogicService()
	if err := c.Run(); err != nil {
		return err
	}
	return nil
}
