package domain

// Exchange
// Pair
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Pair -
*/
type Pair struct {
	MoneyAB           TwoMoney
	MinAmountAndDepth FinData64
}
