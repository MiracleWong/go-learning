package map_test

import "testing"

func TestInitMap(t *testing.T) {

	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len of m1 is %d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Log(m2[2])
	t.Log(m2[4])
	t.Logf("len of m2 is %d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len of m3 is %d", len(m3))
}

func TestAccessNotExistsKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])

	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("key 3 value is %d", v)
	} else {
		t.Logf("key 3 value is not exists")
	}

}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
