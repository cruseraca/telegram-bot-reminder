package bot

type service struct {
}

func (s *service) Ping() string {
	return "PONG"
}

func (s *service) Notify(cronjob, msg string) {
	//TODO implement me
	panic("implement me")
}
