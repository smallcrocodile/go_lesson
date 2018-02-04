package main

import "testing"

func TestTriangle(t *testing.T)  {
	tests :=[]struct{a,b,c int}{
		{3,4,5},
		{5,12,13},
		{8,15,17}
	}
	for _,tt :=range tests{
		if actual := calcTriangle(tt.a,tt.b);actual != tt.c{
			t.Error("calctriangle(%d,%d); got %d;expected:%d",tt.a,tt.b,actual,tt.c)
		}
	}
}

