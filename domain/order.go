package domain

// Exchange
// Order
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Pair -
*/
type Order struct {
	Pair    Pair
	Price   FinData64 // 56-bit rate, 8-bit buy/sell
	Amount  uint64
	Created int64 // time
}
