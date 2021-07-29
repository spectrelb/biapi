package okex

type OKExSpot struct {
	*OKEx
}

func (ok *OKExSpot) GetAccount(name string) (string, error)  {
	_, _ = ok.Client.GetV5AccountBalance(&name)
	return "123", nil
}