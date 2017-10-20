package helpservice
import dlog "debuglogger"

func (p *HelpService) handler(msg string) error {
	dlog.Println("👮 ========>",msg)
	return nil
}