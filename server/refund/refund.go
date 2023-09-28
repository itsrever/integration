package refund

type Refund struct {
	Items []RefundRequestItem
}

type RefundRes struct {
	Items         []RefundRequestItem
	RefundID      string
	TransactionID string
}

type RefundRequestItem struct {
	LineItemId string
	Quantity   int32
	Amount     RefundRequestItemAmount
}

type RefundRequestItemAmount struct {
	Amount   float64
	Currency string
}

type refundManager struct {
	Refunds map[string]Refund
}

type RefundManager interface {
	CreateRefund(orderID string, req Refund)
	GetRefund(orderID string) Refund
}

func New() RefundManager {
	return &refundManager{
		Refunds: make(map[string]Refund),
	}
}

func (m *refundManager) CreateRefund(orderID string, req Refund) {
	curr := m.GetRefund(orderID)
	curr.Items = append(curr.Items, req.Items...)
	m.Refunds[orderID] = curr
}

func (m *refundManager) GetRefund(orderID string) Refund {
	if refund, ok := m.Refunds[orderID]; ok {
		return refund
	}
	return Refund{}
}
