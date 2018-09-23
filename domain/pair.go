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

//type Pair struct {
//	MoneyA            Money
//	MoneyB            Money
//	MinAmountAndDepth FinData64
//}
