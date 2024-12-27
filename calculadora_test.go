package main

import "testing"

func TestSoma(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSoma2(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 7
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMultiplicacao(t *testing.T) {
	teste := multiplicacao(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMultiplicacao2(t *testing.T) {
	teste := multiplicacao(10, 10)
	resultado := 2560
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSubtracao(t *testing.T) {
	teste := subtracao(10, 5)
	resultado := -5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSubtracao2(t *testing.T) {
	teste := subtracao(10, 10)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestDivisao(t *testing.T) {
	teste := divisao(10)
	resultado := 10
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestDivisao2(t *testing.T) {
	teste := divisao(10)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
